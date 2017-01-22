package controller

import (
	"html/template"
	"log"
	"net/http"
)

var ViewBase string = "src/views/"
var TplMap = map[string]string{
	"NavActive": "index",
}

type Handlers struct{}

func (h *Handlers) LoginAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view login")
	t, _ := template.ParseFiles(ViewBase + "login.html")
	t.Execute(w, nil)
}

func (h *Handlers) IndexAction(w http.ResponseWriter, r *http.Request) {
	log.Println(TplMap)
	TplMap["NavActive"] = "index"
	loadCommonTpl(w, "index.html", TplMap)
}

func (h *Handlers) ContentAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view content")
	t, err := template.ParseFiles("src/views/content.html")
	if err != nil {
		log.Println(err.Error())
	}
	t.Execute(w, nil)
}

func (h *Handlers) SystemAction(w http.ResponseWriter, r *http.Request) {
	TplMap["NavActive"] = "system"
	loadCommonTpl(w, "system.html", TplMap)
}

func (h *Handlers) TestAction(w http.ResponseWriter, r *http.Request) {
	checkLogin(w, r)
	log.Println("view test")
	t, err := template.ParseFiles("src/views/test.html")
	if err != nil {
		log.Println("parse view error", err)
	}
	t.Execute(w, nil)
}

func (h *Handlers) NotFoundAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view Not Found")
	t, _ := template.ParseFiles("src/views/notfound.html")
	t.Execute(w, nil)
}

func (h *Handlers) RedislistAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view index")
	TplMap["NavActive"] = "redis"
	loadCommonTpl(w, "redis_list.html", TplMap)
}

func (h *Handlers) RedisaddAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view index")
	TplMap["NavActive"] = "redis"
	loadCommonTpl(w, "redis_add.html", TplMap)
}

func (h *Handlers) KeyslistAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view index")
	TplMap["NavActive"] = "keys"
	loadCommonTpl(w, "keys_list.html", TplMap)
}

func (h *Handlers) KeysaddAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view index")
	TplMap["NavActive"] = "keys"
	loadCommonTpl(w, "keys_add.html", TplMap)
}

func (h *Handlers) SearchAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view index")
	TplMap["NavActive"] = "search"
	loadCommonTpl(w, "search.html", TplMap)
}

/**
 * 加载公共头部、导航和底部
 *
 **/
func loadCommonTpl(w http.ResponseWriter, tpl string, data interface{}) {
	s1, _ := template.ParseFiles(ViewBase+"header.html", ViewBase+"nav.html", ViewBase+tpl, ViewBase+"footer.html")
	s1.ExecuteTemplate(w, "content", data)
}
