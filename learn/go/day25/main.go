package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day25/calc"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day25/server"
)

func main() {
	fmt.Println("Day 25: testing practice")

	result, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("calc error:", err)
		return
	}
	fmt.Println("10 / 2 =", result)
	fmt.Println("2^10 =", calc.Pow(2, 10))

	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(`{"message":"hello"}`))
	rec := httptest.NewRecorder()
	server.EchoHandler(rec, req)
	fmt.Printf("POST /echo -> %d %s", rec.Code, rec.Body.String())

	fmt.Println("\nRun: go test ./day25/...")
}
