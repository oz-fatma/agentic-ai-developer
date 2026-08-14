# Day 8: Version Control with Git — Remote Repositories

**Phase:** Git (Days 6–10)

## Tasks

### 1. Clone

```bash
git clone https://github.com/user/repo.git
cd repo
```

### 2. Push

```bash
git remote add origin https://github.com/user/repo.git
git push -u origin main       # first push sets upstream
git push                      # subsequent pushes
```

### 3. Pull and Fetch

```bash
git fetch origin              # download changes, no merge
git pull origin main          # fetch + merge
git pull                      # if upstream set
```

## Remote Workflow

```
local edit → add → commit → push → remote
remote change → fetch/pull → local
```

## Cheat Sheet

| Command | Purpose |
|---|---|
| `git clone URL` | Copy remote repo locally |
| `git push` | Upload commits to remote |
| `git pull` | Download + merge remote changes |
| `git fetch` | Download only, no merge |

## Summary

Remote = collaboration hub (GitHub/GitLab). Push shares your work; pull brings team changes.
