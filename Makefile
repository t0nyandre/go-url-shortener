.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: docker
docker:
	@docker compose -f ./docker/docker-compose.yml up -d

.PHONY: docker-down
docker-down:
	@docker compose -f ./docker/docker-compose.yml down