package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewCockroachDB(ctx context.Context, cfg PostgresConfig) (*sql.DB, error) {
	database, err := sql.Open("pgx", cfg.string())
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err = database.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return database, nil
}
