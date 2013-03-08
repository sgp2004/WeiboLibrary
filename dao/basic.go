package dao

import (
	"database/sql"
	"fmt"
)
import _ "github.com/Go-SQL-Driver/MySQL"

var dbtype = "mysql"
//var name = "root:root@tcp(localhost:3306)/weibo_library?charset=utf8"
var name = "test:testpwd@tcp(testdb:3306)/weibo_library?charset=utf8"

func NewDB(dbtype, name string) *sql.DB {
	db, err := sql.Open(dbtype, name)
	if err != nil {
		fmt.Print("Open: %v", err)
	}
	return db
}
func NewLibraryDB() *sql.DB {
	db, err := sql.Open(dbtype, name)
	if err != nil {
		fmt.Print("Open: %v", err)
	}
	return db
}

func exec(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		fmt.Print("Exec of %q: %v", query, err)
	}
}

func query(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Print("Exec of %q: %v", query, err)
	}
	return rows
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Print("error closing sql.DB: %v", err)
	}
}
