run: build
	@./bin/main

build:
	@go build -o bin/main cmd/main.go
