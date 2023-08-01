package repository

import (
	"database/sql"
)

type Repository struct {
	*PgRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PgRepo: NewPgRepo(db),
	}
}
