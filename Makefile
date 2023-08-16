.PHONY: dev
dev: ## make clean, creates database container, starts it, make start
	make clean
	docker-compose create database
	docker-compose start
	sleep 10
	make start

.PHONY: start
start: ## starts go service
	go run cmd/main.go

.PHONY: build
build: ## compiles code
	docker-compose build backend

.PHONY: clean
clean: ## stops containers and prunes them
	docker-compose stop database
	docker container prune

.PHONY: help
help:
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'
