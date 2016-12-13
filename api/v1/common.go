package v1

import (
	"log"
	"net/http"
)

type Handlers struct{}

func (h *Handlers) IndexAction(w http.ResponseWriter, r *http.Request) {
	log.Println("API V1, Index")
	r.ParseForm()
	w.Write([]byte("API V1, Index"))
}
