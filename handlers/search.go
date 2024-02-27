package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		str := r.PostFormValue("browsers")
		data, location, _ := fetchData()
		dataExec := Recherche(data, str)
		DataExec1 := map[string]interface{}{"Artists": data, "Locations": location}
		renderTemplate(w, "header", DataExec1)
		renderTemplate(w, "artists", map[string]interface{}{"Artists": dataExec})
	} else {
		Handle405Error(w)
	}
}

func Recherche(data []models.FullArtistData, str string) []models.FullArtistData {
	str = strings.Split(str, " - ")[0]

	var result []models.FullArtistData
	for _, artist := range data {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(str)) && !IsDuplicated(artist.ID, result) {
			result = append(result, artist)
		} else if strings.Contains(fmt.Sprintf("%f", artist.CreationDate), str) && !IsDuplicated(artist.ID, result) {
			result = append(result, artist)
		} else if strings.Contains(artist.FirstAlbum, str) && !IsDuplicated(artist.ID, result) {
			result = append(result, artist)
		} else {
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), strings.ToLower(str)) && !IsDuplicated(artist.ID, result) {
					result = append(result, artist)
					break
				}
			}
			for i, location := range artist.Locations {
				if strings.Contains(strings.ToLower(location), strings.ToLower(str)) && !IsDuplicated(artist.ID, result) {
					result = append(result, artist)
					break
				} else if strings.Contains(artist.ConcertDates[i], str) && !IsDuplicated(artist.ID, result) {
					result = append(result, artist)
					break
				}
			}
		}
	}
	return result
}

func IsDuplicated(value float64, artists []models.FullArtistData) bool {
	for _, artist := range artists {
		if value == artist.ID {
			return true
		}
	}
	return false
}
