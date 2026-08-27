# Day 18: Concurrency Basics — Select and Timeouts

**Project:** Study Buddy — responsive resource fetching with deadlines

## 1. Use select

`select` waits on multiple channel operations — runs the first ready case:

```go
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { time.Sleep(100 * time.Millisecond); ch1 <- "Go resources ready" }()
	go func() { time.Sleep(200 * time.Millisecond); ch2 <- "Math resources ready" }()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
```

**Non-blocking with default:**

```go
select {
case msg := <-notifications:
	fmt.Println(msg)
default:
	fmt.Println("No notifications right now")
}
```

Use `default` sparingly — busy loops burn CPU if misused.

## 2. Add Timeouts

Avoid waiting forever on slow operations:

```go
func fetchStudyMaterial(url string) (string, error) {
	result := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second) // simulate slow fetch
		result <- "Chapter 1: Goroutines"
	}()

	select {
	case data := <-result:
		return data, nil
	case <-time.After(1 * time.Second):
		return "", fmt.Errorf("fetch %s: timed out after 1s", url)
	}
}
```

`time.After` returns a channel that fires after the duration — ideal for one-shot timeouts.

## 3. Cancel Work

Cooperative cancellation with a done channel:

```go
func studyTimer(done <-chan struct{}, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Println("Study session complete!")
	case <-done:
		fmt.Println("Session cancelled")
	}
}

func main() {
	done := make(chan struct{})

	go studyTimer(done, 5*time.Second)

	time.Sleep(1 * time.Second)
	close(done) // broadcast cancel signal
	time.Sleep(100 * time.Millisecond)
}
```

Closing `done` notifies all listeners — multiple goroutines can `<-done`.

Context-based cancellation (preview for HTTP days):

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-ctx.Done():
	return fmt.Errorf("operation cancelled: %w", ctx.Err())
case result := <-workDone:
	return process(result)
}
```

## 4. Handle Backpressure

When consumers are slow, producers must adapt:

```go
jobs := make(chan string, 2) // small buffer — backpressure kicks in early

// Producer blocks when buffer full
go func() {
	for i := 0; i < 10; i++ {
		select {
		case jobs <- fmt.Sprintf("task-%d", i):
			fmt.Println("Queued:", i)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Dropped task", i, "— consumer too slow")
		}
	}
	close(jobs)
}()
```

| Strategy | When to use |
|---|---|
| Block | Consumer will catch up, memory is fine |
| Buffer | Smooth short bursts |
| Drop | Real-time data where stale is useless |
| Timeout | Bound wait time per operation |

Study Buddy applies timeouts to all external fetches — never block forever on a slow API.

## Summary

| Tool | Study Buddy application |
|---|---|
| `select` | Wait on multiple resource fetches |
| `time.After` | Timeout slow study material downloads |
| Done channel | Cancel active study timer |
| Backpressure | Drop or buffer when grading queue is full |

Responsive Study Buddy services degrade gracefully under load instead of hanging.
