# Day 87: Performance — Memory & Allocations

## 1. Benchmark Allocs

benchmem for allocation count.

## 2. Builder vs Concat

Compared string building strategies.

## 3. Reduce Allocations

Prefer pre-sized buffers.

## 4. Read bench Output

ns/op and allocs/op columns.

## Run

```bash
cd learn/go
go test -bench=. ./day87/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Benchmark Allocs | benchmem for allocation count |
| Builder vs Concat | Compared string building strategies |
| Reduce Allocations | Prefer pre-sized buffers |
| Read bench Output | ns/op and allocs/op columns |
