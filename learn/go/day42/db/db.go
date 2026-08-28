package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenMemory() (*sql.DB, error) {
	return sql.Open("sqlite", "file::memory:?cache=shared")
}

func InitProducts(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price REAL NOT NULL
		)
	`)
	return err
}
