package models

type JobRequest struct {
	Question      string         `json:"question"`
	PagesToCrawl  int            `json:"pagesToCrawl"`
	SearchEngines []SearchEngine `json:"searchEngines"`
}
