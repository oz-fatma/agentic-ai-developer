package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day53/auth"
)

func main() {
	fmt.Println("Day 53: RBAC middleware")

	mux := http.NewServeMux()
	mux.Handle("GET /public", http.HandlerFunc(publicHandler))
	mux.Handle("GET /admin", auth.RequireRole("admin", http.HandlerFunc(adminHandler)))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, tc := range []struct {
		path  string
		role  string
		want  int
		label string
	}{
		{"/public", "", http.StatusOK, "public no role"},
		{"/admin", "admin", http.StatusOK, "admin role"},
		{"/admin", "reader", http.StatusForbidden, "reader denied"},
		{"/admin", "", http.StatusUnauthorized, "missing role"},
	} {
		req, _ := http.NewRequest(http.MethodGet, ts.URL+tc.path, nil)
		if tc.role != "" {
			req.Header.Set("X-Role", tc.role)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(tc.label, "error:", err)
			continue
		}
		resp.Body.Close()
		fmt.Printf("%s -> %d\n", tc.label, resp.StatusCode)
	}
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "public ok")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "admin ok")
}
