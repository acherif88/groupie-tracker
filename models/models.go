package models

// Artist représente un artiste/groupe tel que renvoyé par l'API.

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Location représente les lieux de concert d'un artiste.
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Date représente les dates de concert d'un artiste.
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relation associe chaque ville à ses dates de concert pour un artiste donné.
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// ArtistDetail regroupe un artiste et ses informations de concerts,
// pour les transmettre ensemble à un template HTML.
type ArtistDetail struct {
	Artist  Artist
	Relation Relation
	Locations Location
	Dates	  Date
}

// HomeData regroupe la liste des artistes et le texte recherché,
// pour les transmettre ensemble à la page d'accueil.
type HomeData struct {
	Artists []Artist
	Search  string
}

// ErrorData contient les infos affichées sur la page d'erreur.
type ErrorData struct {
	Code    int
	Message string
}