# Day 84: Caching & Messaging — Worker Queue

## 1. Job Queue

Buffered channel of jobs.

## 2. Worker Pool

N goroutines processing jobs.

## 3. Fan-out

Parallel job processing.

## 4. Wait for Completion

WaitGroup drains queue.

## Run

```bash
cd learn/go
go run ./day84
```

## Summary

| Concept | Takeaway |
|---|---|
| Job Queue | Buffered channel of jobs |
| Worker Pool | N goroutines processing jobs |
| Fan-out | Parallel job processing |
| Wait for Completion | WaitGroup drains queue |
