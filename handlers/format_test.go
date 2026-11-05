package handlers

import (
	"testing"

	"groupie-tracker/models"
)

func TestCleanDate(t *testing.T) {
	tests := map[string]string{
		"*16-01-2020":   "16-01-2020",
		"23-05-2020":    "23-05-2020",
		" *07-09-2020 ": "07-09-2020",
	}
	for input, want := range tests {
		if got := cleanDate(input); got != want {
			t.Errorf("cleanDate(%q) = %q, attendu %q", input, got, want)
		}
	}
}

func TestFormatLocation(t *testing.T) {
	tests := map[string]string{
		"new_hampshire-usa":     "New Hampshire, USA",
		"lyon-france":           "Lyon, France",
		"amsterdam-netherlands": "Amsterdam, Netherlands",
		"auckland-new_zealand":  "Auckland, New Zealand",
		"london-uk":             "London, UK",
		"paris":                 "Paris",
	}
	for input, want := range tests {
		if got := formatLocation(input); got != want {
			t.Errorf("formatLocation(%q) = %q, attendu %q", input, got, want)
		}
	}
}

func TestCleanDetail(t *testing.T) {
	detail := models.ArtistDetail{
		Dates:     models.Date{Dates: []string{"*16-01-2020"}},
		Locations: models.Location{Locations: []string{"new_york-usa"}},
		Relation: models.Relation{DatesLocations: map[string][]string{
			"oslo-norway": {"*07-11-2019"},
		}},
	}
	cleanDetail(&detail)

	if detail.Dates.Dates[0] != "16-01-2020" {
		t.Errorf("date non nettoyée : %q", detail.Dates.Dates[0])
	}
	if detail.Locations.Locations[0] != "New York, USA" {
		t.Errorf("lieu non formaté : %q", detail.Locations.Locations[0])
	}
	dates, ok := detail.Relation.DatesLocations["Oslo, Norway"]
	if !ok || dates[0] != "07-11-2019" {
		t.Errorf("relation non nettoyée : %v", detail.Relation.DatesLocations)
	}
}
