package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Create a new ServeMux
    mux := http.NewServeMux()

    // Handle the root route to serve index.html
    mux.Handle("/", http.FileServer(http.Dir(".")))

    // Handle the /assets route to serve the logo.png file
    mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

    // Create the server with the ServeMux as the handler
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    fmt.Println("Starting server on http://localhost:8080.")
    // Start the server
    if err := server.ListenAndServe(); err != nil {
        fmt.Println("Error starting the server:", err)
    }
}
