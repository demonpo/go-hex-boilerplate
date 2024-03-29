package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"goHexBoilerplate/src/application/rest/handlers"
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/domain/contracts"
	domainRepositories "goHexBoilerplate/src/domain/contracts/repositories"
	domainServer "goHexBoilerplate/src/domain/contracts/server"
	"goHexBoilerplate/src/domain/services"
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
			handlers.NewUserHandler,
			infraFx.NewApp,
			fx.Annotate(
				server.NewGinServer,
				fx.As(new(domainServer.Server)),
			),
		),
		fx.Invoke(func(app *contracts.App) {}),
	).Run()
}
