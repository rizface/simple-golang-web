package controller

import "net/http"

type LoginController interface{
	LoginPage(w http.ResponseWriter,r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}
