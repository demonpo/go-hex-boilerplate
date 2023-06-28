package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
	AbstractServer
}

func NewGinServer(port int) *GinServer {
	fmt.Println("Hello, World!")
	server := GinServer{gin.Default(), AbstractServer{port: port}}
	server.AbstractServer.Server = &server
	return &server
}

func (server *GinServer) Listen() {
	server.setAppHandlers(server.router)
	err := server.router.Run(fmt.Sprintf(":%d", server.port))
	if err != nil {
		return
	}
}

func (server *GinServer) setAppHandlers(router *gin.Engine) {
	router.GET("/v1/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
