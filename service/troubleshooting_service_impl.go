package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
	"pbl-orkom/repository"
	"strconv"
)

type troubleshootingServiceImpl struct {
	db *sql.DB
	validate *validator.Validate
	repo repository.TroubleshootingRepository
	componentRepo repository.ComponentRepository
}


func NewTroubleshootingImpl(db *sql.DB,validate *validator.Validate, repo repository.TroubleshootingRepository, componentRepo repository.ComponentRepository) TroubleshootingService {
	return troubleshootingServiceImpl{db:db,validate: validate,repo: repo, componentRepo: componentRepo}
}

func (t troubleshootingServiceImpl) Get(ctx context.Context)[]domain.Troubleshooting {
	tx,err := t.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	result := t.repo.Get(ctx,tx)
	return result;
}

func (t troubleshootingServiceImpl) GetById(ctx context.Context, idTrouble int) {
	panic("implement me")
}

func (t troubleshootingServiceImpl) SaveForm(ctx context.Context)  []domain.Component{
	tx,err := t.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data := t.componentRepo.Get(ctx,tx)
	return data
}

func (t troubleshootingServiceImpl) Save(ctx context.Context, request web.TroubleshootRequest) bool {
	err := t.validate.Struct(request)
	helper.ValidationHandler(err)
	var componentId []int
	tx,err := t.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	id := t.repo.Save(ctx,request,tx)
	if len(request.ComponentId)  == 0 {
		return true
	}
	for _, v := range request.ComponentId {
		cId, _ := strconv.Atoi(v)
		componentId = append(componentId, cId)
	}
	return t.componentRepo.Save(ctx,tx,int(id),componentId)
}

func (t troubleshootingServiceImpl) Update(ctx context.Context, idTrouble int, request web.TroubleshootRequest) bool {
	panic("implement me")
}

func (t troubleshootingServiceImpl) Delete(ctx context.Context, idTrouble int) bool {
	tx,err := t.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	find := t.repo.GetById(ctx,idTrouble,tx)
	if find.Id > 0 {
		return t.repo.Delete(ctx,idTrouble,tx)
	}
	return false
}

