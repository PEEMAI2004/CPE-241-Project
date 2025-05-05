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
        w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // Allow requests from your frontend
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
	const postgrestURL = "https://postgrest.kaminjitt.com"
	const jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpX3VzZXIifQ.4TxmV2vnhZ5YTLw39wURDXQlzTHuAoaXHYhdTiqrNgY"

	// Initialize handlers with dependencies
	harvestHandler := handlers.NewHarvestLogHandler(postgrestURL, jwtToken)
	orderHandler := handlers.NewOrderHandler(postgrestURL, jwtToken)

    // Setup routes
    mux := http.NewServeMux()
    mux.HandleFunc("/api/harvestlog", harvestHandler.HandleHarvestLog)
    mux.HandleFunc("/api/order", orderHandler.HandleCreateOrder)
    mux.HandleFunc("/api/check-stock", orderHandler.HandleCheckStock)

    // Wrap routes with CORS middleware
    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", enableCORS(mux)))
}