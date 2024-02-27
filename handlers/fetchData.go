package handlers

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func fetchData() ([]models.FullArtistData, []string, error) {
	apiResponse, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		log.Println("API not responding")
		return nil, nil, err
	}
	defer apiResponse.Body.Close()

	if apiResponse.StatusCode != http.StatusOK {
		return nil, nil, errors.New("API problemes")
	}

	apiBody, _ := ioutil.ReadAll(apiResponse.Body)
	var apiData map[string]interface{}
	json.Unmarshal(apiBody, &apiData)

	artistApi, locationsApi, relationApi, datesApi := getAPIEndpoints(apiData)

	var artistInfo []models.FullArtistData
	var locations []string

	artistInfo, err = processArtists(artistApi)
	if err != nil {
		log.Println(err)
	}

	locations = processLocations(artistInfo, locationsApi)
	processLocation(&artistInfo, locationsApi)
	processRelations(&artistInfo, relationApi)
	processDates(&artistInfo, datesApi)

	return artistInfo, locations, nil
}

func getAPIEndpoints(data map[string]interface{}) (string, string, string, string) {
	return data["artists"].(string), data["locations"].(string), data["relation"].(string), data["dates"].(string)
}

func processArtists(artistApi string) ([]models.FullArtistData, error) {
	Artresponse, _ := http.Get(artistApi)

	Artbody, _ := ioutil.ReadAll(Artresponse.Body)
	defer Artresponse.Body.Close()
	var data []map[string]interface{}
	json.Unmarshal(Artbody, &data)

	var ArtInfo []models.FullArtistData
	for _, artistData := range data {
		artists := models.FullArtistData{
			ID:           artistData["id"].(float64),
			Image:        artistData["image"].(string),
			Name:         artistData["name"].(string),
			CreationDate: artistData["creationDate"].(float64),
			FirstAlbum:   artistData["firstAlbum"].(string),
			Members:      processMembers(artistData["members"].([]interface{})),
		}
		ArtInfo = append(ArtInfo, artists)
	}

	return ArtInfo, nil
}

func processMembers(memberData []interface{}) []string {
	var members []string
	for _, member := range memberData {
		members = append(members, member.(string))
	}
	return members
}

func processLocation(ArtInfo *[]models.FullArtistData, locationsApi string) {
	Locresponse, _ := http.Get(locationsApi)
	Locbody, _ := ioutil.ReadAll(Locresponse.Body)
	defer Locresponse.Body.Close()
	var Locdata map[string]interface{}
	json.Unmarshal(Locbody, &Locdata)

	for _, item := range Locdata["index"].([]interface{}) {
		location := item.(map[string]interface{})
		loc := location["locations"].([]interface{})
		for _, v := range loc {
			locationStr := v.(string)
			for i, artist := range *ArtInfo {
				if artist.ID == location["id"].(float64) {
					(*ArtInfo)[i].Locations = append(artist.Locations, locationStr)
				}
			}
		}
	}
}

func processLocations(ArtInfo []models.FullArtistData, locationsApi string) []string {
	Locresponse, err := http.Get(locationsApi)
	if err != nil {
		log.Println("Erreur lors de la requête vers l'API de localisation")
		return nil
	}
	defer Locresponse.Body.Close()

	Locbody, _ := ioutil.ReadAll(Locresponse.Body)
	var Locdata map[string]interface{}
	json.Unmarshal(Locbody, &Locdata)

	//---------- Utilisation d'une map pour éviter les duplications------------//
	uniqueLocations := make(map[string]struct{})

	for _, item := range Locdata["index"].([]interface{}) {
		location := item.(map[string]interface{})
		loc := location["locations"].([]interface{})
		for _, v := range loc {
			locationStr := v.(string)
			uniqueLocations[locationStr] = struct{}{}
		}
	}

	var allLocations []string

	//-----------Parcourir les localisations uniques et les ajouter au tableau-------------//
	for location := range uniqueLocations {
		allLocations = append(allLocations, location)
	}

	return allLocations
}

func processDates(ArtInfo *[]models.FullArtistData, datesApi string) {
	Datresponse, _ := http.Get(datesApi)

	Datbody, _ := ioutil.ReadAll(Datresponse.Body)
	defer Datresponse.Body.Close()
	var Datdata map[string]interface{}
	json.Unmarshal(Datbody, &Datdata)

	for _, item := range Datdata["index"].([]interface{}) {
		date := item.(map[string]interface{})
		dates := date["dates"].([]interface{})
		for _, v := range dates {
			DateStr := strings.TrimPrefix(v.(string), "*")
			for i, artist := range *ArtInfo {
				if artist.ID == date["id"].(float64) {
					DateString := DateStr
					(*ArtInfo)[i].ConcertDates = append(artist.ConcertDates, DateString)
				}
			}
		}
	}
}

func processRelations(ArtInfo *[]models.FullArtistData, relationApi string) {
	Relresponse, _ := http.Get(relationApi)
	Relbody, _ := ioutil.ReadAll(Relresponse.Body)
	defer Relresponse.Body.Close()
	var Reldata map[string]interface{}
	json.Unmarshal(Relbody, &Reldata)

	for _, item := range Reldata["index"].([]interface{}) {
		relation := item.(map[string]interface{})
		dateRel := relation["datesLocations"].(map[string]interface{})
		Mapi := map[string][]string{}
		artists := models.FullArtistData{}

		for key, value := range dateRel {
			test := value.([]interface{})
			var tab []string
			for _, v := range test {
				tab = append(tab, v.(string))
			}
			key := (key)
			Mapi[key] = tab
		}

		artists.Relations = Mapi

		for i, artist := range *ArtInfo {
			if artist.ID == relation["id"].(float64) {
				(*ArtInfo)[i].Relations = artists.Relations
			}
		}
	}
}
