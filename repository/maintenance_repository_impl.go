package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type maintenanceRepositoryImpl struct{}

func NewRepoImpl() MaintenanceRepository {
	return maintenanceRepositoryImpl{}
}

func (repo maintenanceRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, spek web.MaintenanceRequest) int64 {
	sql := "INSERT INTO maintenance(motherboard,ram_vendor,ram_amount,gpu,nic,os,arsitektur_os) VALUES(?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, spek.Motherboard, spek.VendorRam, spek.JumlahRam, spek.GraphicCard, spek.NIC, spek.SistemOperasi, spek.ArchOS)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	return id
}

func(repo maintenanceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idMaintenance int) bool {
	sql := "DELETE FROM maintenance WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,idMaintenance)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (repo maintenanceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idMaintenance int) (domain.MaintenanceRequest,error) {
	sql := "SELECT id,motherboard,ram_vendor,ram_amount,gpu,nic,os,arsitektur_os,DATE_FORMAT(tgl_masuk, '%w %M %Y') FROM maintenance WHERE id = ?"
	rows,err := tx.QueryContext(ctx,sql,idMaintenance)
	helper.PanicIfError(err)
	defer rows.Close()
	data := domain.MaintenanceRequest{}
	if rows.Next() {
		err := rows.Scan(&data.Id,&data.Motherboard,&data.VendorRam,&data.JumlahRam,&data.GraphicCard,&data.NIC,&data.SistemOperasi,&data.ArchOS,&data.TglMasuk)
		helper.PanicIfError(err)
		return data,nil
	} else {
		fmt.Println(data)
		return data,errors.New("Data Preventive Maintenace Tidak Ditemukan")
	}
}

func(repo maintenanceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, idMaintenance int, spek web.MaintenanceRequest) bool {
	sql := "UPDATE maintenance SET motherboard = ?,ram_vendor = ?,ram_amount = ?,gpu = ?,nic = ?,os = ?,arsitektur_os = ? WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,spek.Motherboard,spek.VendorRam,spek.JumlahRam,spek.GraphicCard,spek.NIC,
		spek.SistemOperasi,spek.ArchOS,idMaintenance)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}