package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "Message added to conversation successfully!",
	})
}
