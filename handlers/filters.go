package handlers

import (
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"strconv"
	"strings"
)

func Filters(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data, location, _ := fetchData()

		dataExec := FilterData(data, r)
		DataExec1 := map[string]interface{}{"Artists": data, "Locations": location}
		renderTemplate(w, "header", DataExec1)
		renderTemplate(w, "artists", map[string]interface{}{"Artists": dataExec, "Location": DataExec1})
	} else {
		Handle405Error(w)
	}
}

func FilterData(data []models.FullArtistData, r *http.Request) []models.FullArtistData {
	minCreationDate := r.PostFormValue("MinCreationDate")
	maxCreationDate := r.PostFormValue("MaxCreationDate")
	minFirstAlbumDate := r.PostFormValue("MinFirstAlbumDate")
	maxFirstAlbumDate := r.PostFormValue("MaxFirstAlbumDate")
	loc := r.PostFormValue("loc")
	loc = strings.ToLower(loc)
	loc = strings.Replace(loc, ", ", "-", -1)
	loc = strings.Replace(loc, " ", "_", -1)
	members := getMembersNbr(r)

	if len(minCreationDate) != 4 || len(maxCreationDate) != 4 || len(minFirstAlbumDate) != 4 || len(maxFirstAlbumDate) != 4 {
		fmt.Println("Invalid Date")
		return data
	}
	mncreationDate, err := strconv.Atoi(minCreationDate)
	if err != nil {
		fmt.Println("Error convert to int:", err)
	}
	mxCreationDate, err := strconv.Atoi(maxCreationDate)
	if err != nil {
		fmt.Println("Error convert to int:", err)
	}

	mnFirstAlbumYear, err := strconv.Atoi(minFirstAlbumDate[len(minFirstAlbumDate)-4:])
	if err != nil {
		fmt.Println("Error convert to int:", err)
	}
	mxFirstAlbumYear, err := strconv.Atoi(maxFirstAlbumDate[len(maxFirstAlbumDate)-4:])
	if err != nil {
		fmt.Println("Error convert to int:", err)
	}
	if mnFirstAlbumYear < 1963 || mxFirstAlbumYear > 2018 || mncreationDate < 1958 || mxCreationDate > 2015 {
		fmt.Println("Invalid Date")
		return data
	}
	var result []models.FullArtistData
	for _, artist := range data {
		var check = true
		if int(artist.CreationDate) < mncreationDate || int(artist.CreationDate) > mxCreationDate {
			check = false
		}
		firstAlbumYear, err := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
		if err != nil {
			fmt.Println("Error convert to int:", err)
		}

		if firstAlbumYear < mnFirstAlbumYear || firstAlbumYear > mxFirstAlbumYear {
			check = false
		}
		if len(members) != 0 && !checkNbrMembers(members, len(artist.Members)) {
			check = false
		}
		if len(loc) != 0 && !checkLocations(artist.Locations, loc) {
			check = false
		}
		if check {
			result = append(result, artist)
		}
	}
	return result
}

func checkLocations(locations []string, location string) bool {
	for _, loc := range locations {
		if loc == location {
			return true
		}
	}
	return false
}

func checkNbrMembers(tabMembers []int, nbrMembers int) bool {
	for _, num := range tabMembers {
		if num == nbrMembers {
			return true
		}
	}
	return false
}

func getMembersNbr(r *http.Request) []int {
	var tabMembers []int
	if r.PostFormValue("Member1") == "on" {
		tabMembers = append(tabMembers, 1)
	}
	if r.PostFormValue("Member2") == "on" {
		tabMembers = append(tabMembers, 2)
	}
	if r.PostFormValue("Member3") == "on" {
		tabMembers = append(tabMembers, 3)
	}
	if r.PostFormValue("Member4") == "on" {
		tabMembers = append(tabMembers, 4)
	}
	if r.PostFormValue("Member5") == "on" {
		tabMembers = append(tabMembers, 5)
	}
	if r.PostFormValue("Member6") == "on" {
		tabMembers = append(tabMembers, 6)
	}
	if r.PostFormValue("Member7") == "on" {
		tabMembers = append(tabMembers, 7)
	}
	if r.PostFormValue("Member8") == "on" {
		tabMembers = append(tabMembers, 8)
	}
	return tabMembers
}
