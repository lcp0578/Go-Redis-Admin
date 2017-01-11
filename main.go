package main

import (
	"Go-Redis-Admin/src/api/v1"
	"Go-Redis-Admin/src/common/exception"
	"Go-Redis-Admin/src/controller"
	// "io"
	// "fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"
)

func init() {

}

func main() {
	http.Handle("/", http.HandlerFunc(mainRouter))
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

	if strings.HasSuffix(r.URL.Path, ".ico") {
		log.Println(r.URL.Path)
		file := strings.Trim(r.URL.Path, "/")
		f, err := os.Open(file)
		defer f.Close()

		if err != nil && os.IsNotExist(err) {
			file = "favicon.ico"
		}
		http.ServeFile(w, r, file)
		return
	}
	pathinfo := strings.ToLower(strings.Trim(r.URL.Path, "/"))
	log.Println("main pathinfo", pathinfo)

	// if /
	if strings.Contains(pathinfo, "/") {
		patterns := strings.Split(pathinfo, "/")
		log.Println(patterns)
		log.Println("len:", len(patterns))
		if len(patterns) < 3 {
			patterns = append(patterns, "index")
		}
		//fmt.Println(reflect.TypeOf(patterns))
		log.Println("patterns:", patterns)
		switch patterns[0] {
		case "api":
			apiRouter(w, r, patterns)
		case "res":
			resRouter(w, r, patterns)
		default:
			tplRouter(w, r, patterns)
		}
	} else {
		// default tpl router
		patterns := make([]string, 0, 3)
		if pathinfo == "" {
			pathinfo = "index"
		}

		patterns = append(patterns, pathinfo)
		tplRouter(w, r, patterns)
	}
}

// API router
func apiRouter(w http.ResponseWriter, r *http.Request, patterns []string) {
	version := patterns[1]
	var handle interface{}
	if version == "v1" {
		handle = &v1.Handlers{}
	} else {
		// version error
		handle = &exception.Handlers{}
		patterns[2] = "exception"
	}
	controller := reflect.ValueOf(handle)
	log.Println(controller)
	action := strings.Title(patterns[2]) + "Action"
	log.Println("action:", action)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		log.Println("error action:", action)
		method = controller.MethodByName("IndexAction")
	}
	wr := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	log.Println(method)
	method.Call([]reflect.Value{wr, rr})
}

// template router
func tplRouter(w http.ResponseWriter, r *http.Request, patterns []string) {
	handle := &controller.Handlers{}
	controller := reflect.ValueOf(handle)
	log.Println(controller)
	action := strings.Title(patterns[0]) + "Action"
	log.Println("action:", action)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		log.Println("error action:", action)
		method = controller.MethodByName("NotFoundAction")
	}
	wr := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	log.Println(method)
	method.Call([]reflect.Value{wr, rr})
}

// resource router
func resRouter(w http.ResponseWriter, r *http.Request, patterns []string) {
	log.Println("/" + patterns[0] + "/")
	log.Println("Dir:", http.Dir("web/static"))
	http.Handle("/"+patterns[0]+"/", http.StripPrefix("/"+patterns[0]+"/", http.FileServer(http.Dir("web/static"))))
}
