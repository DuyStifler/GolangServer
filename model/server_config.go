package model

import (
	"github.com/labstack/echo/middleware"
)

type ServerConfig struct {
	SessionExpired int                  `json:"sessionExpired"`
	Database       ServerConfigDatabase `json:"database"`
	Cache          ServerConfigCache    `json:"cache"`
}

type ServerConfigDatabase struct {
	UrlMaster  string `json:"url_master"`
	UrlReplica []string
	Port       int    `json:"port"`
	Username   string `json:"user"`
	Password   string `json:"password"`
	Database   string `json:"database"`

	UrlReplicasString string `json:"url_replicas"`
}

type ServerConfigCache struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type MiddlewareAuthConfig struct {
	Skipper middleware.Skipper
}