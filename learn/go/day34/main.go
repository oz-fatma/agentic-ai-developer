package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day34/internal/api"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day34/internal/store"
)

func main() {
	fmt.Println("Day 34: validation and error responses")

	ts := httptest.NewServer(api.NewHandler(store.NewMemoryStore()))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/tasks", "application/json", strings.NewReader(`{"title":""}`))
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST empty title -> %d %s\n", resp.StatusCode, string(body))

	resp, _ = http.Post(ts.URL+"/tasks", "application/json", strings.NewReader(`{"title":"valid task"}`))
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST valid -> %d %s\n", resp.StatusCode, string(body))

	resp, _ = http.Get(ts.URL + "/tasks/99")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("GET missing -> %d %s\n", resp.StatusCode, string(body))
}
