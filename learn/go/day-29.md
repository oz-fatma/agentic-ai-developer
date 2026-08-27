# Day 29: HTTP Basics & Handlers — Middleware Patterns

## 1. Wrap Handlers

Wrote middleware accepting and returning `http.Handler`.

## 2. Log Requests

Added logging middleware for method, path, duration.

## 3. Recover Panics

Recovery middleware returns 500 on panic.

## 4. Chain Middleware

Composed logging, recovery, and request-ID middleware.

## Run

```bash
cd learn/go
go run ./day29
```

## Summary

| Concept | Takeaway |
|---|---|
| Wrap Handlers | Wrote middleware accepting and returning `http |
| Log Requests | Added logging middleware for method, path, duration |
| Recover Panics | Recovery middleware returns 500 on panic |
| Chain Middleware | Composed logging, recovery, and request-ID middleware |
