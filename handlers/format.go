package handlers

import (
	"strings"
	"unicode"

	"groupie-tracker/models"
)

// Codes de pays courts affichés en majuscules (usa → USA).
var countryAcronyms = map[string]bool{"usa": true, "uk": true, "uae": true}

// cleanDate retire le marqueur « * » que l'API ajoute devant certaines dates.
func cleanDate(date string) string {
	return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(date), "*"))
}

// formatLocation transforme un lieu de l'API en texte lisible :
// « new_hampshire-usa » devient « New Hampshire, USA ».
func formatLocation(location string) string {
	location = strings.TrimSpace(location)
	sep := strings.LastIndex(location, "-")
	if sep < 0 {
		return titleWords(location)
	}
	city := titleWords(location[:sep])
	country := location[sep+1:]
	if countryAcronyms[strings.ToLower(country)] {
		country = strings.ToUpper(country)
	} else {
		country = titleWords(country)
	}
	return city + ", " + country
}

// titleWords remplace les « _ » par des espaces et met une majuscule à chaque mot.
func titleWords(text string) string {
	words := strings.Fields(strings.ReplaceAll(text, "_", " "))
	for i, word := range words {
		runes := []rune(strings.ToLower(word))
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

// cleanDetail nettoie les dates et les lieux d'un artiste avant l'affichage.
func cleanDetail(detail *models.ArtistDetail) {
	for i, date := range detail.Dates.Dates {
		detail.Dates.Dates[i] = cleanDate(date)
	}
	for i, location := range detail.Locations.Locations {
		detail.Locations.Locations[i] = formatLocation(location)
	}
	cleaned := make(map[string][]string, len(detail.Relation.DatesLocations))
	for location, dates := range detail.Relation.DatesLocations {
		list := make([]string, len(dates))
		for i, date := range dates {
			list[i] = cleanDate(date)
		}
		cleaned[formatLocation(location)] = list
	}
	detail.Relation.DatesLocations = cleaned
}
