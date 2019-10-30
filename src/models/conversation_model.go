package models

import (
	"jtalk/src/db"
	"jtalk/src/entities"
	"time"
)

// FindConversationsByAuthor Return All the conversation based on the author ID
func FindConversationsByAuthor(authorID int64) []entities.Conversation {

	var conversationid, authorid int64
	var createdat time.Time
	conversations := make([]entities.Conversation, 0)

	iter := db.Session.Query("SELECT * FROM author_conversations WHERE author_id = ?", authorID).Iter()

	for iter.Scan(&authorid, &conversationid, &createdat) {
		c := entities.Conversation{ConversationID: conversationid, AuthorID: authorid, CreatedAt: createdat}
		conversations = append(conversations, c)
	}

	err := iter.Close()
	if err != nil {
		panic(err)
	}
	return conversations
}
