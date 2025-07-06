# Загружаем переменные из .env
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

run: clear
	docker compose up --force-recreate --remove-orphans
clear:
	docker compose down -v --rmi local
migrate-create:
	@if [ -n ${name} ]; then \
		migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq ${name}; \
	else \
		echo "Usage: make migrate-create <migration_name>"; \
		exit 1; \
	fi