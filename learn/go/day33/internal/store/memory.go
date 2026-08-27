package store

import (
	"sync"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day33/internal/models"
)

type MemoryStore struct {
	mu    sync.RWMutex
	next  int
	tasks map[int]models.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tasks: make(map[int]models.Task)}
}

func (s *MemoryStore) Create(title string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	t := models.Task{ID: s.next, Title: title, CreatedAt: time.Now()}
	s.tasks[t.ID] = t
	return t
}

func (s *MemoryStore) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]models.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out
}

func (s *MemoryStore) Get(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	return t, ok
}

func (s *MemoryStore) Update(id int, title *string, done *bool) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tasks[id]
	if !ok {
		return models.Task{}, false
	}
	if title != nil {
		t.Title = *title
	}
	if done != nil {
		t.Done = *done
	}
	s.tasks[id] = t
	return t, true
}

func (s *MemoryStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return false
	}
	delete(s.tasks, id)
	return true
}

var _ TaskStore = (*MemoryStore)(nil)
