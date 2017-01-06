package v1

import (
	"Go-Redis-Admin/src/common/crypto"
	"Go-Redis-Admin/src/common/mysql"
	"Go-Redis-Admin/src/common/response"
	"log"
	"net/http"
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
		jr.Code = 4
		jr.Msg = "密码不能为空"
		response.OuputJson(w, jr)
		return
	}
	// 校验密码
	if !checkPass(username, password) {
		jr.Code = 5
		jr.Msg = "用户名或密码不正确"
		response.OuputJson(w, jr)
		return
	}
	// 更新登录IP
	w.Write([]byte("API V1, login"))
}

func checkPass(username, password string) bool {
	var user *mysql.UserEntity
	user, err := mysql.GetUserPass(username)
	if err != nil {
		log.Println(err)
		return false
	}
	if user.Password != crypto.Md5Double(password, user.Salt) {
		log.Println("password error")
		return false
	}
	return true
}
