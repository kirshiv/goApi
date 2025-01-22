package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"RETAIL/models"
	"RETAIL/services"
)

var jobStore = models.NewJobStore()

// SubmitJob handles job submission
func SubmitJob(c *gin.Context) {
	var payload models.JobRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	if payload.Count != len(payload.Visits) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Count does not match number of visits"})
		return
	}

	jobID := jobStore.CreateJob(payload)
	go services.ProcessJob(jobID, payload, jobStore)

	c.JSON(http.StatusCreated, gin.H{"job_id": jobID})
}

// GetJobStatus handles job status retrieval
func GetJobStatus(c *gin.Context) {
	jobID := c.Query("jobid")
	if jobID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing job ID"})
		return
	}

	job, err := jobStore.GetJob(jobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": job.Status, "job_id": jobID, "error": job.Errors})
}
