package main

import (
	"fmt"
	"net/http"
)

// Handler to reset the hits counter
func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	cfg.fileserverHits = 0
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Hits counter reset"))
}

// Handler for viewing Metrics as Admin
func (cfg *apiConfig) adminMetricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := fmt.Sprintf(`
        <html>
        <body>
            <h1>Welcome, Chirpy Admin</h1>
            <p>Chirpy has been visited %d times!</p>
        </body>
        </html>`, cfg.fileserverHits)

	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte(html))
}
