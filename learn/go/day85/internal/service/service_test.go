package service

import (
	"testing"
	"time"
)

func TestCachedWorkerServiceFetch(t *testing.T) {
	svc := NewCachedWorkerService(2, time.Minute)
	defer svc.Close()

	v1, cached, err := svc.Fetch("item-a")
	if err != nil {
		t.Fatal(err)
	}
	if cached {
		t.Fatal("first fetch should miss cache")
	}
	if v1 != "computed:item-a" {
		t.Fatalf("got %q", v1)
	}

	v2, cached, err := svc.Fetch("item-a")
	if err != nil {
		t.Fatal(err)
	}
	if !cached {
		t.Fatal("second fetch should hit cache")
	}
	if v2 != v1 {
		t.Fatalf("cache mismatch: %q vs %q", v2, v1)
	}
}

func TestCachedWorkerServiceConcurrentFetch(t *testing.T) {
	svc := NewCachedWorkerService(2, time.Minute)
	defer svc.Close()

	type result struct {
		val    string
		cached bool
		err    error
	}
	ch := make(chan result, 2)
	for i := 0; i < 2; i++ {
		go func() {
			v, cached, err := svc.Fetch("shared")
			ch <- result{v, cached, err}
		}()
	}
	r1 := <-ch
	r2 := <-ch
	if r1.err != nil || r2.err != nil {
		t.Fatalf("errors: %v %v", r1.err, r2.err)
	}
	if r1.val != "computed:shared" || r2.val != "computed:shared" {
		t.Fatalf("values mismatch: %q %q", r1.val, r2.val)
	}
	cachedHits := 0
	if r1.cached {
		cachedHits++
	}
	if r2.cached {
		cachedHits++
	}
	if cachedHits > 1 {
		t.Fatal("expected at most one cache hit for concurrent first fetch")
	}
}
