package handler

import (
	"../dao"
	"fmt"
	"net/http"
	"text/template"
    "strconv"
)
func BookNoteHandler(w http.ResponseWriter, r *http.Request) {
     cookie, _ := r.Cookie("user")
    fmt.Println(cookie)
    if cookie.Value == "" {
        msg := "error,please login first"
        tmpl, _ := template.ParseFiles("./static/error.html")
        tmpl.Execute(w, msg)
        return
    } 
    infoId := r.FormValue("info_id")
    note := r.FormValue("note")
    if note == "" || infoId == "" {
     msg := "error,please write note first"
        tmpl, _ := template.ParseFiles("./static/error.html")
        tmpl.Execute(w, msg)
        return
    }
infoIdInt64,_ := strconv.ParseInt(infoId,10,0)
infoIdInt := (int)(infoIdInt64)
    dao.SetBorrowInfoNote(infoIdInt,note)
    msg := "write note success!"
    tmpl, _ := template.ParseFiles("./static/success.html")
    tmpl.Execute(w, msg)
    return
}
func BorrowBookHandler(w http.ResponseWriter, r *http.Request) {
   cookie, _ := r.Cookie("user")
    fmt.Println(cookie)
    if cookie.Value == "" {
        msg := "error,please login first"
        tmpl, _ := template.ParseFiles("./static/error.html")
        tmpl.Execute(w, msg)
        return
    }   
    user := cookie.Value
    bookId := r.FormValue("id") 
    infoId := r.FormValue("info_id")
    status := r.FormValue("status")
    statusInt64,_ := strconv.ParseInt(status,10,0)
    bookIdInt64,_ := strconv.ParseInt(bookId,10,0)
    infoIdInt64,_ := strconv.ParseInt(infoId,10,0)
    statusInt := (int)(statusInt64)
    bookIdInt := (int)(bookIdInt64)
    infoIdInt := (int)(infoIdInt64)
    book := dao.QueryBookById(bookIdInt)
    userInfo := dao.QueryUserByUserName(user)
    owner := dao.QueryUserByUserName(book.UserName)
    if book.UserName == user && status=="1"{
         msg := "error,不能借阅自己的书"
        tmpl, _ := template.ParseFiles("./static/error.html")
        tmpl.Execute(w, msg)
        return
    }
    print("user info ------")
    print(userInfo)
    print(userInfo.BorrowBookCount)
/*    if userInfo.BorrowBookCount>=3 {
            msg := "error,最多只能借阅3本书哦！" 
        tmpl, _ := template.ParseFiles("./static/error.html")
        tmpl.Execute(w, msg)
        return 
    }
*/
    //refuse return book
    if status == "3" && book.UserName == user {
        statusInt = 1
    }
    dao.UpdateBookStatus(bookIdInt,statusInt)
    //accept return book
    if status == "0" && book.UserName == user {
        dao.SetBorrowInfoStatus(infoIdInt,1)
    }
    if status == "1"{
        //borrow count add 1
        dao.IncreaseUserBookInfoCount(1,userInfo.Id,1)
        borrowInfo:= new(dao.BorrowInfo)
        borrowInfo.BookId = bookIdInt
        borrowInfo.OwnerId =  owner.Id
        borrowInfo.BorrowerId = userInfo.Id
        dao.CreateInfo(borrowInfo)
    }
    msg := "operation success!"
    tmpl, _ := template.ParseFiles("./static/success.html")
    tmpl.Execute(w, msg)
    return

}

func AcceptBookHandler(w http.ResponseWriter, r *http.Request){
    cookie, _ := r.Cookie("user")
	fmt.Println(cookie)
	if cookie.Value == "" {
		msg := "error,please login first"
		tmpl, _ := template.ParseFiles("./static/error.html")
		tmpl.Execute(w, msg)
		return
	}
    borrowInfoId := r.FormValue("id")
    IdInt64,_ := strconv.ParseInt(borrowInfoId,10,0) 
    IdInt := (int)(IdInt64)
    aggree := r.FormValue("aggree") 
    if aggree == "1" {
        dao.SetBorrowInfoStatus(IdInt, 2)
    }else{
        dao.SetBorrowInfoStatus(IdInt, 1)  //refuse
    }
}


func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	fmt.Println(cookie)
	if cookie.Value == "" {
		msg := "error,please login first"
		tmpl, _ := template.ParseFiles("./static/error.html")
		tmpl.Execute(w, msg)
		return
	}
	user := cookie.Value
	locals := make(map[string]interface{})
	locals["booklist"] = dao.QueryBooksByUserName(user)

	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("./static/addbook.html")

		locals["user"] = user
		tmpl.Execute(w, locals)
		return
	}
	if r.Method == "POST" {
		user := r.FormValue("user")
		fmt.Println(user)
		bookName := r.FormValue("book")
		fmt.Println(bookName)
		if user== "" || bookName == "" {
			fmt.Println("valid error!")
			msg := "error,please input book name and username"
			tmpl, _ := template.ParseFiles("./static/error.html")
			tmpl.Execute(w, msg)
			return
		}

		fmt.Println("valid ok")
		author := r.FormValue("author")
		describe := r.FormValue("describe")
		book := new(dao.Book)
		book.UserName = user
		book.Name = bookName
		book.Author = author
		book.Describe = describe
		fmt.Println(book)
		dao.CreateBook(book)

        userInfo := dao.QueryUserByUserName(user)
        dao.IncreaseUserBookInfoCount(1, userInfo.Id, 0)  //count+1 , 0 means own_book_cout
		locals["book"] = book
		locals["user"] = user
		locals["msg"] = " 图书创建成功"
		tmpl, _ := template.ParseFiles("./static/addbook.html")
		tmpl.Execute(w, locals)

	}
}

func BookShopShowHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if user == "" {
		msg := "error,please input book name and username"
		tmpl, _ := template.ParseFiles("./static/error.html")
		tmpl.Execute(w, msg)
		return

	}
	userInfo := dao.QueryUserByUserName(user)
	locals := make(map[string]interface{})
	locals["booklist"] = dao.QueryBooksByUserName(user)
	locals["user"] = userInfo
	tmpl, _ := template.ParseFiles("./static/book_shop.html")
	tmpl.Execute(w, locals)

}
