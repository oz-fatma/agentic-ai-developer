package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/internal/greeterpb"
)

func main() {
	fmt.Println("Day 62: gRPC server/client with bufconn (in-memory)")

	ctx := context.Background()
	conn, cleanup, err := greeterpb.DialBufconn(ctx, &greeterpb.Server{})
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	client := greeterpb.NewGreeterClient(conn)
	reply, err := client.SayHello(ctx, &greeterpb.HelloRequest{Name: "bufconn"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SayHello reply:", reply.Message)
}
