// valtar/internal/api/scrapTwitter/scrapTwitter.go
package scrapTwitter

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// ScrapeTwitterMedia scrapes the media URLs from a Twitter page
func ScrapeTwitterMedia(tweetURL string) ([]string, error) {
	log.Printf("Scraping Twitter URL: %s\n", tweetURL)

	resp, err := http.Get(tweetURL)
	if err != nil {
		log.Printf("Error fetching Twitter page: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML: %v\n", err)
		return nil, err
	}

	var mediaURLs []string
	// Example selector - adjust based on actual page structure
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if imgSrc, exists := s.Attr("src"); exists {
			log.Printf("Found media URL: %s\n", imgSrc)
			mediaURLs = append(mediaURLs, imgSrc)
		}
	})

	return mediaURLs, nil
}
