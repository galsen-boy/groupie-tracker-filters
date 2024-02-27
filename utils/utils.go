package utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

var Artists []models.Artist
var Locations models.Location
var Dates models.Date
var Relations models.Relation
var DataExec = map[string]interface{}{}
var Urls = map[string]string{}

func GetJson(w http.ResponseWriter, url string) []byte {
	response, err := http.Get(url)
	var responseData []byte
	if err != nil {
		DataExec := map[string]interface{}{
			"ErrNum":  http.StatusInternalServerError,
			"TextErr": "Internal Server Error"}
		ErrorPages(w, DataExec)
		return responseData 

	} else {

		responseData, _ = ioutil.ReadAll(response.Body)
		return responseData
	}
}
func GetJsons(w http.ResponseWriter) {
	json.Unmarshal(GetJson(w, Urls["artists"]), &Artists)
	json.Unmarshal(GetJson(w, Urls["locations"]), &Locations)
	json.Unmarshal(GetJson(w, Urls["relation"]), &Relations)
	json.Unmarshal(GetJson(w, Urls["dates"]), &Dates)
}

func ErrorPages(w http.ResponseWriter, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, data)
}

func FormatDate(dateStr string) string {
	dateStr = strings.TrimPrefix(dateStr, "*")
	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		os.Exit(1)
	}
	formatedDate := date.Format("Jan 2, 2006")

	return formatedDate
}
func FormatStr(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.Title(str)
	return str
}
