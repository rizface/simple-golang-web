package helper

import (
	"github.com/gorilla/sessions"
	"net/http"
	"pbl-orkom/model/domain"
)

var Store = sessions.NewCookieStore([]byte("simple"))

func SetSession(w http.ResponseWriter, r *http.Request, user domain.User) {
	session,err := Store.Get(r,"admin")
	PanicIfError(err)
	session.Values["username"] = user.Username
	session.Values["login"] = true
	session.Options.MaxAge = 60 * 60
	session.Save(r,w)
}
