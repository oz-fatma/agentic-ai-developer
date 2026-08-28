package repository

import (
	"database/sql"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day48/internal/db"
)

func setupTestRepo(t *testing.T) *TagRepo {
	t.Helper()
	conn := db.OpenTempDB(t)
	if err := db.Migrate(conn); err != nil {
		t.Fatal(err)
	}
	return NewTagRepo(conn)
}

func TestTagRepoCreateAndGet(t *testing.T) {
	repo := setupTestRepo(t)

	created, err := repo.Create("database")
	if err != nil {
		t.Fatal(err)
	}
	if created.ID == 0 || created.Name != "database" {
		t.Fatalf("unexpected tag: %+v", created)
	}

	got, err := repo.GetByName("database")
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != created.ID {
		t.Fatalf("id mismatch: got %d want %d", got.ID, created.ID)
	}
}

func TestTagRepoDuplicateName(t *testing.T) {
	repo := setupTestRepo(t)
	if _, err := repo.Create("sql"); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.Create("sql"); err == nil {
		t.Fatal("expected duplicate error")
	}
}

func TestTagRepoList(t *testing.T) {
	repo := setupTestRepo(t)
	for _, name := range []string{"b", "a", "c"} {
		if _, err := repo.Create(name); err != nil {
			t.Fatal(err)
		}
	}
	tags, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) != 3 {
		t.Fatalf("want 3 tags, got %d", len(tags))
	}
	if tags[0].Name != "a" || tags[2].Name != "c" {
		t.Fatalf("unexpected order: %+v", tags)
	}
}

func TestTagRepoGetNotFound(t *testing.T) {
	repo := setupTestRepo(t)
	_, err := repo.GetByName("missing")
	if err != sql.ErrNoRows {
		t.Fatalf("want ErrNoRows, got %v", err)
	}
}
