# MIME Project Makefile

BINARY_NAME=mime
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean test coverage install run lint help

all: build

build: ## Build binary to ./bin/mime
	@echo "Building $(BINARY_NAME)..."
	@go build -o bin/$(BINARY_NAME) ./cmd/mime

clean: ## Remove build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

install: ## Install binary to GOBIN
	@echo "Installing $(BINARY_NAME)..."
	@go install ./cmd/mime

run: ## Run the application (serve)
	@go run cmd/mime/main.go serve

lint: ## Run golangci-lint
	@echo "Linting..."
	@golangci-lint run

## Help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
