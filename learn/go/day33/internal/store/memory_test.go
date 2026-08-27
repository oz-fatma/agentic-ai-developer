package store

import (
	"sync"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day33/internal/models"
)

func TestMemoryStoreConcurrent(t *testing.T) {
	s := NewMemoryStore()
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			s.Create("task")
			s.List()
		}(i)
	}
	wg.Wait()
	if len(s.List()) != 20 {
		t.Fatalf("expected 20 tasks, got %d", len(s.List()))
	}
}

func TestTaskStoreInterface(t *testing.T) {
	var s TaskStore = NewMemoryStore()
	t1 := s.Create("a")
	got, ok := s.Get(t1.ID)
	if !ok || got.Title != "a" {
		t.Fatalf("Get returned %+v, ok=%v", got, ok)
	}
	done := true
	updated, ok := s.Update(t1.ID, nil, &done)
	if !ok || !updated.Done {
		t.Fatalf("Update failed: %+v", updated)
	}
	if !s.Delete(t1.ID) {
		t.Fatal("Delete failed")
	}
	_, ok = s.Get(t1.ID)
	if ok {
		t.Fatal("task should be deleted")
	}
}

func ExampleMemoryStore() {
	var s TaskStore = NewMemoryStore()
	t := s.Create("demo")
	_ = models.Task(t)
}
