package main

import (
	"net/http"
	"pbl-orkom/app"
	"pbl-orkom/helper"
	"pbl-orkom/helper/setup"
)

func main() {
	mux,guest,auth := setup.MuxSetup()
	maintenanceController := setup.MaintenanceControllerSetup()
	loginController := setup.LoginControllerSetup()
	troubleController := setup.TroubleControllerSetup()


	guest.HandleFunc("/", loginController.LoginPage).Methods(http.MethodGet)
	guest.HandleFunc("/", loginController.Login).Methods(http.MethodPost)

	//Maintenance
	auth.HandleFunc(app.GET_MAINTENANCE, maintenanceController.GetAll).Methods(http.MethodGet)
	auth.HandleFunc(app.GET_MAINTENANCE_FORM, maintenanceController.Form).Methods(http.MethodGet)
	auth.HandleFunc(app.SAVE_MAINTENANCE, maintenanceController.Save).Methods(http.MethodPost)
	auth.HandleFunc(app.UPDATE_MAINTENANCE, maintenanceController.FormUpdate).Methods(http.MethodGet)
	auth.HandleFunc(app.UPDATE_MAINTENANCE, maintenanceController.Update).Methods(http.MethodPost)
	auth.HandleFunc(app.DELETE_MAINTENANCE, maintenanceController.Delete).Methods(http.MethodGet)
	auth.HandleFunc(app.UPDATE_DETAIL_MAINTENANCE, maintenanceController.FormUpdateDetail).Methods(http.MethodGet)
	auth.HandleFunc(app.UPDATE_DETAIL_MAINTENANCE, maintenanceController.UpdateDetail).Methods(http.MethodPost)
	auth.HandleFunc(app.EXPORT_MAINTENANCE, maintenanceController.Export)


	// troubleshooting
	auth.HandleFunc(app.GET_TROUBLESHOOTING, troubleController.Get).Methods(http.MethodGet)
	auth.HandleFunc(app.SAVE_TROUBLESHOOTING_FORM, troubleController.FormSave).Methods(http.MethodGet)
	auth.HandleFunc(app.SAVE_TROUBLESHOOTING,troubleController.Save).Methods(http.MethodPost)
	auth.HandleFunc(app.DELETE_TROUBLESHOOTING, troubleController.Delete).Methods(http.MethodGet)
	auth.HandleFunc(app.UPDATE_TROUBLESHOOTING,troubleController.UpdateForm).Methods(http.MethodGet)
	auth.HandleFunc(app.UPDATE_TROUBLESHOOTING, troubleController.Update).Methods(http.MethodPost)
	auth.HandleFunc(app.EXPORT_TROUBLESHOOTING, troubleController.Export)

	helper.StaticFile(mux)
	mux.NotFoundHandler = helper.NotFoundHandler
	mux.MethodNotAllowedHandler = helper.MethodNotAllowedHandler
	helper.StartServer(mux)
}
