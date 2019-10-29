package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAuthorConversations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"conversations": "List of conversations here!",
	})
}

func getConversation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Here you have the conversation!",
	})
}
