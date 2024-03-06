include .env

DB_CONNECTION="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8"

up:
	@docker compose up

up-no-cache:
	@docker compose up --build

up-detached:
	@docker compose up -d


migration-create:
	@docker compose --profile tools run --rm migrate ${DB_CONNECTION} create -ext sql -seq -dir /migrations/ ${NAME}

migration-up:
	@docker compose --profile tools run --rm migrate ${DB_CONNECTION} up

migration-down:
	@docker compose --profile tools run --rm migrate ${DB_CONNECTION} down
