CONFIG_FILE ?= ./config/local.json
DB_DSN := $(shell sed -n 's/.*"dsn": "\(.*\)",/\1/p' $(CONFIG_FILE))
MIGRATE := docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_DSN)


.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: docker
docker:
	@docker compose -f ./docker/docker-compose.yml up -d

.PHONY: docker-down
docker-down:
	@docker compose -f ./docker/docker-compose.yml down

.PHONY: migrate-new
migrate-new:
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations $${name// /_}
	
.PHONY: migrate
migrate:
	@echo "Running all new database migrations ..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down:
	@echo "Reverting database to the last migration ..."
	@$(MIGRATE) down 1

.PHONY: migrate-reset
migrate-reset:
	@echo "Resetting database ..."
	@$(MIGRATE) drop -f
	@echo "Running all database migrations ..."
	@$(MIGRATE) up