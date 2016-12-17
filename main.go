package main

import (
	"Go-Redis-Admin/api/v1"
	"Go-Redis-Admin/common/exception"
	"Go-Redis-Admin/controller/view"
	"github.com/alexedwards/scs/engine/memstore"
	"github.com/alexedwards/scs/session"
	"io"
	// "fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"text/template"
)

func init() {

}

func main() {
	http.Handle("/", http.HandlerFunc(mainRouter))
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	// Initialise a new storage engine. Here we use the memstore package, but the approach
	// is the same no matter which back-end store you choose.
	engine := memstore.New(0)

	// Initialise the session manager middleware, passing in the storage engine as
	// the first parameter. This middleware will automatically handle loading and
	// saving of session data for you.
	sessionManager := session.Manage(engine)

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)
	http.ListenAndServe(":4000", sessionManager(mux))
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
	handle := &view.Handlers{}
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
	http.Handle("/"+patterns[0]+"/", http.StripPrefix("/"+patterns[0]+"/", http.FileServer(http.Dir("static"))))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Use the PutString helper to store a new key and associated string value in
	// the session data. Helpers are also available for many other data types.
	err := session.PutString(r, "message", "Hello from a session!")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retreive the string value associated with a key.
	// The zero value is returned if the key does not exist.
	msg, err := session.GetString(r, "message")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, msg)
}
