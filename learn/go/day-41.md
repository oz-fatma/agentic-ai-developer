# Day 41: Databases (I) — database/sql Introduction

## 1. Open a DB Connection

Connected to SQLite with `modernc.org/sqlite` driver via `database/sql`.

## 2. Ping the Database

Verified connectivity with `PingContext` on startup.

## 3. Query Rows

Ran `SELECT` with `QueryContext` and scanned into variables.

## 4. Defer Close

Used `defer rows.Close()` and `defer db.Close()` to avoid leaks.

## Run

```bash
cd learn/go
go run ./day41
```

## Summary

| Concept | Takeaway |
|---|---|
| Open a DB Connection | Connected to SQLite with `modernc |
| Ping the Database | Verified connectivity with `PingContext` on startup |
| Query Rows | Ran `SELECT` with `QueryContext` and scanned into varia |
| Defer Close | Used `defer rows |
