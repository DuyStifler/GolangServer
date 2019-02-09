package cache

import (
	"strings"

	"DuyStifler/GolangServer/keys"
)

type CIpCluster struct {
	Addrs []string
}

func (c *Cache) ClearIpCluster() error {
	obj := &CIpCluster{}
	err := c.client.Set(keys.CACHE_IP_CLUSTER, obj.convertFromObjectToString(), 0).Err()
	return err
}

func(c *Cache) AppendIpToCluster(value *CIpCluster) error {
	str := value.convertFromObjectToString()
	stt := c.client.Set(keys.CACHE_IP_CLUSTER, str, 0)
	return stt.Err()
}

func(c *Cache) GetAllIpOnCluster() (*CIpCluster, error) {

	obj := &CIpCluster{}
	str, err := c.client.Get(keys.CACHE_IP_CLUSTER).Result()
	if err != nil {
		return obj, err
	}

	return obj.convertFromStringToObject(str), nil
}

func(obj *CIpCluster) convertFromObjectToString() string {
	return strings.Join(obj.Addrs, ",")
}

func(obj *CIpCluster) convertFromStringToObject(str string) *CIpCluster {
	arr := strings.Split(str, ",")
	for _, value := range arr {
		obj.Addrs = append(obj.Addrs, value)
	}
	return obj
}
