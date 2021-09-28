package middleware

import (
	"net/http"

	"tawesoft.co.uk/go/dialog"
)

func CheckError(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			msg := recover()
			if msg != nil {
				dialog.Alert(msg.(error).Error())
				http.Redirect(rw, r, r.Header.Get("Referer"), http.StatusSeeOther)
			}
		}()
		next.ServeHTTP(rw, r)
	})
}
