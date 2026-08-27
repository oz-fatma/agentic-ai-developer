# Day 12: Errors, Collections & I/O — Slices and Maps

**Project:** Study Buddy — managing courses and word frequencies

## 1. Work with Slices

Slices are flexible views over arrays — Study Buddy's primary list type:

```go
// Create
courses := []string{"Go", "Math", "History"}
courses = make([]string, 0, 10) // len 0, cap 10

// Append
courses = append(courses, "Go", "Math")
courses = append(courses, anotherSlice...)

// Slice expressions
recent := courses[1:3]   // elements at index 1, 2
first := courses[:2]     // first two
last := courses[2:]      // from index 2 to end

// Copy (avoid shared backing array surprises)
dest := make([]string, len(courses))
copy(dest, courses)
```

**Length vs capacity:**

```go
s := make([]int, 3, 5) // len=3, cap=5
fmt.Println(len(s), cap(s)) // 3 5

s = append(s, 4, 5)     // fits in cap — no reallocation
s = append(s, 6)        // exceeds cap — new backing array allocated
```

Study Buddy session list:

```go
sessions := []StudySession{
	{Subject: "Go", Minutes: 60},
	{Subject: "Math", Minutes: 45},
}

totalMinutes := 0
for _, s := range sessions {
	totalMinutes += s.Minutes
}
```

## 2. Iterate Collections

**Slice iteration:**

```go
for i, session := range sessions {
	fmt.Printf("%d: %s (%d min)\n", i, session.Subject, session.Minutes)
}
```

**Map iteration** — order is intentionally randomized:

```go
freq := map[string]int{"Go": 5, "Math": 3}
for subject, count := range freq {
	fmt.Printf("%s: %d sessions\n", subject, count)
}
// Order differs each run — do not depend on it
```

Check map key existence:

```go
count, ok := freq["Go"]
if ok {
	fmt.Println("Go sessions:", count)
}
```

## 3. Use make and Literals

| Initialization | When to use |
|---|---|
| `[]T{}` or `[]T{...}` | Known initial values |
| `make([]T, len)` | Pre-sized slice, zero-filled |
| `make([]T, 0, cap)` | Empty slice with reserved capacity |
| `map[K]V{}` | Small known maps |
| `make(map[K]V)` | Maps built incrementally |

```go
// Study Buddy weekly tracker
weekly := make([]int, 7) // Mon–Sun minutes, all 0

subjectCount := make(map[string]int)
subjectCount["Go"]++
```

## 4. Avoid Shared Mutation Surprises

Slices and maps are passed **by value** but share **backing data**:

```go
func addTag(courses []string, tag string) {
	courses = append(courses, tag) // may not affect caller if reallocated
}

func addTagSafe(courses *[]string, tag string) {
	*courses = append(*courses, tag) // affects caller
}
```

Map mutation is always visible to caller:

```go
func recordSession(freq map[string]int, subject string) {
	freq[subject]++ // caller sees the change
}
```

**Safe pattern** — return new slice instead of mutating:

```go
func filterCompleted(sessions []StudySession) []StudySession {
	result := make([]StudySession, 0)
	for _, s := range sessions {
		if s.Completed {
			result = append(result, s)
		}
	}
	return result
}
```

## Summary

| Collection | Study Buddy use |
|---|---|
| Slice | Course lists, session history |
| Map | Subject frequency, session lookup by ID |
| `append` | Grow session lists dynamically |
| Shared backing | Copy or return new slices when isolating data |

Slices and maps are the workhorse collections for Study Buddy's in-memory data.
