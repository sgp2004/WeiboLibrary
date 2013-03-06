package dao

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Id              int
	UserName        string
	Pwd             string
	Name            string
	OwnBookCount    int
	BorrowBookCount int
	AllowBookCount  int
	AllowBookDays   int
	UserLevel       int
}

func CreateUser(user *User) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "insert into `user`(`username`,`name`) values(?,?)"
	fmt.Println(sql)
	db.Prepare(sql)
	exec(db, sql, user.UserName, user.Name)
}

func RegistUser(user *User) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "insert into `user_pwd`(`username`,`pwd`) values(?,?) ON DUPLICATE KEY UPDATE `pwd`=VALUES(`pwd`)"
	fmt.Println(sql)
	db.Prepare(sql)
	exec(db, sql, user.UserName, user.Pwd)
}

func GetRegisterUser(username, pwd string) *User {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select username,pwd from `user_pwd` where username=? and pwd=?"
	print(sql)
	print(username)
	print(pwd)
	db.Prepare(sql)
	rows, _ := db.Query(sql, username, pwd)
	if rows == nil {
		return nil
	}
	for rows.Next() {
		user := new(User)
		row_err := rows.Scan(&user.UserName, &user.Pwd)
		if row_err != nil {
			print("Row error!!")
			return nil
		}
		return user
	}
	return nil
}

func QueryUserByUserName(username string) *User {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,username,name,own_book_count,borrow_book_count,allow_book_count,allow_borrow_days,user_level from `user` where username=?"
	db.Prepare(sql)
	rows, _ := db.Query(sql, username)
	for rows.Next() {
		user := new(User)
		row_err := rows.Scan(&user.Id, &user.UserName, &user.Name, &user.OwnBookCount, &user.BorrowBookCount, &user.AllowBookCount, &user.AllowBookDays, &user.UserLevel)
		if row_err != nil {
			print("Row error!!")
			return nil
		}
		return user
	}
	return nil
}

func QueryUsers(start, limit int) []*User {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,username,name,own_book_count,borrow_book_count,allow_book_count,allow_borrow_days,user_level from `user` where own_book_count>0 order by allow_borrow_days desc limit ?,? "
	db.Prepare(sql)
	rows, _ := db.Query(sql, start, limit)
	users := make([]*User, 0, limit)

	for rows.Next() {
		user := new(User)
		row_err := rows.Scan(&user.Id, &user.UserName, &user.Name, &user.OwnBookCount, &user.BorrowBookCount, &user.AllowBookCount, &user.AllowBookDays, &user.UserLevel)
		if row_err != nil {
			print("Row error!!")
			return users
		}
		users = append(users, user)
	}
	return users
}

func QueryUserMapsByIds(ids []string) map[string]*User {

	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,username,name,own_book_count,borrow_book_count,allow_book_count,allow_borrow_days,user_level from `user` where id in (" + strings.Join(ids, ",") + ")"
	fmt.Println(sql)
	rows, _ := db.Query(sql)
	users := make(map[string]*User)
	for rows.Next() {
		user := new(User)
		row_err := rows.Scan(&user.Id, &user.UserName, &user.Name, &user.OwnBookCount, &user.BorrowBookCount, &user.AllowBookCount, &user.AllowBookDays, &user.UserLevel)
		if row_err != nil {
			print("Row error!!")
			return users
		}
		users[strconv.Itoa(user.Id)] = user
	}
	return users
}

func AddAllowBorrowDays(count int, uid int) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "update `user` set allow_borrow_days =allow_borrow_days+? where id =?"
	db.Prepare(sql)
	exec(db, sql, count, uid)
}

func IncreaseUserBookInfoCount(count int, uid int, updateType int) {
	dataField := "own_book_count"
	if updateType == 1 {
		dataField = "borrow_book_count"
	} else if updateType == 2 {
		dataField = "allow_book_count"
	}
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "update `user` set " + dataField + "= " + dataField + "+?  where id =?"
	print(sql)
	db.Prepare(sql)
	exec(db, sql, count, uid)
}
