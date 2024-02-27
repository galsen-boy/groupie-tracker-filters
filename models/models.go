package models

type Artist struct {
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationUrl     string   `json:"locations"`
	ConcertDatesUrl string   `json:"concertDates"`
	RelationsUrl    string   `json:"relations"`
}

type FullArtistData struct {
	ID           float64
	Image        string
	Name         string
	Members      []string
	CreationDate float64
	Locations    []string
	FirstAlbum   string
	Relations    map[string][]string
	ConcertDates []string
}
type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
	DatesUrl string   `json:"dates"`
}

type Date struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
