package main

import (
	"fmt"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the fileserverHits counter
		cfg.fileserverHits++

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write the response body
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("OK"))
}

// Handler to write the number of hits
func (cfg *apiConfig) hitsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Write the hits count to the response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hits: %d", cfg.fileserverHits)
}

// Handler to reset the hits counter
func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Reset the hits counter
	cfg.fileserverHits = 0

	// Respond to indicate the counter has been reset
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Hits counter reset"))
}

func main() {
	// Initialize apiConfig
	cfg := &apiConfig{}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register the readiness handler for the /healthz path
	mux.HandleFunc("/api/healthz", readinessHandler)

	// Define the directory to serve files from
	staticDir := "." // Adjust this path to the directory containing your static files

	// Create a FileServer handler for serving static files
	fileServer := http.FileServer(http.Dir(staticDir))

	// Create the middleware handler for the file server
	appHandler := cfg.middlewareMetricsInc(http.StripPrefix("/app/", fileServer))

	// Register the file server handler for the /app/ path
	mux.Handle("/app/", appHandler)

	// Register the hits handler for the /metrics path
	mux.HandleFunc("/api/metrics", cfg.hitsHandler)

	// Register the reset handler for the /reset path
	mux.HandleFunc("/api/reset", cfg.resetHandler)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server
	fmt.Println("Starting server on http://localhost:8080")
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
