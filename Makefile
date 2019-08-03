.PHONY: build

PROJECT_NAME ?= ebox-api
VERSION ?= $(shell git rev-parse HEAD)

build:
	go build -o build/$(PROJECT_NAME) cmd/server/main.go

run: build
	./build/$(PROJECT_NAME)

run-deps:
	docker-compose up -d database

run-all:
	docker-compose up

image-build:
	docker build -t $(PROJECT_NAME):$(VERSION) --target production .
