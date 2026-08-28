package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day50/internal/db"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day50/internal/repository"
)

func main() {
	fmt.Println("Day 50: full repository with tests")
	fmt.Println("Run: go test ./day50/...")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTaskRepo(conn)
	task, err := repo.Create("write tests", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created task: %+v\n", task)

	done := true
	updated, err := repo.Update(task.ID, nil, &done)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated task: %+v\n", updated)
}
