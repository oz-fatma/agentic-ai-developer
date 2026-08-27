# Day 1: Go Fundamentals — Variables, Types, and Your First Program

**Project:** Study Buddy — a CLI app to track study sessions and courses

## 1. Install Go

Installed Go 1.22+ from [go.dev](https://go.dev) and verified the toolchain:

```bash
go version
# go version go1.22.5 darwin/arm64
```

The Go SDK includes the compiler, standard library, `go fmt`, `go test`, and module tools.

## 2. Initialize a Module

Created the project directory and initialized a module for Study Buddy:

```bash
mkdir -p ~/developer/study-buddy
cd ~/developer/study-buddy
go mod init github.com/fatmaoz/study-buddy
```

This created `go.mod`:

```
module github.com/fatmaoz/study-buddy

go 1.22
```

## 3. Write Hello World

Created `main.go` with a Study Buddy greeting:

```go
package main

import "fmt"

func main() {
	fmt.Println("Welcome to Study Buddy!")
}
```

Run it:

```bash
go run .
# Welcome to Study Buddy!
```

**Program structure:**
- `package main` — executable entry package
- `import "fmt"` — standard formatting/printing package
- `func main()` — program entry point

## 4. Explore Core Types

Declared variables with both `var` and short declaration `:=`:

```go
package main

import "fmt"

func main() {
	// Explicit type with var
	var appName string = "Study Buddy"
	var version float64 = 1.0
	var isActive bool = true
	var sessionCount int = 0

	// Type inference with :=
	userName := "Alex"
	dailyGoalHours := 2.5
	hasPremium := false

	fmt.Println(appName, version, isActive, sessionCount)
	fmt.Println(userName, dailyGoalHours, hasPremium)
}
```

| Type | Example | Use in Study Buddy |
|---|---|---|
| `int` | `sessionCount := 3` | Count completed sessions |
| `float64` | `hours := 1.5` | Track study duration |
| `string` | `subject := "Go"` | Course and topic names |
| `bool` | `completed := true` | Session completion flag |

**Notes:**
- `:=` can only be used inside functions
- Uninitialized variables get **zero values**: `0`, `0.0`, `""`, `false`

## 5. Use go fmt

Formatted code with the standard tool:

```bash
go fmt ./...
```

Go enforces consistent style automatically — tabs for indentation, standard spacing around operators, and aligned imports. Running `go fmt` before every commit is standard practice on Go teams.

## Summary

| Concept | Takeaway |
|---|---|
| Module | Versioned unit of Go code (`go.mod`) |
| Package | Organizes source files; `main` is executable |
| `var` vs `:=` | Explicit declaration vs short inference |
| Core types | `int`, `float64`, `string`, `bool` |
| `go fmt` | Enforces one consistent code style |

Study Buddy module is initialized and ready for control flow and functions in upcoming days.
