package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DecodeJSON decodes a single JSON object from an HTTP request into a generic type
func DecodeJSON[T any](r *http.Request) (*T, error) {
	var result T
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&result); err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("empty request body")
		}
		return nil, err
	}
	return &result, nil
}

// DecodeJSONArray decodes an array of JSON objects from an HTTP request into a slice of generic type
func DecodeJSONArray[T any](r *http.Request) ([]T, error) {
	var results []T
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&results); err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("empty request body")
		}
		return nil, err
	}
	return results, nil
}