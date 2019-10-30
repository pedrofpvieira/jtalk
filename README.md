[![CircleCI](https://circleci.com/gh/pedrofpvieira/jtalk/tree/master.svg?style=svg)](https://circleci.com/gh/pedrofpvieira/jtalk/tree/master)

# JTalk Microservice

This microservice intends to create a bridge for communication between 2 peers.

## Technology Stack
* Golang (https://golang.org/)
* ScyllaDB (https://www.scylladb.com/)
* CircleCI (https://circleci.com)

## Setup
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

## Swagger
```
http://127.0.0.1:7070/swagger/
```

## IDE Configurations
### VSCode
* Configure Debug using Palette
```
Ctrl + Shift + P > Go: Install/Update Tools
Search dlv and Select
Enter
Ctrl + Shift + P > Debug: Open launch.json
Enter
```
The above sequence will instal 'delve' and create a configuration file launch.json under .vscode