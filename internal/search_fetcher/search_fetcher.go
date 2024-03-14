package search_fetcher

import (
	"log"

	"chrthal/llm-fact-checker/models"
)

func SearchFetcher(searchEngineData models.SearchEngineData, query string, results int) models.SearchEngineData {
	switch searchEngineData.SearchEngine.String() {
	case "Google":
		urls := googleSearch(query, results)
		log.Printf("Known search engine: %s\n Num Results: %d\n", searchEngineData.SearchEngine.String(), len(urls))
		searchEngineData.Urls = append(searchEngineData.Urls, urls...)
	case "Bing":
		urls := bingSearch(query, results)
		log.Printf("Known search engine: %s\n Num Results: %d\n", searchEngineData.SearchEngine.String(), len(urls))
		searchEngineData.Urls = append(searchEngineData.Urls, urls...)
	default:
		log.Printf("Unknown search engine: %s\n", searchEngineData.SearchEngine.String())
	}
	return searchEngineData
}
