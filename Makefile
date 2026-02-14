APP_ENV ?= dev
DB_URL ?= postgres://postgres:postgres@127.0.0.1:5432/tszh_db?sslmode=disable

MIGRATE_DIR=schema
MIGRATE_EXT=sql

install:
	go mod tidy

run:
	APP_ENV=$(APP_ENV) go run ./cmd/api

test:
	go test ./... -race -cover

recreate-db:
	docker compose -f docker-compose.dev.yml down -v
	docker compose -f docker-compose.dev.yml up -d

db-up:
	docker-compose -f docker-compose.dev.yml down && docker-compose -f docker-compose.dev.yml up -d

migrate-up:
	migrate -path ./schema -database "${DB_URL}" up

migrate-down:
	migrate -path ./schema -database "${DB_URL}" down

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "❌ Usage: make create-migration name=your_migration_name"; \
		exit 1; \
	fi
	migrate create -ext $(MIGRATE_EXT) -dir $(MIGRATE_DIR) -seq $(name)

sqlc:
	sqlc generate

docker:
	docker build -t app-backend:dev .