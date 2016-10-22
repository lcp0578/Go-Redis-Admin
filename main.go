package main

import (
	"html/template"
	"log"
	"net/http"
)

func inti() {

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)
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
