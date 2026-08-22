package api

import (
	"fmt"
	"encoding/json"
	"net/http"

	"groupie-tracker/models"
)

// GetArtists récupère la liste de tous les artistes depuis l'API.
func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

// GetArtistByID récupère un seul artiste directement depuis l'API, via son ID.
func GetArtistByID(id int) (*models.Artist, error) {
	var artist models.Artist

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&artist)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

// GetRelation récupère les lieux et dates de concert liés à un artiste, via son ID.
func GetRelation(id int) (*models.Relation, error) {
	var relation models.Relation

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&relation)
	if err != nil {
		return nil, err
	}

	return &relation, nil
}

// GetLocations récupère les lieux de concert liés à un artiste, via son ID.
func GetLocations(id int) (*models.Location, error) {
	var location models.Location

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

// GetDates récupère les dates de concert liées à un artiste, via son ID.
func GetDates(id int) (*models.Date, error) {
	var date models.Date

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&date)
	if err != nil {
		return nil, err
	}

	return &date, nil
}