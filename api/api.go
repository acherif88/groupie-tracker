package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"groupie-tracker/models"
)

// defaultBaseURL est l'adresse de l'API Groupie Tracker.
// Elle peut être remplacée par la variable d'environnement API_URL.
const defaultBaseURL = "https://groupietrackers.01edu.ai/api"

// baseURL renvoie l'adresse de l'API à utiliser, sans « / » final.
func baseURL() string {
	if url := strings.TrimSpace(os.Getenv("API_URL")); url != "" {
		return strings.TrimRight(url, "/")
	}
	return defaultBaseURL
}

// GetArtists récupère la liste de tous les artistes depuis l'API.
func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist

	response, err := http.Get(baseURL() + "/artists")
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

	url := fmt.Sprintf("%s/artists/%d", baseURL(), id)

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

	url := fmt.Sprintf("%s/relation/%d", baseURL(), id)

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

	url := fmt.Sprintf("%s/locations/%d", baseURL(), id)

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

	url := fmt.Sprintf("%s/dates/%d", baseURL(), id)

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
