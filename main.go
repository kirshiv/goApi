package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"RETAIL/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/api/submit/", handlers.SubmitJob)
	r.GET("/api/status", handlers.GetJobStatus)

	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
