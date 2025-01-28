package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Create a new ServeMux
    mux := http.NewServeMux()

    // Handle the root path and serve the index.html file
    mux.Handle("/", http.FileServer(http.Dir("."))) // Serve files from the current directory

    // Create a new server with the ServeMux and set the Addr to ":8080"
    server := &http.Server{
        Addr:    ":8080", // Bind to localhost:8080
        Handler: mux,     // Set the mux as the handler
    }

    // Start the server
    fmt.Println("Starting server on http://localhost:8080.")
    if err := server.ListenAndServe(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
