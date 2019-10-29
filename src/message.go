package main

import "strconv"

type Message struct {
    conversation_id, message_id, author_id int64
    content string
    created_at int
}

func (message Message) ToMap() map[string]string {
    messageMap := make(map[string]string)

    messageMap["conversation_id"] = strconv.FormatInt(message.conversation_id, 10)
    messageMap["message_id"] = strconv.FormatInt(message.message_id, 10)
    messageMap["author_id"] = strconv.FormatInt(message.author_id, 10)
    messageMap["content"] = message.content
    messageMap["created_at"] = strconv.FormatInt(int64(message.created_at), 10)

    return messageMap
}
