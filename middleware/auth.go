package middleware

import (
	"fmt"
	"net/http"
	"pbl-orkom/helper"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session,err := helper.Store.Get(request,"admin")
		helper.PanicIfError(err)
		login := session.Values["login"]
		if login != nil && login.(bool) == true {
			fmt.Println("bisa")
			next.ServeHTTP(writer,request)
			return
		}
		http.Redirect(writer,request,"/",http.StatusSeeOther)
	})
}