package cookie

import (
	"Go-Redis-Admin/app/config"
	"Go-Redis-Admin/src/common/crypto"
	"log"
	"net/http"
	"strconv"
)

/**
 *   获取用户的信息
 **/
func GetUserInfo(req *http.Request) (int32, string) {
	var auth = Get(req, "gra_auth")
	log.Println("cipherText", auth)
	keyText := config.UserAesKey
	plainTextCopy, err := crypto.AesDecode(auth, keyText)
	log.Println("plainTextCopy", plainTextCopy)
	return 0, auth
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
