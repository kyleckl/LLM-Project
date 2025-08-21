.PHONY: all build run format lint clean

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

clean:
	rm -rf bin/
