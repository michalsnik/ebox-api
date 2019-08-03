.PHONY: build
.PHONY: run run-deps run-all
.PHONY: migrate-up migrate-down migrate-status
.PHONY: image-build

PROJECT_NAME ?= ebox-api
VERSION ?= $(shell git rev-parse HEAD)
DB_CONN ?= "postgres://postgres:supersecret@localhost:5434/postgres?sslmode=disable"

build:
	go build -o build/$(PROJECT_NAME) cmd/server/main.go

run: build
	./build/$(PROJECT_NAME)

run-deps:
	docker-compose up -d database

run-all:
	docker-compose up

migrate-up:
	goose -dir internal/db/migrations postgres $(DB_CONN) up

migrate-down:
	goose -dir internal/db/migrations postgres $(DB_CONN) down

migrate-status:
	goose -dir internal/db/migrations postgres $(DB_CONN) status

migration:
	goose -dir internal/db/migrations postgres $(DB_CONN) create $(name) $(type)

image-build:
	docker build -t $(PROJECT_NAME):$(VERSION) --target production .
