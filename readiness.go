package main

import (
	"net/http"
)

// ReadinessHandler handles GET /api/healthz
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("OK"))
}
