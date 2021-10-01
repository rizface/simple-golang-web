package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
	"sync"
)

type ComponentRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Component
	GetByIdTrouble(ctx context.Context,tx *sql.Tx, idTrouble int) []domain.Component
	Save(ctx context.Context, tx *sql.Tx, idTrouble int, idComponent int, wg *sync.WaitGroup) bool
	DeleteByIdTrouble(ctx context.Context, tx *sql.Tx,idTrouble int, wg interface{})
}
