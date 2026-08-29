package queue

import (
	"context"
	"sync"
)

type Job func()

// WorkerQueue processes jobs with a fixed pool of goroutines.
type WorkerQueue struct {
	jobs    chan Job
	wg      sync.WaitGroup
	workers int
}

func NewWorkerQueue(workers, backlog int) *WorkerQueue {
	return &WorkerQueue{
		jobs:    make(chan Job, backlog),
		workers: workers,
	}
}

func (q *WorkerQueue) Start(ctx context.Context) {
	for i := 0; i < q.workers; i++ {
		q.wg.Add(1)
		go func(id int) {
			defer q.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-q.jobs:
					if !ok {
						return
					}
					job()
				}
			}
		}(i)
	}
}

func (q *WorkerQueue) Submit(job Job) {
	q.jobs <- job
}

func (q *WorkerQueue) CloseAndWait() {
	close(q.jobs)
	q.wg.Wait()
}

func (q *WorkerQueue) Workers() int { return q.workers }
