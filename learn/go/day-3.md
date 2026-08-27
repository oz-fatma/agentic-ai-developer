# Day 3: Go Fundamentals — Functions and Packages

**Project:** Study Buddy — reusable helpers in a separate package

## 1. Write Functions

Defined functions with single and multiple return values:

```go
// Single return
func formatDuration(minutes int) string {
	hours := minutes / 60
	mins := minutes % 60
	return fmt.Sprintf("%dh %dm", hours, mins)
}

// Multiple returns — Go idiom for results + errors (preview)
func parseScore(input string) (int, error) {
	score, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid score %q: %w", input, err)
	}
	return score, nil
}
```

Multiple return values eliminate the need for output parameters and make error handling explicit.

## 2. Export and Import

Created a helper package `studyutil/` inside the Study Buddy module:

**studyutil/duration.go**

```go
package studyutil

// MinutesToHours converts study minutes to fractional hours.
func MinutesToHours(minutes int) float64 {
	return float64(minutes) / 60.0
}

// IsGoalMet reports whether total minutes meet the daily goal.
func IsGoalMet(totalMinutes, goalMinutes int) bool {
	return totalMinutes >= goalMinutes
}
```

**main.go**

```go
package main

import (
	"fmt"

	"github.com/fatmaoz/study-buddy/studyutil"
)

func main() {
	hours := studyutil.MinutesToHours(90)
	fmt.Printf("Studied %.1f hours\n", hours)
	fmt.Println("Goal met:", studyutil.IsGoalMet(90, 60))
}
```

**Export rules:**
- Capitalized names (`MinutesToHours`) are **exported** — visible outside the package
- Lowercase names are **unexported** — package-private

## 3. Use Named Returns

Named result parameters document return values and enable naked returns in small functions:

```go
func splitSession(totalMinutes int) (hours, minutes int) {
	hours = totalMinutes / 60
	minutes = totalMinutes % 60
	return // naked return — returns hours, minutes
}
```

Use naked returns sparingly — they are clearest in very short functions.

## 4. Read godoc Style

Added documentation comments above exported functions:

```go
// AverageScore returns the mean of quiz scores.
// Returns 0 if scores is empty.
func AverageScore(scores []int) float64 {
	if len(scores) == 0 {
		return 0
	}
	sum := 0
	for _, s := range scores {
		sum += s
	}
	return float64(sum) / float64(len(scores))
}
```

View documentation:

```bash
go doc studyutil.AverageScore
```

Comments immediately preceding declarations become the official API documentation.

## Package Layout

```
study-buddy/
├── go.mod
├── main.go
└── studyutil/
    └── duration.go
```

## Summary

| Concept | Study Buddy usage |
|---|---|
| Functions | Format duration, parse scores |
| Multiple returns | `(result, error)` pattern |
| Packages | `studyutil` for shared helpers |
| Export | Capitalize public API names |
| godoc comments | Document exported functions |

Study Buddy now has a clean package structure ready for pointers and practice exercises.
