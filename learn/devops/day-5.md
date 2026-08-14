# Day 5: Linux Fundamentals — Basic Shell Scripting

**Phase:** Linux Fundamentals (Days 1–5)

## Tasks

### 1. First Script

```bash
#!/bin/bash
# hello.sh
echo "Hello DevOps!"
```

### 2. Make Executable

```bash
chmod +x hello.sh
./hello.sh
```

### 3. Variables

```bash
#!/bin/bash
NAME="DevOps"
VERSION=1
echo "Welcome to $NAME v$VERSION"
```

### 4. Input/Output

```bash
#!/bin/bash
read -p "Enter your name: " NAME
echo "Hello, $NAME!"
```

### 5. Conditionals and Loops

```bash
#!/bin/bash
if [ -f "config.txt" ]; then
  echo "Config exists"
else
  echo "Config missing"
fi

for i in 1 2 3; do
  echo "Step $i"
done
```

## Cheat Sheet

| Element | Example |
|---|---|
| Shebang | `#!/bin/bash` |
| Variable | `VAR="value"` |
| Condition | `if [ -f file ]; then` |
| Loop | `for i in 1 2 3; do` |

## Summary

Shell scripts automate repetitive DevOps tasks — backups, deploys, health checks. **Phase 1 complete.**
