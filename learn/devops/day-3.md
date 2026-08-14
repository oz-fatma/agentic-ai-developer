# Day 3: Linux Fundamentals — Permissions and Users

**Phase:** Linux Fundamentals (Days 1–5)

## Tasks

### 1. Understand Permissions

```bash
ls -la
# -rw-r--r--  owner  group  file.txt
#  rwx r-x r-x  = user, group, others
```

### 2. Change Permissions with `chmod`

```bash
chmod u+x script.sh          # symbolic: user + execute
chmod go-w secret.txt        # group & others - write
chmod 755 script.sh          # octal: rwxr-xr-x
chmod 644 config.txt         # rw-r--r--
```

### 3. Change Ownership with `chown`

```bash
sudo chown user:group file.txt
sudo chown -R user:group ./directory
```

### 4. `sudo`, `whoami`, `id`

```bash
whoami
id
sudo apt update              # run as root (Linux)
```

## Octal Reference

| Value | Permission |
|---|---|
| 7 | rwx |
| 6 | rw- |
| 5 | r-x |
| 4 | r-- |

## Summary

Permissions control who can read, write, and execute. Never `chmod 777` on production — use least privilege.
