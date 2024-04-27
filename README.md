# Jira For Peasants 😤 

## Prerequisites

Install Air -> `go install https://github.com/cosmtrek/air`

Install Goose -> `go install github.com/pressly/goose/v3/cmd/goose@latest`

## Commands

Start Dev Server -> `make dev`

Build Prod Server -> `make build`

Create Migration -> `make goose-create NAME=${new_migration_name}`

Migrate Up database -> `make goose-up`

Migrate Down entire database -> `make goose-down`

Kill Server

`lsof -ti :4000 | xargs kill`