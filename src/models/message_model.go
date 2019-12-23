package models

import (
	"jtalk/src/db"
	"jtalk/src/entities"
	"math/rand"
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

// AddMessage Adds a message to an existing conversation
func AddMessage(conversationID int64, authorID int64, message string) {
	query := db.Session.Query("INSERT into conversation_messages (conversation_id, author_id, message_id, created_at) VALUES (?, ?, ?, ?)",
		conversationID,
		authorID,
		rand.Intn(1000000000000), //TODO for now random, needs to be revised
		time.Now(),
	)

	if err := query.Exec(); err != nil {
		panic(err)
	}
}
