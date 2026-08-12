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