package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/auth"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/db"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/repository"
)

func setupServer(t *testing.T) *httptest.Server {
	t.Helper()
	conn := db.OpenTempDB(t)
	if err := db.Migrate(conn); err != nil {
		t.Fatal(err)
	}
	svc := auth.NewService(repository.NewUserRepo(conn), []byte("test-secret"))
	return httptest.NewServer(NewHandler(svc))
}

func TestRegisterAndLogin(t *testing.T) {
	srv := setupServer(t)
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/register", "application/json", strings.NewReader(`{"email":"a@test.com","password":"pass1234"}`))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("register status=%d body=%s", resp.StatusCode, body)
	}

	resp, err = http.Post(srv.URL+"/login", "application/json", strings.NewReader(`{"email":"a@test.com","password":"pass1234"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login status=%d", resp.StatusCode)
	}
}

func TestLoginInvalidPassword(t *testing.T) {
	srv := setupServer(t)
	defer srv.Close()

	http.Post(srv.URL+"/register", "application/json", strings.NewReader(`{"email":"b@test.com","password":"pass1234"}`))

	resp, err := http.Post(srv.URL+"/login", "application/json", strings.NewReader(`{"email":"b@test.com","password":"wrong"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status=%d want 401", resp.StatusCode)
	}
}

func TestMeWithToken(t *testing.T) {
	srv := setupServer(t)
	defer srv.Close()

	http.Post(srv.URL+"/register", "application/json", strings.NewReader(`{"email":"c@test.com","password":"pass1234"}`))
	resp, _ := http.Post(srv.URL+"/login", "application/json", strings.NewReader(`{"email":"c@test.com","password":"pass1234"}`))
	loginBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !strings.Contains(string(loginBody), "token") {
		t.Fatalf("missing token: %s", loginBody)
	}

	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/me", nil)
	req.Header.Set("Authorization", "Bearer "+extractToken(string(loginBody)))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("me status=%d", resp.StatusCode)
	}
}

func extractToken(body string) string {
	const key = `"token":"`
	i := strings.Index(body, key)
	if i < 0 {
		return ""
	}
	rest := body[i+len(key):]
	end := strings.Index(rest, `"`)
	if end < 0 {
		return ""
	}
	return rest[:end]
}
