package fx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	domainServer "goHexBoilerplate/src/domain/contracts/server"
	"goHexBoilerplate/src/infra/server"
)

func NewGinServer(config AppConfig) *server.GinServer {
	fmt.Println("Hello, World!")
	serv := server.GinServer{Router: gin.Default(), AbstractServer: domainServer.AbstractServer{Port: config.Port}}
	serv.AbstractServer.Server = &serv
	return &serv
}
