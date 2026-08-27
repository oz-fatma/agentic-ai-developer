package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
)

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type noteStore struct {
	mu    sync.Mutex
	next  int
	notes []Note
}

func (s *noteStore) add(content string) Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	n := Note{ID: s.next, Content: content}
	s.notes = append(s.notes, n)
	return n
}

func (s *noteStore) list() []Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Note, len(s.notes))
	copy(out, s.notes)
	return out
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func newNotesAPI(store *noteStore) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /notes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.list())
	})
	mux.HandleFunc("POST /notes", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		n := store.add(body.Content)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(n)
	})
	return logging(mux)
}

func main() {
	fmt.Println("Day 30: in-memory notes API with middleware")

	store := &noteStore{}
	ts := httptest.NewServer(newNotesAPI(store))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/notes", "application/json", strings.NewReader(`{"content":"buy milk"}`))
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("POST /notes -> %d %s\n", resp.StatusCode, string(body))

	resp, _ = http.Get(ts.URL + "/notes")
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("GET /notes -> %s\n", string(body))
}
