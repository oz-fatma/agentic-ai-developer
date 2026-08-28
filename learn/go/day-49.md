# Day 49: Databases (II) & Repositories — Connection Pooling

## 1. Tune Pool Settings

Set `MaxOpenConns`, `MaxIdleConns`, connection lifetime.

## 2. Watch Saturation

Logged wait counts when pool is busy.

## 3. Close on Shutdown

Closed DB pool on exit.

## 4. Read Driver Notes

Applied SQLite-specific pooling defaults.

## Run

```bash
cd learn/go
go run ./day49
```

## Summary

| Concept | Takeaway |
|---|---|
| Tune Pool Settings | Set `MaxOpenConns`, `MaxIdleConns`, connection lifetime |
| Watch Saturation | Logged wait counts when pool is busy |
| Close on Shutdown | Closed DB pool on exit |
| Read Driver Notes | Applied SQLite-specific pooling defaults |
