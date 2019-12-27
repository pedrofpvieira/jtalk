package models

import (
	"jtalk/src/db"
	"jtalk/src/entities"
	"math/rand"
	"time"
)

const insertQuery = "INSERT into conversations (author_id, conversation_id, created_at) VALUES (?, ?, ?)"

// FindConversationsByAuthor Return All the conversation based on the author ID
func FindConversationsByAuthor(authorID int64) []entities.Conversation {
	conversations := make([]entities.Conversation, 0)
	row := map[string]interface{}{}

	iter := db.Session.Query("SELECT * FROM conversations WHERE author_id = ?", authorID).Iter()

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

// CreateConversation Creates a new conversation between two authors
func CreateConversation(senderID int64, recipientID int64) int64 {

	createdAt := time.Now()
	conversationID := rand.Int63n(1000000000000) // for now random, needs to revised

	// TODO check lightweith transaction system since we are adding two inserts
	querySender := db.Session.Query(insertQuery, senderID, conversationID, createdAt)
	queryRecipient := db.Session.Query(insertQuery, recipientID, conversationID, createdAt)

	if err := querySender.Exec(); err != nil {
		panic(err)
	}

	if err := queryRecipient.Exec(); err != nil {
		panic(err)
	}

	return conversationID
}
