package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day3/util"
)

func divide(a, b float64) (result float64, ok bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	fmt.Println(util.Greet("Go"))
	fmt.Println(util.FormatDuration(90))

	score, err := util.ParseScore("85")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("score:", score)
	}

	if result, ok := divide(10, 2); ok {
		fmt.Println("10 / 2 =", result)
	}
}
