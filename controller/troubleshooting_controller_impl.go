package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"pbl-orkom/app"
	"pbl-orkom/helper"
	"pbl-orkom/model/web"
	"pbl-orkom/service"
	"strconv"
	"tawesoft.co.uk/go/dialog"
	"time"
)

type troubleshootingControllerImpl struct {
	service service.TroubleshootingService
}

func NewTroubleshootingImpl(service service.TroubleshootingService) TroubleshootingController {
	return troubleshootingControllerImpl{service: service}
}

func (t troubleshootingControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	data := t.service.Get(r.Context())
	helper.ViewParser(w,"troubleshooting",map[string]interface{} {
		"data":data,
	})
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
	params := mux.Vars(r)
	idTrouble,err := strconv.Atoi(params["idTrouble"])
	helper.PanicIfError(err)
	trouble,component := t.service.UpdateFrom(r.Context(),idTrouble)
	helper.ViewParser(w,"troubleshooting-update-form", map[string]interface{} {
		"trouble":trouble,
		"component":component,
	})
}

func (t troubleshootingControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.PanicIfError(err)
	params := mux.Vars(r)
	idTrouble,err := strconv.Atoi(params["idTrouble"])
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
	success := t.service.Update(r.Context(), idTrouble,trouble)
	if success {
		dialog.Alert("Data Troubleshooting Berhasil Diupdate")
	} else {
		dialog.Alert("Data Troubleshooting Gagal Diupdate")
	}
	http.Redirect(w,r, app.GET_TROUBLESHOOTING,http.StatusSeeOther)
}

func (t troubleshootingControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idTrouble,err := strconv.Atoi(params["idTrouble"])
	helper.PanicIfError(err)
	success := t.service.Delete(r.Context(),idTrouble)
	if success {
		dialog.Alert("Data Troubleshooting Berhasil Dihapus")
	} else {
		dialog.Alert("Data Troubleshooting Gagal Dihapus")
	}
	http.Redirect(w,r,r.Header.Get("Referer"),http.StatusSeeOther)
}

func (t troubleshootingControllerImpl) Export(w http.ResponseWriter, r *http.Request)  {
	data := t.service.Export(r.Context())
	w.Header().Add("Content-Disposition", "attachment; filename=\""+data+"\"")
	http.ServeFile(w,r,data)
	time.Sleep(3 * time.Second)
	os.Remove(data)
}
