# Day 2: Linux Fundamentals — Viewing and Searching Files

**Phase:** Linux Fundamentals (Days 1–5)

## Tasks

### 1. View with `cat`

```bash
echo "line1\nline2\nline3" > sample.txt
cat sample.txt
```

### 2. Paginate with `less`

```bash
less sample.txt    # q to quit, Space next page, b previous
```

### 3. Head and Tail

```bash
head -n 3 sample.txt
tail -n 2 sample.txt
tail -f /var/log/system.log   # follow log in real time (Ctrl+C to stop)
```

### 4. Search with `grep`

```bash
grep "line2" sample.txt
grep -i "LINE" sample.txt      # case insensitive
grep -r "error" ./logs/        # recursive search
ps aux | grep node             # filter command output
```

## Cheat Sheet

| Command | Purpose |
|---|---|
| `cat file` | Print entire file |
| `less file` | Page through file |
| `head -n 5` | First 5 lines |
| `tail -n 5` | Last 5 lines |
| `tail -f` | Follow log live |
| `grep pattern file` | Search for text |

## Summary

DevOps engineers read configs and logs daily. `grep` and `tail -f` are your most-used troubleshooting tools.
