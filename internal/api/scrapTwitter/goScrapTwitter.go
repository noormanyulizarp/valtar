// valtar/internal/api/scrapTwitter/goScrapTwitter.go
package scrapTwitter

import (
	"github.com/n0madic/twitter-scraper"
	"log"
)

// GoScrapeTwitterMedia uses twitter-scraper to scrape media URLs from a tweet
func GoScrapeTwitterMedia(tweetID string) ([]string, error) {
	log.Printf("Scraping Twitter URL: %s\n", tweetID)

	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet(tweetID)
	if err != nil {
		log.Printf("Error fetching Twitter page: %v\n", err)
		return nil, err
	}

	var mediaURLs []string
	for _, photo := range tweet.Photos {
		mediaURLs = append(mediaURLs, photo.URL) // Access the URL field from the Photo struct
	}

	return mediaURLs, nil
}
