package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"pbl-orkom/app"
	"pbl-orkom/helper"
	"pbl-orkom/model/web"
	"pbl-orkom/service"
	"strconv"
	"strings"
	"time"

	"tawesoft.co.uk/go/dialog"
)

type maintenanceControllerImpl struct {
	service service.MaintenanceService
}

func NewMaintenanceController(service service.MaintenanceService) MaintenanceController {
	return maintenanceControllerImpl{service: service}
}

func (controller maintenanceControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	data := controller.service.Get(r.Context())
	helper.ViewParser(w, "maintenance", map[string]interface{}{
		"data":data,
	})
}

func (controller maintenanceControllerImpl) Form(w http.ResponseWriter, r *http.Request) {
	helper.ViewParser(w, "maintenance-form", nil)
}

func (controller maintenanceControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	spek := web.MaintenanceRequest{
		Motherboard:   strings.ToLower(r.PostFormValue("motherboard")),
		VendorRam:     strings.ToLower(r.PostFormValue("vendor_ram")),
		JumlahRam:     strings.ToLower(r.PostFormValue("jumlah_ram")),
		GraphicCard:   strings.ToLower(r.PostFormValue("graphic_card")),
		NIC:           strings.ToLower(r.PostFormValue("nic")),
		SistemOperasi: strings.ToLower(r.PostFormValue("sistem_operasi")),
		ArchOS:        strings.ToLower(r.PostFormValue("arsitektur_os")),
	}
	durasi, _ := strconv.Atoi(r.PostFormValue("lama_pengerjaan"))
	detail := web.DetailMaintenance{
		NamaPetugas:      strings.ToLower(r.PostFormValue("nama_petugas")),
		LamaPengerjaan:   durasi,
		InformasiLainnya: strings.ToLower(r.PostFormValue("informasi_lainnya")),
		DetailPengerjaan: strings.ToLower(r.PostFormValue("detail_pengerjaan")),
	}
	result := controller.service.Save(r.Context(), spek, detail)
	if result {
		dialog.Alert("Data Preventive Maintenance Berhasil Ditambahkan")
	} else {
		dialog.Alert("Data Preventice Maintenance Gagal Ditambahkan")
	}
	http.Redirect(w, r, app.GET_MAINTENANCE, http.StatusSeeOther)
}

func (controller maintenanceControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.Atoi(params["idMaintenance"])
	result := controller.service.Delete(r.Context(),id)
	if result == true {
		dialog.Alert("Data Preventive Maintenance Berhasil Dihapus")
	} else {
		dialog.Alert("Data Preventive Maintenance Gagal Dihapus")
	}
	http.Redirect(w,r,app.GET_MAINTENANCE,http.StatusSeeOther)
}

func (controller maintenanceControllerImpl) FormUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMaintenance,_ := strconv.Atoi(params["idMaintenance"])
	data := controller.service.FormUpdate(r.Context(),idMaintenance)
	helper.ViewParser(w,"maintenance-form-update",map[string]interface{}{
		"data":data,
	})
}

func (controller maintenanceControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	idMaintenance,err := strconv.Atoi(r.PostFormValue("idMaintenance"))
	helper.PanicIfError(err)
	spek := web.MaintenanceRequest{
		Motherboard:   strings.ToLower(r.PostFormValue("motherboard")),
		VendorRam:     strings.ToLower(r.PostFormValue("vendor_ram")),
		JumlahRam:     strings.ToLower(r.PostFormValue("jumlah_ram")),
		GraphicCard:   strings.ToLower(r.PostFormValue("graphic_card")),
		NIC:           strings.ToLower(r.PostFormValue("nic")),
		SistemOperasi: strings.ToLower(r.PostFormValue("sistem_operasi")),
		ArchOS:        strings.ToLower(r.PostFormValue("arsitektur_os")),
	}
	result := controller.service.Update(r.Context(),idMaintenance,spek)
	if result {
		dialog.Alert("Spesifikasi Perangkat Berhasil Diupdate")
	} else {
		dialog.Alert("Spesifikasi Perangkat Gagal Diupdate")
	}
	http.Redirect(w,r,app.GET_MAINTENANCE,http.StatusSeeOther)
}

func (controller maintenanceControllerImpl) FormUpdateDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMaintenance,_ := strconv.Atoi(params["idMaintenance"])
	data := controller.service.FormUpdateDetail(r.Context(),idMaintenance)
	helper.ViewParser(w,"detail-maintenance-form",map[string]interface{} {
		"data":data,
	})
}

func (controller maintenanceControllerImpl) UpdateDetail(w http.ResponseWriter, r *http.Request) {
	idMaintenance,_ := strconv.Atoi(r.PostFormValue("id"))
	durasi, _ := strconv.Atoi(r.PostFormValue("lama_pengerjaan"))
	detail := web.DetailMaintenance{
		NamaPetugas:      strings.ToLower(r.PostFormValue("nama_petugas")),
		LamaPengerjaan:   durasi,
		InformasiLainnya: strings.ToLower(r.PostFormValue("informasi_lainnya")),
		DetailPengerjaan: strings.ToLower(r.PostFormValue("detail_pengerjaan")),
	}
	result := controller.service.UpdateDetail(r.Context(),idMaintenance,detail)
	if result {
		dialog.Alert("Detail Maintenance Berhasil Diupdate")
	} else {
		dialog.Alert("Detail Maintenance Gagal Diupdate")
	}
	http.Redirect(w,r,app.GET_MAINTENANCE,http.StatusSeeOther)
}

func (controller maintenanceControllerImpl) Export(w http.ResponseWriter, r *http.Request)  {
 	data := controller.service.Export(r.Context())
 	fmt.Println(data)
	w.Header().Add("Content-Disposition", "attachment; filename=\""+data+"\"")
	http.ServeFile(w,r,data)
 	time.Sleep(3 * time.Second)
 	os.Remove(data)
}
