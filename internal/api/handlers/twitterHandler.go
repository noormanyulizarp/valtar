// valtar/internal/api/handlers/twitterHandler.go
package handlers

import (
	"log"
	"net/http"
	"valtar/internal/api/twitter"
	"valtar/internal/api/utils"
)

func HandleTwitterMediaRequest(w http.ResponseWriter, r *http.Request) {
	tweetURL := r.URL.Query().Get("tweetUrl")
	if tweetURL == "" {
		log.Println("Tweet URL is required")
		utils.RespondWithError(w, "Tweet URL is required", http.StatusBadRequest)
		return
	}

	mediaURLs, err := twitter.FetchMediaContent(tweetURL)
	if err != nil {
		log.Printf("Error fetching media content: %v\n", err)
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(mediaURLs) == 0 {
		log.Println("No media URLs found")
		utils.RespondWithJSON(w, http.StatusOK, []string{})
		return
	}

	log.Println("Media content response sent successfully")
	utils.RespondWithJSON(w, http.StatusOK, mediaURLs)
}
