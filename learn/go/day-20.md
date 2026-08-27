# Day 20: Concurrency Basics — Practice

**Project:** Study Buddy — concurrent downloader, pipeline, and race-free code

## 1. Concurrent Downloader

Fetch several study resource URLs concurrently:

```go
type FetchResult struct {
	URL  string
	Body string
	Err  error
}

func fetchURL(wg *sync.WaitGroup, url string, results chan<- FetchResult) {
	defer wg.Done()

	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)

	results <- FetchResult{
		URL:  url,
		Body: fmt.Sprintf("Content from %s", url),
	}
}

func main() {
	urls := []string{
		"https://go.dev/doc",
		"https://go.dev/blog",
		"https://go.dev/play",
	}

	results := make(chan FetchResult, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(&wg, url, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		if r.Err != nil {
			fmt.Println("Error:", r.URL, r.Err)
		} else {
			fmt.Println("Fetched:", r.URL)
		}
	}
}
```

Fan-out: one goroutine per URL. Fan-in: single results channel.

## 2. Pipeline

Three-stage pipeline for Study Buddy quiz scores:

```go
func generateScores(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			out <- 60 + rand.Intn(40) // scores 60–99
		}
		close(out)
	}()
	return out
}

func squareScores(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for score := range in {
			out <- score * score
		}
		close(out)
	}()
	return out
}

func printResults(in <-chan int) {
	for squared := range in {
		fmt.Println("Squared score:", squared)
	}
}

func main() {
	scores := generateScores(5)
	squared := squareScores(scores)
	printResults(squared)
}
```

```
generate → square → print
  (stage1)   (stage2)  (stage3)
```

Each stage runs in its own goroutine, connected by channels.

## 3. Timeout Fetch

Added timeout to network-style operations:

```go
func fetchWithTimeout(url string, timeout time.Duration) (string, error) {
	result := make(chan FetchResult, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		result <- FetchResult{URL: url, Body: "data"}
	}()

	select {
	case r := <-result:
		return r.Body, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("fetch %s: timed out after %v", url, timeout)
	}
}

func main() {
	data, err := fetchWithTimeout("https://slow.example.com", 100*time.Millisecond)
	if err != nil {
		fmt.Println("Error:", err) // timed out
	}
}
```

Even in exercises, blocking forever is a common beginner mistake — always bound waits.

## 4. Race Check

Ran all solutions with the race detector:

```bash
go run -race .
go test -race ./...
```

**Issue found:** Initial downloader wrote to a shared `[]FetchResult` slice from multiple goroutines.

**Fix:** Use a channel to collect results (each goroutine sends to channel — no shared slice mutation).

After fix:

```bash
go run -race .
# no race detected
```

Checklist:
- No unsynchronized shared variables
- Channels or mutexes protect all shared state
- `WaitGroup` waits before exit
- Channels closed by sender only

## Summary

**Phase checkpoint:** Concurrency Basics complete.

Study Buddy concurrent programs now use:
- Goroutines with `WaitGroup` for parallel fetches
- Channel pipelines for staged processing
- `select` with timeouts for bounded waits
- Mutex/atomic where appropriate
- Clean `-race` output

Ready for testing fundamentals in Days 21–25.
