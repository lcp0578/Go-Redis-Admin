package main

import (
	"Go-Redis-Admin/api/v1"
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func init() {

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	http.Handle("/api/", http.HandlerFunc(mainRouter))
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//	r.ParseForm()                //解析参数，默认是不会解析的
	//	fmt.Println("form:", r.Form) //这些信息是输出到服务器端的打印信息
	//	fmt.Println("path", r.URL.Path)
	//	fmt.Println("scheme", r.URL.Scheme)
	//	fmt.Println("r.Form[\"test\"]", r.Form["test"])
	//	for k, v := range r.Form {
	//		fmt.Println("key:", k)
	//		fmt.Println("val:", strings.Join(v, ""))
	//	}
	log.Println("/")
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
	//w.Write([]byte("Hello, I'm coming..."))
	//fmt.Fprintf(w, "Hello, I'm coming...") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("/login")
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(w, nil)
	//w.Write([]byte("login"))
}

// main router
func mainRouter(w http.ResponseWriter, r *http.Request) {
	pathinfo := strings.Trim(r.URL.Path, "/")
	log.Println(pathinfo)
	// if /
	if strings.Contains(pathinfo, "/") {
		patterns := strings.Split(pathinfo, "/")
		//fmt.Println(reflect.TypeOf(patterns))
		log.Println("patterns:", patterns)
		switch patterns[0] {
		case "api":
			apiRouter(w, r, patterns)
		case "tpl":
			tplRouter(w, r, patterns)
		case "res":
			resRouter(w, r, patterns)
		default:
		}
	} else {
		// default router
	}
}

// API router
func apiRouter(w http.ResponseWriter, r *http.Request, patterns []string) {
	handle := &v1.Handlers{}
	controller := reflect.ValueOf(handle)
	version := patterns[1]
	action := version + "." + strings.ToUpper(patterns[1]) + "Action"
	log.Println("action:", action)
	method := controller.MethodByName(action)
	wr := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	method.Call([]reflect.Value{wr, rr})
}

// template router
func tplRouter(w http.ResponseWriter, r *http.Request, patterns []string) {

}

// resource router
func resRouter(w http.ResponseWriter, r *http.Request, patterns []string) {

}
