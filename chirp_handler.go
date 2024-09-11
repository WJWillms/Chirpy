package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func (cfg *apiConfig) chirpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var requestBody struct {
			Body string `json:"body"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
			return
		}

		if len(requestBody.Body) > 140 {
			http.Error(w, `{"error": "Chirp is too long"}`, http.StatusBadRequest)
			return
		}

		profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
		cleanedBody := requestBody.Body

		for _, word := range profaneWords {
			pattern := "\\b" + word + "\\b"
			re := regexp.MustCompile("(?i)" + pattern)
			cleanedBody = re.ReplaceAllString(cleanedBody, "****")
		}

		chirp, err := cfg.db.CreateChirp(cleanedBody)
		if err != nil {
			http.Error(w, `{"error": "Failed to save chirp"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(chirp)

	case http.MethodGet:
		chirps, err := cfg.db.GetChirps()
		if err != nil {
			http.Error(w, `{"error": "Failed to retrieve chirps"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(chirps)

	default:
		http.Error(w, `{"error": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
	}
}
