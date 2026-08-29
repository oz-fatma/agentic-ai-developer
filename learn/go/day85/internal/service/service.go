package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day85/internal/cache"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day85/internal/queue"
)

// CachedWorkerService combines a TTL cache with a background worker queue.
type CachedWorkerService struct {
	cache  *cache.TTLCache
	queue  *queue.WorkerQueue
	cancel context.CancelFunc
	mu     sync.Mutex
	wait   map[string]chan struct{}
}

func NewCachedWorkerService(workers int, ttl time.Duration) *CachedWorkerService {
	ctx, cancel := context.WithCancel(context.Background())
	wq := queue.NewWorkerQueue(workers, 32)
	queue.Start(ctx, wq)
	return &CachedWorkerService{
		cache:  cache.NewTTLCache(ttl),
		queue:  wq,
		cancel: cancel,
		wait:   make(map[string]chan struct{}),
	}
}

func (s *CachedWorkerService) Close() {
	s.cancel()
}

// Fetch returns a cached value or computes it via the worker queue.
// The second return value is true when the result came from cache.
func (s *CachedWorkerService) Fetch(key string) (string, bool, error) {
	if v, ok := s.cache.Get(key); ok {
		return v, true, nil
	}

	s.mu.Lock()
	if ch, pending := s.wait[key]; pending {
		s.mu.Unlock()
		select {
		case <-ch:
			v, _ := s.cache.Get(key)
			return v, false, nil
		case <-time.After(2 * time.Second):
			return "", false, fmt.Errorf("timeout waiting for %q", key)
		}
	}
	done := make(chan struct{})
	s.wait[key] = done
	s.mu.Unlock()

	s.queue.Submit(func() {
		s.cache.Set(key, fmt.Sprintf("computed:%s", key))
		close(done)
		s.mu.Lock()
		delete(s.wait, key)
		s.mu.Unlock()
	})

	select {
	case <-done:
		v, _ := s.cache.Get(key)
		return v, false, nil
	case <-time.After(2 * time.Second):
		return "", false, fmt.Errorf("compute timeout for %q", key)
	}
}
