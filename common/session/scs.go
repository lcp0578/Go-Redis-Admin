package session

import (
	"github.com/alexedwards/scs/engine/memstore"
	"github.com/alexedwards/scs/session"
	"io"
	"log"
	"net/http"
)

func init() {
	// Initialise a new storage engine. Here we use the memstore package, but the approach
	// is the same no matter which back-end store you choose.
	engine := memstore.New(0)

	// Initialise the session manager middleware, passing in the storage engine as
	// the first parameter. This middleware will automatically handle loading and
	// saving of session data for you.
	// sessionManager := session.Manage(engine)
	session.Manage(engine)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/put", putHandler)
	// mux.HandleFunc("/get", getHandler)
	// http.ListenAndServe(":4000", sessionManager(mux))
}

func Put(r *http.Request, key, val string) bool {
	err := session.PutString(r, key, val)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func Get(r *http.Request, key string) (val string) {
	val, err := session.GetString(r, "message")
	if err != nil {
		val = ""
	}
	return
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
