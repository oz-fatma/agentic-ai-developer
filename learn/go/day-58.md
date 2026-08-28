# Day 58: Project Layout & Architecture — Dependency Injection

## 1. Constructor Wiring

Built services with explicit dependencies.

## 2. Interface Boundaries

Depended on interfaces for repositories and clocks.

## 3. Avoid Service Locator

No hidden global registries.

## 4. Test with Fakes

Injected fakes to test services without databases.

## Run

```bash
cd learn/go
go test ./day58/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Constructor Wiring | Built services with explicit dependencies |
| Interface Boundaries | Depended on interfaces for repositories and clocks |
| Avoid Service Locator | No hidden global registries |
| Test with Fakes | Injected fakes to test services without databases |
