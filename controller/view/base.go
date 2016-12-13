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
	t, _ := template.ParseFiles("views/content.html")
	t.Execute(w, nil)
}

func (h *Handlers) SystemAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view System")
	t, _ := template.ParseFiles("views/system.html")
	t.Execute(w, nil)
}

func (h *Handlers) NotFoundAction(w http.ResponseWriter, r *http.Request) {
	log.Println("view System")
	t, _ := template.ParseFiles("views/notfound.html")
	t.Execute(w, nil)
}
