package handlers

import (
	"github.com/labstack/echo"
	"net/http"

	"DuyStifler/GolangServer/http_server"
	"DuyStifler/GolangServer/manager"
)

type GameHandler struct {
	gameManager *manager.Manager
	server      *http_server.HttpServer
}

func NewGameHandler(gameManager *manager.Manager, server *http_server.HttpServer) *GameHandler {
	return &GameHandler{
		gameManager: gameManager,
		server:      server,
	}
}

func (gh *GameHandler) DemoApi(c echo.Context) (err error) {
	scope := gh.server.GetScope(c)
	scope.UserID()
	return c.NoContent(http.StatusOK)
}
