package repository

import (
	"context"
	"database/sql"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type troubleshootingRepositoryImpl struct{}

func NewTroubleshootingRepo() TroubleshootingRepository {
	return troubleshootingRepositoryImpl{}
}

func (t troubleshootingRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Troubleshooting {
	panic("implement me")
}

func (t troubleshootingRepositoryImpl) GetById(ctx context.Context, idTrouble int, tx *sql.Tx) domain.Troubleshooting {
	panic("implement me")
}

func (t troubleshootingRepositoryImpl) Save(ctx context.Context, request web.TroubleshootRequest, tx *sql.Tx) int64 {
	sql := "INSERT INTO troubleshooting(nama_customer, lama_pengerjaan, biaya, permasalahan, detail_pengerjaan, informasi_lainnya)" +
		"VALUES(?,?,?,?,?,?)"
	result,err := tx.ExecContext(ctx,sql,request.NamaCustomer,request.LamaPengerjaan,request.Biaya,request.Permasalahan,request.DetailPengerjaan,request.InformasiLainnya)
	helper.PanicIfError(err)
	id,_ := result.LastInsertId()
	return id
}

func (t troubleshootingRepositoryImpl) Update(ctx context.Context, idTrouble int, request web.TroubleshootRequest, tx *sql.Tx) bool {
	panic("implement me")
}

func (t troubleshootingRepositoryImpl) Delete(ctx context.Context, idTrouble int, tx *sql.Tx) bool {
	panic("implement me")
}
