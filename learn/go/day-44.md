# Day 44: Databases (I) — Transactions

## 1. Begin Transactions

Started transactions with `BeginTx`, commit or rollback.

## 2. Defer Rollback

Used `defer tx.Rollback()` after begin.

## 3. Multi-Step Updates

Transferred money between accounts atomically.

## 4. Isolation Awareness

Demonstrated rollback on insufficient funds.

## Run

```bash
cd learn/go
go run ./day44
```

## Summary

| Concept | Takeaway |
|---|---|
| Begin Transactions | Started transactions with `BeginTx`, commit or rollback |
| Defer Rollback | Used `defer tx |
| Multi-Step Updates | Transferred money between accounts atomically |
| Isolation Awareness | Demonstrated rollback on insufficient funds |
