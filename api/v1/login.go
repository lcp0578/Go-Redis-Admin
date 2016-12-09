package v1

import (
	"log"
	"net/http"
)

type Handlers struct{}

func (h *Handlers) LoginAction(w http.ResponseWriter, r *http.Request) {
	log.Println("API Login Action")
	r.ParseForm()
	w.Write([]byte("API V1, login"))
}
