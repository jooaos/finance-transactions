include .env

DB_CONNECTION="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8"

attach-app:
	@docker compose exec app sh

attach-db:
	@docker compose exec app sh

migration-create:
	@docker compose --profile tools run --rm migrate ${DB_CONNECTION} create -ext sql -seq -dir /migrations/ ${NAME}

migration-up:
	@docker compose --profile tools run --rm migrate ${DB_CONNECTION} up

migration-down:
	@docker compose --profile tools run --rm migrate ${DB_CONNECTION} down

integration-test-migrate-up:
	@docker compose --profile tools run --rm migrate ${DB_TEST_CONNECTION} up

integration-test-migrate-down:
	@docker compose --profile tools run --rm migrate ${DB_TEST_CONNECTION} down

integration-test-run:
	./scripts/test-integration-run.sh

up:
	@docker compose up

up-no-cache:
	@docker compose up --build

up-detached:
	@docker compose up -d

up-db:
	@docker compose up db