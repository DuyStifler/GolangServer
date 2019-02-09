package cache

import "DuyStifler/GolangServer/keys"

type CUserSession struct {
	UserID  string
	Session string
	Seq     int
}

func (c *Cache) GetUserSession(sessionKey string) (*CUserSession, error) {
	obj := &CUserSession{}
	key := keys.CACHE_USER_SESSION + sessionKey

	value := c.client.HMGet(key, "user_id", "seq")
	if value.Err() != nil {
		return obj, value.Err()
	} else {
		if value.Val()[0] != nil {
			obj.UserID = value.Val()[0].(string)
		}
		if value.Val()[1] != nil {
			obj.Seq = value.Val()[1].(int)
		}
	}

	return obj, nil
}

func (obj *CUserSession) convertFromObjectToString(value CUserSession) string {
	panic("implement me")
}

func (obj *CUserSession) convertFromStringToObject(str string) CUserSession {
	panic("implement me")
}
