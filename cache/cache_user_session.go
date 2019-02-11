package cache

import (
	"DuyStifler/GolangServer/keys"
	"DuyStifler/GolangServer/utils"
)

type CUserSession struct {
	UserID  string
	Seq     int
}

func (c *Cache) GetUserSession(sessionKey string) (*CUserSession, error) {
	obj := &CUserSession{}
	key := keys.CACHE_USER_SESSION + sessionKey

	value := c.client.HMGet(key, "user_id", "seq")
	if value.Err() != nil {
		return obj, value.Err()
	} else {
		//UserID
		if value.Val()[0] != nil {
			obj.UserID = value.Val()[0].(string)
		}
		//Seq
		if value.Val()[1] != nil {
			obj.Seq = value.Val()[1].(int)
		}
	}

	return obj, nil
}

func (c *Cache) UpdateSession(sessionKey string) error {
	key := keys.CACHE_USER_SESSION + sessionKey
	return c.client.Expire(key, keys.DEFAULT_EXPIRE_SESSION).Err()
}

func (c *Cache) CreateSession(userID string) error {
	sessionKey := utils.GenerateSessionToken()
	key := keys.CACHE_USER_SESSION + sessionKey

	obj := &CUserSession{
		UserID: userID,
		Seq: 0,
	}

	stt := c.client.HMSet(key, obj.convertToMap())
	if stt.Err() != nil {
		return nil
	}

	return c.client.Expire(key, keys.DEFAULT_EXPIRE_SESSION).Err()
}

func (c *CUserSession) convertToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["user_id"] = c.UserID
	m["seq"] = c.Seq
	return m
}


