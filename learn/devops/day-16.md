# Day 16: CI/CD Introduction — Pipelines and DevOps Flow

**Phase:** CI/CD — Continuous Integration & Deployment (Days 16-20)

## Phase Context

CI/CD automates build, test, and deploy on every change. Today you map pipeline stages and set up a minimal workflow skeleton in your repository.

## Tasks

### 1. Define pipeline stages

```yaml
# pipeline-stages.yaml (conceptual)
stages:
  - lint
  - test
  - build
  - deploy-staging
  - deploy-production
```

### 2. Create a minimal GitHub Actions workflow

```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]
jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: echo "Pipeline triggered on $GITHUB_REF"
```

### 3. Commit and trigger CI

```bash
mkdir -p .github/workflows
git add .github/workflows/ci.yml
git commit -m "Add CI skeleton"
git push
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| CI | Integrate and test on every commit |
| CD | Deploy automatically after CI passes |
| Pipeline | Automated sequence of stages |
| Artifact | Build output stored between stages |
| Gate | Stage that must pass before next |

## Summary

CI/CD removes manual release steps. Start with lint and test on every PR, then add build and deploy stages incrementally.
