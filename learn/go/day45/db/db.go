package db

import (
	"database/sql"
	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schemaSQL string

type Book struct {
	Title string
	Year  int
}

func OpenMemory() (*sql.DB, error) {
	conn, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}
	if _, err := conn.Exec(`PRAGMA foreign_keys = ON`); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(schemaSQL)
	return err
}

func InsertAuthor(db *sql.DB, name string) (int64, error) {
	res, err := db.Exec(`INSERT INTO authors (name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func InsertBook(db *sql.DB, authorID int64, title string, year int) (int64, error) {
	res, err := db.Exec(`INSERT INTO books (author_id, title, year) VALUES (?, ?, ?)`, authorID, title, year)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func ListBooksByAuthor(db *sql.DB, authorID int64) ([]Book, error) {
	rows, err := db.Query(`SELECT title, year FROM books WHERE author_id = ? ORDER BY year`, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.Title, &b.Year); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, rows.Err()
}
