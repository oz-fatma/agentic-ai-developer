package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func methodCheck(allowed string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowed {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func main() {
	fmt.Println("Day 27: ServeMux routing and method checks")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "healthy")
	})
	mux.HandleFunc("POST /items", methodCheck(http.MethodPost, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "created")
	}))
	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item path: %s", r.URL.Path)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	routes := []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/health"},
		{http.MethodPost, "/items"},
		{http.MethodGet, "/items/42"},
		{http.MethodGet, "/items"},
	}

	for _, route := range routes {
		req, _ := http.NewRequest(route.method, ts.URL+route.path, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s %s -> %d %q\n", route.method, route.path, resp.StatusCode, string(body))
	}
}
