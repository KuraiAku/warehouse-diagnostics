package database

import (
	"database/sql"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func Open() (*sql.DB, error) {
	connectionString := os.Getenv("WAREHOUSE_DB_URL")

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
