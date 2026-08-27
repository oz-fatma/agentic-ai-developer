package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day5/util"
)

func calc(a, b float64, op string) (float64, error) {
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
		return 0, fmt.Errorf("invalid operator %q", op)
	}
}

func main() {
	if len(os.Args) == 4 {
		a, _ := strconv.ParseFloat(os.Args[1], 64)
		b, _ := strconv.ParseFloat(os.Args[3], 64)
		result, err := calc(a, b, os.Args[2])
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(result)
	} else {
		result, err := calc(10, 2, "/")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("10 / 2 =", result)
	}

	c := 25.0
	f := util.ToFahrenheit(c)
	back := util.ToCelsius(f)
	fmt.Printf("%.1f°C = %.1f°F = %.1f°C\n", c, f, back)
}
