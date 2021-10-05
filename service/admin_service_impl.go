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

type adminServiceImpl struct {
	db *sql.DB
	validate *validator.Validate
	repo repository.UserRepository
}

func NewAdminService(db *sql.DB, validate *validator.Validate, repo repository.UserRepository) AdminService {
	return adminServiceImpl{db:db,validate: validate,repo: repo}
}

func (a adminServiceImpl) Get(ctx context.Context) []domain.User {
	tx,err := a.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data := a.repo.Get(ctx,tx)
	return data
}

func (a adminServiceImpl) Save(ctx context.Context, request web.UserRequest) bool {
	err := a.validate.Struct(request)
	helper.ValidationHandler(err)
	tx,err := a.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	result := a.repo.Save(ctx,tx,request)
	return result
}

func (a adminServiceImpl) UpdateForm(ctx context.Context, idAdmin int) domain.User {
	tx,err := a.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := a.repo.FindById(ctx,tx,idAdmin)
	return user
}

func (a adminServiceImpl) Update(ctx context.Context, idAdmin int, request web.UserRequest) bool {
	tx,err := a.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := a.repo.FindById(ctx,tx,idAdmin)
	if len(request.Password) < 1 {
		request.Password = user.Password
	} else {
		hashed,err := helper.HashPassword(request.Password)
		helper.PanicIfError(err)
		request.Password = hashed
	}
	result := a.repo.Update(ctx,tx,idAdmin,request)
	return result
}

func (a adminServiceImpl) Delete(ctx context.Context, idAdmin int) bool {
	tx,err := a.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := a.repo.FindById(ctx,tx,idAdmin)
	result := a.repo.Delete(ctx,tx,user.Id)
	return result
}
