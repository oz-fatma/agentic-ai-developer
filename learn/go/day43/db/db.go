package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB = sql.DB

func OpenMemory() (*sql.DB, error) {
	return sql.Open("sqlite", "file::memory:?cache=shared")
}

func InitAccounts(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			owner TEXT NOT NULL,
			balance INTEGER NOT NULL CHECK(balance >= 0)
		)
	`)
	return err
}

func SeedAccounts(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO accounts (owner, balance) VALUES
			('alice', 500),
			('bob', 200)
	`)
	return err
}
