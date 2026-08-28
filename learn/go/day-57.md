# Day 57: Project Layout & Architecture — Clean Architecture Layers

## 1. Domain Layer

Business rules free of HTTP and SQL details.

## 2. Service Layer

Use cases orchestrating repositories and domain types.

## 3. Transport Layer

HTTP maps to service calls (memory adapter demo).

## 4. Dependency Rule

Inner layers do not import outer layers.

## Run

```bash
cd learn/go
go run ./day57
```

## Summary

| Concept | Takeaway |
|---|---|
| Domain Layer | Business rules free of HTTP and SQL details |
| Service Layer | Use cases orchestrating repositories and domain types |
| Transport Layer | HTTP maps to service calls (memory adapter demo) |
| Dependency Rule | Inner layers do not import outer layers |
