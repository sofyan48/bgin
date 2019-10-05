package moduls

import (
	"fmt"

	"github.com/meongbego/bgin/app/libs"

	"github.com/garyburd/redigo/redis"
)

var Store redis.Conn

func InitRedis() redis.Conn {
	rdhost := libs.GetEnvVariabel("REDIS_HOST", "localhost")
	rdport := libs.GetEnvVariabel("REDIS_PORT", "6379")
	c, err := redis.Dial("tcp", fmt.Sprintf(
		"%s:%s", rdhost, rdport))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	return c
}
