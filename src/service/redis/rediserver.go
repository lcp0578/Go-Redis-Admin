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
func (re *redisapi) Ping() {
	reply,err:=re.redisconn.Do("PING","test")
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(redis.String(reply,err))
	}
}
func (re *redisapi)Set(args ...interface{}){
	reply,err:=re.redisconn.Do("SET",args)
	redis.String(reply,err)
	if err!=nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}
}

func (re *redisapi) Get(args interface{}){
	reply,err:=re.redisconn.Do("GET",args)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}

}

func (re *redisapi) Delete(args ...interface{}){
	reply,err:=re.redisconn.Do("DEL",args)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}
}
