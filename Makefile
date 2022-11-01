ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
.PHONY: help run quality test coverage build up down
GOLANG_VERSION := 1.21-alpine

help: ## Display this help screen
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the application locally
	@echo "Running application..."
	@go run ./cmd/snitch.go

quality: ## Run the quality checks
	@echo "Running Go vet..."
	@docker run --rm -w /app -v "${PWD}":/app golang:$(GOLANG_VERSION) go vet ./...

test: ## Run the tests and generate coverage report
	@echo "Running tests..."
	@docker run --rm -w /app -v "${PWD}":/app golang:$(GOLANG_VERSION) go test -v ./... -cover

coverage: test ## Display HTML coverage report locally
	@echo "Generating coverage report..."
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

build: ## Build the Docker image
	@echo "Building Docker image..."
	@docker build -t snitch:latest .