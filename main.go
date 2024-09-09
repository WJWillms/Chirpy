package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define the directory to serve files from
	staticDir := "." // Adjust this path to the directory containing your static files

	// Create a FileServer handler for serving static files
	fileServer := http.FileServer(http.Dir(staticDir))

	// Use the .Handle() method to register the FileServer for the root path
	mux.Handle("/", http.StripPrefix("/", fileServer))

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
