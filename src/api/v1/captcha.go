package v1

import (
	"Go-Redis-Admin/src/common/cookie"
	"Go-Redis-Admin/src/common/response"
	"errors"
	"github.com/dchest/captcha"
	"log"
	"net/http"
	"time"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 6
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)
const (
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)

var (
	ErrNotFound = errors.New("captcha: id not found")
)

func (h *Handlers) NewcaptchaAction(w http.ResponseWriter, r *http.Request) {
	log.Println("API V1, captcha new")
	var captchaId = captcha.New()
	cookie.Set(w, "captchaId", captchaId, "/", 600)
	log.Println("API V1, captcha captchaId", captchaId)
	captcha.WriteImage(w, captchaId, StdWidth, StdHeight)
}

func (h *Handlers) ReloadcaptchaAction(w http.ResponseWriter, r *http.Request) {
	log.Println("API V1, captcha reload")
	captchaId := cookie.Get(r, "captchaId")
	log.Println("API V1, captcha reload captchaId", captchaId)
	if captcha.Reload(captchaId) {
		captcha.WriteImage(w, captchaId, StdWidth, StdHeight)
	} else {
		h.NewcaptchaAction(w, r)
		log.Println("API V1, captcha reload faild")
	}
}

func (h *Handlers) VerifycaptchaAction(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{}
	jr := &response.JsonResponse{
		0,
		"faild",
		data,
	}
	query := r.URL.Query()
	log.Println(query)
	val, ok := query["val"]
	if !ok {
		jr.Code = 2
		jr.Msg = "val empty"
		response.OuputJson(w, jr)
		return
	}
	captchaVal := val[0]
	log.Println("API V1, captcha verify")
	var captchaId = cookie.Get(r, "captchaId")
	var result = captcha.Verify(captchaId, []byte(captchaVal))

	if result {
		jr.Code = 1
		jr.Msg = "success"
	}
	response.OuputJson(w, jr)
}

func verifyCaptcha(r *http.Request, captchaVal string) bool {
	var captchaId = cookie.Get(r, "captchaId")
	if captchaId == "" {
		return false
	}
	log.Println("captchaId", captchaId)
	log.Println("captchaVal", captchaVal)
	return captcha.VerifyString(captchaId, captchaVal)
}
