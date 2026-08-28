package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenMemory() (*sql.DB, error) {
	return sql.Open("sqlite", "file::memory:?cache=shared")
}

func InitEmployees(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			department TEXT,
			salary REAL
		)
	`)
	return err
}

func SeedEmployees(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO employees (name, department, salary) VALUES
			('Riley', 'Engineering', 95000),
			('Jordan', NULL, NULL),
			('Casey', 'Sales', 72000)
	`)
	return err
}
