package scraper

import (
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeGoogleSearchResults(query string, numResults int) ([]string, error) {
	results := make([]string, 0, numResults)

	// Create a new collector
	c := colly.NewCollector()

	// OnHTML is called when an HTML element matching the provided selector is found.
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Fetch the href attribute from the a element
		href := e.Attr("href")
		if strings.HasPrefix(href, "/url?q=") {
			// Extract the URL from Google's URL format
			u, err := url.Parse(href)
			if err == nil {
				parsedURL, err := url.QueryUnescape(u.Query().Get("q"))
				if err == nil && parsedURL != "" {
					results = append(results, parsedURL)
				}
			}
		}
		/*if len(results) >= numResults {
			// Stop scraping once we reach the desired number of results
			c.Cancel()
		}*/
	})

	// Before making a request, set the user-agent header
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	// Start scraping
	err := c.Visit("https://www.google.com/search?q=" + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}

	return results, nil
}
