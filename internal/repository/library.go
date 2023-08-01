package repository

import "database/sql"

type PgRepo struct {
	db *sql.DB
}

func NewPgRepo(db *sql.DB) *PgRepo {
	return &PgRepo{db: db}
}
