package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

type Artist struct {
	Id             int                 `json:"id"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	Image          string              `json:"image"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Relations      string              `json:"relations"`
	DatesLocations map[string][]string `json:"-"`
}

func GetAPI() error {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return err
	}

	return nil
}

func fetchRelationsForArtist(artist *Artist) error {
	resp, err := http.Get(artist.Relations)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	type RelationsResponse struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
	var relationsResp RelationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&relationsResp); err != nil {
		log.Printf("Error decoding relations data: %v", err)
		return err
	}

	artist.DatesLocations = relationsResp.DatesLocations

	return nil
}
