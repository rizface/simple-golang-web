package app

const (
	DATABASE_DRIVER      = "mysql"
	APP_HOST             = "localhost:3000"
	GET_MAINTENANCE      = "/maintenance"
	GET_MAINTENANCE_FORM = "/maintenance-form"
	SAVE_MAINTENANCE     = "/maintenance/save"
	UPDATE_MAINTENANCE   = "/maintenance/update/{idMaintenance}"
	DELETE_MAINTENANCE   = "/maintenance/delete/{idMaintenance}"
	UPDATE_DETAIL_MAINTENANCE = "/maintenance/detail/update/{idMaintenance}"
	EXPORT_MAINTENANCE = "/maintenance/export"

	//Troubleshooting
	GET_TROUBLESHOOTING = "/troubleshooting"
	SAVE_TROUBLESHOOTING_FORM  = "/troubleshooting/save"
	SAVE_TROUBLESHOOTING = "/troubleshooting/save"
	DELETE_TROUBLESHOOTING = "/troubleshooting/delete/{idTrouble}"
	UPDATE_TROUBLESHOOTING = "/troubleshooting/update/{idTrouble}"
	EXPORT_TROUBLESHOOTING = "/troubleshooting/export"

	// SMS Service
	ACCOUD_SID = "ACe52d1877980d3a04b8a5f496470e2396"
	AUTH_TOKEN = "9f2c9d0f10e18f1d54a01c2dc89bfba5"
)
