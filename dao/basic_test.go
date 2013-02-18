package dao

import "testing"

func TestInsert(t *testing.T) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "insert into `user`(`username`,`name`) values('5','')"
	println(sql)
	exec(db, sql)
}
