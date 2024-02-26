package search_api

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func GoogleSearch(query string, numResults int) []string {
	var results customsearch.Search
	var links []string

	// Set up Google Custom Search API credentials and client
	apiKey := os.Getenv("GCS_KEY") // Your Google API key
	cseID := os.Getenv("GCS_ID")   // Your Custom Search Engine ID
	ctx := context.Background()

	cseService, err := customsearch.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Custom Search service: %v", err)
	}
	// Perform a Google search
	search := cseService.Cse.List().Q(query).Cx(cseID)

	// Execute the search and handle pagination
	for startIndex := 1; startIndex <= numResults; startIndex += 10 {
		search.Start(int64(startIndex))
		results_temp, err := search.Do()

		if err != nil {
			log.Fatalf("Failed to perform search: %v", err)
		}
		results.Items = append(results.Items, results_temp.Items...)
	}

	for index := 0; index < numResults; index++ {
		links = append(links, results.Items[index].Link)
	}
	return links
}
