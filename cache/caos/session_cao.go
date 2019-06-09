package caos

import (
	"fmt"
	"time"

	"DuyStifler/GolangServer/cache"
	"DuyStifler/GolangServer/utils"
)

const (
	CACHE_KEY_USER_SESSION = "user_session"
)

type SessionCAO struct {
	cache *cache.Cache
}

func NewSessionCAO(cache *cache.Cache) *SessionCAO {
	return &SessionCAO{
		cache: cache,
	}
}

func (c *SessionCAO) generateUserSessionKey() (string, error) {
	sessionKey, err := utils.GenerateUUID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s:%s", c.cache.GetPrefixCache(), CACHE_KEY_USER_SESSION, sessionKey), nil
}

func (c *SessionCAO) CreateUseSession(userID string) error {
	sessionKey, err := c.generateUserSessionKey()
	if err != nil {
		return err
	}

	c.cache.Client().Set(sessionKey, userID, c.getExpiredTime())
	return nil
}

func (c *SessionCAO) RemoveUserSession(sessionKey string) {
	c.cache.Client().Del(sessionKey)
}

func (c *SessionCAO) ExtendUserSession(sessionKey string, userID string) {
	c.cache.Client().Expire(sessionKey, c.getExpiredTime())
}

func (c *SessionCAO) getExpiredTime() time.Duration {
	return time.Duration(c.cache.ServerConfig().ExpiredTime) * time.Second
}