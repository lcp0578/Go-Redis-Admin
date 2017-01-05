package redisserver_test

import (
	"testing"
	"github.com/garyburd/redigo/redis"
)

func TestConnect(t *testing.T) {
	redis.Dial("tcp", "127.0.0.1:6379")

}
