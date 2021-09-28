package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
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
		err := rows.Scan(&each.Id,&each.Component)
		helper.PanicIfError(err)
		data = append(data, each)
	}
	return data
}

func (c componentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, idTrouble int, componentId []int) bool {
	sql := "INSERT INTO change_component(id_troubleshooting,id_component) VALUES(?, ?)"
	stmt,_ := tx.PrepareContext(ctx,sql)
	for _, v := range componentId {
		_,err := stmt.ExecContext(ctx,idTrouble,v)
		helper.PanicIfError(err)
	}
	return true
}
