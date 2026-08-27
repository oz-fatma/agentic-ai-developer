# Day 6: Structs, Methods & Interfaces — Structs and Fields

**Project:** Study Buddy — domain models with structs

## 1. Define a Struct

Created struct types representing Study Buddy entities:

```go
package studybuddy

import "time"

type Course struct {
	ID       string
	Name     string
	Subject  string
	Credits  int
}

type StudySession struct {
	ID        string
	CourseID  string
	Subject   string
	Minutes   int
	Notes     string
	StartedAt time.Time
	Completed bool
}
```

Structs group related data under one type — Study Buddy's backbone for courses and sessions.

## 2. Create Struct Values

Three initialization styles:

**Positional literal** (field order must match):

```go
session := StudySession{"s1", "c1", "Go", 60, "", time.Now(), false}
```

**Keyed literal** (recommended — order-independent, self-documenting):

```go
session := StudySession{
	ID:        "s1",
	CourseID:  "go-101",
	Subject:   "Go",
	Minutes:   60,
	StartedAt: time.Now(),
	Completed: false,
}
```

**Pointer to struct:**

```go
course := &Course{
	ID:      "go-101",
	Name:    "Go Fundamentals",
	Subject: "Programming",
	Credits: 3,
}
```

**Zero value** — all fields get defaults:

```go
var empty StudySession
// empty.Minutes == 0, empty.Completed == false, empty.Subject == ""
```

## 3. Add Struct Methods Preview

Attached a simple method to preview tomorrow's deeper dive:

```go
func (s StudySession) Summary() string {
	status := "in progress"
	if s.Completed {
		status = "completed"
	}
	return fmt.Sprintf("%s: %d min (%s)", s.Subject, s.Minutes, status)
}
```

Usage:

```go
session := StudySession{Subject: "Go", Minutes: 45}
fmt.Println(session.Summary()) // Go: 45 min (in progress)
```

## 4. Use fmt and String Representation

Default formatting prints field names and values:

```go
fmt.Printf("%v\n", session)  // {s1 go-101 Go 60  2026-08-21 ... false}
fmt.Printf("%+v\n", session)  // includes field names
fmt.Printf("%#v\n", session)  // Go syntax representation
```

For custom output, implement `String()` (implements `fmt.Stringer`):

```go
func (s StudySession) String() string {
	return fmt.Sprintf("StudySession[%s] %s — %d min", s.ID, s.Subject, s.Minutes)
}
```

Now `fmt.Println(session)` uses the custom format automatically.

## Domain Model Sketch

```
Course (1) ──< (many) StudySession
  ├── ID, Name, Subject
  └── Sessions track Minutes, Notes, Completed
```

## Summary

| Concept | Study Buddy example |
|---|---|
| Struct | `Course`, `StudySession` |
| Keyed literal | Clear initialization with field names |
| Pointer struct | `&Course{...}` for shared references |
| Zero value | Empty session before first use |
| `String()` | Readable log output |

Structs model Study Buddy's core data. Tomorrow: methods and receiver types.
