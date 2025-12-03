.PHONY: run build test clean install dev

run:
	go run main.go

build:
	go build -o bin/app main.go

test:
	go test -v ./...

test-cover:
	go test -cover ./...

clean:
	rm -rf bin/
	rm -rf data/*.db

install:
	go mod download
	go mod tidy

dev:
	air

.DEFAULT_GOAL := run
