# Day 16: Concurrency Basics — Goroutines

**Project:** Study Buddy — concurrent reminders and parallel fetches

## 1. Launch Goroutines

Goroutines are lightweight threads managed by the Go runtime:

```go
func sendReminder(subject string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Printf("Reminder: time to study %s!\n", subject)
}

func main() {
	go sendReminder("Go", 100*time.Millisecond)
	go sendReminder("Math", 200*time.Millisecond)
	go sendReminder("History", 150*time.Millisecond)

	time.Sleep(300 * time.Millisecond) // wait — better pattern tomorrow
}
```

Output order is **non-deterministic** — goroutines run concurrently:

```
Reminder: time to study Go!
Reminder: time to study History!
Reminder: time to study Math!
```

The `go` keyword launches a function concurrently. Main must not exit before goroutines finish.

## 2. Use sync.WaitGroup

Proper way to wait for goroutines:

```go
func fetchCourseInfo(wg *sync.WaitGroup, courseID string) {
	defer wg.Done()
	fmt.Printf("Loading course %s...\n", courseID)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Course %s loaded\n", courseID)
}

func main() {
	courses := []string{"go-101", "math-201", "hist-301"}
	var wg sync.WaitGroup

	for _, id := range courses {
		wg.Add(1)
		go fetchCourseInfo(&wg, id)
	}

	wg.Wait()
	fmt.Println("All courses loaded")
}
```

Pattern: `Add(1)` before launch, `defer Done()` inside goroutine, `Wait()` in main.

## 3. Share Data Safely

Naive counter — **data race**:

```go
func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // UNSAFE — 1000 goroutines, no synchronization
		}()
	}
	wg.Wait()
	fmt.Println(counter) // unpredictable — often < 1000
}
```

Multiple goroutines read and write `counter` without coordination — undefined behavior.

Study Buddy symptom: session count wrong when logging concurrently.

## 4. Run the Race Detector

Enable with `-race` flag:

```bash
go run -race .
```

Output for the naive counter:

```
WARNING: DATA RACE
Write at 0x... by goroutine 7
Previous write at 0x... by goroutine 6
```

Fix options (Days 17–19):
- Channels for communication
- `sync.Mutex` for shared state
- `sync/atomic` for simple counters

```bash
go test -race ./...
```

Run the race detector early and often — it catches bugs that are hard to reproduce.

## Summary

| Concept | Study Buddy application |
|---|---|
| `go func()` | Concurrent course loading, reminders |
| `WaitGroup` | Wait for all fetches before proceeding |
| Data race | Unsafe shared session counter |
| `-race` | Detect races before production |

Goroutines make Study Buddy responsive for I/O-bound work — but shared state requires synchronization.
