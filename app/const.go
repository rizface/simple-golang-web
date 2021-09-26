package app

const (
	DATABASE_DRIVER      = "mysql"
	APP_HOST             = "localhost:3000"
	GET_MAINTENANCE      = "/maintenance/"
	GET_MAINTENANCE_FORM = "/maintenance-form/"
	SAVE_MAINTENANCE     = "/maintenance/save/"
	UPDATE_MAINTENANCE   = "/maintenance/update"
	DELETE_MAINTENANCE   = "/maintenance/delete/{idMaintenance}/"
	UPDATE_DETAIL_MAINTENANCE = "/maintenance/detail/update"
	EXPORT_MAINTENANCE = "/maintenance/export/"
)
