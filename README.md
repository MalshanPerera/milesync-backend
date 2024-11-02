# Jira For Peasants 😤

## Folder Structure

The project follows a standard Go project layout:

- `/cmd` - Main applications for this project
  - `/main.go` - Entry point of the application
- `/internal` - Private application and library code
  - `/api` - API request/response models and validation
  - `/handlers` - HTTP handlers for each route
  - `/repositories` - Database access layer
  - `/services` - Business logic layer
- `/migrations` - Database migration files
- `/pkg` - Library code that can be used by external applications
- `/bin` - Compiled application binaries

## Prerequisites

Install Air -> `go install https://github.com/cosmtrek/air`

Install Goose -> `go install github.com/pressly/goose/v3/cmd/goose@latest`

## Commands

Start Dev Server -> `make dev`

Build Prod Server -> `make build`

Generate SQLC -> `make sqlc`

Create Migration -> `make goose-create NAME=${new_migration_name}`

Migrate Up database -> `make goose-up`

Migrate Down entire database -> `make goose-down`

Kill Server -> `lsof -ti :4000 | xargs kill`
