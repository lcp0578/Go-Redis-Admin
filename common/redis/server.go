package redis

import (
	"github.com/garyburd/redigo/redis"
)

func test() {
	redis.Conn().Do("APPEND", "test", 1024)
}
