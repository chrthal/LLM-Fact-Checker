package models

type CrawledData struct {
	WebScrape  []string `json:"webScrape"`
	LLMScrape  string   `json:"llmScrape"`
	Similarity float64  `json:"similarity"`
}
