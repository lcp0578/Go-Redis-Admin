package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
type redisapi struct{
}
func (c *redisapi) Connect() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect error:", err.Error())
	} else {
		fmt.Println("connect redis success")
	}
	r, err_do := conn.Do("set", "age", 24)
	if err_do != nil {
		fmt.Println("do error:", err.Error())
	} else {
		fmt.Println(r)
	}
	//conn.Close()
}
