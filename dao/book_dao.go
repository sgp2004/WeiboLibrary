package dao

import (
	"fmt"
	"strconv"
	"strings"
)

type Book struct {
	Id           int
	UserName     string
	Name         string
	Author       string
	Describe     string
	Img          string
	ModifiedTime string
	Count        int
	Status       int  
    /*0 init or returned ; 1 send borrow request; 
        2 accept borrow ; 3 refuse; 4 send return request;5 refuse accept
    */
    SendBorrowReq bool
    AcceptBorrow bool
    SendReturnReq bool
    AcceptReturn bool   
}

func CreateBook(book *Book) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "insert into books(username,name,author,info,img,count) values(?,?,?,?,?,?)"
	fmt.Println("-------book name:" + book.Name)
	fmt.Println(book)
	db.Prepare(sql)
	exec(db, sql, book.UserName, book.Name, book.Author, book.Describe, book.Img, book.Count)
}
func SetStatusForTemplate(book *Book){
    print("set status")
    print(book)
    if book.Status ==0 {
        book.AcceptReturn = true
        book.SendBorrowReq = false
        book.AcceptBorrow = false
        book.SendReturnReq = false
    }
     if book.Status ==1 {
        book.AcceptReturn = false
        book.SendBorrowReq = true
        book.AcceptBorrow = false
        book.SendReturnReq = false
    }
    if book.Status ==2 {
        book.AcceptReturn = false
        book.SendBorrowReq = false
        book.AcceptBorrow = false
        book.SendReturnReq = true
    }
    print(book)
}
func QueryBooksByName(name string) []*Book {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,username,name,author,info,img,modified_time,count,status from books where name like ?"
	db.Prepare(sql)
	rows, _ := db.Query(sql, name+"%")
	books := make([]*Book, 0, 100)
	for rows.Next() {
		book := new(Book)
		row_err := rows.Scan(&book.Id, &book.UserName, &book.Name, &book.Author, &book.Describe, &book.Img, &book.ModifiedTime, &book.Count, &book.Status)
		if row_err != nil {
			print("Row error!!")
			return books
		}
        SetStatusForTemplate(book)
		books = append(books, book)
	}
	return books
}

func QueryBooksByUserName(username string) []*Book {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,username,name,author,info,img,modified_time,count,status from books where username = ? order by modified_time desc"
	db.Prepare(sql)
	rows, _ := db.Query(sql, username)
	books := make([]*Book, 0, 100)
	for rows.Next() {
		book := new(Book)
		row_err := rows.Scan(&book.Id, &book.UserName, &book.Name, &book.Author, &book.Describe, &book.Img, &book.ModifiedTime, &book.Count, &book.Status)
		if row_err != nil {
			print("Row error!!")
			return books
		}
        SetStatusForTemplate(book)
		books = append(books, book)
	}
	return books
}

func QueryBookById(id int) *Book {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,username,name,author,info,img,modified_time,count,status from books where id=?"
	db.Prepare(sql)
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		book := new(Book)
		row_err := rows.Scan(&book.Id, &book.UserName, &book.Name, &book.Author, &book.Describe, &book.Img, &book.ModifiedTime, &book.Count, &book.Status)
		if row_err != nil {
			print("Row error!!")
			return nil
		}
        SetStatusForTemplate(book)
		return book
	}
	return nil
}

func QueryBookMapsByIds(ids []string) map[string]*Book{
	db := NewLibraryDB()
	defer closeDB(db)
	books := make(map[string]*Book)
	sql := "select id,username,name,author,info,img,modified_time,count,status from books where id in (" + strings.Join(ids, ",") + ")"
	rows, _ := db.Query(sql)
	for rows.Next() {
		book := new(Book)
		row_err := rows.Scan(&book.Id, &book.UserName, &book.Name, &book.Author, &book.Describe, &book.Img, &book.ModifiedTime, &book.Count, &book.Status)
		if row_err != nil {
			print("Row error!!")
			return nil
		}
        
        SetStatusForTemplate(book)
		books[strconv.Itoa(book.Id)] = book
	}
	return books
}

func UpdateBook(book *Book) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "update books set name=? ,author=? ,info=? ,img=?  where id =?"
	db.Prepare(sql)
	exec(db, sql, book.Name, book.Author, book.Describe, book.Img, book.Id)
}

func UpdateBookStatus(id int, status int) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "update books set status=? where id =?"
	db.Prepare(sql)
    print("in update book status")
    print(id)
    print(status)
	exec(db, sql, status, id)
}

func DeleteBookById(id int) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "delete from books where id=?"
	db.Prepare(sql)
	exec(db, sql, id)
}

