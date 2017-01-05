package v1

import (
	"Go-Redis-Admin/src/common/mysql"
	"Go-Redis-Admin/src/common/response"
	"fmt"
	"log"
	"net/http"
	// "reflect"
	//"strings"
)

func (h *Handlers) LoginAction(w http.ResponseWriter, r *http.Request) {
	// init return
	var data = map[string]string{}
	jr := &response.JsonResponse{
		0,
		"faild",
		data,
	}
	if r.Method != "POST" {
		jr.Code = 2
		jr.Msg = "请求出错啦"
		response.OuputJson(w, jr)
		return
	}
	log.Println("API Login Action")
	r.ParseForm() //解析参数，默认是不会解析的
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	// fmt.Println("username", r.FormValue("username"))
	// fmt.Println("password", r.FormValue("password"))
	var username = r.FormValue("username")
	var password = r.FormValue("password")
	if username == "" {
		jr.Code = 3
		jr.Msg = "用户名不能为空"
		response.OuputJson(w, jr)
		return
	}
	if password == "" {
		jr.Code = 3
		jr.Msg = "密码不能为空"
		response.OuputJson(w, jr)
		return
	}
	// 校验密码
	checkPass(username, password)
	fmt.Println(r.Method)
	w.Write([]byte("API V1, login"))
}

func checkPass(username, password string) {
	db := mysql.Connet()
	var user *mysql.UserEntity
	user = mysql.GetUser(db, username)
	fmt.Println(user)
	// 关闭链接
	mysql.Close(db)
}
