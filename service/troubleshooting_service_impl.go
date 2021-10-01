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
	"sync"
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

func saveComponent(t troubleshootingServiceImpl, request web.TroubleshootRequest, ctx context.Context, idTrouble int, tx *sql.Tx, wg *sync.WaitGroup) {
	defer wg.Done()

	componentWG := sync.WaitGroup{}
	componentWG.Add(len(request.ComponentId))
	for _, v := range request.ComponentId {
		componentId,err := strconv.Atoi(v)
		helper.PanicIfError(err)
		go t.componentRepo.Save(ctx,tx,idTrouble,componentId, &componentWG)
	}
	componentWG.Wait()
}

func (t troubleshootingServiceImpl) Save(ctx context.Context, request web.TroubleshootRequest) bool {
	wg := sync.WaitGroup{}

	err := t.validate.Struct(request)
	helper.ValidationHandler(err)

	tx,err := t.db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	id := t.repo.Save(ctx,request,tx)
	if len(request.ComponentId)  == 0 {
		return true
	}
	wg.Add(1)
		go saveComponent(t,request,ctx,int(id),tx,&wg)
	wg.Wait()
	return true
}

func (t troubleshootingServiceImpl) UpdateFrom(ctx context.Context, idTrouble int) (domain.Troubleshooting, []domain.Component) {
	tx, err := t.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	trouble := t.repo.GetById(ctx,idTrouble,tx)
	component := t.componentRepo.GetByIdTrouble(ctx,tx,idTrouble)
	return trouble,component
}

func (t troubleshootingServiceImpl) Update(ctx context.Context, idTrouble int, request web.TroubleshootRequest) bool {
	wg := sync.WaitGroup{}

	err := t.validate.Struct(request)
	helper.ValidationHandler(err)

	tx,err := t.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	trouble := t.repo.GetById(ctx, idTrouble, tx)

	t.componentRepo.DeleteByIdTrouble(ctx,tx,trouble.Id,nil)
	wg.Add(2)
		go t.repo.Update(ctx,trouble.Id,request,tx,&wg)
		go saveComponent(t,request,ctx,trouble.Id,tx,&wg)
	wg.Wait()
	return true
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

