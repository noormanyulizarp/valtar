// valtar/internal/api/scrapTwitter/scrapTwitter.go
package scrapTwitter

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// ScrapeTwitterMedia scrapes the media URLs from a Twitter page
func ScrapeTwitterMedia(tweetURL string) ([]string, error) {
	resp, err := http.Get(tweetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var mediaURLs []string
	// This selector should be adjusted according to the specific structure of the Twitter page
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if imgSrc, exists := s.Attr("src"); exists {
			mediaURLs = append(mediaURLs, imgSrc)
		}
	})

	return mediaURLs, nil
}
