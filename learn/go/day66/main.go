package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day66/internal/store"
)

func main() {
	fmt.Println("Day 66: integration tests with httptest")
	fmt.Println("Run: go test ./day66/...")

	s := store.New()
	ts := httptest.NewServer(store.NewHandler(s))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/items", "application/json", strings.NewReader(`{"title":"demo item"}`))
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST /items -> %d %s\n", resp.StatusCode, body)
}
