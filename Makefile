#
# Project Makefile
#
export GO111MODULE=on

.PHONY: setup help dep install format lint vet build build-docker test test-coverage
.DEFAULT: help

dev-setup: ## Install Development dependencies
	@pip install pre-commit
	@go install github.com/mgechev/revive@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@pre-commit install
	

download: # Download go modules
	@go mod download

fmt: ## Formats the go code using gofmt
	@gofmt -w -s .

lint: ## Lint code
	@revive -config revive.toml -formatter friendly $(PACKAGE)

vet: ## Run go vet
	@go vet $(PACKAGE)

build: ## Build the app
	@go build -o build/main main.go

docs: ## generate swagger docs
	@swag init

run: ## Build and run the app
	@go run main.go

test: ## Run package unit tests
	@go test -v -short ./...

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ./...

help: ## Displays help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
