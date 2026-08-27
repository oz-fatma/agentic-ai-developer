# Day 23: Testing Fundamentals — Benchmarks and Examples

## 1. Write Benchmarks

Added `BenchmarkSum` and ran with `go test -bench`.

## 2. Reset Timer

Used `b.ResetTimer()` after setup in benchmarks.

## 3. Add Examples

Wrote `ExampleUnique` with `// Output:` comment.

## 4. Read ns/op

Interpreted benchmark output for `Sum` and `FilterEven`.

## Run

```bash
cd learn/go
go test ./day23/... -bench=. -benchmem
```

## Summary

| Concept | Takeaway |
|---|---|
| Write Benchmarks | Added `BenchmarkSum` and ran with `go test -bench` |
| Reset Timer | Used `b |
| Add Examples | Wrote `ExampleUnique` with `// Output:` comment |
| Read ns/op | Interpreted benchmark output for `Sum` and `FilterEven` |
