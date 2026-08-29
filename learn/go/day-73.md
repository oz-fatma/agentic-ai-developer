# Day 73: Observability & Resilience — Retries

## 1. Exponential Backoff

Retry failed operations with increasing delay.

## 2. Max Attempts

Stop after N retries.

## 3. Retryable Errors

Only retry transient failures.

## 4. Context Aware

Respect cancellation during backoff.

## Run

```bash
cd learn/go
go run ./day73
```

## Summary

| Concept | Takeaway |
|---|---|
| Exponential Backoff | Retry failed operations with increasing delay |
| Max Attempts | Stop after N retries |
| Retryable Errors | Only retry transient failures |
| Context Aware | Respect cancellation during backoff |
