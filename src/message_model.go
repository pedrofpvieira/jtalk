package main

func FindMessagesByConversation(conversation_id int64) []Message {
    return []Message{
        Message{conversation_id: 1234, content: "abc"},
        Message{conversation_id: conversation_id, content: "def"},
    }
}

