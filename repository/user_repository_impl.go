package repository

import (
	"context"
	"database/sql"
	"errors"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type userRepositoryImpl struct {}

func NewUserRepository() UserRepository {
	return userRepositoryImpl{}
}

func (u userRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, request web.LoginRequest) bool {
	sql := "INSERT INTO users(username,password) VALUES(?,?)"
	result,err := tx.ExecContext(ctx,sql,request.Username,request.Password)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (u userRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idUser int) bool {
	sql := "DELETE FROM users WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql ,idUser)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (u userRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, request web.LoginRequest) domain.User {
	sql := "SELECT id,username,password FROM users WHERE username = ?"
	user := domain.User{}
	rows,err := tx.QueryContext(ctx,sql,request.Username)
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id,&user.Username,&user.Password)
		helper.PanicIfError(err)
		return user
	}
	panic(errors.New("email kamu tidak terdaftar"))
}

