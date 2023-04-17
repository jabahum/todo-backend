DOCKER_CONTAINER_NAME=todo-backend_postgres_1
DB_NAME=todo
DB_URL=postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable

postgres:
	docker run --name $(DOCKER_CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.4-alpine

createdb:
	docker exec -it $(DOCKER_CONTAINER_NAME) createdb --username=root --owner=root $(DB_NAME)

dropdb:
	docker exec -it $(DOCKER_CONTAINER_NAME) dropdb $(DB_NAME)

backupdb:
	docker exec -it $(DOCKER_CONTAINER_NAME) pg_dump $(DB_NAME) > backup.sql
	
restoredb:
	docker exec -i $(DOCKER_CONTAINER_NAME) psql $(DB_NAME) < backup.sql

build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: watch run build postgres dropdb
.DEFAULT_GOAL := help


