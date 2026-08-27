# Day 17: Concurrency Basics — Channels

**Project:** Study Buddy — worker pattern for quiz grading

## 1. Create Channels

Unbuffered channels synchronize sender and receiver at the handoff:

```go
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Session complete: Go, 60 min"
	}()

	msg := <-ch
	fmt.Println(msg)
}
```

Send blocks until receive is ready — and vice versa. This synchronizes goroutines.

## 2. Buffered Channels

Buffer decouples producer and consumer:

```go
jobs := make(chan string, 3) // buffer size 3

jobs <- "Grade quiz 1"
jobs <- "Grade quiz 2"
jobs <- "Grade quiz 3"
// 4th send would block until someone receives

result := <-jobs
fmt.Println(result)
```

Use small buffers to smooth bursts — large buffers hide backpressure problems.

## 3. Close and Range

Close signals no more values; receivers drain remaining:

```go
func producer(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func main() {
	nums := make(chan int)
	go producer(nums)

	for n := range nums {
		fmt.Println("Received:", n)
	}
	// loop exits when channel closed and drained
}
```

Rules:
- Only the **sender** should close
- Sending on a closed channel panics
- Receiving from closed channel returns zero value immediately

## 4. Coordinate Workers

Study Buddy quiz grading worker pool:

```go
type QuizResult struct {
	StudentID string
	Score     int
}

func worker(id int, jobs <-chan string, results chan<- QuizResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for studentID := range jobs {
		time.Sleep(50 * time.Millisecond) // simulate grading
		score := 70 + id*5               // pretend scoring
		results <- QuizResult{StudentID: studentID, Score: score}
	}
}

func main() {
	jobs := make(chan string, 10)
	results := make(chan QuizResult, 10)

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	students := []string{"s1", "s2", "s3", "s4", "s5"}
	for _, id := range students {
		jobs <- id
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Student %s scored %d\n", r.StudentID, r.Score)
	}
}
```

Channel direction types enforce usage:
- `chan<- T` — send-only
- `<-chan T` — receive-only

## Summary

| Pattern | Study Buddy use |
|---|---|
| Unbuffered channel | Synchronize session notifications |
| Buffered channel | Queue quiz jobs |
| `close` + `range` | Drain results cleanly |
| Worker pool | Parallel quiz grading |

Channels embody Go's motto: *share memory by communicating*.
