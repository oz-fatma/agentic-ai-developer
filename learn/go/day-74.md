# Day 74: Observability & Resilience — Circuit Breaker

## 1. Open Circuit

Stop calling failing upstream after threshold.

## 2. Half-open Probe

Test recovery before closing circuit.

## 3. Fail Fast

Return error immediately when open.

## 4. Protect Downstream

Prevent cascade failures.

## Run

```bash
cd learn/go
go run ./day74
```

## Summary

| Concept | Takeaway |
|---|---|
| Open Circuit | Stop calling failing upstream after threshold |
| Half-open Probe | Test recovery before closing circuit |
| Fail Fast | Return error immediately when open |
| Protect Downstream | Prevent cascade failures |
