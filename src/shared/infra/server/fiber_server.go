package server

import (
	"fmt"
	"goHexBoilerplate/src/modules/user/application/rest/handlers"
	"goHexBoilerplate/src/shared/contracts/server"
	"goHexBoilerplate/src/shared/infra/fx"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	App *fiber.App
	server.AbstractServer
	UserHandler *handlers.UserHandler
}

func NewFiberServer(config fx.AppConfig, userHandler *handlers.UserHandler) *FiberServer {
	app := fiber.New()
	serv := FiberServer{App: app, AbstractServer: server.AbstractServer{Port: config.Port}, UserHandler: userHandler}
	serv.AbstractServer.Server = &serv
	return &serv
}

func (server *FiberServer) Listen() {
	server.setAppHandlers(server.App)
	if err := server.App.Listen(fmt.Sprintf(":%d", server.Port)); err != nil {
		return
	}
}

func (server *FiberServer) setAppHandlers(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/users/:id", func(c *fiber.Ctx) error { return server.UserHandler.ReadUser(c) })
	v1.Post("/users", func(c *fiber.Ctx) error { return server.UserHandler.CreateUser(c) })
	v1.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "OK"})
	})
}
