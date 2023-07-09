.PHONY: start
start:
	go run cmd/main.go

.PHONY: build
build:
	go mod tidy