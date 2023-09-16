package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Option func(*pgxpool.Config)

func WithConnMaxIdleTime(t time.Duration) Option {
	return func(config *pgxpool.Config) {
		config.MaxConnIdleTime = t
	}
}

func WithMaxOpenConns(n int32) Option {
	return func(config *pgxpool.Config) {
		config.MaxConns = n
	}
}
