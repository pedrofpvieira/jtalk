package main

// FindMessagesByConversation List the paged conversation based on the conversationID
func FindMessagesByConversation(conversationID int64) []Message {
	return []Message{
		Message{conversationID: 1234, content: "abc"},
		Message{conversationID: conversationID, content: "def"},
	}
}
