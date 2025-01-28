package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the hit counter safely
		cfg.fileserverHits.Add(1)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	// Fetch the current count
	hits := cfg.fileserverHits.Load()

	// Respond with the hits count
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hits: %d", hits)
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	// Reset the counter safely
	cfg.fileserverHits.Store(0)

	// Respond with success message
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hits reset to 0")
}

func main() {
	// Create API config instance
	apiCfg := apiConfig{}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Readiness check endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	// Fileserver path
	fs := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(fs)) // Wrap fileserver with middleware

	// Register /metrics endpoint
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)

	// Register /reset endpoint
	mux.HandleFunc("/reset", apiCfg.handlerReset)

	// Create the HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server
	fmt.Println("Starting server on http://localhost:8080")
	server.ListenAndServe()
}
