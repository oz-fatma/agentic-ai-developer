package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day57/internal/adapter/memory"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day57/internal/usecase"
)

func main() {
	fmt.Println("Day 57: clean architecture layers")

	repo := memory.NewTaskRepo()
	svc := usecase.NewTaskService(repo)

	task, err := svc.Add("design layers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added task: %+v\n", task)

	tasks, err := svc.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tasks: %d\n", len(tasks))
}
