package main

import (
	"DuyStifler/GolangServer/cache"
	"github.com/labstack/gommon/log"
	"os"

	httpServerPackage "DuyStifler/GolangServer/http_server"
	managerPackage "DuyStifler/GolangServer/manager"
	"DuyStifler/GolangServer/models"
	"DuyStifler/GolangServer/server/game/routers"
)

const (
	CONFIG_FILE_DIR = "server/game/config.json"
)

var (
	httpServer *httpServerPackage.HttpServer
	manager    *managerPackage.Manager
)

func main() {
	serverConfig, err := getConfig()
	if err != nil {
		log.Fatal("error ", err)
	}

	serverCache := cache.NewCache(serverConfig)
	manager = managerPackage.NewManager(serverCache)
	httpServer = httpServerPackage.NewHttpServer(serverConfig)

	router := routers.NewGameRouter(serverConfig.GroupUrl, httpServer, manager)
	router.InitApi()

	httpServer.Run()
}

func getConfig() (*models.ServerConfig, error) {
	configFileJson, err := os.Open(CONFIG_FILE_DIR)
	if err != nil {
		return nil, err
	}

	serverConfig, err := models.NewServerConfig(configFileJson)
	if err != nil {
		return nil, err
	}

	return serverConfig, nil
}
