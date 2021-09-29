package helper

import (
	"github.com/gorilla/mux"
	"net/http"
)

var NotFoundHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	ViewParser(w,"not-found", nil)
}

var MethodNotAllowedHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	ViewParser(w,"method-not-allowed", nil)
}

func StartServer(mux *mux.Router) {
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
