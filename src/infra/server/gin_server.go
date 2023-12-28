package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handler "goHexBoilerplate/src/adapters/handlers"
	"goHexBoilerplate/src/domain/contracts/server"
)

type GinServer struct {
	Router *gin.Engine
	server.AbstractServer
	userHandler handler.UserHandler
}

func NewGinServer(port int, userHandler handler.UserHandler) *GinServer {
	fmt.Println("Hello, World!")
	server := GinServer{gin.Default(), server.AbstractServer{Port: port}, userHandler}
	server.AbstractServer.Server = &server
	return &server
}

func (server *GinServer) Listen() {
	server.setAppHandlers(server.Router)
	err := server.Router.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		return
	}
}

func (server *GinServer) setAppHandlers(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/users/:id", server.userHandler.ReadUser)
	v1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
