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

func MuxSetup() *mux.Router {
	mux := mux.NewRouter()
	mux.Use(middleware.CheckError)
	return mux
}

func GuestSetup(mux *mux.Router) *mux.Router {
	guest := mux.NewRoute().Subrouter()
	guest.Use(middleware.Guest)
	return guest
}

func AuthSetup(mux *mux.Router) *mux.Router {
	auth := mux.NewRoute().Subrouter()
	auth.Use(middleware.Auth)
	return auth
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

func AdminControllerSetup() controller.AdminController {
	adminRepo := repository.NewUserRepository()
	adminService := service.NewAdminService(db,validate,adminRepo)
	adminController := controller.NewAdminController(adminService)
	return adminController
}