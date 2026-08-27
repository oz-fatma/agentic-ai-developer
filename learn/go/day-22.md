# Day 22: Testing Fundamentals — Table-Driven Tests

## 1. Build Test Tables

Defined slices of input/expected pairs in `stringutil_test.go`.

## 2. Use t.Run

Created subtests with `t.Run` for each table row.

## 3. Cover Edge Cases

Included empty strings, single chars, and truncate boundaries.

## 4. Keep Tables Readable

Kept expectations visible in the table struct.

## Run

```bash
cd learn/go
go test ./day22/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Build Test Tables | Defined slices of input/expected pairs in `stringutil_test |
| Use t.Run | Created subtests with `t |
| Cover Edge Cases | Included empty strings, single chars, and truncate boundarie |
| Keep Tables Readable | Kept expectations visible in the table struct |
