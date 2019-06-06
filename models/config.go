package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ServerConfig struct {
	ExpiredTime int             `json:"expired_time"`
	Version     string          `json:"version"`
	ServerSSl   ServerSSlConfig `json:"ssl"`
	Port        int             `json:"port"`
	Database    DatabaseConfig  `json:"database"`
	Cache       CacheConfig     `json:"cache"`
	Log         LogConfig       `json:"log_config"`
}

type ServerSSlConfig struct {
	IsEnabled bool   `json:"is_enabled"`
	KeyDir    string `json:"key_dir"`
	CertDir   string `json:"cert_dir"`
}

type DatabaseConfig struct {
	UrlMaster   string   `json:"url_master"`
	UrlReplicas []string `json:"url_replicas"`
	UserName    string   `json:"user_name"`
	Password    string   `json:"password"`
	Schema      string   `json:"schema"`
	Port        string   `json:"string"`
}

type CacheConfig struct {
	Url  string `json:"url"`
	Port string `json:"port"`
}

type LogConfig struct {
	LogInfoDir  string `json:"log_info_dir"`
	LogErrorDir string `json:"log_error_dir"`
}

func NewServerConfig(jsonFile *os.File) (*ServerConfig, error) {
	bytesValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	serverConfig := &ServerConfig{}
	err = json.Unmarshal(bytesValue, serverConfig)
	if err != nil {
		return nil, err
	}

	return serverConfig, nil
}
