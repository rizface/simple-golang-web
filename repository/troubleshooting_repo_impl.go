package repository

import (
	"context"
	"database/sql"
	"errors"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
	"sync"
)

type troubleshootingRepositoryImpl struct{}

func NewTroubleshootingRepo() TroubleshootingRepository {
	return troubleshootingRepositoryImpl{}
}

func (t troubleshootingRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Troubleshooting {
	sql := "SELECT id,nama_customer,biaya,DATE_FORMAT(created_at,'%w %M %Y') FROM troubleshooting ORDER BY id DESC"
	data := []domain.Troubleshooting{}
	rows,err := tx.QueryContext(ctx,sql);
	helper.PanicIfError(err)
	for rows.Next() {
		each := domain.Troubleshooting{}
		err := rows.Scan(&each.Id,&each.NamaCustomer,&each.Biaya,&each.TglMasuk)
		helper.PanicIfError(err)
		data = append(data,each)
	}
	return data
}

func (t troubleshootingRepositoryImpl) GetById(ctx context.Context, idTrouble int, tx *sql.Tx) domain.Troubleshooting {
	sql := "SELECT id,nama_customer,lama_pengerjaan,biaya,permasalahan,detail_pengerjaan,informasi_lainnya FROM troubleshooting WHERE id = ?"
	rows,err := tx.QueryContext(ctx,sql,idTrouble)
	helper.PanicIfError(err)
	defer rows.Close()
	each := domain.Troubleshooting{}
	if rows.Next() {
		err := rows.Scan(&each.Id,&each.NamaCustomer,&each.LamaPengerjaan,&each.Biaya,&each.Permasalahan,&each.DetailPengerjaan,&each.InformasiLainnya)
		helper.PanicIfError(err)
	} else {
		helper.PanicIfError(errors.New("Data Troubleshooting Tidak Ditemukan"))
	}
	return each
}

func (t troubleshootingRepositoryImpl) Save(ctx context.Context, request web.TroubleshootRequest, tx *sql.Tx) int64 {
	sql := "INSERT INTO troubleshooting(nama_customer, lama_pengerjaan, biaya, permasalahan, detail_pengerjaan, informasi_lainnya)" +
		"VALUES(?,?,?,?,?,?)"
	result,err := tx.ExecContext(ctx,sql,request.NamaCustomer,request.LamaPengerjaan,request.Biaya,request.Permasalahan,request.DetailPengerjaan,request.InformasiLainnya)
	helper.PanicIfError(err)
	id,_ := result.LastInsertId()
	return id
}

func (t troubleshootingRepositoryImpl) Update(ctx context.Context, idTrouble int, request web.TroubleshootRequest, tx *sql.Tx, wg interface{}) bool {
	rest,ok := wg.(*sync.WaitGroup)
	if ok {
		defer rest.Done()
	}
	sql := "UPDATE troubleshooting SET nama_customer = ?, lama_pengerjaan = ?, biaya = ?, permasalahan = ?, detail_pengerjaan = ?, informasi_lainnya = ? WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,request.NamaCustomer,request.LamaPengerjaan,request.Biaya,request.Permasalahan,request.DetailPengerjaan, request.InformasiLainnya,idTrouble)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (t troubleshootingRepositoryImpl) Delete(ctx context.Context, idTrouble int, tx *sql.Tx) bool {
	sql := "DELETE FROM troubleshooting WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,idTrouble)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (t troubleshootingRepositoryImpl) Export(ctx context.Context, tx *sql.Tx) []domain.Export {
	sql :=  "SELECT nama_customer, biaya, DATE_FORMAT(troubleshooting.created_at, '%w %M %Y'),(SELECT GROUP_CONCAT(component SEPARATOR \", \") FROM change_component INNER JOIN components ON components.id = change_component.id_component WHERE id_troubleshooting = troubleshooting.id) FROM troubleshooting\n"
	rows,err := tx.QueryContext(ctx,sql)
	helper.PanicIfError(err)
	defer rows.Close()
	data := []domain.Export{}
	for rows.Next() {
		each := domain.Export{}
		err := rows.Scan(&each.NamaCustomer,&each.Biaya,&each.TglMasuk,&each.ChangeComponent)
		helper.PanicIfError(err)
		data = append(data, each)
	}
	return data
}

