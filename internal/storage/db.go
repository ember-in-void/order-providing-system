package storage

import (
	"database/sql"

	"order-providing-system/pkg/logger"

	_ "github.com/lib/pq"
)

func SetupDataBase(logg *logger.CustomLogger) *sql.DB {
	connStr := "host=db port=5432 user=latte password=latte dbname=order-providing-system sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logg.Error(err)
	}

	logg.Info("Successfully connected to database")
	return db
}
