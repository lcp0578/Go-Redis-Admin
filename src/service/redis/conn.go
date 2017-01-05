package redisserver

import (
	"log"

	"github.com/garyburd/redigo/redis"

)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

type connection struct {
	//The redis connection.
	redisconn redis.Conn


}

func NewConnection(redisconn redis.Conn) *connection {
	return &connection{
		redisconn: redisconn,
	}
}

