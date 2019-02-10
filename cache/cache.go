package cache

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

type Cache struct {
	Config *ConfigCache
	client *redis.Client
}

type ConfigCache struct {
	Ip string
	Port int
}

func NewRedisCache(config *ConfigCache) (Cache, error) {
	cache := &Cache{
		Config: config,
	}

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cache.Config.Ip, cache.Config.Port),
	})

	_, err := client.Ping().Result()
	if err != nil {
		return Cache{}, errors.New(fmt.Sprintf("[ERROR - cache] - ping cache %s ", err))
	}

	cache.client = client

	return cache, nil
}
