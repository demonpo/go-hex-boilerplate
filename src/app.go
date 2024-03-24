package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	handler "goHexBoilerplate/src/adapters/handlers"
	"goHexBoilerplate/src/domain/contracts"
	domainRepositories "goHexBoilerplate/src/domain/contracts/repositories"
	domainServer "goHexBoilerplate/src/domain/contracts/server"
	"goHexBoilerplate/src/domain/services"
	"goHexBoilerplate/src/infra/db"
	infraFx "goHexBoilerplate/src/infra/fx"
	"goHexBoilerplate/src/infra/repositories"
	"goHexBoilerplate/src/infra/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fx.New(
		fx.Provide(
			func() infraFx.AppConfig { return infraFx.AppConfig{Port: 3000} },
			db.NewDB,
			fx.Annotate(
				repositories.NewPostgresUserRepository,
				fx.As(new(domainRepositories.UserRepository)),
			),
			services.NewUserService,
			handler.NewUserHandler,
			infraFx.NewApp,
			fx.Annotate(
				server.NewGinServer,
				fx.As(new(domainServer.Server)),
			),
		),
		fx.Invoke(func(app *contracts.App) {}),
	).Run()
}
