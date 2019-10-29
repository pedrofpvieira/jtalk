package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAuthorConversations(c *gin.Context) {
	authorID, _ := strconv.ParseInt(c.Param("author_id"), 10, 64)

	c.JSON(http.StatusOK, gin.H{
		"conversations": FindConversationsByAuthor(authorID),
	})
}

func getConversation(c *gin.Context) {
	conversationID, _ := strconv.ParseInt(c.Param("conversation_id"), 10, 64)

	messages := FindMessagesByConversation(conversationID)

	var messagesAsMap []map[string]string
	for _, message := range messages {
		messagesAsMap = append(messagesAsMap, message.ToMap())
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messagesAsMap,
	})
}
