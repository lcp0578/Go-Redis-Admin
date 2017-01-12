package mysql

import (
	"database/sql"
	// "fmt"
	"Go-Redis-Admin/src/common/crypto"
	"Go-Redis-Admin/src/common/lib"
	"time"
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

func CreateUser(username, password, ip string) (int32, error) {
	db := Connet()
	var stmt *sql.Stmt
	stmt, err := db.Prepare("INSERT INTO gra_user(username, password, salt, add_time, status, last_ip) VALUES (?,?,?,?,?,?)")
	if err != nil {
		//log.Println(err)
		return 0, err
	}
	var salt = lib.StrRand(8, 3)
	var addTime = time.Now().Format("2006-01-02")
	password = crypto.Md5Double(password, salt)
	result, err := stmt.Exec(username, password, salt, addTime, 1, ip)
	if err != nil {
		//log.Println(err)
		return 0, err
	}
	lastId, _ := result.LastInsertId()
	return int32(lastId), nil
}
