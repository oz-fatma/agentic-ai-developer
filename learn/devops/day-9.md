# Day 9: Version Control with Git — Conflict Resolution and Rebase

**Phase:** Git (Days 6–10)

## Tasks

### 1. Merge Conflicts

When two branches edit the same lines:

```bash
git merge feature/conflict
# CONFLICT in file.txt
# edit file — remove <<<<<<, ======, >>>>>> markers
git add file.txt
git commit -m "Resolve merge conflict"
```

### 2. Conflict Markers

```
<<<<<<< HEAD
main branch text
=======
feature branch text
>>>>>>> feature/conflict
```

Keep the correct version (or combine), then `git add`.

### 3. Rebase vs Merge

```bash
git switch feature/login
git rebase main               # replay commits on top of main
```

| | Merge | Rebase |
|---|---|---|
| History | Preserves branches | Linear history |
| Use when | Shared/public branches | Local feature cleanup |

**Rule:** Never rebase commits already pushed to shared main.

## Summary

Conflicts are normal — resolve manually, then commit. Rebase for clean history on local branches only.
