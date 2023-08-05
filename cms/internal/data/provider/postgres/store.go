// internal/data/provider/postgres/store.go
package postgres

import (
	"context"

	"database/sql"
	_ "github.com/lib/pq"

	"cms/main/internal/models"
)

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) CreateSchema(ctx context.Context, schema models.Schema) error {
	// implement your logic to create schema here
	// for now, just return nil
	return nil
}

func (s *PostgresStore) GetSchema(ctx context.Context, name string) (models.Schema, error) {
	// implement your logic to get schema here
	// for now, just return an empty schema and nil error
	return models.Schema{}, nil
}


func NewPostgresStore(ctx context.Context, connString string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}
