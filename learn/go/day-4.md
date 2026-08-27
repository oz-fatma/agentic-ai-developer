# Day 4: Go Fundamentals — Pointers and Memory Basics

**Project:** Study Buddy — mutating session state through pointers

## 1. Understand Addresses

Pointers hold memory addresses. Use `&` to get an address and `*` to dereference:

```go
package main

import "fmt"

func main() {
	minutes := 45
	ptr := &minutes // address of minutes

	fmt.Println("Value:", minutes)   // 45
	fmt.Println("Address:", ptr)     // 0x...
	fmt.Println("Dereferenced:", *ptr) // 45

	*ptr = 60 // mutate original through pointer
	fmt.Println("Updated:", minutes) // 60
}
```

Pointer type syntax: `*int` means "pointer to int".

## 2. Pass by Value vs Pointer

**Pass by value** — function gets a copy; original is unchanged:

```go
func addMinutes(session struct{ Duration int }, extra int) {
	session.Duration += extra // modifies copy only
}

type Session struct{ Duration int }

func main() {
	s := Session{Duration: 30}
	addMinutes(s, 15)
	fmt.Println(s.Duration) // still 30
}
```

**Pass by pointer** — function can mutate the original:

```go
func addMinutesPtr(session *Session, extra int) {
	session.Duration += extra // modifies original
}

func main() {
	s := Session{Duration: 30}
	addMinutesPtr(&s, 15)
	fmt.Println(s.Duration) // 45
}
```

Go automatically dereferences pointers for struct field access — `session.Duration` works on `*Session`.

## 3. Use new

`new(T)` allocates zero-valued memory and returns `*T`:

```go
counter := new(int)   // *int pointing to 0
*counter = 5

session := new(Session) // *Session with Duration: 0
session.Duration = 90
```

Compare with composite literal address:

```go
session := &Session{Duration: 90} // preferred when you know initial values
```

Use `&T{...}` when initializing with values; use `new` when you need a zero-valued pointer.

## 4. Avoid Nil Panics

A nil pointer holds no address. Dereferencing it causes a runtime panic:

```go
var session *Session // nil

// session.Duration = 30  // PANIC: nil pointer dereference

if session != nil {
	session.Duration = 30 // safe
} else {
	fmt.Println("No session loaded")
}
```

**Nil is the zero value for:** pointers, slices, maps, channels, interfaces, and function values.

Study Buddy helper to safely extend a session:

```go
func ExtendSession(session *Session, extra int) error {
	if session == nil {
		return fmt.Errorf("session is nil")
	}
	session.Duration += extra
	return nil
}
```

## When to Use Pointers

| Use pointer when | Use value when |
|---|---|
| Function must mutate caller's data | Data is small and immutable |
| Struct is large (avoid copying) | Simple value types (`int`, `bool`) |
| Method needs to modify receiver | Read-only methods on small structs |

## Summary

Pointers let Study Buddy update session duration in place without copying entire structs. Always check for `nil` before dereferencing — missing sessions should return errors, not panic.
