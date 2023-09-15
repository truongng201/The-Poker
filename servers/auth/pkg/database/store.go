package database

import (
	database "auth-service/pkg/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	database.Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*database.Queries
	connPool *pgxpool.Pool
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  database.New(connPool),
	}
}
