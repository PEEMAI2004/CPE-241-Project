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

// New version: returns status code and response body
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

// // ForwardToPostgREST forwards an object to PostgREST and writes the response to the http.ResponseWriter
// func (s *PostgRESTService) ForwardToPostgREST(obj interface{}, w http.ResponseWriter, tablePath string) error {
// 	jsonData, err := json.Marshal(obj)
// 	if err != nil {
// 		return fmt.Errorf("failed to encode JSON: %w", err)
// 	}

// 	url := s.BaseURL + tablePath
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return fmt.Errorf("failed to create request to PostgREST: %w", err)
// 	}
	
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+s.JWTToken)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to contact PostgREST: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read response body: %w", err)
// 	}
	
// 	w.WriteHeader(resp.StatusCode)
// 	w.Write(body)
// 	return nil
// }



