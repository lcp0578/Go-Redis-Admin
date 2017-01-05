package mysql

import (
	"database/sql"
	"fmt"
)

type UserEntity struct {
	id       int32  `用户自增ID`
	username string `用户名`
	password string `用户密码，32位`
	salt     string `用户的随机盐值，8位`
	add_time string `创建时间`
	status   int8   `用户状态，0禁用;1正常`
	last_ip  string `最后一次登录IP`
}

func GetUser(db *sql.DB, username string) (u *UserEntity) {
	// GET ONE ROW
	var row *sql.Row
	row = db.QueryRow("SELECT * FROM gra_user WHERE username=?", username)
	var ue *UserEntity
	// GET VLAUES
	var err = row.Scan(ue.id, ue.username, ue.password, ue.salt, ue.add_time, ue.status, ue.last_ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	return ue
}
