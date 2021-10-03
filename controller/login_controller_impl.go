package controller

import (
	"fmt"
	"net/http"
	"pbl-orkom/helper"
	"pbl-orkom/model/web"
	"pbl-orkom/service"
)

type loginControllerImpl struct{
	service service.LoginService
}

func NewLoginController(service service.LoginService) LoginController {
	return loginControllerImpl{service:service}
}

func (l loginControllerImpl) LoginPage(w http.ResponseWriter, r *http.Request) {
	session,err :=  helper.Store.Get(r,"admin")
	helper.PanicIfError(err)
	fmt.Println(session)
	helper.ViewParser(w,"login-page",nil)
}

func (l loginControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	data := web.LoginRequest{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	user := l.service.Login(r.Context(),data)
	helper.SetSession(w,r,user)
	http.Redirect(w,r,"/maintenance", http.StatusSeeOther)
}
