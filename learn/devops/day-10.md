# Day 10: Version Control with Git — Practice Project

**Phase:** Git (Days 6–10)

## Tasks — Full Git Workflow

### 1. Setup

```bash
mkdir devops-site && cd devops-site
echo "<h1>Hello</h1>" > index.html
git init
git add .
git commit -m "Initial commit"
```

### 2. Feature Branch

```bash
git switch -c feature/navbar
echo "<nav>Home | About</nav>" >> index.html
git add index.html
git commit -m "Add navbar"
git switch main
git merge feature/navbar
```

### 3. Remote

```bash
git remote add origin https://github.com/user/devops-site.git
git push -u origin main
git pull
```

### 4. Checklist

- [ ] `git init` and first commit
- [ ] Feature branch created and merged
- [ ] (Optional) Conflict resolved manually
- [ ] Pushed to remote
- [ ] Pulled remote changes

## Summary

You now know the full Git cycle: init → branch → commit → merge → push → pull. **Git phase complete.**
