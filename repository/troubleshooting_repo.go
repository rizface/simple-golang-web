package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type TroubleshootingRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Troubleshooting
	GetById(ctx context.Context, idTrouble int, tx *sql.Tx) domain.Troubleshooting
	Save(ctx context.Context,request web.TroubleshootRequest, tx *sql.Tx) int64
	Update(ctx context.Context, idTrouble int, request web.TroubleshootRequest, tx *sql.Tx, wg interface{}) bool
	Delete(ctx context.Context, idTrouble int, tx *sql.Tx) bool
}
