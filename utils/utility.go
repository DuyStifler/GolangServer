package utils

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"

	"DuyStifler/GolangServer/cache"
	"DuyStifler/GolangServer/database"
	"DuyStifler/GolangServer/model"
)

func GetConfig() (*cache.ConfigCache, *database.ServerDatabase, error) {
	serverConfig, err := ReadJsonConfig()
	if err != nil {
		return &cache.ConfigCache{}, &database.ServerDatabase{}, err
	}

	objCache := cache.ConfigCache{
		Ip: serverConfig.Cache.Ip,
		Port: serverConfig.Cache.Port,
	}

	arrReplicas := strings.Split(serverConfig.Database.UrlReplicasString, ",")
	db := database.ServerDatabase{
		MySQLUser: serverConfig.Database.Username,
		MySQLReplicaURL: arrReplicas,
		MySQLDatabase: serverConfig.Database.Database,
		MySQLMasterURL: serverConfig.Database.UrlMaster,
		MySQLPort: strconv.Itoa(serverConfig.Database.Port),
		MySQLPassword: serverConfig.Database.Password,
	}

	return &objCache, &db, nil
}

func ReadJsonConfig() (*model.ServerConfig, error) {
	str, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		return &model.ServerConfig{}, err
	}

	config := &model.ServerConfig{}
	err = json.Unmarshal([]byte(str), config)
	if err != nil {
		return &model.ServerConfig{}, err
	}

	return config, nil
}

func GenerateSessionToken() string {
	return uuid.Must(uuid.NewV4()).String()
}