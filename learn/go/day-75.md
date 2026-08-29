# Day 75: Observability & Resilience — Timeouts

## 1. context.WithTimeout

Bound HTTP fetch duration.

## 2. Deadline Exceeded

Handle timeout errors gracefully.

## 3. Fast Path

Successful requests within deadline.

## 4. No Goroutine Leaks

Clean up on timeout.

## Run

```bash
cd learn/go
go run ./day75
```

## Summary

| Concept | Takeaway |
|---|---|
| context.WithTimeout | Bound HTTP fetch duration |
| Deadline Exceeded | Handle timeout errors gracefully |
| Fast Path | Successful requests within deadline |
| No Goroutine Leaks | Clean up on timeout |
