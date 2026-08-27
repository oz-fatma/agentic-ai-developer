# Day 8: Structs, Methods & Interfaces — Interfaces and Polymorphism

**Project:** Study Buddy — pluggable notifiers and formatters

## 1. Define Small Interfaces

Go favors **narrow interfaces** — one or two methods:

```go
type Notifier interface {
	Notify(message string) error
}

type SessionFormatter interface {
	Format(session StudySession) string
}
```

Compare with large interfaces — small contracts are easier to implement and test.

## 2. Implement Implicitly

No `implements` keyword. Types satisfy interfaces by having the required methods:

```go
type ConsoleNotifier struct{}

func (ConsoleNotifier) Notify(message string) error {
	fmt.Println("[Study Buddy]", message)
	return nil
}

type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Notify(message string) error {
	fmt.Printf("Email to %s: %s\n", e.Address, message)
	return nil
}
```

Both satisfy `Notifier` automatically.

## 3. Use Interface Values

Store concrete types behind interface variables for polymorphism:

```go
func sendReminder(n Notifier, msg string) {
	if err := n.Notify(msg); err != nil {
		fmt.Println("Notification failed:", err)
	}
}

func main() {
	notifiers := []Notifier{
		ConsoleNotifier{},
		EmailNotifier{Address: "alex@example.com"},
	}

	for _, n := range notifiers {
		sendReminder(n, "Time to study Go!")
	}
}
```

`sendReminder` accepts any type with a `Notify(string) error` method — console, email, or future Slack notifier.

**Text formatter example:**

```go
type PlainFormatter struct{}

func (PlainFormatter) Format(s StudySession) string {
	return fmt.Sprintf("%s: %d min", s.Subject, s.Minutes)
}

type MarkdownFormatter struct{}

func (MarkdownFormatter) Format(s StudySession) string {
	return fmt.Sprintf("**%s**: %d min", s.Subject, s.Minutes)
}

func printSession(f SessionFormatter, s StudySession) {
	fmt.Println(f.Format(s))
}
```

## 4. Handle nil Interfaces

Two distinct nil situations:

```go
var n Notifier            // nil interface — no type, no value
fmt.Println(n == nil)     // true

var cn *ConsoleNotifier   // typed nil pointer
var n2 Notifier = cn      // non-nil interface holding nil concrete value!
fmt.Println(n2 == nil)    // false — gotcha!
```

Safe check before calling:

```go
func sendReminder(n Notifier, msg string) {
	if n == nil {
		return
	}
	n.Notify(msg)
}
```

For typed nil inside interface, type assertion helps:

```go
if cn, ok := n.(*ConsoleNotifier); ok && cn != nil {
	cn.Notify(msg)
}
```

## Summary

| Concept | Study Buddy application |
|---|---|
| Small interfaces | `Notifier`, `SessionFormatter` |
| Implicit satisfaction | Console and Email notifiers |
| Polymorphism | One `sendReminder` for all notifiers |
| Nil gotcha | Check interface before method calls |

Interfaces decouple Study Buddy's core logic from delivery mechanisms — essential for testable service design.
