package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"pbl-orkom/helper"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
	"pbl-orkom/repository"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
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

func (service maintenanceServiceImpl) createExcel(maintenances []domain.DetailJoinMaintenance) string {
	fileName:= strconv.Itoa(int(time.Now().Unix()))
	xlsx := excelize.NewFile()
	xlsx.NewSheet("Preventive Maintenance")
	xlsx.SetColWidth("Preventive Maintenance", "A","K", 20)
	xlsx.AutoFilter("Preventive Maintenance","A1","K1","")
	xlsx.SetCellValue("Preventive Maintenance", "A1", "No")
	xlsx.SetCellValue("Preventive Maintenance", "B1", "Petugas")
	xlsx.SetCellValue("Preventive Maintenance", "C1", "Lama Pengerjaan")
	xlsx.SetCellValue("Preventive Maintenance", "D1", "Motherboard")
	xlsx.SetCellValue("Preventive Maintenance", "E1", "Vendor RAM")
	xlsx.SetCellValue("Preventive Maintenance", "F1", "Jumlah RAM")
	xlsx.SetCellValue("Preventive Maintenance", "G1", "Graphic Card")
	xlsx.SetCellValue("Preventive Maintenance", "H1", "NIC")
	xlsx.SetCellValue("Preventive Maintenance", "I1", "Arsitektur OS")
	xlsx.SetCellValue("Preventive Maintenance", "J1", "Detail Pengerjaan")
	xlsx.SetCellValue("Preventive Maintenance", "K1", "Tanggal Masuk")

	id := 1
	row := 2

	for _, v := range maintenances {
		xlsx.SetCellValue("Preventive Maintenance",  fmt.Sprintf("A%d",row), id)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("B%d",row), v.Detail.NamaPetugas)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("C%d",row), strconv.Itoa(v.Detail.LamaPengerjaan) + " Hari")
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("D%d",row), v.Spek.Motherboard)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("E%d",row), v.Spek.VendorRam)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("F%d",row), v.Spek.JumlahRam)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("G%d",row), v.Spek.GraphicCard)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("H%d",row), v.Spek.NIC)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("I%d",row), v.Spek.ArchOS)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("J%d",row), v.Detail.DetailPengerjaan)
		xlsx.SetCellValue("Preventive Maintenance", fmt.Sprintf("K%d",row), v.Spek.TglMasuk)
		id += 1
		row += 1
	}
	err := xlsx.SaveAs(fileName+".xlsx")
	helper.PanicIfError(err)
	return fileName+".xlsx"
}

func (service maintenanceServiceImpl) Export(ctx context.Context) string {
	tx,err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	data := service.detail.Export(ctx,tx)
	fileName := service.createExcel(data)
	return fileName
}
