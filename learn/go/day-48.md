# Day 48: Databases (II) & Repositories — Testing with Databases

## 1. Use Test Database

Tests run against disposable in-memory SQLite.

## 2. Reset State

Each test gets fresh schema.

## 3. Test Repositories

Integration tests for create, read, not-found.

## 4. Parallel Safety

Tests isolated with separate DB instances.

## Run

```bash
cd learn/go
go test ./day48/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Use Test Database | Tests run against disposable in-memory SQLite |
| Reset State | Each test gets fresh schema |
| Test Repositories | Integration tests for create, read, not-found |
| Parallel Safety | Tests isolated with separate DB instances |
