package main

import "net/http"

// Middleware to increment the fileserverHits counter
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the fileserverHits counter
		cfg.fileserverHits++

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
