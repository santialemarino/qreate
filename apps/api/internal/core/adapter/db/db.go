package db

import (
	"context"
	"errors"
	"fmt"
	"qreate/config"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database wraps the database connection and query builder
type Database struct {
	DB      *sqlx.DB
	Dialect goqu.DialectWrapper
}

func New(ctx context.Context, cfg *config.Settings) (*Database, error) {
	dbConfig := cfg.Database

	if dbConfig.Driver != "postgres" {
		return nil, errors.New("only postgres driver is supported")
	}

	dialect := goqu.Dialect("postgres")

	db, err := CreatePostgresConnection(ctx, &dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	if db == nil {
		return nil, errors.New("failed to create database connection")
	}

	return &Database{
		DB:      db,
		Dialect: dialect,
	}, nil
}

func CreatePostgresConnection(ctx context.Context, dbConfig *config.DatabaseConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.SSLMode,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
