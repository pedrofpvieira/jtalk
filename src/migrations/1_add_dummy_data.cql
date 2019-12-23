CREATE TABLE IF NOT EXISTS conversation_messages
(
    conversation_id bigint,
    author_id       bigint,
    message_id      bigint,
    content         text,
    created_at      timestamp,
    PRIMARY KEY (conversation_id, author_id)
);

CREATE TABLE IF NOT EXISTS author_conversations
(
    author_id       bigint,
    conversation_id bigint,
    created_at      timestamp,
    PRIMARY KEY (author_id, conversation_id)
);

INSERT INTO author_conversations (author_id, conversation_id, created_at) VALUES (1,100,'2019-10-29T17:40:16+00:00');
INSERT INTO author_conversations (author_id, conversation_id, created_at) VALUES (2,100,'2019-10-29T17:40:16+00:00');
INSERT INTO author_conversations (author_id, conversation_id, created_at) VALUES (3,101,'2019-10-29T17:40:16+00:00');
INSERT INTO author_conversations (author_id, conversation_id, created_at) VALUES (4,101,'2019-10-29T17:40:16+00:00');

INSERT INTO conversation_messages (conversation_id,author_id,message_id,content,created_at) VALUES (100,1000,1,'teste','2019-10-29T17:40:16+00:00');
INSERT INTO conversation_messages (conversation_id,author_id,message_id,content,created_at) VALUES (100,1001,2,'test2','2019-10-29T17:40:16+00:00');

