# Day 7: Structs, Methods & Interfaces — Methods and Receivers

**Project:** Study Buddy — behavior on session and course types

## 1. Write Value Receivers

Value receivers operate on a **copy** — mutations do not persist:

```go
func (s StudySession) IsLongSession() bool {
	return s.Minutes >= 60
}

func (s StudySession) DisplayDuration() string {
	return fmt.Sprintf("%d minutes", s.Minutes)
}
```

These read-only methods work on both values and pointers.

## 2. Write Pointer Receivers

Pointer receivers can **mutate** the struct and avoid copying large values:

```go
func (s *StudySession) Complete() {
	s.Completed = true
}

func (s *StudySession) AddMinutes(extra int) {
	s.Minutes += extra
}

func (s *StudySession) AppendNote(note string) {
	if s.Notes == "" {
		s.Notes = note
	} else {
		s.Notes += "; " + note
	}
}
```

Usage:

```go
session := &StudySession{Subject: "Go", Minutes: 30}
session.AddMinutes(15)
session.Complete()
// session.Minutes == 45, session.Completed == true
```

## 3. Choose Receiver Type

**Use pointer receiver when:**
- Method mutates the receiver
- Struct is large (avoid copy overhead)
- Consistency — if one method uses pointer, all should

**Use value receiver when:**
- Method is read-only and struct is small
- Receiver is a basic type alias or tiny struct

Study Buddy rule: `StudySession` methods all use pointer receivers because sessions are mutated (add minutes, complete, append notes).

```go
// Course is mostly read-only after creation — mixed is OK
func (c Course) FullName() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.Subject)
}

func (c *Course) AddCredit() {
	c.Credits++
}
```

## 4. Method Sets

The **method set** determines which methods are callable and which interfaces a type satisfies.

| Variable type | Can call value receivers | Can call pointer receivers |
|---|---|---|
| `StudySession` (value) | Yes | Yes (Go auto-takes address) |
| `*StudySession` (pointer) | Yes | Yes |
| `Course` stored in interface | Value methods only | — |

Important for interfaces (Day 8): if an interface requires a pointer-receiver method, only `*StudySession` satisfies it — not `StudySession` value.

Example:

```go
type Completer interface {
	Complete()
}

// Only *StudySession satisfies Completer, not StudySession value
var c Completer = &StudySession{}
```

## Summary

| Receiver | Semantics | Study Buddy use |
|---|---|---|
| Value `(s StudySession)` | Read-only, copy | `IsLongSession()`, `DisplayDuration()` |
| Pointer `(s *StudySession)` | Mutate in place | `Complete()`, `AddMinutes()` |
| Consistency | Same type, same receiver style | All mutating methods → pointer |

Methods attach behavior to Study Buddy types without classes or inheritance.
