package server

type Server interface {
	Listen()
}

type AbstractServer struct {
	port int
	Server
}

type App struct {
	appName string
	server  *AbstractServer
}

func NewApp() *App {
	return &App{}
}

func (app *App) Name(name string) *App {
	app.appName = name
	return app
}

func (app *App) Server(server *AbstractServer) *App {
	app.server = server
	return app
}

func (app *App) Start() {
	println("Listening")
	app.server.Listen()
}
