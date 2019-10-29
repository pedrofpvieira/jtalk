#!/bin/bash
CQL="CREATE KEYSPACE jtalk WITH replication ={'class': 'SimpleStrategy', 'replication_factor' : 3};

CREATE TABLE jtalk.conversation_messages
(
    conversation_id bigint,
    message_id      bigint,
    author_id       bigint,
    content         text,
    created_at      timestamp,
    PRIMARY KEY ((conversation_id, message_id), created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);

CREATE TABLE jtalk.author_conversations
(
    author_id       bigint,
    conversation_id bigint,
    created_at      timestamp,
    PRIMARY KEY (author_id, conversation_id)
) WITH CLUSTERING ORDER BY (author_id DESC);"

until echo $CQL | cqlsh; do
  echo "cqlsh: Scylla is unavailable to initialize - will retry later"
  sleep 2
done &

exec /docker-entrypoint.py "$@"