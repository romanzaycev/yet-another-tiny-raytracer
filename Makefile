MKDIR_P = mkdir -p

.PHONY: build
build: clean ## Build executable.
	go build -v -o tinyraytracer ./cmd

.PHONY: test
test: ## Run unit tests.
	go test -v -race -timeout 30s ./src/*

.PHONY: coverage
coverage: ## Run test with generating coverge report.
	${MKDIR_P} ./build
	go test -v -race -coverprofile ./build/coverage.out -timeout 30s ./src/*
	go tool cover -html ./build/coverage.out -o ./build/coverage.html

.PHONY: clean
clean: ## Cleanup
	rm -f ./tinyraytracer

.PHONY: run
run: build ## Run executable.
	./tinyraytracer

.PHONY: help
help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
