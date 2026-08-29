# Day 76: Containers & CI/CD — Docker Multi-stage

## 1. Multi-stage Dockerfile

Build stage + minimal runtime image.

## 2. Copy Binary

Only artifact in final stage.

## 3. Small Image

Alpine/scratch-based runtime.

## 4. Build Command

docker build -f day76/Dockerfile .

## Run

```bash
cd learn/go
go run ./day76
```

## Summary

| Concept | Takeaway |
|---|---|
| Multi-stage Dockerfile | Build stage + minimal runtime image |
| Copy Binary | Only artifact in final stage |
| Small Image | Alpine/scratch-based runtime |
| Build Command | docker build -f day76/Dockerfile  |
