package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"sync"
)

type componentRepositoryImpl struct{}

func NewComponentRepoImpl() ComponentRepository {
	return componentRepositoryImpl{}
}

func (c componentRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Component{
	sql := "SELECT id,component FROM components"
	data := []domain.Component{}
	rows,err := tx.QueryContext(ctx,sql)
	helper.PanicIfError(err)
	defer rows.Close()
	for rows.Next() {
		each := domain.Component{}
		err := rows.Scan(&each.Id.Int64,&each.Component)
		helper.PanicIfError(err)
		data = append(data, each)
	}
	return data
}

func (c componentRepositoryImpl) GetByIdTrouble(ctx context.Context, tx *sql.Tx, idTrouble int) []domain.Component {
	sql := "SELECT id,component, (SELECT id_component FROM change_component WHERE id_troubleshooting = ? AND id_component = components.id) AS include  FROM components"
	data := []domain.Component{}
	rows,err := tx.QueryContext(ctx,sql, idTrouble)
	helper.PanicIfError(err)
	defer rows.Close()
	for rows.Next() {
		each := domain.Component{}
		err := rows.Scan(&each.Id.Int64,&each.Component,&each.Include)
		helper.PanicIfError(err)
		data = append(data, each)
	}
	return data
}

func (c componentRepositoryImpl) DeleteByIdTrouble(ctx context.Context, tx *sql.Tx, idTrouble int, wg interface{})  {
	rest,ok := wg.(*sync.WaitGroup)
	if ok {
		rest.Done()
	}
	sql := "DELETE FROM change_component WHERE id_troubleshooting = ?"
	_,err := tx.ExecContext(ctx,sql,idTrouble)
	helper.PanicIfError(err)
}

func (c componentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, idTrouble int, idComponent int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	sql := "INSERT INTO change_component(id_troubleshooting,id_component) VALUES(?, ?)"
	result,err := tx.ExecContext(ctx,sql,idTrouble,idComponent)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}
