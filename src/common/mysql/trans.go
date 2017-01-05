package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func transaction() {
	db, err := sql.Open("mysql", "root:lcp0578@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var trans *sql.Tx
	trans, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = trans.Exec("INSERT INTO users(name, age, sex) VALUES ('lcp0578', 26, 2)")
	if err != nil {
		fmt.Println("Rollback")
		trans.Rollback()
	} else {
		fmt.Println("Commit")
		trans.Commit()
	}

}
