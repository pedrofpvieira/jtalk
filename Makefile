# Go parameters
GENERATEDFOLDER=_generated/
DOCKERCOMPOSE=docker-compose
RUNSERVER=docker-compose run --rm --no-deps server sh

test:
	$(RUNSERVER) -c "go test -v ./src"

test-coverage: 
	$(RUNSERVER) -c "go test -coverprofile=$(GENERATEDFOLDER)c.out -v ./src"

test-coverage-html:
	$(RUNSERVER) -c "go tool cover -html=$(GENERATEDFOLDER)c.out -o $(GENERATEDFOLDER)coverage.html"

# TODO This reloads server container to rebuild GO (temp solution until hot-reload)
reload-server:
	$(DOCKERCOMPOSE) stop server && $(DOCKERCOMPOSE) start server

reload-scylla:
	$(DOCKERCOMPOSE) stop scylla && $(DOCKERCOMPOSE) start scylla

cqlsh:
	$(DOCKERCOMPOSE) exec scylla cqlsh -k jtalk

staticcheck:
	$(RUNSERVER) -c "cd src/ && staticcheck -f stylish jtalk/src"
