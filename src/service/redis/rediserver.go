package redisserver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"Go-Redis-Admin/src/common/response"
)

type redisapi struct {
	user string
	redisconn redis.Conn
}

var data = map[string]string{}
var result = &response.JsonResponse{
	0,
	"faild",
	data,
}
func (re *redisapi) Connect(w http.ResponseWriter,r *http.Request) {

	if r.Method!="POST"{
		result.Code=-1
		result.Msg="请求错误"
		response.OuputJson(w,result)
		fmt.Println("请求错误")
	}
	r.ParseForm()
	var network string=r.FormValue("network")
	var addredd string=r.FormValue("addredd")
	var password string=r.FormValue("password")
	option:=redis.DialOption{}
	if password!=""{
		option=redis.DialPassword(password)
	}
	conn, err := redis.Dial(network, addredd,option)
	if err != nil {
		fmt.Println("connect error:", err.Error())
	} else {
		fmt.Println("connect redis success")
	}
	re.redisconn = conn
	result.Code=1
	result.Msg="sucess"
	response.OuputJson(w,result)
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
