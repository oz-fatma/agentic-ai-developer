package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenMemory() (*sql.DB, error) {
	return sql.Open("sqlite", "file::memory:?cache=shared")
}

func InitUsers(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		)
	`)
	return err
}
