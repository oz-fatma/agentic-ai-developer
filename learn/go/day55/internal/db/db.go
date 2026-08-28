package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenMemory() (*sql.DB, error) {
	return sql.Open("sqlite", "file::memory:?cache=shared")
}

func OpenTempDB(t interface {
	TempDir() string
	Cleanup(func())
}) *sql.DB {
	path := t.TempDir() + "/users.db"
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() { conn.Close() })
	return conn
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'user'
		)
	`)
	return err
}
