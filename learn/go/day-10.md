# Day 10: Structs, Methods & Interfaces — Practice

**Project:** Study Buddy — shape library kata + logger abstraction

## 1. Shape Library

Classic Go kata adapted for Study Buddy's "study progress shapes" visualization:

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64  { return 2 * (r.Width + r.Height) }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
```

Both `Rectangle` and `Circle` satisfy `Shape` implicitly.

## 2. Logger Abstraction

Mirrors real service design — depend on interfaces at boundaries:

```go
type Logger interface {
	Info(msg string)
	Error(msg string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Info(msg string)  { fmt.Println("[INFO]", msg) }
func (ConsoleLogger) Error(msg string) { fmt.Println("[ERROR]", msg) }

type NoopLogger struct{}

func (NoopLogger) Info(msg string)  {} // silent — for tests
func (NoopLogger) Error(msg string) {}
```

Study Buddy service using injected logger:

```go
type SessionService struct {
	log Logger
}

func (svc *SessionService) StartSession(subject string, log Logger) {
	log.Info(fmt.Sprintf("Starting session: %s", subject))
}

func NewSessionService(log Logger) *SessionService {
	if log == nil {
		log = NoopLogger{}
	}
	return &SessionService{log: log}
}
```

Tests inject `NoopLogger`; production uses `ConsoleLogger`.

## 3. Manual Test Cases

Table-style inputs and expected outputs for shape methods:

| Shape | Input | Area (want) | Perimeter (want) |
|---|---|---|---|
| Rectangle | W=4, H=5 | 20.0 | 18.0 |
| Rectangle | W=0, H=5 | 0.0 | 10.0 |
| Circle | R=3 | ~28.27 | ~18.85 |
| Circle | R=0 | 0.0 | 0.0 |

Verified manually:

```go
r := Rectangle{Width: 4, Height: 5}
fmt.Println(r.Area())      // 20
fmt.Println(r.Perimeter()) // 18
```

## 4. Code Review

Self-review checklist for Study Buddy code:

| Check | Finding | Action |
|---|---|---|
| Naming | `StudySession` not `SS` | Clear domain names |
| Receiver consistency | All mutating methods use pointer receiver on `StudySession` | Consistent |
| Interface size | `Logger` has 2 methods, `Shape` has 2 | Narrow — good |
| Export boundaries | `SessionService` exported, internals lowercase | Clean API |
| Nil safety | `NewSessionService` defaults nil logger to `NoopLogger` | Defensive |

## Summary

**Phase checkpoint:** Structs, Methods & Interfaces complete.

Study Buddy now has:
- Domain structs (`Course`, `StudySession`, `User`, `Student`)
- Methods with appropriate receivers
- Interfaces for logging and formatting
- Composition for user roles
- Practice with polymorphic shape and logger patterns

Ready for errors, collections, and I/O in Days 11–15.
