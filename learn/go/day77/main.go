package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	fmt.Println("Day 77: docker-compose demo")
	fmt.Println("Run: docker compose -f day77/docker-compose.yml up --build")
	fmt.Println("Local demo uses httptest (no blocking server in main)")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Day 77 app OK")
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET / -> %d %s", resp.StatusCode, body)
}
