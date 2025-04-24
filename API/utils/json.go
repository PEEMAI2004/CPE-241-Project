package utils

import (
	"encoding/json"
	"net/http"
)

// DecodeJSON decodes JSON from an HTTP request into a given type
func DecodeJSON[T any](r *http.Request) (*T, error) {
	var obj T
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
