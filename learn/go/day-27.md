# Day 27: HTTP Basics & Handlers — Routing and Handlers

## 1. Use ServeMux

Registered `/health`, `/items`, and `/items/` patterns.

## 2. Path Parameters

Parsed IDs from paths manually.

## 3. Method Checks

Returned 405 for unsupported HTTP methods.

## 4. Organize Handlers

Split handlers by resource in separate functions.

## Run

```bash
cd learn/go
go run ./day27
```

## Summary

| Concept | Takeaway |
|---|---|
| Use ServeMux | Registered `/health`, `/items`, and `/items/` patterns |
| Path Parameters | Parsed IDs from paths manually |
| Method Checks | Returned 405 for unsupported HTTP methods |
| Organize Handlers | Split handlers by resource in separate functions |
