package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func startConversation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Conversation started successfully!",
	})
}

func getConversation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Here you have the conversation!",
	})
}
