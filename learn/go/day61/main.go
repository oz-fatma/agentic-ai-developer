package main

import (
	"fmt"
	"os"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/internal/greeterpb"
)

func main() {
	fmt.Println("Day 61: Protobuf definitions and wire encoding")
	fmt.Println("Proto: internal/greeterpb/greeter.proto")
	fmt.Println("Hand-written .pb.go equivalent: internal/greeterpb/messages.go")

	req := &greeterpb.HelloRequest{Name: "Go learner"}
	data, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Printf("HelloRequest wire bytes (%d): % x\n", len(data), data)

	got := &greeterpb.HelloRequest{}
	if err := got.Unmarshal(data); err != nil {
		panic(err)
	}
	fmt.Printf("Round-trip name: %q\n", got.Name)

	protoBytes, _ := os.ReadFile("internal/greeterpb/greeter.proto")
	fmt.Printf("greeter.proto (%d bytes) defines Greeter service + messages\n", len(protoBytes))
}
