#!make
include .env

# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./cmd/main.go
BINARY_NAME := app

build:
	go build -o bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

dev:
	# Custom port
	# air -- --port=:4002
	air \
		--build.cmd "make build" --build.bin "bin/${BINARY_NAME}" --build.delay "100" \
		--misc.clean_on_exit "true"

run: build
	./bin/app

tidy:
	@echo "Tidying up..."
	go fmt ./...
	go mod tidy -v

test:
	go test -v ./... -count=1

sqlc:
	@echo "Generating SQLC..."
	sqlc generate

goose-status:
	goose -dir ./migrations postgres "user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" status 

goose-create:
ifdef NAME
	@echo "Creating migration $(NAME)"
	goose -dir ./migrations create $(NAME) sql
else
	$(error NAME is not set. Please provide a name using NAME=<migration-name>)
endif

goose-up:
	@echo "Running migrations..."
	goose -dir ./migrations postgres "user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" up

goose-down:
	@echo "Rolling back migrations..."
	goose -dir ./migrations postgres "user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" down