// veltern/internal/api/handler.go
package api

import (
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.APIInfo(w, r)
	case "/rectangle/area":
		h.CalculateRectangleArea(w, r)
	default:
		respondWithError(w, "Not Found", http.StatusNotFound)
	}
}
