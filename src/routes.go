package main

import (
	"jtalk/src/handlers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRoutes() {
	// TODO For development purposes, to be reviewed, for now all requests are accepted
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusNotImplemented, "")
	})

	router.Static("/swagger/", "./swagger-ui")

	conversationRoutes := router.Group("/conversations")
	{
		// For a given author id, return a list of available conversations
		conversationRoutes.GET("/author/:author_id", handlers.GetAuthorConversations)

		// Get Conversation
		conversationRoutes.GET("/messages/:conversation_id", handlers.GetConversation)
	}
}
