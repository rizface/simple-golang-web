package main

import (
	"net/http"
	"pbl-orkom/app"
	"pbl-orkom/controller"
	"pbl-orkom/helper"
	"pbl-orkom/middleware"
	"pbl-orkom/repository"
	"pbl-orkom/service"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	db := helper.Connection()
	validate := validator.New()

	spekImpl := repository.NewRepoImpl()
	detailImpl := repository.NewDetailRepositoryImpl()
	serviceImpl := service.NewMaintenanceServiceImpl(db, validate, spekImpl, detailImpl)
	controllerImpl := controller.NewMaintenanceController(serviceImpl)

	mux.Use(middleware.CheckError)
	mux.HandleFunc(app.GET_MAINTENANCE, controllerImpl.GetAll).Methods(http.MethodGet)
	mux.HandleFunc(app.GET_MAINTENANCE_FORM, controllerImpl.Form).Methods(http.MethodGet)
	mux.HandleFunc(app.SAVE_MAINTENANCE, controllerImpl.Save).Methods(http.MethodPost)
	mux.HandleFunc(app.UPDATE_MAINTENANCE, controllerImpl.FormUpdate).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_MAINTENANCE, controllerImpl.Update).Methods(http.MethodPost)
	mux.HandleFunc(app.DELETE_MAINTENANCE, controllerImpl.Delete).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_DETAIL_MAINTENANCE, controllerImpl.FormUpdateDetail).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_DETAIL_MAINTENANCE,controllerImpl.UpdateDetail).Methods(http.MethodPost)
	mux.HandleFunc(app.EXPORT_MAINTENANCE,controllerImpl.Export)


	helper.StaticFile(mux)
	server := http.Server{
		Addr:    app.APP_HOST,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

