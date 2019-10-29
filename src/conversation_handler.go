package main

import (
    "strconv"
    "net/http"

    "github.com/gin-gonic/gin"
)

func getAuthorConversations(c *gin.Context) {
    author_id, _ := strconv.ParseInt(c.Param("author_id"), 10, 64)

    c.JSON(http.StatusOK, gin.H{
        "conversations": FindConversationsByAuthor(author_id),
    })
}

func getConversation(c *gin.Context) {
    conversation_id, _ := strconv.ParseInt(c.Param("conversation_id"), 10, 64)

    messages := FindMessagesByConversation(conversation_id)

    var messagesAsMap []map[string]string
    for _, message := range messages {
        messagesAsMap = append(messagesAsMap, message.ToMap())
    }

    c.JSON(http.StatusOK, gin.H{
        "messages": messagesAsMap,
    })
}



