package api

import (
	"chrthal/llm-fact-checker/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var resolvedJobs []models.Job

var test_job = models.Job{
	ID:            1,
	Question:      "Whats the capital of france?",
	PagesToCrawl:  15,
	SearchEngines: 2,
	Status:        "In Progress",
}

func SetupRoutes() {
	router := gin.Default()

	resolvedJobs = append(resolvedJobs, test_job)

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
			c.JSON(http.StatusOK, resolvedJobs)
		})

		// Endpoint to add new job
		api.POST("/addJob", func(c *gin.Context) {
			var newJob models.Job
			if err := c.BindJSON(&newJob); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Assign ID and add new job to the list of resolved jobs
			newJob.ID = len(resolvedJobs) + 1
			resolvedJobs = append(resolvedJobs, newJob)

			fmt.Printf("New job added: %+v\n", newJob)

			c.JSON(http.StatusOK, gin.H{"success": true})
		})
	}

	router.Run(":8080")
}
