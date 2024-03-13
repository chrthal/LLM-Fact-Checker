package api

import (
	"chrthal/llm-fact-checker/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	id = 0
)

func SetupRoutes(queue *models.JobQueue, resolvedJobs *models.JobQueue) {
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
			}

			for i, engine := range newRequest.SearchEngines {
				newJob.SearchEngineData[i] = models.SearchEngineData{
					SearchEngine: engine,
					Urls:         make([]string, 0),
				}
			}

			queue.Mu.Lock()
			queue.Jobs = append(queue.Jobs, newJob)
			queue.Mu.Unlock()

			fmt.Printf("New job added: %+v\n", newJob)

			c.JSON(http.StatusOK, gin.H{"success": true})
			id += 1
		})
	}

	router.Run(":8080")
}
