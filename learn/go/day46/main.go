package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day46/internal/db"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day46/internal/models"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day46/internal/repository"
)

func main() {
	fmt.Println("Day 46: repository interface pattern")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	var users repository.UserRepository = repository.NewSQLiteUserRepo(conn)

	u, err := users.Create("Grace", "grace@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created: %+v\n", u)

	found, err := users.GetByEmail("grace@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found by email: %+v\n", found)

	all, err := users.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total users: %d\n", len(all))
	_ = models.User{}
}
