migrate:
	@docker compose  -f docker/docker-compose.yml up -d
	@echo "### Importing rows, pleas waiting... ###"
	@go run ./import/main.go
	@echo "### Import finish ###"
	@docker compose  -f docker/docker-compose.yml down
	@echo "### Run -> make run ###"

run:
	@echo "### Runing docker ###"
	@docker compose  -f docker/docker-compose.yml up -d
	@echo "### Runing Go app ###"
	@go run main.go

off:
	@docker-compose -f docker/docker-compose.yml down
