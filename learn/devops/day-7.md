# Day 7: Version Control with Git — Branching

**Phase:** Git (Days 6–10)

## Tasks

### 1. Create and List Branches

```bash
git branch                    # list local branches
git branch feature/login      # create branch
git branch -a                 # all branches
```

### 2. Switch Branches

```bash
git checkout feature/login    # switch (legacy)
git switch feature/login      # switch (modern)
git switch -c feature/signup  # create + switch
```

### 3. Merge Branches

```bash
git switch main
git merge feature/login
git log --oneline --graph
```

## Workflow Example

```bash
git switch -c feature/navbar
# make changes, commit
git switch main
git merge feature/navbar
git branch -d feature/navbar   # delete merged branch
```

## Summary

Branches isolate work. Feature branch → develop → merge to main. Never develop directly on main in teams.
