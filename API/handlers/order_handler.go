package handlers

import (
	"cpe241/models"
	"cpe241/services"
	"cpe241/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// OrderHandler handles order HTTP requests
type OrderHandler struct {
	postgrestService *services.PostgRESTService
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(postgrestURL, jwtToken string) *OrderHandler {
	return &OrderHandler{
		postgrestService: services.NewPostgRESTService(postgrestURL, jwtToken),
	}
}

// HandleCheckStock is a debug endpoint to check if a stock exists and its status
func (h *OrderHandler) HandleCheckStock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	stockID := r.URL.Query().Get("id")
	if stockID == "" {
		http.Error(w, "Stock ID is required", http.StatusBadRequest)
		return
	}

	var stockIDInt int
	fmt.Sscanf(stockID, "%d", &stockIDInt)

	// Check directly with the database
	isSold, err := h.postgrestService.GetStockSoldStatus(stockIDInt)
	if err != nil {
		http.Error(w, "Error checking stock: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status":   "success",
		"stock_id": stockIDInt,
		"is_sold":  isSold,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


// HandleCreateOrder processes a new order
func (h *OrderHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the order request
	orderRequest, err := utils.DecodeJSON[models.OrderRequest](r)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	// Log the received order request
	fmt.Printf("Received order request: UserID=%d, CustomerID=%d, Items=%d\n", 
		orderRequest.UserID, orderRequest.CustomerID, len(orderRequest.Items))

	// Validate the request
	if err := h.validateOrderRequest(orderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the order
	orderID, err := h.processOrder(orderRequest)
	if err != nil {
		fmt.Printf("Failed to process order: %v\n", err)
		http.Error(w, "Failed to process order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	fmt.Printf("Order processed successfully: OrderID=%d\n", orderID)

	// Return success response
	response := map[string]interface{}{
		"status":   "success",
		"order_id": orderID,
		"message":  "Order created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// validateOrderRequest checks if the order request is valid
func (h *OrderHandler) validateOrderRequest(req *models.OrderRequest) error {
	// Check if user exists
	userExists, err := h.postgrestService.CheckRecordExists("/webuser", "user_id", req.UserID)
	if err != nil {
		return fmt.Errorf("error checking user: %w", err)
	}
	if !userExists {
		return fmt.Errorf("user ID %d does not exist", req.UserID)
	}

	// Check if customer exists
	customerExists, err := h.postgrestService.CheckRecordExists("/customer", "customer_id", req.CustomerID)
	if err != nil {
		return fmt.Errorf("error checking customer: %w", err)
	}
	if !customerExists {
		return fmt.Errorf("customer ID %d does not exist", req.CustomerID)
	}

	// Check if there are any items
	if len(req.Items) == 0 {
		return fmt.Errorf("order must contain at least one item")
	}

	// Check if stocks exist and are not sold
	for _, item := range req.Items {
		isSold, err := h.postgrestService.GetStockSoldStatus(item.StockID)
		if err != nil {
			return fmt.Errorf("error checking stock %d: %w", item.StockID, err)
		}
		if isSold {
			return fmt.Errorf("stock ID %d is already sold", item.StockID)
		}
	}

	return nil
}

// processOrder handles the order processing logic
func (h *OrderHandler) processOrder(req *models.OrderRequest) (int, error) {
	fmt.Printf("Processing order: UserID=%d, CustomerID=%d, Items=%d\n", 
		req.UserID, req.CustomerID, len(req.Items))
		
	// 1. Create an entry in the orderlist table
	orderList := &models.OrderList{
		CustomerID: req.CustomerID,
		UserID:     req.UserID,
		OrderDate:  time.Now().Format("2006-01-02"),
	}

	statusCode, body, err := h.postgrestService.ForwardToPostgREST(orderList, "/orderlist")
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}
	if statusCode != http.StatusCreated && statusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to create order: unexpected status code %d: %s", statusCode, string(body))
	}

	// Try to get order ID from response first
	var createdOrder []map[string]interface{}
	err = json.Unmarshal(body, &createdOrder)
	
	var orderID int
	if err == nil && len(createdOrder) > 0 && createdOrder[0]["order_id"] != nil {
		// If we got the order ID directly from the creation response
		orderIDFloat, ok := createdOrder[0]["order_id"].(float64)
		if ok {
			orderID = int(orderIDFloat)
			fmt.Printf("Extracted order ID from response: %d\n", orderID)
		}
	}
	
	// If we couldn't get it from the response, get the latest order ID
	if orderID == 0 {
		fmt.Println("Fetching latest order ID from database")
		latestOrderID, err := h.postgrestService.GetLatestPrimaryKey("/orderlist", "order_id")
		if err != nil {
			return 0, fmt.Errorf("failed to get order ID: %w", err)
		}

		orderIDFloat, ok := latestOrderID.(float64)
		if !ok {
			return 0, fmt.Errorf("invalid order ID type: %T", latestOrderID)
		}
		
		orderID = int(orderIDFloat)
		fmt.Printf("Got latest order ID: %d\n", orderID)
	}

	// 2. Create entries in the orderitem table for each item
	for _, item := range req.Items {
		// Update the honeystock record to mark it as sold
		updateStock := map[string]interface{}{
			"is_sold": true,
		}
		statusCode, respBody, err := h.postgrestService.UpdatePostgREST(updateStock, "/honeystock", "stock_id", item.StockID)
		if err != nil {
			return 0, fmt.Errorf("failed to update stock %d: %w", item.StockID, err)
		}
		if statusCode != http.StatusOK && statusCode != http.StatusCreated {
			return 0, fmt.Errorf("failed to update stock %d: unexpected status code %d: %s", item.StockID, statusCode, string(respBody))
		}

		// Create order item record
		orderItem := &models.OrderItemDB{
			OrderID: orderID,
			StockID: item.StockID,
			Price:   item.Price,
		}

		statusCode, respBody, err = h.postgrestService.ForwardToPostgREST(orderItem, "/orderitem")
		if err != nil {
			return 0, fmt.Errorf("failed to create order item for stock %d: %w", item.StockID, err)
		}
		if statusCode != http.StatusCreated && statusCode != http.StatusOK {
			return 0, fmt.Errorf("failed to create order item for stock %d: unexpected status code %d: %s", item.StockID, statusCode, string(respBody))
		}
	}

	return orderID, nil
}