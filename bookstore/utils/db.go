package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db   *sql.DB
	err  error
	Rows *sql.Rows
)

func init() {
	Db, err = sql.Open("mysql", "root:root@/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
