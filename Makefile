# Makefile for a Go project

# Go related variables
GO ?= go
GORUN = $(GO) run
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOTEST = $(GO) test
GOFMT = $(GO) fmt
BUILD_DIR = dist
APP_NAME = connect4
TARGET_PACKAGE = cmd/main.go
DOCKER_TAG = latest

# Default target executed when no arguments are given to make.
all: build

# Build the project
build: 
	$(GOBUILD) -o ${BUILD_DIR}/$(APP_NAME) -v $(TARGET_PACKAGE)

# Run tests
test: 
	$(GOTEST) -v ./...

# Format the code
fmt: 
	$(GOFMT) ./...

# Clean the binary and other build artifacts
clean: 
	$(GOCLEAN)
	rm -f ${BUILD_DIR}

run-dev:
	$(GORUN) $(TARGET_PACKAGE)

# Run the application
run: build
	./${BUILD_DIR}/$(APP_NAME)

docker-build:
	rm -f ${BUILD_DIR}
	docker build -t ${APP_NAME}:latest .

docker-run:
	docker run -it ${APP_NAME}:${DOCKER_TAG}

.PHONY: all build test fmt clean run-dev run docker-build docker-run