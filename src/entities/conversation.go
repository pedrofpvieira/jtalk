package entities

import (
	"strconv"
	"time"
)

// Conversation Entity to represent a Conversation from the DB
type Conversation struct {
	ConversationID int64
	CreatedAt      time.Time
}

// ToMap Maps the conversation to a map for Json Serialization
func (conversation Conversation) ToMap() map[string]string {
	conversationMap := make(map[string]string)

	conversationMap["conversation_id"] = strconv.FormatInt(conversation.ConversationID, 10)
	conversationMap["created_at"] = conversation.CreatedAt.String()

	return conversationMap
}
