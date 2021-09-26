package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type MaintenanceRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, idMaintenance int) (domain.MaintenanceRequest,error)
	Save(ctx context.Context, tx *sql.Tx, spek web.MaintenanceRequest) int64
	Delete(ctx context.Context, tx *sql.Tx, idMaintenance int) bool
	Update(ctx context.Context, tx *sql.Tx, idMaintenance int, spek web.MaintenanceRequest) bool
}
