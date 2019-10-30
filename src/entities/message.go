package entities

import (
	"strconv"
	"time"
)

// Message Entity to represent a Message from the DB
type Message struct {
	ConversationID, MessageID, AuthorID int64
	Content                             string
	CreatedAt                           time.Time
}

// ToMap Maps the message to a map for Json Serialization
func (message Message) ToMap() map[string]string {
	messageMap := make(map[string]string)

	messageMap["conversation_id"] = strconv.FormatInt(message.ConversationID, 10)
	messageMap["message_id"] = strconv.FormatInt(message.MessageID, 10)
	messageMap["author_id"] = strconv.FormatInt(message.AuthorID, 10)
	messageMap["content"] = message.Content
	messageMap["created_at"] = message.CreatedAt.String()

	return messageMap
}
