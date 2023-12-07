// valtar/internal/twitter/twitter.go
package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"valtar/config"
)

type TweetMedia struct {
	MediaURL string `json:"media_url_https"`
}

type TweetEntities struct {
	Media []TweetMedia `json:"media"`
}

type Tweet struct {
	ExtendedEntities *TweetEntities `json:"extended_entities"`
}

func FetchMediaContent(tweetURL string) ([]string, error) {
	cfg, err := config.LoadConfig("config.json") // Ensure this path is correct
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
		return nil, err
	}

	oauthConfig := &clientcredentials.Config{
		ClientID:     cfg.TwitterApiKey,
		ClientSecret: cfg.TwitterApiSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := oauthConfig.Client(context.Background())

	parts := strings.Split(tweetURL, "/")
	if len(parts) < 6 {
		log.Println("Invalid tweet URL format")
		return nil, fmt.Errorf("invalid tweet URL format")
	}
	tweetID := parts[5]
	tweetID = strings.Split(tweetID, "?")[0] // Removing any query parameters

	endpoint := fmt.Sprintf("https://api.twitter.com/1.1/statuses/show.json?id=%s&tweet_mode=extended", tweetID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error making request to Twitter API: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	// Log the raw Twitter API response
	log.Println("Raw Twitter API Response:", string(body))

	var tweet Tweet
	if err := json.Unmarshal(body, &tweet); err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return nil, err
	}

	if tweet.ExtendedEntities == nil || len(tweet.ExtendedEntities.Media) == 0 {
		log.Println("No media found in the tweet")
		return []string{}, nil
	}

	var mediaURLs []string
	for _, media := range tweet.ExtendedEntities.Media {
		mediaURLs = append(mediaURLs, media.MediaURL)
	}

	log.Println("Media URLs fetched successfully")
	return mediaURLs, nil
}
