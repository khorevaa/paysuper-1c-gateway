package main

import (
	"github.com/kelseyhightower/envconfig"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/paysuper/paysuper-1c-gateway/internal"
	"go.uber.org/zap"
)

const ApplicationName = "PAYSUPER-1C-API"

type Config struct {
	ServerPort  int `envconfig:"SERVER_PORT" required:"false" default:"80"`
	MetricsPort int `envconfig:"METRICS_PORT" required:"false" default:"81"`
}

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger.Named(ApplicationName))

	defer logger.Sync()

	config := &Config{}
	if err := envconfig.Process("GW", config); err != nil {
		zap.L().Fatal("Config init failed with error", zap.Error(err))
	}

	service := internal.NewService(config.ServerPort, config.MetricsPort)
	service.Run()
}
