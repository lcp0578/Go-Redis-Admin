package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

func Md5(str, salt string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str + salt))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func Md5Double(str, salt string) string {
	return Md5(Md5(str, salt), salt)
}

func DesEncode() {
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

/**
 * plainText 明文
 * keyText, 必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
 *
 **/
func AesEncode(plainText, keyText string) (string, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		return "", err
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, []byte(plainText))
	return hex.EncodeToString(cipherText), nil
}

func AesDecode(cipherText, keyText string) (string, error) {
	cipherByte, err := hex.DecodeString(cipherText)
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		return "", err
	}
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plainText := make([]byte, len(cipherText))
	cfbdec.XORKeyStream(plainText, cipherByte)
	return string(plainText), nil
}
