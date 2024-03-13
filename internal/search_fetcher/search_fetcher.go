package search_fetcher

import (
	"fmt"

	"chrthal/llm-fact-checker/models"
)

func SearchFetcher(searchEngineData models.SearchEngineData, query string, results int) models.SearchEngineData {
	switch searchEngineData.SearchEngine.String() {
	case "Google":
		urls := googleSearch(query, results)
		fmt.Printf("Known search engine: %s\n Result 0: %s", searchEngineData.SearchEngine.String(), urls[0])
		searchEngineData.Urls = append(searchEngineData.Urls, urls...)
	case "Bing":
		urls := bingSearch(query, results)
		fmt.Printf("Known search engine: %s\n Result 0: %s", searchEngineData.SearchEngine.String(), urls[0])
		searchEngineData.Urls = append(searchEngineData.Urls, urls...)
	default:
		fmt.Printf("Unknown search engine: %s\n", searchEngineData.SearchEngine.String())
	}
	return searchEngineData
}
