package api

import (
	"chrthal/llm-fact-checker/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	id = 0
)

func SetupRoutes(queue *models.JobQueue, resolvedJobs *models.JobQueue, runningJobs *int) {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./web/build", true)))

	api := router.Group("/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// Endpoint to fetch resolved jobs
		api.GET("/resolvedJobs", func(c *gin.Context) {
			queue.Mu.Lock()
			resolvedJobs.Mu.Lock()
			allJobs := append(resolvedJobs.Jobs, queue.Jobs...)
			resolvedJobs.Mu.Unlock()
			queue.Mu.Unlock()
			c.JSON(http.StatusOK, allJobs)
		})
		api.GET("/status", func(c *gin.Context) {
			status := map[string]interface{}{
				"queuedJobs":   len(queue.Jobs),
				"runningJobs":  runningJobs,
				"resolvedJobs": len(resolvedJobs.Jobs),
			}
			c.JSON(http.StatusOK, status)
		})
		// Endpoint to add new job
		api.POST("/addJob", func(c *gin.Context) {
			var newRequest models.JobRequest
			if err := c.BindJSON(&newRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			newJob := models.Job{
				ID:               id,
				Question:         newRequest.Question,
				PagesToCrawl:     newRequest.PagesToCrawl,
				SearchEngineData: make([]models.SearchEngineData, len(newRequest.SearchEngines)),
				LLMData:          make([]models.LLMData, len(newRequest.LLMs)),
				StartDate:        time.Now(),
				LastUpdate:       time.Now(),
			}

			for i, engine := range newRequest.SearchEngines {
				newJob.SearchEngineData[i] = models.SearchEngineData{
					SearchEngine: engine,
					Urls:         make([]string, 0),
				}
			}

			for i, llm := range newRequest.LLMs {
				newJob.LLMData[i] = models.LLMData{
					LLM: llm,
				}
			}

			queue.Mu.Lock()
			queue.Jobs = append(queue.Jobs, newJob)
			queue.Mu.Unlock()

			fmt.Printf("New job added: %+v\n", newJob)

			c.JSON(http.StatusOK, gin.H{"success": true})
			id += 1
		})

		api.POST("/config", func(c *gin.Context) {
			var newConfig models.Config
			if err := c.BindJSON(&newConfig); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			os.Setenv("OPENAI_API_KEY", newConfig.OpenAIKey)
			os.Setenv("GCS_KEY", newConfig.GCSKey)
			os.Setenv("GCS_ID", newConfig.GCSID)
			os.Setenv("BING_KEY", newConfig.BingKey)
			os.Setenv("OLLAMA_HOST", newConfig.OllamaHost)
			os.Setenv("OLLAMA_VERBOSE", newConfig.OllamaVerbose)

			c.JSON(http.StatusOK, gin.H{"status": "Environment variables set successfully"})
		})
	}

	router.Run(":8080")
}
