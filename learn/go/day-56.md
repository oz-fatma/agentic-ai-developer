# Day 56: Project Layout & Architecture — Standard Go Project Layout

## 1. Organize cmd and internal

Binaries under `cmd/`, private code under `internal/`.

## 2. Separate internal packages

Split domain, app, and adapter packages.

## 3. Keep main Thin

Wired dependencies in main without business logic.

## 4. Add README Structure

Documented package responsibilities.

## Run

```bash
cd learn/go
go run ./day56
```

## Summary

| Concept | Takeaway |
|---|---|
| Organize cmd and internal | Binaries under `cmd/`, private code under `internal/` |
| Separate internal packages | Split domain, app, and adapter packages |
| Keep main Thin | Wired dependencies in main without business logic |
| Add README Structure | Documented package responsibilities |
