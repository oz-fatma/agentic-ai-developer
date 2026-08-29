package perf

import (
	"strings"
	"sync"
)

// Summarize builds a report from string fragments using a worker pool.
func Summarize(parts []string, workers int) string {
	if workers < 1 {
		workers = 1
	}
	jobs := make(chan int, len(parts))
	var wg sync.WaitGroup
	var mu sync.Mutex
	rows := make([]string, len(parts))

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				rows[idx] = strings.ToUpper(parts[idx])
			}
		}()
	}

	for i := range parts {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	var b strings.Builder
	b.Grow(len(parts) * 8)
	mu.Lock()
	defer mu.Unlock()
	for _, row := range rows {
		b.WriteString(row)
		b.WriteByte('\n')
	}
	return b.String()
}
