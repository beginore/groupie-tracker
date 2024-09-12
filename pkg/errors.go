package pkg

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type error_status struct {
	Message string
	Status  string
}

func ErrorHandler(w http.ResponseWriter, status int) {
	var Error error_status = error_status{}

	if status == 400 {
		Error.Status = "400"
		Error.Message = "Bad Request"
		w.WriteHeader(http.StatusBadRequest)
	} else if status == 404 {
		Error.Status = "404"
		Error.Message = "Not Found"
		w.WriteHeader(http.StatusNotFound)
	} else if status == 405 {
		Error.Status = "405"
		Error.Message = "Method Not Allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else if status == 500 {
		Error.Status = "500"
		Error.Message = "Internal Server Error"
		w.WriteHeader(http.StatusInternalServerError)
	}

	tmplPath := filepath.Join(TemplateDir, "error.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, Error); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
