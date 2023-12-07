// veltern/internal/api/calculateArea.go
package api

import (
	"net/http"
	"strconv"
	"valtar/internal/calculator"
)

func (h *Handler) CalculateRectangleArea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	width, err := strconv.ParseFloat(r.URL.Query().Get("width"), 64)
	if err != nil {
		respondWithError(w, "Invalid width", http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseFloat(r.URL.Query().Get("height"), 64)
	if err != nil {
		respondWithError(w, "Invalid height", http.StatusBadRequest)
		return
	}

	area := calculator.CalculateRectangleArea(width, height)
	respondWithJSON(w, http.StatusOK, map[string]float64{"area": area})
}
