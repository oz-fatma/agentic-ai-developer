# Day 21: Testing Fundamentals — Writing Your First Tests

## 1. Create _test.go Files

Added `day21/mathutil/` with production code and `mathutil_test.go` in the same package.

## 2. Use testing Package

Wrote `TestAdd`, `TestMultiply`, and other `TestXxx` functions.

## 3. Assert with Comparisons

Compared got/want values using `t.Errorf` on mismatch.

## 4. Test Exported Behavior

Focused tests on exported functions: `Add`, `Multiply`, `IsEven`, `Abs`.

## Run

```bash
cd learn/go
go test ./day21/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Create _test.go Files | Added `day21/mathutil/` with production code and `mathutil_t |
| Use testing Package | Wrote `TestAdd`, `TestMultiply`, and other `TestXxx` functio |
| Assert with Comparisons | Compared got/want values using `t |
| Test Exported Behavior | Focused tests on exported functions: `Add`, `Multiply`, `IsE |
