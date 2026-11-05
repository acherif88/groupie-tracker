package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

// renderError affiche une page d'erreur stylée avec le code et le message donnés.
func renderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Si même la page d'erreur ne charge pas, on retombe sur un message basique.
		http.Error(w, message, code)
		return
	}

	tmpl.Execute(w, models.ErrorData{Code: code, Message: message})
}

// filterArtists renvoie uniquement les artistes dont le nom contient le texte recherché.
func filterArtists(artists []models.Artist, search string) []models.Artist {
	if search == "" {
		return artists
	}

	var filtered []models.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(search)) {
			filtered = append(filtered, artist)
		}
	}
	return filtered
}

// HomeHandler répond aux demandes sur la page d'accueil ("/"), avec recherche optionnelle.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	artists, err := api.GetArtists()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Erreur lors de la récupération des artistes")
		return
	}

	filteredArtists := filterArtists(artists, search)

	data := models.HomeData{
		Artists: filteredArtists,
		Search:  search,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Erreur lors de la récupération des artistes")
		return
	}
	tmpl.Execute(w, data)
}

// ArtistHandler répond aux demandes sur la page de détail d'un artiste ("/artist?id=...").
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		renderError(w, http.StatusBadRequest, "Identifiant invalide")
		return
	}

	selectedArtist, err := api.GetArtistByID(id)
	if err != nil || selectedArtist.Name == "" {
		renderError(w, http.StatusNotFound, "Artiste introuvable")
		return
	}
	relation, err := api.GetRelation(id)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Erreur lors de la récupération de la relation")
		return
	}

	locations, err := api.GetLocations(id)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Erreur lors de la récupération des lieux")
		return
	}

	dates, err := api.GetDates(id)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Erreur lors de la récupération des dates")
		return
	}

	detail := models.ArtistDetail{
		Artist:    *selectedArtist,
		Relation:  *relation,
		Locations: *locations,
		Dates:     *dates,
	}
	// Nettoie les données brutes de l'API (« * » devant les dates, lieux au format technique).
	cleanDetail(&detail)

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Erreur lors de la récupération des artistes")
		return
	}
	tmpl.Execute(w, detail)
}
