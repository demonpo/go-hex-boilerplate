package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goHexBoilerplate/src/domain/contracts/server"
)

type GinServer struct {
	Router *gin.Engine
	server.AbstractServer
}

func NewGinServer(port int) *GinServer {
	fmt.Println("Hello, World!")
	server := GinServer{gin.Default(), server.AbstractServer{Port: port}}
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
	router.GET("/v1/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
