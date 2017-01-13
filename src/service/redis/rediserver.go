package redisserver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type redisapi struct {
	redisconn redis.Conn
}

func (re *redisapi) Connect(network string, addredd string) {
	conn, err := redis.Dial(network, addredd)
	if err != nil {
		fmt.Println("connect error:", err.Error())
	} else {
		fmt.Println("connect redis success")
	}
	re.redisconn = conn
	//c:=NewConnection(conn)
	//conn.Close()
}
func (re *redisapi) ping() {
	err := redis.PubSubConn{}.Ping("test")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (re *redisapi) Get(args ...interface{}){
	reply,err:=re.redisconn.Do("GET",args)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}

}
