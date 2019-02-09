package cache

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/logger"
)

type Cache struct {
	Logger *logger.Logger
	Config *ConfigCache
	client *redis.Client
}

type ConfigCache struct {
	Ip string
	Port int
}

func NewRedisCache(config *ConfigCache, logger *logger.Logger) error {
	cache := &Cache{
		Logger: logger,
		Config: config,
	}

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cache.Config.Ip, cache.Config.Port),
	})

	_, err := client.Ping().Result()
	if err != nil {
		cache.Logger.Error("[ERROR - cache] - ping cache %s ", err)
		return errors.New(fmt.Sprintf("[ERROR - cache] - ping cache %s ", err))
	}

	cache.client = client

	return nil
}
