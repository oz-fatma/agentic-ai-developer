package repository

import (
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day50/internal/db"
)

func setupRepo(t *testing.T) TaskRepository {
	t.Helper()
	conn := db.OpenTempDB(t)
	if err := db.Migrate(conn); err != nil {
		t.Fatal(err)
	}
	return NewTaskRepo(conn)
}

func TestTaskRepoCRUD(t *testing.T) {
	repo := setupRepo(t)

	created, err := repo.Create("learn sql", false)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.Get(created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "learn sql" || got.Done {
		t.Fatalf("unexpected task: %+v", got)
	}

	done := true
	updated, err := repo.Update(created.ID, nil, &done)
	if err != nil {
		t.Fatal(err)
	}
	if !updated.Done {
		t.Fatal("expected done=true")
	}

	title := "learn repositories"
	updated, err = repo.Update(created.ID, &title, nil)
	if err != nil {
		t.Fatal(err)
	}
	if updated.Title != title {
		t.Fatalf("title = %q", updated.Title)
	}

	if err := repo.Delete(created.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.Get(created.ID); err != ErrNotFound {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestTaskRepoList(t *testing.T) {
	repo := setupRepo(t)
	for _, title := range []string{"a", "b"} {
		if _, err := repo.Create(title, false); err != nil {
			t.Fatal(err)
		}
	}
	tasks, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Fatalf("want 2 tasks, got %d", len(tasks))
	}
}

func TestTaskRepoNotFound(t *testing.T) {
	repo := setupRepo(t)
	if _, err := repo.Get(99); err != ErrNotFound {
		t.Fatalf("Get: want ErrNotFound, got %v", err)
	}
	if err := repo.Delete(99); err != ErrNotFound {
		t.Fatalf("Delete: want ErrNotFound, got %v", err)
	}
}
