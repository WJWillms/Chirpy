package main

import (
	"fmt"
	"net/http"

	"github.com/WJWillms/chirpy/database"
)

var (
	dbPath = "./database/database.json"
	db     *database.DB
)

func main() {
	cfg := &apiConfig{}

	mux := http.NewServeMux()

	// Register the chirp handler
	mux.HandleFunc("/api/chirps", cfg.chirpHandler)

	// Register other handlers
	mux.HandleFunc("/api/healthz", readinessHandler)
	mux.HandleFunc("/admin/metrics", cfg.adminMetricsHandler)
	mux.HandleFunc("/api/reset", cfg.resetHandler)

	staticDir := "." // Adjust this path as needed
	fileServer := http.FileServer(http.Dir(staticDir))
	appHandler := cfg.middlewareMetricsInc(http.StripPrefix("/app/", fileServer))
	mux.Handle("/app/", appHandler)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting server on http://localhost:8080")
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
