run:
	docker compose up --force-recreate --remove-orphans
clear:
	docker compose down -v --rmi local