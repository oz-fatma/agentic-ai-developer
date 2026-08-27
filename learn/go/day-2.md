# Day 2: Go Fundamentals — Control Flow

**Project:** Study Buddy — branching and loops for session logic

## 1. Learn if/else

Wrote conditional logic to classify study session quality based on duration:

```go
package main

import "fmt"

func sessionQuality(minutes int) string {
	if minutes >= 90 {
		return "Excellent — deep focus session"
	} else if minutes >= 45 {
		return "Good — solid progress"
	} else if minutes >= 15 {
		return "Short — better than nothing"
	} else {
		return "Too brief — try again"
	}
}

func main() {
	fmt.Println(sessionQuality(120)) // Excellent
	fmt.Println(sessionQuality(30))  // Short
}
```

Go allows an initialization statement before the condition:

```go
if minutes := 60; minutes >= 45 {
	fmt.Println("Goal met")
}
```

## 2. Use for Loops

Go has one loop keyword — `for` — used in three ways:

**Classic C-style:**

```go
for i := 1; i <= 5; i++ {
	fmt.Printf("Pomodoro round %d\n", i)
}
```

**While-style (condition only):**

```go
remaining := 25
for remaining > 0 {
	fmt.Printf("%d minutes left\n", remaining)
	remaining -= 5
}
```

**Range over a slice:**

```go
subjects := []string{"Go", "Math", "History"}
for index, subject := range subjects {
	fmt.Printf("%d: %s\n", index, subject)
}

// Ignore index with _
for _, subject := range subjects {
	fmt.Println(subject)
}
```

## 3. Try switch

**Value switch** — pick a study command:

```go
func handleCommand(cmd string) {
	switch cmd {
	case "start":
		fmt.Println("Starting study session...")
	case "stop":
		fmt.Println("Session ended.")
	case "stats":
		fmt.Println("Showing weekly stats...")
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
```

**Tagless switch** — replaces long if/else chains:

```go
func gradeLabel(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "Needs improvement"
	}
}
```

No fall-through by default — each case breaks automatically (unlike C).

## 4. Handle break and continue

**continue** — skip incomplete sessions in a weekly report:

```go
sessions := []int{45, 0, 60, 10, 90}
total := 0

for _, mins := range sessions {
	if mins == 0 {
		continue // skip days with no study
	}
	total += mins
}
fmt.Println("Total minutes:", total) // 205
```

**break with label** — exit nested loops when daily goal is met:

```go
goal := 120
studied := 0

outer:
for day := 1; day <= 7; day++ {
	for session := 1; session <= 3; session++ {
		studied += 45
		if studied >= goal {
			fmt.Printf("Goal reached on day %d\n", day)
			break outer
		}
	}
}
```

## Summary

| Construct | Purpose in Study Buddy |
|---|---|
| `if/else` | Grade sessions, validate input |
| `for` | Iterate subjects, pomodoro rounds |
| `switch` | Route CLI commands |
| `break` / `continue` | Control nested loops, skip empty days |

Control flow makes Study Buddy react to different inputs and repeat tasks efficiently.
