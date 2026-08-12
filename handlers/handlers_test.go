package handlers

import (
	"testing"

	"groupie-tracker/models"
)

func TestFilterArtists(t *testing.T) {
	artists := []models.Artist{
		{ID: 1, Name: "Queen"},
		{ID: 2, Name: "Pink Floyd"},
		{ID: 3, Name: "Led Zeppelin"},
	}

	result := filterArtists(artists, "queen")

	if len(result) != 1 {
		t.Errorf("attendu 1 résultat, obtenu %d", len(result))
	}

	if result[0].Name != "Queen" {
		t.Errorf("attendu Queen, obtenu %s", result[0].Name)
	}
}