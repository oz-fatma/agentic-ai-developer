# Day 37: Context, Config & Middleware — Environment Config

## 1. Read env Vars

Loaded PORT, LOG_LEVEL from environment with defaults.

## 2. Support .env Locally

Documented env vars; defaults for local dev.

## 3. Validate on Startup

Fail fast when required settings are invalid.

## 4. Struct Config

Mapped env vars into typed `Config` struct.

## Run

```bash
cd learn/go
go test ./day37/...
```

## Summary

| Concept | Takeaway |
|---|---|
| Read env Vars | Loaded PORT, LOG_LEVEL from environment with defaults |
| Support .env Locally | Documented env vars; defaults for local dev |
| Validate on Startup | Fail fast when required settings are invalid |
| Struct Config | Mapped env vars into typed `Config` struct |
