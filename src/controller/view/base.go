package view

import (
	"html/template"
	"log"
	"net/http"
)

type Handlers struct{}

func (h *Handlers) LoginAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view login")
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(w, nil)
}

func (h *Handlers) IndexAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view index")
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}

func (h *Handlers) ContentAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view content")
	t, err := template.ParseFiles("views/content.html")
	if err != nil {
		log.Println(err.Error())
	}
	t.Execute(w, nil)
}

func (h *Handlers) SystemAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view System")
	t, err := template.ParseFiles("views/system.html")
	if err != nil {
		log.Println("parse view error", err)
	}
	t.Execute(w, nil)
}

func (h *Handlers) TestAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view test")
	t, err := template.ParseFiles("views/test.html")
	if err != nil {
		log.Println("parse view error", err)
	}
	t.Execute(w, nil)
}

func (h *Handlers) NotFoundAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view Not Found")
	t, _ := template.ParseFiles("views/notfound.html")
	t.Execute(w, nil)
}
