# Day 46: Databases (II) & Repositories — Repository Pattern

## 1. Define Repository Interfaces

Exposed domain operations, not SQL strings.

## 2. Implement SQLite Repository

Concrete repository using `database/sql`.

## 3. Map Models to Domain

Converted DB models to domain types at boundaries.

## 4. Inject Repositories

Passed repositories via constructors.

## Run

```bash
cd learn/go
go run ./day46
```

## Summary

| Concept | Takeaway |
|---|---|
| Define Repository Interfaces | Exposed domain operations, not SQL strings |
| Implement SQLite Repository | Concrete repository using `database/sql` |
| Map Models to Domain | Converted DB models to domain types at boundaries |
| Inject Repositories | Passed repositories via constructors |
