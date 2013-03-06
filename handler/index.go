package handler

import (
	"../dao"
	"net/http"
	//"reflect"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./static/index.html")
	infos := dao.TimeLine("1=1",0, 10)
	locals := make(map[string]interface{})
	locals["infos"] = infos
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		locals["user"] = cookie.Value
		//locals["user"] = reflect.ValueOf(cookie.Value)
	}
	tmpl.Execute(w, locals)
}
