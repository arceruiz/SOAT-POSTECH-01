run:
	docker compose -f ./build/db-docker-compose.yaml up -d
build-app:
	docker build --no-cache -t soat_postech_database . -f Dockerfile.db
build-db:
	docker build --no-cache -t soat_postech_app . -f Dockerfile
build-and-run: build-db build-app run