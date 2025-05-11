package main

import (
    "fmt"
    "log"
    "net/http"

    "cpe241/handlers"
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "https://app.kaminjitt.com") // Allow requests from your frontend
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
	// Load configuration (could be moved to a config package)
	const postgrestURL = "http://192.168.1.225:4000"
	const jwtToken = ""

	// Initialize handlers with dependencies
	harvestHandler := handlers.NewHarvestLogHandler(postgrestURL, jwtToken)
	orderHandler := handlers.NewOrderHandler(postgrestURL, jwtToken)

    // Setup routes
    mux := http.NewServeMux()
    mux.HandleFunc("/api/harvestlog", harvestHandler.HandleHarvestLog)
    mux.HandleFunc("/api/order", orderHandler.HandleCreateOrder)
    mux.HandleFunc("/api/check-stock", orderHandler.HandleCheckStock)

    // Wrap routes with CORS middleware
    fmt.Println("Server is running on http://localhost:6000")
    log.Fatal(http.ListenAndServe(":6000", enableCORS(mux)))
}