package psql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgresDb(ctx context.Context, connString string) (*Postgres, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	return &Postgres{
		Db: db,
	}, nil
}

func (db *Postgres) Close() error {
	if err := db.Db.Close(); err != nil {
		return err
	}
	return nil
}
