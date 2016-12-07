package v1

import (
	"net/http"
)

func LoginAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}
