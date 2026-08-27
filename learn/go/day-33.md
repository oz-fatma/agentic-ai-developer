# Day 33: First REST API Project (MVP) — In-Memory Storage Layer

## 1. Create a Store Interface

Defined storage operations behind an interface.

## 2. Implement Mutex Protection

Guarded map with `sync.RWMutex`.

## 3. Generate IDs

Assigned incremental IDs on create.

## 4. Return Copies

Returned copies to avoid leaking internal state.

## Run

```bash
cd learn/go
go test ./day33/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Create a Store Interface | Defined storage operations behind an interface |
| Implement Mutex Protection | Guarded map with `sync |
| Generate IDs | Assigned incremental IDs on create |
| Return Copies | Returned copies to avoid leaking internal state |
