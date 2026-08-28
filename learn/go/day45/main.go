package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day45/db"
)

func main() {
	fmt.Println("Day 45: schema with foreign keys")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	authorID, err := db.InsertAuthor(conn, "Octavia Butler")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.InsertBook(conn, authorID, "Parable of the Sower", 1993); err != nil {
		log.Fatal(err)
	}
	if _, err := db.InsertBook(conn, authorID, "Kindred", 1979); err != nil {
		log.Fatal(err)
	}

	books, err := db.ListBooksByAuthor(conn, authorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Books by author %d:\n", authorID)
	for _, b := range books {
		fmt.Printf("  %s (%d)\n", b.Title, b.Year)
	}
}
