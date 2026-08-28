package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day47/internal/db"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day47/internal/repository"
)

func main() {
	fmt.Println("Day 47: SQL organization")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewNoteRepo(conn)
	n, err := repo.Create("Learn Go", "Keep SQL in one place")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created note %d\n", n.ID)

	notes, err := repo.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, note := range notes {
		fmt.Printf("  [%d] %s — %s\n", note.ID, note.Title, note.Body)
	}
}
