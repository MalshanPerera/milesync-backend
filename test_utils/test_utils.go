package test_utils

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func CreatePostgresContainer(ctx context.Context, db_scripts []string) (*PostgresContainer, error) {
	scripts := make([]string, 0, len(db_scripts))
	defer RemoveTempMigrations()
	migrations, err := GetMigrationScripts()
	if err != nil {
		return nil, err
	}

	for _, migration := range migrations {
		scripts = append(scripts, filepath.Join("..", "migrations", migration))
	}

	for _, script := range db_scripts {
		scripts = append(scripts, filepath.Join("..", "testdata", script))
	}
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(scripts...),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

// GetMigrationScripts retrieves the migration scripts from the "migrations" directory and extracts the content between "-- +goose Up" and "-- +goose Down" comments.
// It creates a temporary directory "temp_migrations" and saves the extracted content into separate files.
// The function returns a slice of file paths to the extracted migration scripts.
// If an error occurs during the process, it returns nil and the error.
func GetMigrationScripts() ([]string, error) {
	files, err := os.ReadDir(filepath.Join("..", "migrations"))
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, len(files))
	os.Mkdir(filepath.Join("..", "temp_migrations"), 0755)
	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(filepath.Join("..", "migrations", file.Name()))
			if err != nil {
				return nil, err
			}
			start := strings.Index(string(content), "-- +goose Up")
			end := strings.Index(string(content), "-- +goose Down")
			if start == -1 || end == -1 || start >= end {
				continue
			}

			extracted := content[start:end]
			err = os.WriteFile(filepath.Join("..", "temp_migrations", file.Name()), extracted, 0644)
			if err != nil {
				return nil, err
			}

			paths = append(paths, filepath.Join("..", "temp_migrations", file.Name()))
		}
	}

	return paths, nil
}

func RemoveTempMigrations() error {
	err := os.RemoveAll(filepath.Join("..", "temp_migrations"))
	if err != nil {
		return err
	}

	return nil
}
