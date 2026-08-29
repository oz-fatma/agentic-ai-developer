package store_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day66/internal/store"
)

func setupServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(store.NewHandler(store.New()))
}

func TestIntegrationCreateAndGet(t *testing.T) {
	ts := setupServer(t)
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/items", "application/json", bytes.NewReader([]byte(`{"title":"integration test"}`)))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("POST status = %d", resp.StatusCode)
	}

	var created store.Item
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatal(err)
	}
	if created.ID != 1 || created.Title != "integration test" {
		t.Fatalf("unexpected item: %+v", created)
	}

	resp, err = http.Get(ts.URL + "/items/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET status = %d body = %s", resp.StatusCode, body)
	}
}

func TestIntegrationList(t *testing.T) {
	ts := setupServer(t)
	defer ts.Close()

	for _, title := range []string{"a", "b"} {
		http.Post(ts.URL+"/items", "application/json", bytes.NewReader([]byte(`{"title":"`+title+`"}`)))
	}

	resp, err := http.Get(ts.URL + "/items")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var items []store.Item
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 items, got %d", len(items))
	}
}
