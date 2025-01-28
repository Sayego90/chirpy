package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080", // Bind to localhost:8080
		Handler: mux,     // Use the new ServeMux
	}

	fmt.Println("Starting server on http://localhost:8080")

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
