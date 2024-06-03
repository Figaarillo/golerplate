.PHONY: all clean security test build run run.build docker.build docker.run docker.stop docker.run.test docker.stop.test docker.clean migrate.up migrate.down migrate.force swag

### VARIABLES ###
APP_NAME = apiserver
BUILD_DIR = ./build
MIGRATIONS_FOLDER = $(PWD)/migrations
DATABASE_URL = postgres://${DATABASE_USER}:${DATABASE_PASS}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable

### COMMANDS ###
clean:
	rm -rf $(BUILD_DIR)

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out


test.unit:
	@echo "Running unit test..."
	go test -v ./internal/domain/entity/

test.unit.category:
	@echo "Running unit test for category..."
	go test -v ./internal/domain/entity/category_test.go

test.unit.product:
	@echo "Running unit test for product..."
	go test -v ./internal/domain/entity/product_test.go

test.e2e:
	@echo "Running all E2E tests..."
	go test -v ./internal/test/

test.e2e.category:
	@echo "Running E2E tests for Category..."
	go test ./internal/test/ -run Category -v

test.e2e.product:
	@echo "Running E2E tests for Product..."
	go test ./internal/test/ -run Product -v

build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api

run: docker.run.db
	@echo "Running local server..."
	DATABASE_HOST=localhost go run ./cmd/api/main.go

run.build:
	@echo "Running build app..."
	$(BUILD_DIR)/$(APP_NAME)

docker.build:
	docker-compose build

docker.run: swag docker.clean docker.build
	@echo "Runnung server in docker container..."
	docker-compose up -d database apiserver

docker.run.db:
	docker-compose up -d database

docker.run.test:
	@echo "Running database for testing..."
	docker-compose up -d database-test

docker.stop:
	@echo "Stop docker container..."
	docker-compose stop database apiserver database-test

docker.stop.test:
	@echo "Stop docker container for testing..."
	docker-compose stop database-test
	docker-compose rm -f database-test

docker.clean:
	docker-compose down --volumes

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

swag:
	./scripts/swag init -g cmd/api/main.go -d ./ -o ./docs
