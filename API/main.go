package main

import (
	"fmt"
	"log"
	"net/http"

	"cpe241/handlers"
)

func main() {
	// Load configuration (could be moved to a config package)
	const postgrestURL = "https://postgrest.kaminjitt.com"
	const jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpX3VzZXIifQ.4TxmV2vnhZ5YTLw39wURDXQlzTHuAoaXHYhdTiqrNgY"

	// Initialize handlers with dependencies
	harvestHandler := handlers.NewHarvestLogHandler(postgrestURL, jwtToken)

	// Setup routes
	http.HandleFunc("/api/harvestlog", harvestHandler.HandleHarvestLog)

	// Start server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
