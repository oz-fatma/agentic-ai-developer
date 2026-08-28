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
	path := t.TempDir() + "/tasks.db"
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() { conn.Close() })
	return conn
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			done INTEGER NOT NULL DEFAULT 0
		)
	`)
	return err
}
