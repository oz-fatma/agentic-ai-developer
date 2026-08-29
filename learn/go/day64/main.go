package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/internal/greeterpb"
	"google.golang.org/grpc/metadata"
)

func main() {
	fmt.Println("Day 64: gRPC metadata (headers)")

	ctx := context.Background()
	conn, cleanup, err := greeterpb.DialBufconn(ctx, &greeterpb.Server{})
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	md := metadata.Pairs("request-id", "day64-demo", "client", "learn-go")
	ctx = metadata.NewOutgoingContext(ctx, md)

	client := greeterpb.NewGreeterClient(conn)
	reply, err := client.SayHello(ctx, &greeterpb.HelloRequest{Name: "metadata"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reply:", reply.Message)
	fmt.Println("Sent metadata: request-id=day64-demo, client=learn-go")
}
