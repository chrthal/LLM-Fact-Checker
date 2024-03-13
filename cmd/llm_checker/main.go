package main

import (
	api "chrthal/llm-fact-checker/api/handlers"
	"chrthal/llm-fact-checker/internal/search_fetcher"
	"chrthal/llm-fact-checker/models"
	"fmt"

	"log"
	"net/http"
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

			// Fetch search result urls
			for i := range job.SearchEngineData {
				searchEngineData := &job.SearchEngineData[i]
				*searchEngineData = search_fetcher.SearchFetcher(*searchEngineData, job.Question, job.PagesToCrawl)
			}

			resolvedJobs.Mu.Lock()
			resolvedJobs.Jobs = append(resolvedJobs.Jobs, job)
			resolvedJobs.Mu.Unlock()
		}
		time.Sleep(time.Second)
	}
}
