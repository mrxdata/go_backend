.PHONY: clean critic security lint test build run swag

include .env
include .env.db
include .env.redis
export

APP_NAME = apiserver
BUILD_DIR = $(CURDIR)/build
MIGRATIONS_FOLDER = $(CURDIR)/platform/migrations
PG_DB_URL = $(PG_DB_DRIVER)://$(PG_DB_USER):$(PG_DB_PASSWORD)@$(PG_DB_HOST):$(PG_DB_PORT)/$(PG_DB_NAME)?sslmode=$(PG_DB_SSL_MODE)

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.pg.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(PG_DB_URL)" up

migrate.pg.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(PG_DB_URL)" down

migrate.pg.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(PG_DB_URL)" force $(version)

docker.run: docker.network docker.postgres swag docker.fiber docker.redis migrate.up

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name flutty-fiber \
		--network dev-network \
		-p $(SERVER_PORT):$(SERVER_PORT) \
		fiber

docker.postgres:
	docker run --rm -d \
		--name flutty-postgres \
		--network dev-network \
		-e POSTGRES_USER=$(PG_DB_USER) \
		-e POSTGRES_PASSWORD=$(PG_DB_PASSWORD) \
		-e POSTGRES_DB=$(PG_DB_NAME) \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p $(PG_DB_PORT):$(PG_DB_PORT) \
		postgres

docker.redis:
	docker run --rm -d \
		--name flutty-redis \
		--network dev-network \
		-p $(REDIS_PORT):$(REDIS_PORT) \
		redis

docker.stop: docker.stop.fiber docker.stop.postgres docker.stop.redis

docker.stop.fiber:
	docker stop flutty-fiber

docker.stop.postgres:
	docker stop flutty-postgres

docker.stop.redis:
	docker stop flutty-redis

swag:
	swag init
