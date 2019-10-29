package main

import "strconv"

// Message Entity to represent a Message from the DB
type Message struct {
	conversationID, messageID, authorID int64
	content                             string
	createdAt                           int
}

// ToMap Maps the message to a map for Json Serialization
func (message Message) ToMap() map[string]string {
	messageMap := make(map[string]string)

	messageMap["conversation_id"] = strconv.FormatInt(message.conversationID, 10)
	messageMap["message_id"] = strconv.FormatInt(message.messageID, 10)
	messageMap["author_id"] = strconv.FormatInt(message.authorID, 10)
	messageMap["content"] = message.content
	messageMap["created_at"] = strconv.FormatInt(int64(message.createdAt), 10)

	return messageMap
}
