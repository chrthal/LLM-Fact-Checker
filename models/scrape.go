package models

type Scrape struct {
	Id      int    `json:"scrapeId"`
	Url     string `json:"Url"`
	Content string `json:"Content"`
	JobId   int    `json:"jobId"`
}
