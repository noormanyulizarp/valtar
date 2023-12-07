// veltern/internal/api/apiInfo.go
package api

import (
	"net/http"
)

func (h *Handler) APIInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{
		"message":   "Welcome to the Veltern API",
		"endpoints": "/rectangle/area - Calculate the area of a rectangle",
	}
	respondWithJSON(w, http.StatusOK, info)
}
