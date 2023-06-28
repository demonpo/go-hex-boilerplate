package main

import (
	"goHexBoilerplate/src/infra/server"
)

func main() {
	var app = server.NewApp()
	var ginServer = server.NewGinServer(3000)
	app.Name("Boilerplate Service")
	app.Server(&ginServer.AbstractServer)
	app.Start()
	//r := gin.Default()
	//var cam int
	//fmt.Println(cam)
	//fmt.Println(cam)
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()
}
