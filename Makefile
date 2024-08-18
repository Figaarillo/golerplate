.PHONY: all clean build test test.unit test.unit.cover test.e2e run run.build docker.build docker.run docker.stop docker.run.test docker.stop.test docker.clean docs

# ##################### VARIABLES ##################### #

APP_NAME = apiserver
BUILD_DIR = ./build
DATABASE_URL = postgres://${DATABASE_USER}:${DATABASE_PASS}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable

# ###################### COMMANDS ##################### #

clean:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │             CLEANNIG BUILD             │ "
	@echo " ╰────────────────────────────────────────╯ "
	rm -rf $(BUILD_DIR)

build: clean
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │              BUILDING APP              │ "
	@echo " ╰────────────────────────────────────────╯ "
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api

test.unit:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │         RUNNING ALL UNIT TESTS         │ "
	@echo " ╰────────────────────────────────────────╯ "
	go test -v -timeout 30s -coverprofile=cover.out -cover ./internal/domain/entity/

test.unit.cover:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │         COVEARAGE OF ALL TESTS         │ "
	@echo " ╰────────────────────────────────────────╯ "
	go tool cover -func=cover.out

test.unit.category:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │     RUNNING UNIT TEST FOR CATEGORY     │ "
	@echo " ╰────────────────────────────────────────╯ "
	go test -v ./internal/domain/entity/category_test.go

test.unit.product:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │      RUNNING UNIT TEST FOR PRODUCT     │ "
	@echo " ╰────────────────────────────────────────╯ "
	go test -v ./internal/domain/entity/product_test.go

test.e2e:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │          RUNNING ALL E2E TESTS         │ "
	@echo " ╰────────────────────────────────────────╯ "
	go test -v ./internal/test/

test.e2e.category:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │     RUNNING E2E TESTS FOR CATEGORY     │ "
	@echo " ╰────────────────────────────────────────╯ "
	go test ./internal/test/ -run Category -v

test.e2e.product:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │      RUNNING E2E TESTS FOR PRODUCT     │ "
	@echo " ╰────────────────────────────────────────╯ "
	go test ./internal/test/ -run Product -v

run: docker.run.db
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │         RUNNING SERVER IN LOCAL        │ "
	@echo " ╰────────────────────────────────────────╯ "
	DATABASE_HOST=localhost go run ./cmd/api/main.go

run.build: docker.run.db build
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │    RUNNING BUILD OF SERVER IN LOCAL    │ "
	@echo " ╰────────────────────────────────────────╯ "
	DATABASE_HOST=localhost ./build/apiserver

docker.build:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │       BUILDING DOCKER CONTAINERS       │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose build

docker.run: docs docker.clean docker.build
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │        RUNNING SERVER IN DOCKER        │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose up -d database apiserver

docker.run.db:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │       RUNNING DATABASE CONTAINER       │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose up -d database

docker.run.test:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │  RUNNING DATABASE CONTAINER FOR TESTS  │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose up -d database-test

docker.stop:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │     STOPPING ALL DOCKER CONTAINERS     │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose stop database apiserver database-test

docker.stop.test:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │      STOPPING CONTAINER FOR TESTS      │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose stop database-test
	docker-compose rm -f database-test

docker.clean:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │        CLEANNING DOCKER VOLUMES        │ "
	@echo " ╰────────────────────────────────────────╯ "
	docker-compose down --volumes

docs:
	@echo " ╭────────────────────────────────────────╮ "
	@echo " │         GENERATING SWAGGER DOC         │ "
	@echo " ╰────────────────────────────────────────╯ "
	./scripts/swag init -g cmd/api/main.go -d ./ -o ./docs
