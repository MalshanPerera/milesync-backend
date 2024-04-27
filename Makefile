#!make
include .env

build:
	go build -o bin/app

dev:
	# Custom port
	# air -- --port=:4002
	air

run: build
	./bin/app

test:
	go test -v ./... -count=1

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