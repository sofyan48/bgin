package moduls

import (
	"fmt"

	"github.com/meongbego/bgin/app/libs"
	"github.com/sirupsen/logrus"

	"github.com/garyburd/redigo/redis"
)

// Store for global connect to cache
var Store redis.Conn

// Data type
type Data struct{}

// InitRedis func
func InitRedis() redis.Conn {
	rdhost := libs.GetEnvVariabel("REDIS_HOST", "localhost")
	rdport := libs.GetEnvVariabel("REDIS_PORT", "6379")
	c, err := redis.Dial("tcp", fmt.Sprintf(
		"%s:%s", rdhost, rdport))
	if err != nil {
		logrus.Errorf("failed to connect to database: %v", err)
	}
	return c
}

// RowsCached func
func RowsCached(keys string, d []byte, ttl int) ([]byte, error) {
	_, err := redis.String(Store.Do("SET", keys, d))
	if err != nil {
		return nil, err
	}
	redis.String(Store.Do("EXPIRE", keys, ttl))
	return d, nil
}

// GetRowsCached func
func GetRowsCached(keys string) (string, error) {
	value, err := redis.String(Store.Do("GET", keys))
	if err != nil {
		return "", err
	}
	return value, nil
}
