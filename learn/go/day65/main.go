package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/internal/greeterpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func chainInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Println("Server received metadata:", md)
	}
	start := time.Now()
	resp, err := handler(ctx, req)
	fmt.Printf("RPC %s completed in %v\n", info.FullMethod, time.Since(start))
	return resp, err
}

func main() {
	fmt.Println("Day 65: gRPC recap — proto, bufconn, interceptors, metadata")

	ctx := context.Background()
	conn, cleanup, err := greeterpb.DialBufconn(ctx, &greeterpb.Server{},
		grpc.UnaryInterceptor(chainInterceptor),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("trace", "day65"))
	client := greeterpb.NewGreeterClient(conn)

	for _, name := range []string{"Ada", "Bob"} {
		reply, err := client.SayHello(ctx, &greeterpb.HelloRequest{Name: name})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply.Message)
	}
}
