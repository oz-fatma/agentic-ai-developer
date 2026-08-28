package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/api"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/auth"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/db"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/repository"
)

func main() {
	fmt.Println("Day 55: secure MVP")
	fmt.Println("Run: go test ./day55/...")

	conn, err := db.OpenMemory()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if err := db.Migrate(conn); err != nil {
		panic(err)
	}

	svc := auth.NewService(repository.NewUserRepo(conn), []byte("day55-demo-secret"))
	ts := httptest.NewServer(api.NewHandler(svc))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/register", "application/json", strings.NewReader(`{"email":"demo@example.com","password":"secret123"}`))
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST /register -> %d %s\n", resp.StatusCode, body)

	resp, _ = http.Post(ts.URL+"/login", "application/json", strings.NewReader(`{"email":"demo@example.com","password":"secret123"}`))
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST /login -> %d %s\n", resp.StatusCode, body)
}
