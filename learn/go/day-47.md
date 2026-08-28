# Day 47: Databases (II) & Repositories — Query Organization

## 1. Centralize SQL

Kept SQL strings in `queries.go` constants.

## 2. Explore sqlc Optionally

Documented sqlc as optional compile-time tool.

## 3. Name Queries Clearly

Used descriptive names like `GetNoteByID`.

## 4. Avoid N+1 Queries

Used joins instead of per-row queries.

## Run

```bash
cd learn/go
go run ./day47
```

## Summary

| Concept | Takeaway |
|---|---|
| Centralize SQL | Kept SQL strings in `queries |
| Explore sqlc Optionally | Documented sqlc as optional compile-time tool |
| Name Queries Clearly | Used descriptive names like `GetNoteByID` |
| Avoid N+1 Queries | Used joins instead of per-row queries |
