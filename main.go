package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create a semaphore with a buffer of 10
	semaphore := make(chan struct{}, 10)

	r.POST("/send_email", func(c *gin.Context) {
		// Acquire a semaphore token
		semaphore <- struct{}{}
		defer func() {
			// Release the semaphore token
			<-semaphore
		}()

		// Simulate sending email (replace with your actual logic)
		time.Sleep(300 * time.Millisecond)

		c.JSON(200, gin.H{"message": "Email sent successfully"})
	})

	r.Run(":8080")
}
