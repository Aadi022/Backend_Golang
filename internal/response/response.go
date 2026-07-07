package response

import (
	"encoding/json"
	"net/http"
)

// Success writes a successful JSON response.
func Success(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data":    data,
	})
}

// Error writes an error JSON response.
func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"error":   message,
	})
}
