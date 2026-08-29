package tune

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

func Process(workers, jobs int) (time.Duration, int32) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobsCh := make(chan int, jobs)
	var processed atomic.Int32
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case _, ok := <-jobsCh:
					if !ok {
						return
					}
					time.Sleep(time.Millisecond)
					processed.Add(1)
				}
			}
		}()
	}

	start := time.Now()
	for j := 0; j < jobs; j++ {
		jobsCh <- j
	}
	close(jobsCh)
	wg.Wait()
	return time.Since(start), processed.Load()
}
