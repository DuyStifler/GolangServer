package cache

import (
	"fmt"

	"DuyStifler/GolangServer/models"

	"github.com/go-redis/redis"
)

type Cache struct {
	serverConfig *models.ServerConfig
	client       *redis.Client
}

func NewCache(serverConfig *models.ServerConfig) *Cache {
	cache := &Cache{
		serverConfig: serverConfig,
	}

	return cache
}

func (c *Cache) Connect() {
	c.client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", c.serverConfig.Cache.Url, c.serverConfig.Cache.Port),
		Password: "",
		DB: 0,
	})
}


func (c *Cache) Client() *redis.Client {
	return c.client
}