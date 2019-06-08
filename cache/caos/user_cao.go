package caos

import "DuyStifler/GolangServer/cache"

type UserCAO struct {
	cache *cache.Cache
}

func NewUserCAO(cache *cache.Cache) *UserCAO {
	return &UserCAO{
		cache: cache,
	}
}