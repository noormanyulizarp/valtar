// valtar/internal/api/utils/utils.go
package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, message string, statusCode int) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
