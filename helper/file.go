package helper

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StaticFile(mux *mux.Router) {
	dir := http.Dir("public")
	fileServer := http.FileServer(dir)
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
}
