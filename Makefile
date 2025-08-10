MIGRATIONS_DIR=migrations

run:
	MIGRATIONS_DIR=${MIGRATIONS_DIR} \
		docker compose up --force-recreate --remove-orphans
clear:
	docker compose down -v --rmi local
migrate-create:
	@if [ -n ${name} ]; then \
		migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq ${name}; \
	else \
		echo "Usage: make migrate-create name=<migration_name>"; \
		exit 1; \
	fi
