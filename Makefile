run: build
	@./bin/ssh/main

build:
	@go build -o bin/ssh/main cmd/ssh/main.go
