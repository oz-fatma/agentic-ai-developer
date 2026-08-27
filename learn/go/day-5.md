# Day 5: Go Fundamentals — Practice

**Project:** Study Buddy — small CLI utilities combining Days 1–4

## 1. CLI Calculator

Built a command-line calculator that reads two numbers and an operator:

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: study-buddy calc <num1> <op> <num2>")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid first number")
		os.Exit(1)
	}
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid second number")
		os.Exit(1)
	}

	result, err := calculate(a, b, os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("Result: %.2f\n", result)
}
```

**Test cases:**

| Input | Expected |
|---|---|
| `calc 10 + 5` | `15.00` |
| `calc 10 / 0` | Error: division by zero |
| `calc 10 % 5` | Error: unknown operator |

## 2. Temperature Converter

Added conversion functions to `studyutil/convert.go`:

```go
package studyutil

// CelsiusToFahrenheit converts °C to °F.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// FahrenheitToCelsius converts °F to °C.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
```

Used from `main`:

```go
fmt.Println(studyutil.CelsiusToFahrenheit(20))  // 68
fmt.Println(studyutil.FahrenheitToCelsius(68))  // 20
```

Handy for Study Buddy's "optimal study room temperature" feature (21°C = 69.8°F).

## 3. Package Refactor

Reorganized the module:

```
study-buddy/
├── go.mod
├── main.go
└── studyutil/
    ├── duration.go    // MinutesToHours, IsGoalMet
    └── convert.go     // temperature helpers
```

Moved all reusable logic out of `main`. `main` only parses CLI args and calls package functions — clean separation.

## 4. Self-Review

Ran quality checks:

```bash
go fmt ./...    # format all files
go vet ./...    # static analysis — no issues
go run . calc 90 / 2   # Result: 45.00
```

**Edge cases tested manually:**

| Case | Handled? |
|---|---|
| Division by zero | Yes — returns error |
| Invalid operator | Yes — returns error |
| Missing CLI args | Yes — usage message + exit 1 |
| Non-numeric input | Yes — parse error message |

## Summary

Days 1–4 skills combined into working Study Buddy utilities: variables, control flow, functions, packages, and pointers. The module is structured for structs and domain modeling starting Day 6.

**Phase checkpoint:** Go Fundamentals complete.
