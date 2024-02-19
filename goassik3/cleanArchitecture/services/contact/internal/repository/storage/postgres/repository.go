package postgres

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	var r = &Repository{
		db: db,
	}

	return r
}
