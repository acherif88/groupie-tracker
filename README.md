# Groupie Tracker

Site web en Go qui affiche des artistes et leurs concerts à partir d'une API JSON publique.
Projet réalisé en équipe à **Zone 01 Rouen** (titre Concepteur Développeur d'Applications).

## Fonctionnalités

- Liste des artistes avec recherche par nom.
- Page détaillée d'un artiste : membres, date de création, premier album, lieux et dates de concerts.
- Pages d'erreur propres : `400` si l'identifiant n'est pas un nombre, `404` si l'artiste n'existe pas, `500` si l'API ne répond pas.

## Ma partie : le back-end

- `api/` : une fonction par route de l'API (`artists`, `relation`, `locations`, `dates`), avec décodage JSON vers des structures Go.
- `models/` : les structures de données.
- `handlers/` : les handlers HTTP. La page artiste assemble les données des 4 routes, et les erreurs sont gérées sans faire planter le serveur.
- `handlers/handlers_test.go` : test unitaire de la recherche d'artistes.

## Technologies

Go (bibliothèque standard : `net/http`, `html/template`, `encoding/json`), HTML, CSS.

## Lancer le projet

```bash
git clone https://github.com/acherif88/groupie-tracker.git
cd groupie-tracker
go run .
```

Puis ouvrir http://localhost:8080

## Lancer les tests

```bash
go test ./...
```

## Auteur

Abdou Cherif Hono Digo – [portfolio](https://acherif88.github.io)
