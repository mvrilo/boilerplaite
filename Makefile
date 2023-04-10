.PHONY: build
build:
	go mod tidy
	go build -o boilerplaite ./cmd/boilerplaite/main.go
