package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day51/auth"
)

func main() {
	fmt.Println("Day 51: bcrypt password hashing")

	password := "s3cret-p@ss"
	hash, err := auth.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash:", hash[:20]+"...")

	if err := auth.CheckPassword(hash, password); err != nil {
		log.Fatal("valid password rejected:", err)
	}
	fmt.Println("Password check: ok")

	if err := auth.CheckPassword(hash, "wrong"); err == nil {
		log.Fatal("invalid password accepted")
	}
	fmt.Println("Wrong password rejected: ok")
}
