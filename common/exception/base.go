package exception

import (
	"log"
	"net/http"
)

type Handlers struct{}

func (h *Handlers) ExceptionAction(w http.ResponseWriter, r *http.Request) {
	log.Println("Exception")
	w.Write([]byte("Exception"))
}
