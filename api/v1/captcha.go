package v1

import (
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

func (h *Handlers) NewCaptchaAction(w http.ResponseWriter, r *http.Request) {
	log.Println("API V1, captcha Index")
	var captchaId = captcha.New()
	captcha.WriteImage(w, captchaId, StdWidth, StdHeight)
}
