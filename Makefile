include .env

run:
	docker compose up --force-recreate --remove-orphans
clear:
	docker compose down -v --rmi local
migrate-create:
	@if [ -n ${name} ]; then \
		migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq ${name}; \
	else \
		echo "Usage: make migrate-create <migration_name>"; \
		exit 1; \
	fi