// valtar/internal/api/handlers/scrapTwitterHandler.go
package handlers

import (
	"log"
	"net/http"
	"valtar/internal/api/scrapTwitter"
	"valtar/internal/api/utils"
)

func HandleScrapTwitterRequest(w http.ResponseWriter, r *http.Request) {
	tweetURL := r.URL.Query().Get("tweetUrl")
	if tweetURL == "" {
		log.Println("Tweet URL is required")
		utils.RespondWithError(w, "Tweet URL is required", http.StatusBadRequest)
		return
	}

	mediaURLs, err := scrapTwitter.ScrapeTwitterMedia(tweetURL)
	if err != nil {
		log.Printf("Error scraping Twitter media: %v\n", err)
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(mediaURLs) == 0 {
		log.Println("No media URLs found in the tweet")
		utils.RespondWithJSON(w, http.StatusOK, []string{})
		return
	}

	log.Println("Media content response sent successfully")
	utils.RespondWithJSON(w, http.StatusOK, mediaURLs)
}
