// Package postgres provides repository layer implementations using postgres
package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	dbpkg "github.com/4nd3rs0n/oapi-test/pkg/db"
)

type DB struct {
	*pgxpool.Pool
	*Users
}

// NewDB initializes a new connection pool.
func NewDB(ctx context.Context, url string) (*DB, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{
		Pool:  pool,
		Users: &Users{q: dbpkg.New(pool)},
	}, nil
}
