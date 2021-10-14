package controller

import "net/http"

type AdminController interface {
	Dashboard(w http.ResponseWriter, r *http.Request)
	SseDashboard(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	SaveForm(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	UpdateForm(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
