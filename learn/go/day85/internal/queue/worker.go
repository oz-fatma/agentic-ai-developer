package queue

import (
	"context"
	"sync"
)

type Job func()

type WorkerQueue struct {
	jobs chan Job
	wg   sync.WaitGroup
}

func NewWorkerQueue(workers, backlog int) *WorkerQueue {
	q := &WorkerQueue{jobs: make(chan Job, backlog)}
	for i := 0; i < workers; i++ {
		q.wg.Add(1)
		go func() {
			defer q.wg.Done()
			for job := range q.jobs {
				job()
			}
		}()
	}
	return q
}

func (q *WorkerQueue) Submit(job Job) { q.jobs <- job }

func (q *WorkerQueue) Close() {
	close(q.jobs)
	q.wg.Wait()
}

func Start(ctx context.Context, q *WorkerQueue) {
	go func() {
		<-ctx.Done()
		q.Close()
	}()
}
