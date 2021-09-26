package helper

import (
	"html/template"
	"net/http"
)

func ViewParser(rw http.ResponseWriter, view string, data map[string]interface{}) {
	tmp := template.Must(template.ParseGlob("view/*/*.gohtml"))
	err := tmp.ExecuteTemplate(rw, view, data)
	PanicIfError(err)
}
