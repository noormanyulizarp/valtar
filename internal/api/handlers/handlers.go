// valtar/internal/api/handlers/handlers.go
package handlers

import (
	"net/http"
	"valtar/internal/api/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		APIInfo(w, r)
	case "/rectangle/area":
		CalculateRectangleArea(w, r)
	case "/twitter/media":
		HandleTwitterMediaRequest(w, r)
	case "/scrapTwitter":
		HandleScrapTwitterRequest(w, r)
	case "/goScrapTwitter":
		HandleGoScrapTwitterRequest(w, r)
	default:
		utils.RespondWithError(w, "Not Found", http.StatusNotFound)
	}
}
