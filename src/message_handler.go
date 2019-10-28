package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "Message added to conversation successfully!",
	})
}

func deleteMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "Message deleted from conversation successfully!",
	})
}
