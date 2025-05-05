package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PostgRESTService handles communication with PostgREST
type PostgRESTService struct {
	BaseURL  string
	JWTToken string
}

// NewPostgRESTService creates a new PostgREST service
func NewPostgRESTService(baseURL, jwtToken string) *PostgRESTService {
	return &PostgRESTService{
		BaseURL:  baseURL,
		JWTToken: jwtToken,
	}
}

// ForwardToPostgREST forwards a POST request to PostgREST
func (s *PostgRESTService) ForwardToPostgREST(obj interface{}, tablePath string) (int, []byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	url := s.BaseURL + tablePath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request to PostgREST: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.JWTToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to contact PostgREST: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return resp.StatusCode, body, nil
}

// UpdatePostgREST sends a PATCH request to update a record in PostgREST
func (s *PostgRESTService) UpdatePostgREST(obj interface{}, tablePath string, filterColumn string, filterValue interface{}) (int, []byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	url := fmt.Sprintf("%s%s?%s=eq.%v", s.BaseURL, tablePath, filterColumn, filterValue)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request to PostgREST: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.JWTToken)
	req.Header.Set("Prefer", "return=representation")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to contact PostgREST: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return resp.StatusCode, body, nil
}

// CheckRecordExists checks if a record exists in the specified table
func (s *PostgRESTService) CheckRecordExists(tablePath, filterColumn string, filterValue interface{}) (bool, error) {
	url := fmt.Sprintf("%s%s?%s=eq.%v", s.BaseURL, tablePath, filterColumn, filterValue)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request to PostgREST: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.JWTToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to contact PostgREST: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode response body: %w", err)
	}

	return len(result) > 0, nil
}

// GetStockSoldStatus checks if a stock is already sold
func (s *PostgRESTService) GetStockSoldStatus(stockID int) (bool, error) {
	url := fmt.Sprintf("%s/honeystock?stock_id=eq.%d", s.BaseURL, stockID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request to PostgREST: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.JWTToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to contact PostgREST: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Stock check response for ID %d: Status=%d, Body=%s\n", stockID, resp.StatusCode, string(body))

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, fmt.Errorf("failed to decode response body: %w", err)
	}

	if len(result) == 0 {
		return false, fmt.Errorf("stock ID %d not found", stockID)
	}

	// Check if is_sold field exists and handle it appropriately
	isSoldVal, exists := result[0]["is_sold"]
	if !exists {
		// If the field doesn't exist, treat it as not sold
		fmt.Printf("Warning: is_sold field not found for stock ID %d\n", stockID)
		return false, nil
	}

	// Handle null value
	if isSoldVal == nil {
		fmt.Printf("Warning: is_sold is null for stock ID %d\n", stockID)
		return false, nil
	}

	// Try to convert to boolean
	isSold, ok := isSoldVal.(bool)
	if !ok {
		// Try to convert from string if needed
		if strVal, ok := isSoldVal.(string); ok {
			return strVal == "true", nil
		}
		fmt.Printf("Warning: invalid is_sold value type for stock ID %d: %T\n", stockID, isSoldVal)
		return false, nil
	}

	return isSold, nil
}

// GetLatestPrimaryKey gets the latest primary key from a table
func (s *PostgRESTService) GetLatestPrimaryKey(tablePath, primaryKeyColumn string) (interface{}, error) {
	url := fmt.Sprintf("%s%s?select=%s&order=%s.desc&limit=1", s.BaseURL, tablePath, primaryKeyColumn, primaryKeyColumn)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to PostgREST: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.JWTToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to contact PostgREST: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no records found in table %s", tablePath)
	}

	return result[0][primaryKeyColumn], nil
}