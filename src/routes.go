package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes() {

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusNotImplemented, "")
	})

	// Start a new conversation
	router.POST("/conversations", startConversation)

	// Get Conversation
	router.GET("/conversations/:conversation_id", getConversation)

	// Add a message to a conversation
	router.POST("/conversations/:conversation_id/messages", addMessage)

	// Delete a message from a conversation
	router.DELETE("/conversations/:conversation_id/messages/:message_id", deleteMessage)
}
