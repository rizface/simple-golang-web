package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"github.com/go-playground/validator/v10"
	"os"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
	"pbl-orkom/repository"
	"strconv"
)

type maintenanceServiceImpl struct {
	db       *sql.DB
	validate *validator.Validate
	spek     repository.MaintenanceRepository
	detail   repository.DetailRepository
}

func NewMaintenanceServiceImpl(db *sql.DB, validate *validator.Validate, spek repository.MaintenanceRepository, detail repository.DetailRepository) MaintenanceService {
	return maintenanceServiceImpl{db: db, validate: validate, spek: spek, detail: detail}
}

func (service maintenanceServiceImpl) Get(ctx context.Context) []domain.DetailJoinMaintenance {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	detail := service.detail.Get(ctx,tx)
	return detail
}

func (service maintenanceServiceImpl) Save(ctx context.Context, spek web.MaintenanceRequest, detail web.DetailMaintenance) bool {
	err := service.validate.Struct(spek)
	helper.ValidationHandler(err)
	err = service.validate.Struct(detail)
	helper.ValidationHandler(err)
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	maintenance_id := service.spek.Save(ctx, tx, spek)
	detail.IdMaintenance = maintenance_id
	detail_maintenance := service.detail.Save(ctx, tx, detail)
	return detail_maintenance
}

func(service maintenanceServiceImpl) Delete(ctx context.Context,idMaintenance int) bool {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	result := service.spek.Delete(ctx,tx,idMaintenance)
	return result
}

func (service maintenanceServiceImpl) FormUpdate(ctx context.Context, idMaintenance int) domain.MaintenanceRequest {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data,err := service.spek.FindById(ctx,tx,idMaintenance)
	helper.PanicIfError(err)
	return data
}

func (service maintenanceServiceImpl) Update(ctx context.Context, idMaintenance int, spek web.MaintenanceRequest) bool {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data,err := service.spek.FindById(ctx,tx,idMaintenance)
	helper.PanicIfError(err)
	if spek.JumlahRam == "" {
		spek.JumlahRam = data.JumlahRam
	}
	if spek.ArchOS == "" {
		spek.ArchOS = data.ArchOS
	}
	result := service.spek.Update(ctx,tx,idMaintenance,spek)
	return result
}

func (service maintenanceServiceImpl) FormUpdateDetail(ctx context.Context, idMaintenance int) domain.DetailMaintenance {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data,err := service.detail.FindByIdMaintenance(ctx,tx,idMaintenance)
	helper.PanicIfError(err)
	return data
}

func (service maintenanceServiceImpl) UpdateDetail(ctx context.Context, idMaintenance int, detail web.DetailMaintenance) bool {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	result := service.detail.Update(ctx,tx,idMaintenance,detail)
	return result
}


func (service maintenanceServiceImpl) createCSV(maintenances []domain.DetailJoinMaintenance) string {
	data := [][]string{
		{"Petugas", "Lama Pengerjaan", "Motherboard", "Vendor RAM", "Jumlah RAM", "Graphic Card", "NIC","Sistem Opeasi","Arsitektur OS", "Detail Pengerjaan", "Tanggal Masuk"},
	}
	for _, v := range maintenances {
		lamaPengerjaan := strconv.Itoa(v.Detail.LamaPengerjaan)
		each := []string{
			v.Detail.NamaPetugas,lamaPengerjaan,v.Spek.Motherboard,v.Spek.VendorRam,v.Spek.JumlahRam,
			v.Spek.GraphicCard,v.Spek.NIC,v.Spek.SistemOperasi,v.Spek.ArchOS, v.Detail.DetailPengerjaan,v.Spek.TglMasuk,
		}
		data = append(data,each)
	}

	csvFile,err := os.Create("maintenance.csv")
	helper.PanicIfError(err)
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	for _, v := range data {
		err = csvWriter.Write(v)
		helper.PanicIfError(err)
	}
	return csvFile.Name()
}

func (service maintenanceServiceImpl) Export(ctx context.Context) string {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data := service.detail.Export(ctx,tx)
	fileName := service.createCSV(data)
	return fileName
}
