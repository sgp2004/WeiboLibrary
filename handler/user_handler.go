package handler

import (
	"../dao"
	"../service"
	"../wb_util"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	print("start login out----")
	cookie, _ := r.Cookie("user")
	user := cookie.Value
	if user != "" {
		cookie := http.Cookie{Name: "user", Value: "", Path: "/"}
		http.SetCookie(w, &cookie)
	}
	msg := "logout success!!!"
	tmpl, _ := template.ParseFiles("./static/error.html")
	tmpl.Execute(w, msg)
	return
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	print("start login----")
	username := r.FormValue("user")
	pwd := r.FormValue("pwd")
	pwdMd5 := wb_util.CreateMD5String(pwd)
	print("------md5 pwd: ---")
	print(pwd)
	if !service.Auth(username, pwdMd5, pwd) {
		io.WriteString(w, "valid user error")
		return
	}

	//add user if not exists
	loginUser := dao.QueryUserByUserName(username)
	if loginUser == nil {
		fmt.Println("create user:" + username)
		user := new(dao.User)
		user.UserName = username
		user.Name = username
		dao.CreateUser(user)
		loginUser = dao.QueryUserByUserName(username)
	}
	fmt.Println("login success,user:" + username)

	//set cookie 

	cookie := http.Cookie{Name: "user", Value: username, Path: "/"}
	http.SetCookie(w, &cookie)
	fmt.Println("add user to cookie :" + username)
	tmpl, _ := template.ParseFiles("./static/user.html")
	tmpl.Execute(w, loginUser)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("./static/adduser.html")

		tmpl.Execute(w, "")
		return
	}
	username := r.FormValue("user")
	pwd := r.FormValue("pwd")
	pwdConfirm := r.FormValue("pwd_confirm")

	user := dao.QueryUserByUserName(username)
	if user != nil {
		io.WriteString(w, "user "+username+"alread exists")
		return
	}
	if pwd == "" || pwd != pwdConfirm {
		io.WriteString(w, "password confirm error, not the same one!")
		return
	}

	user = new(dao.User)
	user.UserName = username
	user.Name = username
	user.Pwd = wb_util.CreateMD5String(pwd)
	dao.CreateUser(user)
	dao.RegistUser(user)
	msg := "register success!!!"
	tmpl, _ := template.ParseFiles("./static/success.html")
	tmpl.Execute(w, msg)
}

func UsersShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show shops")
	users := dao.QueryUsers(0, 20)
	fmt.Println(users[0])
	locals := make(map[string]interface{})
	locals["users"] = users
	tmpl, _ := template.ParseFiles("./static/book_shops.html")
	tmpl.Execute(w, locals)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./static/user.html")
	tmpl.Execute(w, "")
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	user := cookie.Value
	if user == "" {
		msg := "error,please login first"
		tmpl, _ := template.ParseFiles("./static/error.html")
		tmpl.Execute(w, msg)
		return
	}
	borrowInfos, lendInfos := dao.BooksBorrowInfoByUserName(user)
	userInfo := dao.QueryUserByUserName(user)
	locals := make(map[string]interface{})
	locals["borrowList"] = borrowInfos
	locals["lendList"] = lendInfos
	locals["user"] = userInfo
	tmpl, _ := template.ParseFiles("./static/profile.html")
	tmpl.Execute(w, locals)
}
