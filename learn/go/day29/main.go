package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day29/middleware"
)

func main() {
	fmt.Println("Day 29: middleware (logging, recovery, chain)")

	handler := middleware.Chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "ok")
		}),
		middleware.Recovery,
		middleware.Logging,
		middleware.RequestID,
	)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("GET / -> %d %q\n", resp.StatusCode, string(body))
	fmt.Println("X-Request-ID:", resp.Header.Get("X-Request-ID"))
}
