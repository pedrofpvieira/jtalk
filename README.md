# JTalk Microservice

This microservice intends to create a bridge for communication between 2 peers.

## Technology Stack
* Golang (https://golang.org/)
* ScyllaDB (https://www.scylladb.com/)

## Setup
Just run:
```
git clone git@github.com:pedrofpvieira/jtalk.git
cd jtalk
docker-compose up
```

## Endpoints

* POST /conversations
* GET /conversations/:conversation_id
* POST /conversations/:conversation_id/messages
* DELETE /conversations/:conversation_id/messages/:message_id