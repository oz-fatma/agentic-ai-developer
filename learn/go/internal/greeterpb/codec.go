package greeterpb

import (
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterCodec(JSONCodec{})
}

// JSONCodec lets gRPC use plain structs without protoc-generated types.
type JSONCodec struct{}

func (JSONCodec) Name() string { return "json" }

func (JSONCodec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (JSONCodec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// DefaultCallOptions returns client options that use the JSON codec.
func DefaultCallOptions() []grpc.CallOption {
	return []grpc.CallOption{grpc.CallContentSubtype(JSONCodec{}.Name())}
}

// DefaultServerOptions returns server options that use the JSON codec.
func DefaultServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{grpc.ForceServerCodec(JSONCodec{})}
}
