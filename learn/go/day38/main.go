package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day38/middleware"
)

func main() {
	fmt.Println("Day 38: middleware chaining with request ID in context")

	handler := middleware.Chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := middleware.RequestIDFromContext(r.Context())
			fmt.Fprintf(w, "request id: %s", id)
		}),
		middleware.RequestID,
		middleware.Logging,
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

	fmt.Printf("Response: %q\n", string(body))
	fmt.Println("X-Request-ID header:", resp.Header.Get("X-Request-ID"))
}
