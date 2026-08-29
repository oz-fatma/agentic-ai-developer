package greeterpb

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protowire"
)

// HelloRequest mirrors greeter.proto (hand-written; protoc would generate .pb.go).
type HelloRequest struct {
	Name string
}

// Marshal encodes HelloRequest to protobuf wire format (field 1 = name).
func (m *HelloRequest) Marshal() ([]byte, error) {
	var b []byte
	if m.Name != "" {
		b = protowire.AppendTag(b, 1, protowire.BytesType)
		b = protowire.AppendString(b, m.Name)
	}
	return b, nil
}

// Unmarshal decodes protobuf wire format into HelloRequest.
func (m *HelloRequest) Unmarshal(b []byte) error {
	*m = HelloRequest{}
	for len(b) > 0 {
		num, wt, n := protowire.ConsumeTag(b)
		if n < 0 {
			return fmt.Errorf("tag: %w", protowire.ParseError(n))
		}
		b = b[n:]
		switch num {
		case 1:
			if wt != protowire.BytesType {
				return fmt.Errorf("field name: unexpected wire type %d", wt)
			}
			v, n := protowire.ConsumeString(b)
			if n < 0 {
				return fmt.Errorf("field name: %w", protowire.ParseError(n))
			}
			m.Name = v
			b = b[n:]
		default:
			n = protowire.ConsumeFieldValue(num, wt, b)
			if n < 0 {
				return fmt.Errorf("skip field %d: %w", num, protowire.ParseError(n))
			}
			b = b[n:]
		}
	}
	return nil
}

// HelloReply mirrors greeter.proto.
type HelloReply struct {
	Message string
}

// Marshal encodes HelloReply to protobuf wire format (field 1 = message).
func (m *HelloReply) Marshal() ([]byte, error) {
	var b []byte
	if m.Message != "" {
		b = protowire.AppendTag(b, 1, protowire.BytesType)
		b = protowire.AppendString(b, m.Message)
	}
	return b, nil
}

// Unmarshal decodes protobuf wire format into HelloReply.
func (m *HelloReply) Unmarshal(b []byte) error {
	*m = HelloReply{}
	for len(b) > 0 {
		num, wt, n := protowire.ConsumeTag(b)
		if n < 0 {
			return fmt.Errorf("tag: %w", protowire.ParseError(n))
		}
		b = b[n:]
		switch num {
		case 1:
			if wt != protowire.BytesType {
				return fmt.Errorf("field message: unexpected wire type %d", wt)
			}
			v, n := protowire.ConsumeString(b)
			if n < 0 {
				return fmt.Errorf("field message: %w", protowire.ParseError(n))
			}
			m.Message = v
			b = b[n:]
		default:
			n = protowire.ConsumeFieldValue(num, wt, b)
			if n < 0 {
				return fmt.Errorf("skip field %d: %w", num, protowire.ParseError(n))
			}
			b = b[n:]
		}
	}
	return nil
}
