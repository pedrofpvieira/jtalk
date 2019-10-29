# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GENERATEDFOLDER=_generated/

test:
	$(GOTEST) -v ./src

test-coverage: 
	$(GOTEST) -coverprofile=$(GENERATEDFOLDER)c.out -v ./src

test-coverage-html: test-coverage
	$(GOTOOL) cover -html=$(GENERATEDFOLDER)c.out -o $(GENERATEDFOLDER)coverage.html