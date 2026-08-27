# Day 11: Errors, Collections & I/O — Error Values and Handling

**Project:** Study Buddy — idiomatic error handling for sessions and courses

## 1. Return Errors

Go functions return `(T, error)` — callers must check every error:

```go
var (
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidDuration = errors.New("duration must be positive")
)

func GetSession(id string, sessions map[string]StudySession) (StudySession, error) {
	s, ok := sessions[id]
	if !ok {
		return StudySession{}, ErrSessionNotFound
	}
	return s, nil
}

func ValidateDuration(minutes int) error {
	if minutes <= 0 {
		return ErrInvalidDuration
	}
	return nil
}
```

Call site — always check:

```go
session, err := GetSession("s1", sessions)
if err != nil {
	fmt.Println("Failed:", err)
	return
}
fmt.Println(session.Subject)
```

## 2. Create Error Values

Two main constructors:

```go
// Static sentinel errors
var ErrCourseNotFound = errors.New("course not found")

// Dynamic errors with context
func LoadCourse(id string) (*Course, error) {
	if id == "" {
		return nil, fmt.Errorf("load course: empty id")
	}
	// ...
	return nil, fmt.Errorf("load course %q: not in database", id)
}
```

Sentinel errors (`ErrCourseNotFound`) are for known, checkable conditions. `fmt.Errorf` adds context for specific failures.

## 3. Wrap Errors

Use `%w` to wrap errors while preserving the chain:

```go
func OpenSessionFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open session file %q: %w", path, err)
	}
	return data, nil
}
```

Inspect the chain:

```go
_, err := OpenSessionFile("missing.json")
if errors.Is(err, os.ErrNotExist) {
	fmt.Println("File does not exist")
}

var pathErr *os.PathError
if errors.As(err, &pathErr) {
	fmt.Println("Failed path:", pathErr.Path)
}
```

| Function | Purpose |
|---|---|
| `errors.Is(err, target)` | Check if error matches sentinel in chain |
| `errors.As(err, &target)` | Extract typed error from chain |

Avoid string matching on `err.Error()` — use `Is` and `As`.

## 4. Fail Fast

Early returns keep handlers readable — no deep nesting:

```go
// Bad — nested
func ProcessSession(id string) error {
	s, err := GetSession(id)
	if err == nil {
		if err := ValidateDuration(s.Minutes); err == nil {
			return saveSession(s)
		} else {
			return err
		}
	} else {
		return err
	}
}

// Good — fail fast
func ProcessSession(id string) error {
	s, err := GetSession(id)
	if err != nil {
		return fmt.Errorf("process session: %w", err)
	}
	if err := ValidateDuration(s.Minutes); err != nil {
		return fmt.Errorf("process session %q: %w", id, err)
	}
	return saveSession(s)
}
```

Each failure returns immediately with wrapped context.

## Summary

| Pattern | Study Buddy example |
|---|---|
| `(T, error)` return | `GetSession`, `LoadCourse` |
| Sentinel errors | `ErrSessionNotFound` |
| `fmt.Errorf` + `%w` | Wrap file read failures |
| `errors.Is` / `errors.As` | Detect missing files |
| Early return | Flat, readable handlers |

Explicit error handling is a core Study Buddy habit — every boundary checks and wraps failures.
