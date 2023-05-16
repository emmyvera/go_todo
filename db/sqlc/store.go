package db

import (
	"database/sql"
)

// Store provides all functions to execute DB queries
type Store struct {
	*Queries
	db *sql.DB
}

// Store provides all functions to execute SQL queries
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
