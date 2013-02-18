package dao

import "fmt"
import "testing"

func TestCreateBook(t *testing.T) {
	book := new(Book)
	book.UserName = "daodao"
	book.Name = "Hadoop"
	fmt.Println(book)
	CreateBook(book)
}
func TestQueryBooksByName(t *testing.T) {
	books := QueryBooksByName("Hadoop")
	fmt.Println("books = \n")
	fmt.Println(books)
}

func TestUpdate(t *testing.T) {
	books := QueryBooksByName("Hadoop")
	fmt.Println("-----books[0]-----")
	fmt.Println(books)
	books[0].Author = "Chunk"
	UpdateBook(books[0])
	UpdateBookStatus(books[0].Id, 1)
	books = QueryBooksByName("Hadoop")
	fmt.Println("book after update= \n")
	fmt.Println(books)
}

func TestDeleteBookById(t *testing.T) {
	books := QueryBooksByName("Hadoop")
	DeleteBookById(books[0].Id)
	fmt.Println("after delete :\n")
	books = QueryBooksByName("Hadoop")
	fmt.Println(books)
}
