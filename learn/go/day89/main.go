package main

import (
	"fmt"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day89/internal/stringsdemo"
)

func main() {
	fmt.Println("Day 89: strings.Builder vs concat")

	parts := make([]string, 500)
	for i := range parts {
		parts[i] = strings.Repeat("go", 4)
	}

	concat, concatDur := stringsdemo.Concat(parts)
	builder, builderDur := stringsdemo.WithBuilder(parts)

	fmt.Println("concat len:", len(concat), "elapsed:", concatDur)
	fmt.Println("builder len:", len(builder), "elapsed:", builderDur)
	if builderDur < concatDur {
		fmt.Println("Builder was faster for this input size")
	}
}
