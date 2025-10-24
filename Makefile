.PHONY: help test test-unit test-acceptance test-coverage run clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run all tests
	go test ./... -v

test-unit: ## Run unit tests only
	go test ./internal/... -v

test-coverage: ## Run tests with coverage report
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

run: ## Run the CLI application
	go run cmd/api/main.go --weight=5.5 --distance=150 --delivery-type=EXPRESS

build: ## Build the application
	go build -o bin/shipping-calculator cmd/api/main.go

clean: ## Clean build artifacts and test cache
	go clean
	rm -rf bin/
	rm -f coverage.out coverage.html

deps: ## Download dependencies
	go mod download
	go mod tidy

fmt: ## Format code
	go fmt ./...

lint: ## Run linter (requires golangci-lint)
	golangci-lint run

.DEFAULT_GOAL := help
