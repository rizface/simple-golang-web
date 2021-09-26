package repository

import (
	"context"
	"database/sql"
	"errors"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type detailRepositoryImpl struct{}

func NewDetailRepositoryImpl() DetailRepository {
	return detailRepositoryImpl{}
}

func (repo detailRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.DetailJoinMaintenance {
	data := []domain.DetailJoinMaintenance{}
	sql := `SELECT id_maintenance ,nama_petugas,lama_pengerjaan, DATE_FORMAT(tgl_masuk,'%w %M %Y') FROM detail_maintenance
	INNER JOIN maintenance ON maintenance.id = detail_maintenance.id_maintenance`
	rows,err := tx.QueryContext(ctx,sql)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		each := domain.DetailJoinMaintenance{}
		err := rows.Scan(&each.Detail.IdMaintenance,&each.Detail.NamaPetugas,&each.Detail.LamaPengerjaan,&each.Spek.TglMasuk)
		helper.PanicIfError(err)
		data = append(data, each)
	}
	return data
}

func (repo detailRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, detail web.DetailMaintenance) bool {
	sql := "INSERT INTO detail_maintenance(id_maintenance,nama_petugas,lama_pengerjaan,informasi_lainnya,detail_pengerjaan) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, detail.IdMaintenance, detail.NamaPetugas, detail.LamaPengerjaan, detail.InformasiLainnya, detail.DetailPengerjaan)
	helper.PanicIfError(err)
	affected, err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (repo detailRepositoryImpl) FindByIdMaintenance(ctx context.Context, tx *sql.Tx, idMaintenance int) (domain.DetailMaintenance,error) {
	sql := "SELECT id,nama_petugas,lama_pengerjaan,informasi_lainnya,detail_pengerjaan FROM detail_maintenance WHERE id_maintenance = ?"
	rows,err := tx.QueryContext(ctx,sql,idMaintenance)
	helper.PanicIfError(err)
	defer rows.Close()
	data := domain.DetailMaintenance{}
	if rows.Next() {
		err := rows.Scan(&data.Id,&data.NamaPetugas,&data.LamaPengerjaan,&data.InformasiLainnya,&data.DetailPengerjaan)
		helper.PanicIfError(err)
		return data,nil
	} else {
		return data,errors.New("Data Detail Maintenance Tidak Ditemukan")
	}
}

func (repo detailRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, idMaintenance int, detail web.DetailMaintenance) bool {
	sql := "UPDATE detail_maintenance SET nama_petugas = ?, lama_pengerjaan = ?, informasi_lainnya = ?, detail_pengerjaan = ? WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,detail.NamaPetugas,detail.LamaPengerjaan,detail.InformasiLainnya,detail.DetailPengerjaan,idMaintenance)
	helper.PanicIfError(err)
	affected,_ := result.RowsAffected()
	return  affected > 0
}

func (repo detailRepositoryImpl) Export(ctx context.Context, tx *sql.Tx) []domain.DetailJoinMaintenance {
	sql := "SELECT nama_petugas,lama_pengerjaan, maintenance.motherboard,maintenance.ram_vendor,CONCAT(maintenance.ram_amount, \" GB\") AS jumlah_ram,maintenance.gpu,maintenance.nic,CONCAT(maintenance.arsitektur_os, \" BIT \") AS arsitektur_os,detail_pengerjaan, DATE_FORMAT(maintenance.tgl_masuk,'%w %M %Y') AS tgl_mulai FROM detail_maintenance INNER JOIN maintenance ON maintenance.id = detail_maintenance.id_maintenance \n"
	rows,err := tx.QueryContext(ctx,sql)
	helper.PanicIfError(err)
	defer rows.Close()
	data := []domain.DetailJoinMaintenance{}
	for rows.Next() {
		each := domain.DetailJoinMaintenance{}
		err := rows.Scan(&each.Detail.NamaPetugas,&each.Detail.LamaPengerjaan,
			&each.Spek.Motherboard,&each.Spek.VendorRam,&each.Spek.JumlahRam,&each.Spek.GraphicCard,
			&each.Spek.NIC,&each.Spek.ArchOS,&each.Detail.DetailPengerjaan,&each.Spek.TglMasuk)
		helper.PanicIfError(err)
		data = append(data,each)
	}
	return data
}
