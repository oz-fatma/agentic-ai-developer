# Day 34: First REST API Project (MVP) — Validation and Error Responses

## 1. Validate Input

Rejected empty titles and invalid IDs.

## 2. Map Errors to Status

400 for validation, 404 for missing, 500 for internal.

## 3. Centralize Error Writing

Created `writeError` helper for JSON problem responses.

## 4. Test Unhappy Paths

Demo covers validation and not-found cases.

## Run

```bash
cd learn/go
go run ./day34
```

## Summary

| Concept | Takeaway |
|---|---|
| Validate Input | Rejected empty titles and invalid IDs |
| Map Errors to Status | 400 for validation, 404 for missing, 500 for internal |
| Centralize Error Writing | Created `writeError` helper for JSON problem responses |
| Test Unhappy Paths | Demo covers validation and not-found cases |
