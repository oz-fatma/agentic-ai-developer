package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day35/internal/api"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day35/internal/store"
)

func main() {
	fmt.Println("Day 35: polished MVP with tests")

	ts := httptest.NewServer(api.NewHandler(store.NewMemoryStore()))
	defer ts.Close()

	resp, _ := http.Get(ts.URL + "/health")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("GET /health -> %s\n", string(body))

	resp, _ = http.Post(ts.URL+"/tasks", "application/json", strings.NewReader(`{"title":"ship MVP"}`))
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST /tasks -> %s\n", string(body))

	fmt.Println("Run: go test ./day35/...")
}
