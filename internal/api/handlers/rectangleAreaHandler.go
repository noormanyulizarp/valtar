// valtar/internal/api/handlers/rectangleAreaHandler.go
package handlers

import (
	"net/http"
	"strconv"
	"valtar/internal/api/utils"
	"valtar/internal/calculator"
)

func CalculateRectangleArea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithError(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	width, err := strconv.ParseFloat(r.URL.Query().Get("width"), 64)
	if err != nil {
		utils.RespondWithError(w, "Invalid width", http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseFloat(r.URL.Query().Get("height"), 64)
	if err != nil {
		utils.RespondWithError(w, "Invalid height", http.StatusBadRequest)
		return
	}

	area := calculator.CalculateRectangleArea(width, height)
	utils.RespondWithJSON(w, http.StatusOK, map[string]float64{"area": area})
}
