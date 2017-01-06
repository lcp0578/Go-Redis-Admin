package mysql

import (
	"database/sql"
	// "fmt"
)

type UserEntity struct {
	Id       int32  `用户自增ID`
	Username string `用户名`
	Password string `用户密码，32位`
	Salt     string `用户的随机盐值，8位`
	Add_time string `创建时间`
	Status   int8   `用户状态，0禁用;1正常`
	Last_ip  string `最后一次登录IP`
}

func GetUserPass(username string) (*UserEntity, error) {
	db := Connet()
	var ue UserEntity
	err := db.QueryRow("SELECT id,password,salt FROM gra_user WHERE username=?", username).Scan(&ue.Id, &ue.Password, &ue.Salt)
	Close(db)
	if err != nil {
		return &UserEntity{}, err
	} else {
		return &ue, nil
	}
}

func SetLastLoginIp(id int32, ip string) (bool, error) {
	db := Connet()
	// prepare SQL
	var stmt *sql.Stmt
	stmt, err := db.Prepare("UPDATE gra_user SET `last_ip` = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(ip, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
