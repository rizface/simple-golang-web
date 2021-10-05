package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type UserRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.User
	Save(ctx context.Context, tx *sql.Tx, request web.UserRequest) bool
	Delete(ctx context.Context, tx *sql.Tx, idUser int) bool
	FindByUsername(ctx context.Context, tx *sql.Tx, request web.UserRequest) domain.User
	FindById(ctx context.Context, tx *sql.Tx, idAdmin int) domain.User
	Update(ctx context.Context, tx *sql.Tx, idAdmin int, request web.UserRequest) bool
}
