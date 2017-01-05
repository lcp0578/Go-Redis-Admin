package redisserver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
type redisapi struct{

}
func (this *redisapi) Connect(network string,addredd string) {
	conn, err := redis.Dial(network, addredd)
	if err != nil {
		fmt.Println("connect error:", err.Error())
	} else {
		fmt.Println("connect redis success")
	}
	c:=NewConnection(conn)
	_=c
	//conn.Close()
}
