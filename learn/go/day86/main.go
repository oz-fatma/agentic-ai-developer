package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func hotLoop(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i % 7
	}
	return sum
}

func main() {
	fmt.Println("Day 86: pprof demo")
	fmt.Println("Tip: go tool pprof cpu.prof")

	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove("cpu.prof")

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	result := hotLoop(5_000_000)
	pprof.StopCPUProfile()

	fmt.Println("hotLoop result:", result)
	fmt.Println("wrote cpu.prof (removed on exit)")
}
