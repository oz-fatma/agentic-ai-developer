package main

import (
	"fmt"
	"log"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day52/auth"
)

func main() {
	fmt.Println("Day 52: JWT create and validate")

	secret := []byte("day52-demo-secret")
	token, err := auth.CreateToken("user-42", "reader", secret, time.Hour)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token:", token[:32]+"...")

	claims, err := auth.ParseToken(token, secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Claims: sub=%s role=%s exp=%v\n", claims.Subject, claims.Role, claims.ExpiresAt.Time)

	if _, err := auth.ParseToken(token, []byte("wrong-secret")); err == nil {
		log.Fatal("expected invalid signature error")
	}
	fmt.Println("Invalid signature rejected: ok")
}
