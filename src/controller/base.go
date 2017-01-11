package controller

import (
	"Go-Redis-Admin/src/common/cookie"
	"log"
	"net/http"
)

func checkLogin(w http.ResponseWriter, r *http.Request) {
	var userId, _ = cookie.GetUserInfo(r)
	log.Println("checkLogin:", userId)
	if userId < 1 {
		// redirect login
		http.Redirect(w, r, "login", http.StatusMovedPermanently)
	}
}
