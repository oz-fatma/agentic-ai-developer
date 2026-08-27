# Day 1: Go Fundamentals — Variables, Types, and Your First Program

## 1. Install Go

Installed Go 1.22+ from [go.dev](https://go.dev) and verified the toolchain:

```bash
go version
# go version go1.26.5 darwin/arm64
```

The Go SDK includes the compiler, standard library, `go fmt`, `go test`, and module tools.

## 2. Initialize a Module

Initialized a module in `learn/go`:

```bash
cd learn/go
go mod init github.com/oz-fatma/agentic-ai-developer/learn/go
```

This created `go.mod`:

```
module github.com/oz-fatma/agentic-ai-developer/learn/go

go 1.26
```

## 3. Write Hello World

Created `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run .
# Hello, Go!
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
	var count int = 0
	var name string = "Go"
	var score float64 = 9.5
	var active bool = true

	label := "Day 1"
	hours := 2.5
	done := false

	fmt.Println("Hello, Go!")
	fmt.Println(name, count, score, active)
	fmt.Println(label, hours, done)
}
```

| Type | Example | Typical use |
|---|---|---|
| `int` | `count := 3` | Counters, indexes |
| `float64` | `score := 9.5` | Decimals, measurements |
| `string` | `name := "Go"` | Text |
| `bool` | `active := true` | Flags, on/off state |

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

Module is initialized and ready for control flow in Day 2.
