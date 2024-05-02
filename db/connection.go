package datastore

import (
	"context"
	"fmt"
	sqlc "jira-for-peasants/db/sqlc"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool  *pgxpool.Pool
	query *sqlc.Queries
}

// Use pgxpool for concurrent connections
// if you have multiple threads working with a DB at the same time, you must use pgxpool
func NewDB(
	username string,
	password string,
	host string,
	port string,
	database string,
) *DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	var err error
	d := &DB{}
	d.pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = d.pool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	d.query = sqlc.New(d.pool)
	fmt.Println("Connected to database")

	return d
}

func NewDBFromConnectionString(connStr string) *DB {
	var err error
	d := &DB{}
	d.pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = d.pool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	d.query = sqlc.New(d.pool)
	fmt.Println("Connected to database")

	return d
}

func (d *DB) Close() {
	d.pool.Close()
}

func (d *DB) GetDB() *pgxpool.Pool {
	return d.pool
}

func (d *DB) GetQuery() *sqlc.Queries {
	return d.query
}

func (d *DB) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return d.pool.Begin(ctx)
}

func (d *DB) RollbackTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}

func (d *DB) CommitTx(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}
