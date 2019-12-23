package models

import (
	"jtalk/src/db"
	"jtalk/src/entities"
	"time"
)

// FindConversationsByAuthor Return All the conversation based on the author ID
func FindConversationsByAuthor(authorID int64) []entities.Conversation {
	conversations := make([]entities.Conversation, 0)
	row := map[string]interface{}{}

	iter := db.Session.Query("SELECT * FROM author_conversations WHERE author_id = ?", authorID).Iter()

	for iter.MapScan(row) {
		conversations = append(conversations, entities.Conversation{
			ConversationID: row["conversation_id"].(int64),
			CreatedAt:      row["created_at"].(time.Time),
		})
		row = map[string]interface{}{}
	}

	err := iter.Close()
	if err != nil {
		panic(err)
	}
	return conversations
}
