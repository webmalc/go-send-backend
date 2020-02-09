
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run .
GOCOV=$(GOCMD) tool cover -html=coverage.out
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go_send_backend

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./... -coverprofile=coverage.out
coverage:
	$(GOCOV)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GORUN)

