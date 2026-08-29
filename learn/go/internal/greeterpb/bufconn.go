package greeterpb

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// DialBufconn starts an in-memory gRPC server and returns a client connection.
// The server goroutine exits when ctx is cancelled.
func DialBufconn(ctx context.Context, srv GreeterServer, opts ...grpc.ServerOption) (*grpc.ClientConn, func(), error) {
	lis := bufconn.Listen(bufSize)
	base := DefaultServerOptions()
	all := append(base, opts...)
	gs := grpc.NewServer(all...)
	RegisterGreeterServer(gs, srv)

	errCh := make(chan error, 1)
	go func() {
		errCh <- gs.Serve(lis)
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.NewClient("passthrough:///bufconn",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(DefaultCallOptions()...),
	)
	if err != nil {
		gs.Stop()
		return nil, nil, err
	}

	cleanup := func() {
		_ = conn.Close()
		gs.Stop()
		<-errCh
	}
	return conn, cleanup, nil
}
