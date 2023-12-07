// valtar/internal/api/handlers/apiInfoHandler.go
package handlers

import (
	"net/http"
	"valtar/internal/api/utils"
)

func APIInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{
		"message":   "example for /rectangle/area API hit -> /rectangle/area?width=10&height=5",
		"endpoints": "/rectangle/area?width=[value]&heights=[value]",
	}
	utils.RespondWithJSON(w, http.StatusOK, info)
}
