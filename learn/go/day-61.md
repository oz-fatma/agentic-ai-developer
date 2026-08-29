# Day 61: gRPC & Protocol Buffers — Protobuf Definitions

## 1. Write .proto Files

Defined Greeter service and messages in `internal/greeterpb/greeter.proto`.

## 2. Hand-written Wire Types

Implemented message encoding in `messages.go` without protoc.

## 3. Inspect Wire Bytes

Printed protobuf wire format for HelloRequest.

## 4. Plan gRPC Services

Prepared shared greeter package for days 62-65.

## Run

```bash
cd learn/go
go run ./day61
```

## Summary

| Concept | Takeaway |
|---|---|
| Write .proto Files | Defined Greeter service and messages in `internal/ |
| Hand-written Wire Types | Implemented message encoding in `messages |
| Inspect Wire Bytes | Printed protobuf wire format for HelloRequest |
| Plan gRPC Services | Prepared shared greeter package for days 62-65 |
