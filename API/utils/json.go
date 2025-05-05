package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DecodeJSON decodes JSON from an HTTP request into a generic type
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