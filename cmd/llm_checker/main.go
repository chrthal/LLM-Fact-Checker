package main

import (
	api "chrthal/llm-fact-checker/api/handlers"
	"chrthal/llm-fact-checker/internal/html_grabber"
	"chrthal/llm-fact-checker/internal/search_fetcher"
	"chrthal/llm-fact-checker/models"
	"fmt"

	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

var (
	jobQueue = models.JobQueue{
		Jobs: make([]models.Job, 0),
		Mu:   sync.Mutex{},
	}
	resolvedJobs = models.JobQueue{
		Jobs: make([]models.Job, 0),
		Mu:   sync.Mutex{},
	}
)

func main() {
	go queueWatchdog()

	api.SetupRoutes(&jobQueue, &resolvedJobs)
	log.Println("Starting server on :8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func queueWatchdog() {
	for {
		if len(jobQueue.Jobs) > 0 {
			// Lock the queue
			jobQueue.Mu.Lock()
			// Get the first job in the queue
			job := jobQueue.Jobs[0]

			// Remove the first job from the queue
			jobQueue.Jobs = jobQueue.Jobs[1:]
			jobQueue.Mu.Unlock()

			fmt.Printf("Processing job: %+v\n", job)

			job.Status = "Resolved"

			// Fetch search results (URLs)
			log.Println("Start fetching urls from search engines...")
			for i := range job.SearchEngineData {
				searchEngineData := &job.SearchEngineData[i]
				*searchEngineData = search_fetcher.SearchFetcher(*searchEngineData, job.Question, job.PagesToCrawl)
			}

			log.Println("Start fetching results from llm...")
			/*for i := range job.LLMData {
				llmData := &job.LLMData[i]
				*llmData = llm_fetcher.LLMFetcher(*llmData, job.Question)
			}*/

			urls := make([]string, 0)

			for i := range job.SearchEngineData {
				searchEngineData := &job.SearchEngineData[i]

				urls = append(urls, searchEngineData.Urls...)
			}

			// Fetch and prepare body from urls
			log.Println("Start fetching and preparing body from urls...")
			for _, url := range urls {
				body, err := html_grabber.FetchAndPrepareBody(url)
				if err != nil {
					log.Printf("Failed to fetch body from url: %s\n", url)
					continue
				}
				log.Printf("Fetched body from url: %s\n", url)
				job.CrawledData.WebScrape = append(job.CrawledData.WebScrape, body)
			}
			log.Println("Done fetching and preparing body from urls...")

			cmd := exec.Command("python3", "/root/python/compare_texts.py", "claim", "article")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			print(output)

			resolvedJobs.Mu.Lock()
			resolvedJobs.Jobs = append(resolvedJobs.Jobs, job)
			resolvedJobs.Mu.Unlock()
		}
		time.Sleep(time.Second)
	}
}
