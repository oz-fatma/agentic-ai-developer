package store

import (
	"sync"
)

type Item struct {
	ID    int
	Title string
}

type Store struct {
	mu    sync.Mutex
	next  int
	items map[int]Item
}

func New() *Store {
	return &Store{items: make(map[int]Item)}
}

func (s *Store) Add(title string) Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	item := Item{ID: s.next, Title: title}
	s.items[item.ID] = item
	return item
}

func (s *Store) Get(id int) (Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	return item, ok
}

func (s *Store) List() []Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		out = append(out, item)
	}
	return out
}
