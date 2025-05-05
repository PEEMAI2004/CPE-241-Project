package handlers

import (
	"fmt"
	"net/http"
	"math"
	"cpe241/models"
	"cpe241/services"
	"cpe241/utils"
)

// HarvestLogHandler handles harvest log HTTP requests
type HarvestLogHandler struct {
	postgrestService *services.PostgRESTService
}

// NewHarvestLogHandler creates a new harvest log handler
func NewHarvestLogHandler(postgrestURL, jwtToken string) *HarvestLogHandler {
	return &HarvestLogHandler{
		postgrestService: services.NewPostgRESTService(postgrestURL, jwtToken),
	}
}

// PrintHarvestLog prints a harvest log to stdout for debugging
func (h *HarvestLogHandler) PrintHarvestLog(log *models.HarvestLog) {
	fmt.Println("Received Harvest Log:")
	// fmt.Printf("  Harvest ID: %s\n", log.HarvestID)
	fmt.Printf("  BeeHive ID: %s\n", log.BeeHiveID)
	fmt.Printf("  Harvest Date: %s\n", log.HarvestDate)
	fmt.Printf("  Production: %d %s\n", log.Production, log.Unit)
	fmt.Printf("  Notes: %s\n", log.ProductionNote)
}

func (h *HarvestLogHandler) TurnHarvest2Stock(log *models.HarvestLog, portion float64) {
	// Number of Item = Production / Portion 
	quantity := int(log.Production / portion)
	remainder := math.Mod(log.Production, portion)
	fmt.Println("Quantity:", quantity, "Remainder:", remainder)
	// Loop through the quantity and create HoneyStock objects
	for i := 0; i < quantity; i++ {
		honeyStock := h.CreateHoneyStock(log, portion)
		fmt.Println("Created HoneyStock:", honeyStock)
		// Post the honey stock to PostgREST
		// StatusCode, body, err := h.postgrestService.ForwardToPostgREST(honeyStock, nil, "/honeystock")
		StatusCode, body, err := h.postgrestService.ForwardToPostgREST(honeyStock, "/honeystock")
		fmt.Println("StatusCode:", StatusCode, "Body:", string(body))
		if err != nil {
			fmt.Println("Error posting HoneyStock:", err)
		}
	}
	// If there is a remainder, create a HoneyStock object for it
	if remainder > 0 {
		honeyStock := h.CreateHoneyStock(log, remainder)
		fmt.Println("Created HoneyStock with remainder:", honeyStock)
		// Post the honey stock to PostgREST
		StatusCode, body, err := h.postgrestService.ForwardToPostgREST(honeyStock, "/honeystock")
		fmt.Println("StatusCode:", StatusCode, "Body:", string(body))
		if err != nil {
			fmt.Println("Error posting HoneyStock with remainder:", err)
		}
	}
}

// CreateHoneyStock creates a HoneyStock object from a HarvestLog
func (h *HarvestLogHandler) CreateHoneyStock(log *models.HarvestLog, quantity float64) *models.HoneyStock {
	// Get the latest primary key from the harvest table
	latestPK, err := h.postgrestService.GetLatestPrimaryKey("/harvestlog", "harvest_id")
	if err != nil {
		fmt.Println("Error getting latest primary key:", err)
		return nil
	}
	HarvestID, ok := latestPK.(float64)
	if !ok {
		fmt.Println("Error converting latest primary key to int")
		return nil
	}

	return &models.HoneyStock{
		// StockID:   0, // This will be set by the database
		BeeHiveID: log.BeeHiveID,
		Quantity:  float64(quantity),
		Unit:      log.Unit,
		HarvestID: int(HarvestID),
		StockDate: log.HarvestDate,
	}
}

// HandleHarvestLog handles POST requests for harvest logs
func (h *HarvestLogHandler) HandleHarvestLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	logEntry, err := utils.DecodeJSON[models.HarvestLog](r)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	h.PrintHarvestLog(logEntry)

	// err = h.postgrestService.ForwardToPostgREST(logEntry, w, "/harvestlog")
	StatusCode, body, err := h.postgrestService.ForwardToPostgREST(logEntry, "/harvestlog")
	fmt.Println("StatusCode:", StatusCode, "Body:", string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	h.TurnHarvest2Stock(logEntry, 1)
}
