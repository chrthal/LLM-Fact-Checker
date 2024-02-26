package models

import "time"

type Job struct {
	ID            int       `json:"jobId"`
	Question      string    `json:"question"`
	PagesToCrawl  int       `json:"pagesToCrawl"`
	SearchEngines int       `json:"searchEngines"`
	Status        string    `json:"status"`
	StartDate     time.Time `json:"startDate"`
	LastUpdate    time.Time `json:"lastUpdate"`
}
