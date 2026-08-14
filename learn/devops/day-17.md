# Day 17: GitHub Actions Basics — Workflows, Jobs, and Steps

**Phase:** CI/CD — Continuous Integration & Deployment (Days 16-20)

## Phase Context

GitHub Actions runs workflows triggered by repository events. Today you build a workflow with jobs, reusable steps, and manual dispatch triggers.

## Tasks

### 1. Workflow with triggers and permissions

```yaml
name: Build
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
permissions:
  contents: read
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: "20"
      - run: npm ci && npm run build
```

### 2. Share data between steps

```yaml
      - run: echo "version=1.2.3" >> $GITHUB_OUTPUT
        id: meta
      - run: echo "Deploying ${{ steps.meta.outputs.version }}"
```

### 3. Manual workflow dispatch

```yaml
on:
  workflow_dispatch:
    inputs:
      environment:
        description: Target env
        required: true
        default: staging
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| `on:` | Events that trigger workflow |
| `jobs` | Parallel or sequential units |
| `runs-on` | Runner OS image |
| `uses:` | Reusable action |
| `secrets.` | Encrypted repo/org variables |

## Summary

Workflows combine triggers, jobs, and steps. Use official actions for checkout and tool setup; keep secrets out of logs.
