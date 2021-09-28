package controller

import (
	"net/http"
	"pbl-orkom/helper"
	"pbl-orkom/model/web"
	"pbl-orkom/service"
	"strconv"
	"tawesoft.co.uk/go/dialog"
)

type troubleshootingControllerImpl struct {
	service service.TroubleshootingService
}

func NewTroubleshootingImpl(service service.TroubleshootingService) TroubleshootingController {
	return troubleshootingControllerImpl{service: service}
}

func (t troubleshootingControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (t troubleshootingControllerImpl) GetById(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (t troubleshootingControllerImpl) FormSave(w http.ResponseWriter, r *http.Request) {
	component := t.service.SaveForm(r.Context())
	helper.ViewParser(w,"troubleshooting-save-form",map[string]interface{} {
		"component":component,
	})
}

func (t troubleshootingControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.PanicIfError(err)
	durasi,_ := strconv.Atoi(r.PostFormValue("lama_pengerjaan"))
	biaya,_ := strconv.Atoi(r.PostFormValue("biaya"))
	trouble := web.TroubleshootRequest{
		NamaCustomer: r.PostFormValue("nama_customer"),
		LamaPengerjaan: durasi,
		Biaya: biaya,
		Permasalahan: r.PostFormValue("permasalahan"),
		InformasiLainnya: r.PostFormValue("informasi_lainnya"),
		DetailPengerjaan: r.PostFormValue("detail_pengerjaan"),
		ComponentId: r.PostForm["components"],
	}
	success := t.service.Save(r.Context(),trouble)

	if success {
		dialog.Alert("Data Troubleshooting Berhasil Ditambahkan")
	} else {
		dialog.Alert("Data Troubleshooting Gagal Ditambahkan")
	}
	http.Redirect(w,r,r.Header.Get("Referer"),http.StatusSeeOther)
}

func (t troubleshootingControllerImpl) UpdateForm(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (t troubleshootingControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (t troubleshootingControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

