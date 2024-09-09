package main

import (
	"fmt"
	"net/http"
)

type server struct {
	Addr    string
	Handler http.Handler
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	srv := &server{
		Addr:    ":8080",
		Handler: mux,
	}

	httpServer := &http.Server{
		Addr:    srv.Addr,
		Handler: srv.Handler,
	}

	// Start the server
	fmt.Println("Starting server on http://localhost:8080")
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
