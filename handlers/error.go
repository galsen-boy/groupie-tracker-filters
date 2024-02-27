package handlers

import (
	"groupie-tracker/utils"
	"net/http"
	"text/template"
)

func Handle404Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusNotFound,
		"TextErr": "Page Not Found"}
	ErrorPages(w, utils.DataExec)
}

func Handle405Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusMethodNotAllowed,
		"TextErr": "Method Not Allowed"}
	ErrorPages(w, utils.DataExec)
}

func Handle400Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusBadRequest,
		"TextErr": "Bad Request"}
	ErrorPages(w, utils.DataExec)
}

func Handle500Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	utils.DataExec = map[string]interface{}{
		"ErrNum":  http.StatusInternalServerError,
		"TextErr": "Internal Server Error"}
	ErrorPages(w, utils.DataExec)
}

func ErrorPages(w http.ResponseWriter, data map[string]interface{}) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl.ExecuteTemplate(w, "header", data)
	tmpl.ExecuteTemplate(w, "error", data)
}
