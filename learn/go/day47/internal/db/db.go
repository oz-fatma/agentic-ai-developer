package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenMemory() (*sql.DB, error) {
	return sql.Open("sqlite", "file::memory:?cache=shared")
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			body TEXT NOT NULL
		)
	`)
	return err
}
