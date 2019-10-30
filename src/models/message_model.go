package models

import (
	"jtalk/src/db"
	"jtalk/src/entities"
	"time"
)

// FindMessagesByConversation List the paged conversation based on the conversationID
func FindMessagesByConversation(conversationID int64) []entities.Message {
	var conversationid, messageid, authorid int64
	var content string
	var createdat time.Time
	messages := make([]entities.Message, 0)

	iter := db.Session.Query("SELECT * FROM conversation_messages WHERE conversation_id = ?", conversationID).Iter()

	for iter.Scan(&conversationid, &messageid, &createdat, &authorid, &content) {
		m := entities.Message{
			ConversationID: conversationid,
			Content:        content,
			AuthorID:       authorid,
			CreatedAt:      createdat,
		}
		messages = append(messages, m)
	}

	err := iter.Close()
	if err != nil {
		panic(err)
	}
	return messages
}
