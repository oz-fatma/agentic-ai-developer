package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day41/db"
)

func main() {
	fmt.Println("Day 41: database/sql basics with SQLite")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.InitUsers(conn); err != nil {
		log.Fatal(err)
	}

	id, err := insertUser(conn, "Ada", "ada@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user id:", id)

	name, email, err := getUser(conn, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fetched user: id=%d name=%s email=%s\n", id, name, email)

	count, err := countUsers(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total users:", count)
}

func insertUser(db *sql.DB, name, email string) (int64, error) {
	res, err := db.Exec(`INSERT INTO users (name, email) VALUES (?, ?)`, name, email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func getUser(db *sql.DB, id int64) (string, string, error) {
	var name, email string
	err := db.QueryRow(`SELECT name, email FROM users WHERE id = ?`, id).Scan(&name, &email)
	return name, email, err
}

func countUsers(db *sql.DB) (int, error) {
	var n int
	err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&n)
	return n, err
}
