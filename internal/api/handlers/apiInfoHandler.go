// valtar/internal/api/handlers/apiInfoHandler.go
package handlers

import (
	"net/http"
	"valtar/internal/api/utils"
)

func APIInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{
		"message":   "Welcome to the Valtar API",
		"endpoints": "/rectangle/area - Calculate the area of a rectangle",
	}
	utils.RespondWithJSON(w, http.StatusOK, info)
}
