// valtar/internal/api/handlers/goScrapTwitterHandler.go
package handlers

import (
	"net/http"
	"valtar/internal/api/scrapTwitter"
	"valtar/internal/api/utils"
)

// HandleGoScrapTwitterRequest handles requests to the GoScrapTwitter endpoint
func HandleGoScrapTwitterRequest(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweetId")
	if tweetID == "" {
		utils.RespondWithError(w, "Tweet ID is required", http.StatusBadRequest)
		return
	}

	mediaURLs, err := scrapTwitter.GoScrapeTwitterMedia(tweetID)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, mediaURLs)
}
