run: build
	@./bin/term/main

build:
	@go build -o bin/term/main cmd/term/main.go
