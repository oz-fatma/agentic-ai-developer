package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day54/limit"
)

func main() {
	fmt.Println("Day 54: rate limiting")

	rl := limit.NewRateLimiter(3, time.Second)
	handler := rl.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}))

	ts := httptest.NewServer(handler)
	defer ts.Close()

	for i := 1; i <= 5; i++ {
		resp, err := http.Get(ts.URL)
		if err != nil {
			fmt.Println("request", i, "error:", err)
			continue
		}
		fmt.Printf("request %d -> %d\n", i, resp.StatusCode)
		resp.Body.Close()
	}
}
