package main

import (
	api "chrthal/llm-fact-checker/api/handlers"

	"log"
	"net/http"
)

func main() {
	// var question string = "Whats the capital of france?"
	// var count_results int = 15

	// var links_google []string = search_api.GoogleSearch(question, count_results)
	// var links_bing []string = search_api.BingSearch(question, count_results)

	// for link_index, link := range links_google {
	// 	log.Println(link_index+1, link)
	// }

	// for link_index, link := range links_bing {
	// 	log.Println(link_index+1, link)
	// }
	api.SetupRoutes()
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
