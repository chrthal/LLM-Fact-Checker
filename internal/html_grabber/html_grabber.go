package html_grabber

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func FetchAndPrepareBody(url string) (string, error) {
	c := colly.NewCollector()

	p := ""

	// Set up callbacks to handle scraping events
	c.OnHTML("body", func(e *colly.HTMLElement) {
		p = strings.ReplaceAll(e.ChildText("p"), "\n", "")
		p = strings.ReplaceAll(p, "\t", "")

		// Remove special characters using regular expressions
		re := regexp.MustCompile(`[^a-zA-Z0-9\s.]+`)
		p = re.ReplaceAllString(p, "")
	})

	// Visit the URL and start scraping
	err := c.Visit(url)

	return p, err
}
