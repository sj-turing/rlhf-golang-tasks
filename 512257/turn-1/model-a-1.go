package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/subscription-event", handleSubscriptionEvent)
	r.Run(":8080")
}

func handleSubscriptionEvent(c *gin.Context) {
	userID := c.Query("user_id")
	notificationType := c.Query("notification_type")

	// Validate inputs
	if userID == "" || notificationType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	// Process the event in a goroutine
	go processSubscriptionEvent(userID, notificationType)

	c.JSON(http.StatusOK, gin.H{"message": "Event received and processed"})
}
