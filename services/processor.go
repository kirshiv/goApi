package services

import (
	"log"
	"math/rand"
	"time"

	"RETAIL/models"
	"RETAIL/utils"
)
func ProcessJob(jobID string, request models.JobRequest, store *models.JobStore) {
	errors := []string{}
	for _, visit := range request.Visits {
		for _, url := range visit.ImageURLs {
			perimeter, err := utils.CalculatePerimeter(url)
			if err != nil {
				log.Printf("Failed to process image %s: %v", url, err)
				errors = append(errors, visit.StoreID)
				continue
			}

			// Log the calculated perimeter
			log.Printf("Processed image %s: Perimeter = %d\n", url, perimeter)

			// Simulate processing time
			time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
		}
	}

	status := "completed"
	if len(errors) > 0 {
		status = "failed"
	}

	store.UpdateJobStatus(jobID, status, errors)
}

