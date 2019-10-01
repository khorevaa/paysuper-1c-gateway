package internal

import (
	"context"
	"fmt"
	"github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/client/selector/static"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

const (
	LimitDefault  = 1000
	OffsetDefault = 0
)

type Service struct {
	echoServer   *echo.Echo
	httpServer   *http.Server
	billing      grpc.BillingService
	httpPort     int
	metricPort   int
	authUser     string
	authPassword string
}

func NewService(httpPort int, metricPort int, user string, password string) *Service {
	return &Service{httpPort: httpPort, metricPort: metricPort, authUser: user, authPassword: password}
}

func (s *Service) Run() {
	s.echoServer = echo.New()
	s.echoServer.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		return username == s.authUser && password == s.authPassword, nil
	}))
	s.echoServer.GET("/transactions", s.listOrders)

	s.httpServer = s.initMetricsAndHealthCheck()

	zap.L().Info("Initialize micro service")
	options := s.createMicroOptions()

	clientService := micro.NewService(options...)
	clientService.Init()

	s.billing = grpc.NewBillingService(pkg.ServiceName, clientService.Client())
	if err := clientService.Run(); err != nil {
		zap.L().Fatal("Can`t run service", zap.Error(err))
	}
}

func (s *Service) listOrders(ctx echo.Context) error {
	req := &grpc.ListOrdersRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorRequestParamsIncorrect)
	}

	if req.Limit <= 0 || req.Limit > LimitDefault {
		req.Limit = LimitDefault
	}

	if req.Offset <= 0 {
		req.Offset = OffsetDefault
	}

	err = ctx.Validate(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorRequestParamsIncorrect)
	}

	rsp, err := s.billing.FindAllOrdersPrivate(ctx.Request().Context(), req)

	if err != nil {
		zap.L().Error(
			pkg.ErrorGrpcServiceCallFailed,
			zap.Error(err),
			zap.Any("request", req),
		)

		return echo.NewHTTPError(http.StatusInternalServerError, errorUnknown)
	}

	if rsp.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(rsp.Status), rsp.Message)
	}

	return ctx.JSON(http.StatusOK, rsp.Item)
}

func (s *Service) createMicroOptions() []micro.Option {
	options := []micro.Option{
		micro.BeforeStart(func() error {
			go s.beforeRun()
			return nil
		}),
		micro.AfterStop(func() error {
			s.afterStop()
			return nil
		}),
	}

	if os.Getenv("MICRO_SELECTOR") == "static" {
		zap.L().Info("Use micro selector `static`")
		options = append(options, micro.Selector(static.NewSelector()))
	}

	return options
}

func (s *Service) beforeRun() {
	zap.L().Info("Starting server", zap.Int("port", s.httpPort))
	if err := s.echoServer.Start(fmt.Sprintf(":%d", s.httpPort)); err != nil {
		zap.L().Info("Starting API server failed", zap.Error(err))
	}

	zap.L().Info("Starting metrics and health check listening", zap.Int("port", s.metricPort))
	if err := s.httpServer.ListenAndServe(); err != nil {
		zap.L().Fatal("Metrics and health check listen failed", zap.Error(err))
	}
}

func (s *Service) afterStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.echoServer.Shutdown(ctx)
	if err != nil {
		zap.L().Fatal("Echo server shutdown failed", zap.Error(err))
	}
	zap.L().Info("Echo server stopped")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		zap.L().Fatal("Http server shutdown failed", zap.Error(err))
	}

	zap.L().Info("Metric server stopped")
}

func (s *Service) Status() (interface{}, error) {
	return "OK", nil
}

func (s *Service) initMetricsAndHealthCheck() *http.Server {
	h := health.New()
	err := h.AddChecks([]*health.Config{
		{
			Name:     "health-check",
			Checker:  s,
			Interval: time.Duration(1) * time.Second,
			Fatal:    true,
		},
	})
	if err != nil {
		zap.L().Fatal("Health check register failed", zap.Error(err))
	}

	router := http.NewServeMux()
	router.HandleFunc("/health", handlers.NewJSONHandlerFunc(h, nil))

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.metricPort),
		Handler: router,
	}
}
