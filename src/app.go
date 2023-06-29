package main

import (
	"go.uber.org/fx"
	domainServer "goHexBoilerplate/src/domain/contracts/server"
	infraFx "goHexBoilerplate/src/infra/fx"
)

func main() {
	fx.New(
		fx.Provide(
			func() infraFx.AppConfig { return infraFx.AppConfig{Port: 3000} },
			infraFx.NewApp,
			fx.Annotate(
				infraFx.NewGinServer,
				fx.As(new(domainServer.Server)),
			),
		),
		fx.Invoke(func(app *domainServer.App) {}),
	).Run()
}
