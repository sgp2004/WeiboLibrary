package main

import (
	"./handler"
	"fmt"
	"log"
	"net/http"
)

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		fmt.Println(file)
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./static")
	mux.HandleFunc("/index", handler.IndexHandler)
	mux.HandleFunc("/book/create", handler.AddBookHandler)
	mux.HandleFunc("/book/borrow", handler.BorrowBookHandler)
	mux.HandleFunc("/book/note/create", handler.BookNoteHandler)
	mux.HandleFunc("/user/login", handler.LoginHandler)
	mux.HandleFunc("/user/register", handler.RegisterHandler)
	mux.HandleFunc("/user/logout", handler.LogOutHandler)
	mux.HandleFunc("/feed", handler.LoginHandler)
	mux.HandleFunc("/book_shop", handler.BookShopShowHandler)
	mux.HandleFunc("/book_shops", handler.UsersShowHandler)
	mux.HandleFunc("/my_shop", handler.ProfileHandler)
	mux.HandleFunc("/about", handler.LoginHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", mux))
}
