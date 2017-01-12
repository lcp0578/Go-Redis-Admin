package cookie

import (
	"Go-Redis-Admin/app/config"
	"Go-Redis-Admin/src/common/crypto"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
 *   获取用户的信息
 **/
func GetUserInfo(req *http.Request) (int32, string) {
	var auth = Get(req, "grd_auth")
	if auth == "" {
		return 0, ""
	}
	log.Println("cipherText", auth)
	keyText := config.UserAesKey
	plainTextCopy, err := crypto.AesDecode(auth, keyText)
	log.Println("plainTextCopy", plainTextCopy)
	plainTextSlice := strings.Split(plainTextCopy, "|")
	userId, err := strconv.ParseInt(plainTextSlice[0], 10, 32)
	if err != nil {
		return 0, ""
	}
	return int32(userId), plainTextSlice[1]
}

/**
 *
 *  设置用户登录信息
 *
 **/
func SetUserInfo(w http.ResponseWriter, userId int32, username string) (bool, error) {
	// 设置用户的cookie
	plainText := strconv.Itoa(int(userId)) + "|" + username
	keyText := config.UserAesKey
	cipherText, err := crypto.AesEncode(plainText, keyText)
	if err != nil {
		return false, err
	}
	Set(w, "gra_username", username, "/", 8600)
	Set(w, "gra_auth", cipherText, "/", 8600)
	return true, nil
}

func DelUserInfo(w http.ResponseWriter) {
	Del(w, "gra_username", "/")
	Del(w, "gra_auth", "/")
}
