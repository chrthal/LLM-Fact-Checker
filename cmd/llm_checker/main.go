package main

import (
	api "chrthal/llm-fact-checker/api/handlers"
	"chrthal/llm-fact-checker/internal/html_grabber"
	"chrthal/llm-fact-checker/internal/llm_fetcher"
	"chrthal/llm-fact-checker/internal/search_fetcher"
	"chrthal/llm-fact-checker/models"

	"fmt"
	"log"
	"net/http"
	//"os/exec"
	"sync"
	"time"

	strutil "github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

type PythonOutput struct {
	Similarity float64 `json:"similarity"`
}

var (
	jobQueue = models.JobQueue{
		Jobs: make([]models.Job, 0),
		Mu:   sync.Mutex{},
	}
	resolvedJobs = models.JobQueue{
		Jobs: make([]models.Job, 0),
		Mu:   sync.Mutex{},
	}
	runningJobs = 0
)

func main() {
	go queueWatchdog()

	api.SetupRoutes(&jobQueue, &resolvedJobs, &runningJobs)
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
			runningJobs++
			fmt.Printf("Processing job: %+v\n", job)

			job.Status = "Resolved"

			// Fetch search results (URLs)
			log.Println("Start fetching urls from search engines...")
			for i := range job.SearchEngineData {
				searchEngineData := &job.SearchEngineData[i]
				*searchEngineData = search_fetcher.SearchFetcher(*searchEngineData, job.Question, job.PagesToCrawl)
			}

			log.Println("Start fetching results from llm...")
			for i := range job.LLMData {
				llmData := &job.LLMData[i]
				*llmData = llm_fetcher.LLMFetcher(*llmData, job.Question)
				job.CrawledData.LLMScrape = llmData.Response
			}

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

			/* ---------- Start Comparison ----- */
			log.Println("Start comparison")
			sum_similarity := 0.00
			for _, webScrape := range job.CrawledData.WebScrape {
				log.Println(webScrape)
				log.Println(job.CrawledData.LLMScrape)

				//cmd := exec.Command("python", "/root/python/compare_texts.py", string(llmScrapeArg), string(webScrapeArg))
				similarity := strutil.Similarity(job.CrawledData.LLMScrape, webScrape, metrics.NewJaccard())
				// Get the output from the command
				// output, err := cmd.CombinedOutput()
				// if err != nil {
				// 	log.Println(err)
				// }

				// Parse the output into the predefined structure
				// var result PythonOutput
				// if err := json.Unmarshal(output, &result); err != nil {
				// 	log.Println(err)
				// }
				sum_similarity += similarity
			}

			job.CrawledData.Similarity = sum_similarity / float64(len(job.CrawledData.WebScrape))
			resolvedJobs.Mu.Lock()
			resolvedJobs.Jobs = append(resolvedJobs.Jobs, job)
			resolvedJobs.Mu.Unlock()
			log.Println("Job resolved...")
			runningJobs--
		}
		time.Sleep(time.Second)
	}
}
