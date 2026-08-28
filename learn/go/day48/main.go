package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day48/internal/db"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day48/internal/repository"
)

func main() {
	fmt.Println("Day 48: database tests")
	fmt.Println("Run: go test ./day48/...")

	conn, err := db.OpenTestDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTagRepo(conn)
	tag, err := repo.Create("go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created tag: %+v\n", tag)
}
