package routers

import (
	"DuyStifler/GolangServer/http_server"
	"DuyStifler/GolangServer/manager"
	"DuyStifler/GolangServer/server/game/handlers"
)

type GameRouter struct {
	groupUrl    string
	server      *http_server.HttpServer
	gameHandler *handlers.GameHandler

}

func NewGameRouter(groupUrl string, server *http_server.HttpServer, gameManager *manager.Manager) *GameRouter {
	gameHandler := handlers.NewGameHandler(gameManager, server)

	return &GameRouter{
		groupUrl,
		server,
		gameHandler,
	}
}

func (r *GameRouter) InitApi() {
	group := r.server.E().Group(r.groupUrl)

	//demo api
	group.GET("/demo", r.gameHandler.DemoApi)
}
