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

func (u userRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.User {
	data := []domain.User{}
	sql := "SELECT id,username FROM users"
	rows,err := tx.QueryContext(ctx,sql)
	helper.PanicIfError(err)
	defer rows.Close()
	for rows.Next() {
		each := domain.User{}
		err := rows.Scan(&each.Id, &each.Username)
		helper.PanicIfError(err)
		data = append(data,each)
	}
	return data
}

func (u userRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, request web.UserRequest) bool {
	sql := "INSERT INTO users(username,password) VALUES(?,?)"
	hashed,err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)
	result,err := tx.ExecContext(ctx,sql,request.Username,hashed)
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

func findByFilter(ctx context.Context, tx *sql.Tx, query string, idAdmin interface{}, username interface{}) domain.User {
	var rows *sql.Rows
	var err error

	if idAdmin == nil {
		rows,err = tx.QueryContext(ctx,query,username)
	} else {
		rows,err = tx.QueryContext(ctx,query,idAdmin.(int))
	}
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id,&user.Username,&user.Password)
		helper.PanicIfError(err)
		return user
	} else {
		err := errors.New("Data Tidak Ditemukan")
		panic(err)
	}
}

func (u userRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, request web.UserRequest) domain.User {
	sql := "SELECT id,username,password FROM users WHERE username = ?"
	user := findByFilter(ctx,tx,sql,nil,request.Username)
	return user
}

func (u userRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idAdmin int) domain.User {
	sql := "SELECT id,username,password FROM users WHERE id = ?"
	user := findByFilter(ctx,tx,sql,idAdmin,nil)
	return user
}

func (u userRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, idAdmin int, request web.UserRequest) bool {
	sql := "UPDATE users SET username = ?, password = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, sql, request.Username, request.Password, idAdmin)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return  affected > 0
}
