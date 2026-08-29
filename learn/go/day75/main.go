package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day75/internal/fetch"
)

func main() {
	fmt.Println("Day 75: context timeouts")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		fmt.Fprint(w, "slow response")
	}))
	defer ts.Close()

	ctx, cancel := fetch.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := fetch.Get(ctx, nil, ts.URL)
	if err != nil {
		fmt.Println("Expected timeout:", err)
	}

	fast := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	}))
	defer fast.Close()

	ctx2, cancel2 := fetch.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	body, err := fetch.Get(ctx2, nil, fast.URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fast response:", body)
}
