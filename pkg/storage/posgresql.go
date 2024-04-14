package storage

import (
	"context"
	"database/sql"
	"fmt"
)

func NewPostgreSQL(ctx context.Context, cfg PostgresConfig) (*sql.DB, error) {
	database, err := sql.Open("postgres", cfg.string())
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}
	if err := database.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("could not ping database")
	}
	return database, nil
}
