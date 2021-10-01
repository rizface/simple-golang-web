package main

import (
	"fmt"
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

	// maintenance
	mux.Use(middleware.CheckError)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Ini Akan Jadi Halaman Login")
	})
	mux.HandleFunc(app.GET_MAINTENANCE, controllerImpl.GetAll).Methods(http.MethodGet)
	mux.HandleFunc(app.GET_MAINTENANCE_FORM, controllerImpl.Form).Methods(http.MethodGet)
	mux.HandleFunc(app.SAVE_MAINTENANCE, controllerImpl.Save).Methods(http.MethodPost)
	mux.HandleFunc(app.UPDATE_MAINTENANCE, controllerImpl.FormUpdate).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_MAINTENANCE, controllerImpl.Update).Methods(http.MethodPost)
	mux.HandleFunc(app.DELETE_MAINTENANCE, controllerImpl.Delete).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_DETAIL_MAINTENANCE, controllerImpl.FormUpdateDetail).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_DETAIL_MAINTENANCE, controllerImpl.UpdateDetail).Methods(http.MethodPost)
	mux.HandleFunc(app.EXPORT_MAINTENANCE, controllerImpl.Export)


	// troubleshooting
	troubleRepo := repository.NewTroubleshootingRepo()
	componentRepo := repository.NewComponentRepoImpl()
	troubleService := service.NewTroubleshootingImpl(db,validate,troubleRepo,componentRepo)
	troubleController := controller.NewTroubleshootingImpl(troubleService)

	mux.HandleFunc(app.GET_TROUBLESHOOTING, troubleController.Get).Methods(http.MethodGet)
	mux.HandleFunc(app.SAVE_TROUBLESHOOTING_FORM, troubleController.FormSave).Methods(http.MethodGet)
	mux.HandleFunc(app.SAVE_TROUBLESHOOTING,troubleController.Save).Methods(http.MethodPost)
	mux.HandleFunc(app.DELETE_TROUBLESHOOTING, troubleController.Delete).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_TROUBLESHOOTING,troubleController.UpdateForm).Methods(http.MethodGet)
	mux.HandleFunc(app.UPDATE_TROUBLESHOOTING, troubleController.Update).Methods(http.MethodPost)
	mux.HandleFunc(app.EXPORT_TROUBLESHOOTING, troubleController.Export)

	helper.StaticFile(mux)
	mux.NotFoundHandler = helper.NotFoundHandler
	mux.MethodNotAllowedHandler = helper.MethodNotAllowedHandler
	helper.StartServer(mux)
}
