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
