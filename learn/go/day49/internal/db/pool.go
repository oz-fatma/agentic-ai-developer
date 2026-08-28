package db

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

func OpenPooled(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(4)
	conn.SetMaxIdleConns(2)
	conn.SetConnMaxLifetime(5 * time.Minute)
	return conn, nil
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS metrics (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			source TEXT NOT NULL,
			created_at TEXT NOT NULL DEFAULT (datetime('now'))
		)
	`)
	return err
}

func InsertMetric(db *sql.DB, source string) error {
	_, err := db.Exec(`INSERT INTO metrics (source) VALUES (?)`, source)
	return err
}

func CountMetrics(db *sql.DB) (int, error) {
	var n int
	err := db.QueryRow(`SELECT COUNT(*) FROM metrics`).Scan(&n)
	return n, err
}
