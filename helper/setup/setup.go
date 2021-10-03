package setup

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"pbl-orkom/controller"
	"pbl-orkom/helper"
	"pbl-orkom/middleware"
	"pbl-orkom/repository"
	"pbl-orkom/service"
)

var db = helper.Connection()
var validate = validator.New()

func MuxSetup() (*mux.Router,*mux.Router,*mux.Router) {
	mux := mux.NewRouter()
	mux.Use(middleware.CheckError)
	guest := mux.NewRoute().Subrouter()
	auth := mux.NewRoute().Subrouter()
	guest.Use(middleware.Guest)
	auth.Use(middleware.Auth)
	return mux,guest,auth
}
func MaintenanceControllerSetup() controller.MaintenanceController{
	spekImpl := repository.NewRepoImpl()
	detailImpl := repository.NewDetailRepositoryImpl()
	serviceImpl := service.NewMaintenanceServiceImpl(db, validate, spekImpl, detailImpl)
	maintenanceController := controller.NewMaintenanceController(serviceImpl)
	return maintenanceController
}

func LoginControllerSetup() controller.LoginController {
	loginRepo := repository.NewUserRepository()
	loginService := service.NewLoginService(db, validate,loginRepo)
	loginController := controller.NewLoginController(loginService)
	return loginController
}

func TroubleControllerSetup() controller.TroubleshootingController {
	troubleRepo := repository.NewTroubleshootingRepo()
	componentRepo := repository.NewComponentRepoImpl()
	troubleService := service.NewTroubleshootingImpl(db, validate,troubleRepo,componentRepo)
	troubleController := controller.NewTroubleshootingImpl(troubleService)
	return troubleController
}