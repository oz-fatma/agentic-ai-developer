package main

import "fmt"

func main() {
	fmt.Println("Day 76: multi-stage Dockerfile")
	fmt.Println("Build: docker build -f day76/Dockerfile -t learn-go-day76 .")
	fmt.Println("This binary is the runtime artifact copied into the final image stage.")
}
