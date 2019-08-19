# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=twitter-clone
BINARY_UNIX=$(BINARY_NAME)_unix
all: test build
fmt:
		$(GOCMD) fmt ./...
build:

		$(GOBUILD) -o $(BINARY_NAME) -v
linux:
		GOOS=linux $(GOBUILD) -o $(BINARY_NAME) -v	
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
getdeps:
	    $(GOCMD) mod vendor	
# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker:
		docker build -t twitter-clone .
docker-run:

		docker run -d  -p 8090:8090 twitter-clone:latest 

deploy:
		kubectl apply -f deploy/mongo.yaml
		kubectl create secret generic configfile  --from-file=properties.json
		kubectl apply -f deploy/backend_service.yaml
		kubectl apply -f deploy/backend.yaml


.PHONY: all fmt build linux test clean run getdeps build-linux docker docker-run deploy
