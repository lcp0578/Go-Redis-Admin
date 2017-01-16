package redisserver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"Go-Redis-Admin/src/common/request"
	"Go-Redis-Admin/src/common/response"
)

type redisapi struct {
	input request.Input
	output response.Output
	Redisconn redis.Conn
}


func (re *redisapi) Connect(w http.ResponseWriter,r *http.Request) {
	if r.Method!="POST"{
		fmt.Println("请求错误")
	}
	body,_:=re.input.InputBody()
	fmt.Println(body)
	re.input.Request.ParseForm()
	var network string=r.FormValue("network")
	var address string=r.FormValue("address")
	var password string=r.FormValue("password")
	option:=redis.DialOption{}
	if password!=""{
		option=redis.DialPassword(password)
	}
	conn, err := redis.Dial(network, address,option)
	if err != nil {
		fmt.Println("connect error:", err.Error())
	} else {
		fmt.Println("connect redis success")
	}
	re.Redisconn = conn
	re.output.ResponseWriter.WriteHeader(200)
	//c:=NewConnection(conn)
	//conn.Close()
}
func (re *redisapi) Ping() {
	reply,err:=re.Redisconn.Do("PING","test")
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(redis.String(reply,err))
	}
}
func (re *redisapi)Set(){
	var args []interface{}
	reply,err:=re.Redisconn.Do("SET",args)
	redis.String(reply,err)
	if err!=nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}
}

func (re *redisapi) Get(){
	var key string
	reply,err:=re.Redisconn.Do("GET",key)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}

}

func (re *redisapi) Delete(){
	var args []interface{}
	reply,err:=re.Redisconn.Do("DEL",args)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(reply)
	}
}
