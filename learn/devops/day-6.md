# Day 6: Version Control with Git — Basics

**Phase:** Git (Days 6–10)

## Tasks

### 1. Initialize Repository

```bash
mkdir my-project && cd my-project
git init
git status
```

### 2. Stage and Commit

```bash
echo "# My Project" > README.md
git add README.md
git commit -m "Initial commit"
git log --oneline
```

### 3. Workflow

```bash
# Edit file
echo "New line" >> README.md
git status          # modified, not staged
git add README.md   # staged
git commit -m "Add new line"
git log --oneline
```

## Cheat Sheet

| Command | Purpose |
|---|---|
| `git init` | Start tracking directory |
| `git add file` | Stage changes |
| `git commit -m "msg"` | Save snapshot |
| `git status` | See working tree state |
| `git log` | View history |

## Summary

Git tracks every change. Workflow: edit → `add` → `commit`. Always check `git status` first.
