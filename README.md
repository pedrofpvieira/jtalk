[![CircleCI](https://circleci.com/gh/pedrofpvieira/jtalk/tree/master.svg?style=svg)](https://circleci.com/gh/pedrofpvieira/jtalk/tree/master)

# JTalk Microservice

This microservice intends to create a bridge for communication between 2 peers.

## Technology Stack
* Golang (https://golang.org/)
* ScyllaDB (https://www.scylladb.com/)
* CircleCI (https://circleci.com)

## Setup
Just run:
```
git clone git@github.com:pedrofpvieira/jtalk.git
cd jtalk
docker-compose up
```
## Run Tests
```
cd jtalk
make test (Basic unit test validation)
make test-coverage (Unit Test + Coverage)
make test-coverage-html (Unit Test + Coverage in HTML, exported to _generated/ folder)
```
## Endpoints

* GET /conversations/author/:author_id
* GET /conversations/messages/:conversation_id