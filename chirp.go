package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

// Handler for validating Chirp length
func (cfg *apiConfig) validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	if len(requestBody.Body) > 140 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Chirp is too long"})
		return
	}

	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	cleanedBody := requestBody.Body

	for _, word := range profaneWords {
		pattern := "\\b" + word + "\\b"
		re := regexp.MustCompile("(?i)" + pattern)
		cleanedBody = re.ReplaceAllString(cleanedBody, "****")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"cleaned_body": cleanedBody})
}
