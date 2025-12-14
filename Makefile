APP_ENV ?= dev

install:
	go mod tidy

run:
	APP_ENV=$(APP_ENV) go run ./cmd/api

test:
	go test ./... -race -cover

recreate-db:
	docker compose -f dev/docker-compose.yml down -v
	docker compose -f dev/docker-compose.yml up -d

db-up:
	cd dev && docker-compose down -v && docker-compose up -d

migrate-up:
	goose -dir ./schema postgres "$(DB_DSN)" up

migrate-down:
	goose -dir ./schema postgres "$(DB_DSN)" down

sqlc:
	sqlc generate

docker:
	docker build -t app-backend:dev .