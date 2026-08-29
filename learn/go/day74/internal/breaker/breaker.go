package breaker

import (
	"errors"
	"sync"
	"time"
)

var ErrOpen = errors.New("circuit breaker open")

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type Breaker struct {
	mu           sync.Mutex
	state        State
	failures     int
	maxFailures  int
	openUntil    time.Time
	resetTimeout time.Duration
}

func New(maxFailures int, resetTimeout time.Duration) *Breaker {
	return &Breaker{maxFailures: maxFailures, resetTimeout: resetTimeout, state: Closed}
}

func (b *Breaker) Call(fn func() error) error {
	b.mu.Lock()
	switch b.state {
	case Open:
		if time.Now().Before(b.openUntil) {
			b.mu.Unlock()
			return ErrOpen
		}
		b.state = HalfOpen
	case HalfOpen, Closed:
	}
	b.mu.Unlock()

	err := fn()

	b.mu.Lock()
	defer b.mu.Unlock()
	if err != nil {
		b.failures++
		if b.failures >= b.maxFailures {
			b.state = Open
			b.openUntil = time.Now().Add(b.resetTimeout)
		}
		return err
	}
	b.failures = 0
	b.state = Closed
	return nil
}

func (b *Breaker) State() State {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}
