# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GENERATEDFOLDER=_generated/
DOCKERCOMPOSE=docker-compose

test:
	$(GOTEST) -v ./src

test-coverage: 
	$(GOTEST) -coverprofile=$(GENERATEDFOLDER)c.out -v ./src

test-coverage-html: test-coverage
	$(GOTOOL) cover -html=$(GENERATEDFOLDER)c.out -o $(GENERATEDFOLDER)coverage.html

# This reloads server container to rebuild GO (temp solution until hot-reload)
reload-server:
	$(DOCKERCOMPOSE) stop server && $(DOCKERCOMPOSE) start server

reload-scylla:
	$(DOCKERCOMPOSE) stop scylla && $(DOCKERCOMPOSE) start scylla

cqlsh:
	$(DOCKERCOMPOSE) exec scylla cqlsh -k jtalk