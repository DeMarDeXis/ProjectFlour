package httplib

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(w http.ResponseWriter, logg *slog.Logger, statusCode int, message string) {
	logg.Error("Response status: " + strconv.Itoa(statusCode) + " message: " + message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := errorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}

func NewStatusResponse(w http.ResponseWriter, logg *slog.Logger, statusCode int, status string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := statusResponse{Status: status}
	json.NewEncoder(w).Encode(response)
}
