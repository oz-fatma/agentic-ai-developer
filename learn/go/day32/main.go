package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day32/internal/api"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day32/internal/store"
)

func main() {
	fmt.Println("Day 32: CRUD endpoints for tasks")

	ts := httptest.NewServer(api.NewHandler(store.NewMemoryStore()))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/tasks", "application/json", strings.NewReader(`{"title":"learn Go"}`))
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST /tasks -> %d %s\n", resp.StatusCode, string(body))

	resp, _ = http.Get(ts.URL + "/tasks")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("GET /tasks -> %s\n", string(body))

	resp, _ = http.Get(ts.URL + "/tasks/1")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("GET /tasks/1 -> %s\n", string(body))
}
