package controller

import "net/http"

type MaintenanceController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Form(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	FormUpdate(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FormUpdateDetail(w http.ResponseWriter,r *http.Request)
	UpdateDetail(w http.ResponseWriter, r *http.Request)
	Export(w http.ResponseWriter, r *http.Request)
}
