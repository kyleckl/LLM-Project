.PHONY: all build run format unit-test lint clean

all: build

build:
	go build -o bin/llmApp ./internal

run: build
	./bin/llmApp

format:
	go fmt ./...

lint:
	go vet ./...
	# Add any other linting tools here, e.g., golangci-lint
	# golangci-lint run

unit-test:
	@if [ -z "$$GOOGLE_API_KEY" ]; then \
		echo "Error: GOOGLE_API_KEY is not set. Google LLM testing will be skipped"; \
	fi
	ginkgo ./...

clean:
	rm -rf bin/
