package handlers

import (
	"groupie-tracker/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HandleHome(w http.ResponseWriter, r *http.Request) {
	data, location, err := fetchData()
	utils.GetJsons(w)
	if r.URL.Path != "/" {
		Handle404Error(w)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		Handle405Error(w)
		return
	}

	if err != nil {
		log.Println(err)
	}
	dataArtist := map[string]interface{}{"Artists": data,
		"Locations": location}

	renderTemplate(w, "index", dataArtist)
}

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	data, location, err := fetchData()
	if err != nil {
		log.Println(err)
	}
	utils.GetJsons(w)
	if r.URL.Path != "/artists" {
		Handle404Error(w)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		Handle405Error(w)
		return
	}
	DataExec1 := map[string]interface{}{"Artists": data, "Locations": location}
	renderTemplate(w, "header", DataExec1)
	renderTemplate(w, "artists", DataExec1)
}

func ViewArtist(w http.ResponseWriter, r *http.Request) {
	tabURL := strings.Split(r.URL.String(), "/")
	if len(tabURL) > 3 {
		Handle404Error(w)
		return
	}

	id, err := strconv.Atoi(tabURL[len(tabURL)-1])

	if id < 1 || id > 52 || err != nil {
		Handle400Error(w)
		return
	}
	if id > 0 {
		id = id - 1
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		Handle405Error(w)
		return
	}
	data, location, err := fetchData()
	if err != nil {
		log.Println(err)
	}
	DataExec1 := map[string]interface{}{"Artists": data, "Artist": data[id], "Locations": location}

	renderTemplate(w, "viewArtist", DataExec1)
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		Handle500Error(w)
	}
}
