# Day 19: Concurrency Basics — sync Package Primitives

**Project:** Study Buddy — protecting shared stats and lazy initialization

## 1. Use Mutex

Protect shared counters with mutual exclusion:

```go
type SessionTracker struct {
	mu       sync.Mutex
	sessions int
	minutes  int
}

func (t *SessionTracker) RecordSession(mins int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.sessions++
	t.minutes += mins
}

func (t *SessionTracker) Stats() (int, int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.sessions, t.minutes
}
```

`defer Unlock()` ensures the lock is released even if the function panics.

**RWMutex** for read-heavy caches:

```go
type CourseCache struct {
	mu      sync.RWMutex
	courses map[string]Course
}

func (c *CourseCache) Get(id string) (Course, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	course, ok := c.courses[id]
	return course, ok
}

func (c *CourseCache) Set(id string, course Course) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.courses[id] = course
}
```

Many readers OR one writer — efficient for Study Buddy's course catalog cache.

## 2. Try Once

Initialize expensive resources exactly once:

```go
var (
	dbInstance *sql.DB
	dbOnce     sync.Once
)

func GetDB() *sql.DB {
	dbOnce.Do(func() {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}
		dbInstance = db
	})
	return dbInstance
}
```

Even with 100 concurrent callers, `Do` runs the function exactly once.

Study Buddy use: lazy-init JSON config loader, database connection, HTTP client.

## 3. Use atomic

Simple counters without full mutex overhead:

```go
var activeSessions int64

func startSession() {
	atomic.AddInt64(&activeSessions, 1)
}

func endSession() {
	atomic.AddInt64(&activeSessions, -1)
}

func getActiveCount() int64 {
	return atomic.LoadInt64(&activeSessions)
}
```

Atomics suit numeric stats. For struct maps or compound updates, use mutex.

## 4. Compare Approaches

| Scenario | Best tool | Why |
|---|---|---|
| Pass data between goroutines | Channel | Clear ownership transfer |
| Protect shared map/counter | Mutex | Simple critical section |
| Read-heavy cache | RWMutex | Parallel reads |
| Simple int counter | atomic | Lower overhead |
| One-time init | sync.Once | Guaranteed single execution |
| Pipeline stages | Channels | Natural stage coupling |

Study Buddy guidelines:
- **Channels** for job queues (quiz grading pipeline)
- **Mutex** for session tracker with multiple fields
- **atomic** for active session count displayed on dashboard
- **sync.Once** for database and config initialization

Go developers mix both pragmatically — choose what makes correctness obvious.

## Summary

Not every concurrency problem needs a channel. Mutexes, atomics, and `sync.Once` are straightforward tools for Study Buddy's shared in-memory state.
