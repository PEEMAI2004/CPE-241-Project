package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HarvestLog struct {
	HarvestID      int `json:"harvest_id"`
	BeeHiveID      int `json:"beehive_id"`
	HarvestDate    string `json:"harvestdate"`
	Production     int    `json:"production"`
	Unit           string `json:"unit"`
	ProductionNote string `json:"production_note"`
}

// Replace this with your actual JWT token
const jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpX3VzZXIifQ.4TxmV2vnhZ5YTLw39wURDXQlzTHuAoaXHYhdTiqrNgY"

// PostgREST endpoint for the HarvestLog table
const postgrestURL = "https://postgrest.kaminjitt.com"

// func handleHarvestLog(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var logEntry HarvestLog
// 	// Print to console
// 	fmt.Println("Received Harvest Log:")
// 	fmt.Printf("  Harvest ID: %s\n", logEntry.HarvestID)
// 	fmt.Printf("  BeeHive ID: %s\n", logEntry.BeeHiveID)
// 	fmt.Printf("  Harvest Date: %s\n", logEntry.HarvestDate)
// 	fmt.Printf("  Production: %d %s\n", logEntry.Production, logEntry.Unit)
// 	fmt.Printf("  Notes: %s\n", logEntry.ProductionNote)
// 	// Decode the JSON body
// 	err := json.NewDecoder(r.Body).Decode(&logEntry)
// 	fmt.Println(err)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
// 		return
// 	}

// 	// Print to console
// 	fmt.Println("Received Harvest Log:")
// 	fmt.Printf("  Harvest ID: %s\n", logEntry.HarvestID)
// 	fmt.Printf("  BeeHive ID: %s\n", logEntry.BeeHiveID)
// 	fmt.Printf("  Harvest Date: %s\n", logEntry.HarvestDate)
// 	fmt.Printf("  Production: %d %s\n", logEntry.Production, logEntry.Unit)
// 	fmt.Printf("  Notes: %s\n", logEntry.ProductionNote)

// 	// Convert to JSON
// 	jsonData, err := json.Marshal(logEntry)
// 	if err != nil {
// 		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Forward to PostgREST
// 	req, err := http.NewRequest("POST", postgrestURL, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		http.Error(w, "Failed to create request to PostgREST", http.StatusInternalServerError)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+jwtToken)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		http.Error(w, "Failed to contact PostgREST", http.StatusBadGateway)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Relay the response back to the client
// 	body, _ := io.ReadAll(resp.Body)
// 	w.WriteHeader(resp.StatusCode)
// 	w.Write(body)
// }

///////////////

func decodeJSON[T any](r *http.Request) (*T, error) {
	var obj T
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func forwardToPostgREST[T any](obj *T, w http.ResponseWriter, tableURL string) error {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("Failed to encode JSON")
	}

	req, err := http.NewRequest("POST", tableURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Failed to create request to PostgREST")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
	return nil
}

func printHarvestLog(log *HarvestLog) {
	fmt.Println("Received Harvest Log:")
	fmt.Printf("  Harvest ID: %s\n", log.HarvestID)
	fmt.Printf("  BeeHive ID: %s\n", log.BeeHiveID)
	fmt.Printf("  Harvest Date: %s\n", log.HarvestDate)
	fmt.Printf("  Production: %d %s\n", log.Production, log.Unit)
	fmt.Printf("  Notes: %s\n", log.ProductionNote)
}

func handleHarvestLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	logEntry, err := decodeJSON[HarvestLog](r)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	printHarvestLog(logEntry)

	err = forwardToPostgREST(logEntry, w, postgrestURL+"/harvestlog")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
}

func main() {
	http.HandleFunc("/api/harvestlog", handleHarvestLog)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
if err != nil {
		return fmt.Errorf("Failed to contact PostgREST")
	}
	