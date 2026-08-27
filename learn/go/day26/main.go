package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from net/http!")
}

func main() {
	fmt.Println("Day 26: net/http basics (httptest demo, no blocking server)")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Printf("GET %s/ -> %d %q\n", ts.URL, resp.StatusCode, string(body))
}
