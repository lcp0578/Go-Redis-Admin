package crypto

import (
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
