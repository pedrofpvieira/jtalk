package handlers

import (
	"jtalk/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAuthorConversations Handler to return all conversations per author_id
func GetAuthorConversations(c *gin.Context) {
	authorID, _ := strconv.ParseInt(c.Param("author_id"), 10, 64)

	conversations := models.FindConversationsByAuthor(authorID)

	var conversationsAsMap []map[string]string
	var status int

	for _, conversation := range conversations {
		conversationsAsMap = append(conversationsAsMap, conversation.ToMap())
	}

	if conversationsAsMap == nil {
		status = http.StatusNotFound
	} else {
		status = http.StatusOK
	}

	c.JSON(status, gin.H{
		"conversations": conversationsAsMap,
	})
}

// GetConversation Handler to return all mesages from a conversation_id
func GetConversation(c *gin.Context) {
	conversationID, _ := strconv.ParseInt(c.Param("conversation_id"), 10, 64)

	messages := models.FindMessagesByConversation(conversationID)

	var messagesAsMap []map[string]string
	var status int

	for _, message := range messages {
		messagesAsMap = append(messagesAsMap, message.ToMap())
	}

	if messagesAsMap == nil {
		status = http.StatusNotFound
	} else {
		status = http.StatusOK
	}

	c.JSON(status, gin.H{
		"messages": messagesAsMap,
	})
}

// CreateConversation Creates a new conversation between two authors. If they already have a conversation
// do not create and simply return the existing conversation ID
func CreateConversation(c *gin.Context) {
	senderID, _ := strconv.ParseInt(c.PostForm("sender_id"), 10, 64)
	recipientID, _ := strconv.ParseInt(c.PostForm("recipient_id"), 10, 64)
	message := c.PostForm("message")

	conversationID := models.CreateConversation(senderID, recipientID)

	models.AddMessage(conversationID, senderID, message)
}
