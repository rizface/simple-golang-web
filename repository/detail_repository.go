package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type DetailRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.DetailJoinMaintenance
	FindByIdMaintenance(ctx context.Context, tx *sql.Tx, idMaintenance int) (domain.DetailMaintenance,error)
	Save(ctx context.Context, tx *sql.Tx, detail web.DetailMaintenance) bool
	Update(ctx context.Context,tx *sql.Tx, idMaintenance int, detail web.DetailMaintenance) bool
	Export(ctx context.Context, tx *sql.Tx)[]domain.DetailJoinMaintenance
}
