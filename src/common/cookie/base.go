package cookie

import (
	"net/http"
)

/**
 * set cookie
 **/
func Set(w http.ResponseWriter, key, value, path string, maxAge int) {
	cookie := http.Cookie{Name: key, Value: value, Path: path, MaxAge: maxAge}
	http.SetCookie(w, &cookie)
}

/**
 * get cookie
 **/
func Get(req *http.Request, key string) string {
	cookie, err := req.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}

/**
 * del cookie
 **/
func Del(w http.ResponseWriter, key, path string) {
	cookie := http.Cookie{Name: key, Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
}
