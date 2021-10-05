package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"pbl-orkom/app"
	"pbl-orkom/helper"
	"pbl-orkom/model/web"
	"pbl-orkom/service"
	"strconv"
	"tawesoft.co.uk/go/dialog"
)

type adminControllerImpl struct {
	service service.AdminService
}

func NewAdminController(service service.AdminService) AdminController {
	return adminControllerImpl{service: service}
}

func (a adminControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	data := a.service.Get(r.Context())
	helper.ViewParser(w,"admin", map[string]interface{} {
		"data":data,
	})
}

func (a adminControllerImpl) SaveForm(w http.ResponseWriter, r *http.Request) {
	helper.ViewParser(w,"admin-form",nil)
}

func (a adminControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	request := web.UserRequest{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	success := a.service.Save(r.Context(),request)
	if success {
		dialog.Alert("Data Admin Berhasil Ditambahkan")
	} else {
		dialog.Alert("Data Admin Gagal Ditambahkan")
	}
	http.Redirect(w,r,app.GET_ADMIN,http.StatusSeeOther)
}

func (a adminControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idAdmin,err := strconv.Atoi(params["idAdmin"])
	helper.PanicIfError(err)
	success := a.service.Delete(r.Context(),idAdmin)
	if success {
		dialog.Alert("Data Admin Berhasil Dihapus")
	} else {
		dialog.Alert("Data Admin Gagal Dihapus")
	}
	http.Redirect(w,r,app.GET_ADMIN,http.StatusSeeOther)
}

func (a adminControllerImpl) UpdateForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idAdmin,err := strconv.Atoi(params["idAdmin"])
	helper.PanicIfError(err)
	data := a.service.UpdateForm(r.Context(),idAdmin)
	helper.ViewParser(w,"admin-form-update", map[string]interface{} {
		"data":data,
	})
}

func (a adminControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idAdmin,err := strconv.Atoi(params["idAdmin"])
	helper.PanicIfError(err)
	request := web.UserRequest{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	success := a.service.Update(r.Context(),idAdmin,request)
	if success {
		dialog.Alert("Data Admin Berhasil Diupdate")
	} else {
		dialog.Alert("Data Admin Gagal Diupdate")
	}
	http.Redirect(w,r,app.LOGOUT,http.StatusSeeOther)
}


