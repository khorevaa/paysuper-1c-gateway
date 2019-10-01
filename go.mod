module github.com/paysuper/paysuper-1c-gateway

go 1.12

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/go-playground/validator v9.30.0+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/echo/v4 v4.1.6
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
	github.com/paysuper/paysuper-billing-server v0.0.0-20190930122133-08e4a0cd855b
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190927123631-a832865fa7ad // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
