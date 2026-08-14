# Day 4: Linux Fundamentals — Environment Variables and Processes

**Phase:** Linux Fundamentals (Days 1–5)

## Tasks

### 1. Environment Variables

```bash
echo $HOME
echo $PATH
export APP_ENV=production
env | grep APP
```

### 2. View and Set

```bash
export DB_HOST=localhost
export DB_PORT=5432
echo "Connecting to $DB_HOST:$DB_PORT"
```

### 3. Monitor Processes

```bash
ps aux                       # all processes
ps aux | grep nginx
top                          # real-time (q to quit)
```

### 4. Manage with `kill`

```bash
kill PID                     # SIGTERM — graceful stop
kill -9 PID                  # SIGKILL — force stop
killall nginx                # by name (careful)
```

## Cheat Sheet

| Command | Purpose |
|---|---|
| `export VAR=value` | Set env var for child processes |
| `env` | List all environment variables |
| `ps aux` | List running processes |
| `top` | Live process monitor |
| `kill PID` | Send signal to process |

## Summary

Env vars configure apps without hardcoding. Process management is essential for debugging stuck services.
