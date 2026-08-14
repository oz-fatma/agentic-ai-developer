# Day 20: End-to-End CI/CD Pipeline — Build, Test, Deploy

**Phase:** CI/CD — Continuous Integration & Deployment (Days 16-20)

## Phase Context

Capstone of the CI/CD phase: a full pipeline from commit to deployed container. Today you wire lint, test, build, and deploy together in one workflow.

## Tasks

### 1. Complete pipeline workflow

```yaml
name: E2E Pipeline
on:
  push:
    branches: [main]
  pull_request:

jobs:
  lint-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: pip install ruff pytest
      - run: ruff check .
      - run: pytest

  build-push:
    needs: lint-test
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: docker build -t myapp:${{ github.sha }} .
```

### 2. Tag releases on main

```bash
git tag -a v1.0.0 -m "First release"
git push origin v1.0.0
```

### 3. Verify pipeline on a PR

```bash
git checkout -b feature/ci-pipeline
git push -u origin feature/ci-pipeline
# Open PR — confirm lint-test runs; merge to trigger deploy
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| E2E pipeline | Lint → test → build → deploy |
| `needs:` | Serialize jobs after tests pass |
| Immutable tags | Deploy by SHA or semver tag |
| PR checks | Block merge if CI fails |
| Pipeline as code | Workflow YAML lives in repo |

## Summary

A production pipeline gates deploys on green tests. You completed the CI/CD phase — next up: AWS fundamentals.
