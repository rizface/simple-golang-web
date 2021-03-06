package controller

import (
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
	helper.ViewParser(w,"login-page",nil)
}

func (l loginControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	data := web.UserRequest{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	user := l.service.Login(r.Context(),data)
	helper.SetSession(w,r,user)
	http.Redirect(w,r,"/maintenance", http.StatusSeeOther)
}

func (l loginControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	session,err := helper.Store.Get(r,"admin")
	helper.PanicIfError(err)
	session.Options.MaxAge = -1
	session.Save(r,w)
	http.Redirect(w,r,"/",http.StatusSeeOther)
}
