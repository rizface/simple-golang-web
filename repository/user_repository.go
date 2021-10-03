package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, request web.LoginRequest) bool
	Delete(ctx context.Context, tx *sql.Tx, idUser int) bool
	FindByUsername(ctx context.Context, tx *sql.Tx, request web.LoginRequest) domain.User
}
