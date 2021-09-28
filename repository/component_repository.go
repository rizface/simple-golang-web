package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
)

type ComponentRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Component
	Save(ctx context.Context, tx *sql.Tx, idTrouble int, componentId []int) bool
}
