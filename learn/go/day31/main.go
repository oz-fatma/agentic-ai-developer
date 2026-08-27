package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day31/internal/api"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day31/internal/models"
)

func main() {
	fmt.Println("Day 31: MVP scaffold (cmd layout plan)")
	fmt.Println("  internal/models ->", models.Task{})
	fmt.Println("  internal/api    -> health endpoint only")

	ts := httptest.NewServer(api.NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("GET /health -> %d %q\n", resp.StatusCode, string(body))
}
