module github.com/paysuper/paysuper-1c-gateway

go 1.13

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/go-playground/validator v9.30.0+incompatible // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/client/selector/static v0.0.0-20200119172437-4fe21aa238fd
	github.com/paysuper/paysuper-proto/go/billingpb v0.0.0-20200213094344-132dbbef912e
	go.uber.org/zap v1.13.0
	gopkg.in/go-playground/validator.v9 v9.30.0
)

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
