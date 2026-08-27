# Day 39: Context, Config & Middleware — Graceful Shutdown

## 1. Trap Signals

Listened for SIGINT/SIGTERM via `os/signal`.

## 2. Use http.Server Shutdown

Called `Shutdown` with context to drain requests.

## 3. Set Timeouts

Configured read/write timeouts on `http.Server`.

## 4. Close Resources

Demonstrated clean exit after drain.

## Run

```bash
cd learn/go
go run ./day39
```

## Summary

| Concept | Takeaway |
|---|---|
| Trap Signals | Listened for SIGINT/SIGTERM via `os/signal` |
| Use http.Server Shutdown | Called `Shutdown` with context to drain requests |
| Set Timeouts | Configured read/write timeouts on `http |
| Close Resources | Demonstrated clean exit after drain |
