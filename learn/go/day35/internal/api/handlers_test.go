package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day35/internal/store"
)

func TestCreateAndGetTask(t *testing.T) {
	srv := httptest.NewServer(NewHandler(store.NewMemoryStore()))
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/tasks", "application/json", strings.NewReader(`{"title":"test task"}`))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201", resp.StatusCode)
	}
	resp.Body.Close()

	resp, err = http.Get(srv.URL + "/tasks/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
}

func TestValidationError(t *testing.T) {
	srv := httptest.NewServer(NewHandler(store.NewMemoryStore()))
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/tasks", "application/json", strings.NewReader(`{"title":""}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", resp.StatusCode)
	}
}

func TestNotFound(t *testing.T) {
	srv := httptest.NewServer(NewHandler(store.NewMemoryStore()))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/tasks/404")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", resp.StatusCode)
	}
}

func TestUpdateAndDelete(t *testing.T) {
	srv := httptest.NewServer(NewHandler(store.NewMemoryStore()))
	defer srv.Close()

	http.Post(srv.URL+"/tasks", "application/json", strings.NewReader(`{"title":"x"}`))

	req, _ := http.NewRequest(http.MethodPut, srv.URL+"/tasks/1", bytes.NewBufferString(`{"done":true}`))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("update status = %d", resp.StatusCode)
	}
	resp.Body.Close()

	req, _ = http.NewRequest(http.MethodDelete, srv.URL+"/tasks/1", nil)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("delete status = %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestHealth(t *testing.T) {
	srv := httptest.NewServer(NewHandler(store.NewMemoryStore()))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
}
