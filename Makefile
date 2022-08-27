up:
	DOCKER_BUILDKIT=1 docker compose up --build

down:
	docker compose down

db-in:
	docker compose exec db bash -c "psql \"user=postgres password=postgres_pw host=localhost port=5432 dbname=postgres\""

prune:
	docker system prune

gen:
	cd app; \
	go generate ./ent
