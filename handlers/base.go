package handlers

import (
	"net/http"
	"encoding/json"
	"log"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type APIError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewAPIError(e *APIError, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.Status)

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Println("[API ERROR]: The website encountered an unexpected error.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewAPIResponse(res *APIResponse, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("[API RESPONSE]: The website encountered an unexpected error.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
