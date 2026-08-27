# Day 24: Testing Fundamentals — httptest and Test Helpers

## 1. Use httptest.NewRecorder

Tested handlers without opening real ports.

## 2. Build Request Fixtures

Constructed `http.Request` values with query params.

## 3. Extract Helpers

Created helper functions marked with `t.Helper()`.

## 4. Table-Test Handlers

Table-driven tests for status codes and response bodies.

## Run

```bash
cd learn/go
go test ./day24/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Use httptest.NewRecorder | Tested handlers without opening real ports |
| Build Request Fixtures | Constructed `http |
| Extract Helpers | Created helper functions marked with `t |
| Table-Test Handlers | Table-driven tests for status codes and response bodies |
