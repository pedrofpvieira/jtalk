package models

import (
	"jtalk/src/db"
	"jtalk/src/entities"
	"time"
)

// FindMessagesByConversation List the paged conversation based on the conversationID
func FindMessagesByConversation(conversationID int64) []entities.Message {
	messages := make([]entities.Message, 0)
	row := map[string]interface{}{}

	iter := db.Session.Query("SELECT * FROM conversation_messages WHERE conversation_id = ?", conversationID).Iter()

	for iter.MapScan(row) {
		messages = append(messages, entities.Message{
			AuthorID:  row["author_id"].(int64),
			MessageID: row["message_id"].(int64),
			Content:   row["content"].(string),
			CreatedAt: row["created_at"].(time.Time),
		})
		row = map[string]interface{}{}
	}

	err := iter.Close()
	if err != nil {
		panic(err)
	}
	return messages
}
