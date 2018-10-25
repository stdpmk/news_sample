package app

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error codes
const (
	Ok             string = ""
	ErrInvalidJson string = "invalid_json"
	ErrInternal    string = "internal"
	ErrNotFound    string = "not_found"
)

func responseJson(w http.ResponseWriter, httpCode int, obj ApiResponse) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(obj)
}
