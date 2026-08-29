package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/internal/greeterpb"
	"google.golang.org/grpc"
)

func loggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	fmt.Printf("[interceptor] -> %s\n", info.FullMethod)
	resp, err := handler(ctx, req)
	fmt.Printf("[interceptor] <- %s (%v)\n", info.FullMethod, time.Since(start))
	return resp, err
}

func main() {
	fmt.Println("Day 63: gRPC unary interceptors")

	ctx := context.Background()
	conn, cleanup, err := greeterpb.DialBufconn(ctx, &greeterpb.Server{},
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	client := greeterpb.NewGreeterClient(conn)
	reply, err := client.SayHello(ctx, &greeterpb.HelloRequest{Name: "interceptor"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reply:", reply.Message)
}
