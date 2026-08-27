package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	fmt.Println("Day 24: httptest handler tests")

	req := httptest.NewRequest(http.MethodGet, "/greet?name=Go", nil)
	rec := httptest.NewRecorder()
	greetHandler(rec, req)
	fmt.Println("GET /greet?name=Go ->", rec.Body.String())

	req = httptest.NewRequest(http.MethodGet, "/health", nil)
	rec = httptest.NewRecorder()
	healthHandler(rec, req)
	fmt.Println("GET /health ->", rec.Body.String())

	fmt.Println("Run: go test ./day24/...")
}
