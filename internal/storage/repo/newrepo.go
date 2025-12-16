package repo

import (
	"database/sql"

	"order-providing-system/internal/storage"
)

type Repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) storage.RepoModule {
	return &Repo{db: db}
}
