package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/alphakta/URLShortenerProject/domain/url"
)

func AddURLHandler(service *url.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST requests are accepted", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		longURL := string(body)
		shortURL, err := service.Create(longURL)

		if err != nil {
			http.Error(w, "Error adding URL", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(shortURL))
	}
}

func RedirectHandler(service *url.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extrait l'identifiant (path) de l'URL courte à partir de l'URL de la requête
		path := r.URL.Path[1:] // Supprime le slash de début

		// Incrémente les clics pour l'URL courte
		err := service.IncrementClicks(path)
		if err != nil {
			// Gérer l'erreur d'incrémentation ici si nécessaire
			log.Printf("Erreur lors de l'incrémentation des clics pour %s: %v", path, err)
		}

		// Utilisation du service pour trouver l'URL longue correspondante
		longURL, err := service.FindByShortURL(path)
		if err != nil {
			if err == sql.ErrNoRows {
				// Si aucune URL correspondante n'est trouvée, renvoie une erreur 404.
				http.NotFound(w, r)
				return
			}
			// Gère d'autres erreurs potentielles
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		// Redirige vers l'URL longue
		http.Redirect(w, r, longURL, http.StatusFound) // 302 Found
	}
}

func StatsHandler(service *url.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total, err := service.GetTotalLinks()
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du nombre total de liens", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]int{"total_links": total})
	}
}

func ClicksHandler(service *url.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/stats/"):]
		clicks, err := service.GetClicksByID(id)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des clics pour l'ID donné", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]int{"clicks": clicks})
	}
}
