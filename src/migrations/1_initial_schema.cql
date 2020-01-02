CREATE TABLE IF NOT EXISTS messages
(
    conversation_id bigint,
    author_id       bigint,
    message_id      bigint,
    content         text,
    created_at      timestamp,
    PRIMARY KEY (conversation_id, author_id)
);

CREATE TABLE IF NOT EXISTS conversations
(
    author_id       bigint,
    conversation_id bigint,
    created_at      timestamp,
    PRIMARY KEY (author_id, conversation_id)
);