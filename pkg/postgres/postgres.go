// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"fmt"
	"github.com/ivanov-gv/color-picker-test-app/pkg/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func New(cfg config.PG) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DbConnString)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create pg pool: %w", err)
	}

	return pool, nil
}
