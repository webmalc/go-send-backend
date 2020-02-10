
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run .
GOCOV=$(GOCMD) tool cover -html=coverage.out
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go_send_backend

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	GOENV=test $(GOTEST) ./... -coverprofile=coverage.out

testv:
	GOENV=test $(GOTEST) -v ./... -coverprofile=coverage.out

testl: testv lint

coverage:
	$(GOCOV)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

lint: 
	golangci-lint run ./...
run:
	$(GORUN)

