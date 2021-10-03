package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
	"pbl-orkom/repository"
)

type loginServiceImpl struct {
	db *sql.DB
	validate *validator.Validate
	repo repository.UserRepository
}

func NewLoginService(db *sql.DB, validate *validator.Validate, repo repository.UserRepository) LoginService {
	return loginServiceImpl{db:db, validate:validate, repo:repo}
}

func (l loginServiceImpl) Login(ctx context.Context, request web.LoginRequest) domain.User {
	tx,err := l.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := l.repo.FindByUsername(ctx,tx,request)
	err = helper.ComparePassword(request.Password,user.Password)
	helper.PanicIfError(err)
	return user
}


