package controller

import "net/http"

type TroubleshootingController interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	FormSave(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	UpdateForm(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http. ResponseWriter, r *http.Request)
}