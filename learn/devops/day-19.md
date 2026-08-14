# Day 19: Continuous Deployment Basics — Environments and Secrets

**Phase:** CI/CD — Continuous Integration & Deployment (Days 16-20)

## Phase Context

CD pushes verified artifacts to staging or production. Today you add deploy jobs gated by branch, protected environments, and GitHub secrets.

## Tasks

### 1. Environment-gated deploy job

```yaml
jobs:
  deploy-staging:
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    environment: staging
    steps:
      - uses: actions/checkout@v4
      - run: ./scripts/deploy.sh staging
```

### 2. Use GitHub secrets safely

```yaml
      - env:
          AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
          AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        run: aws s3 sync ./dist s3://my-bucket-staging/
```

### 3. Manual approval for production

```yaml
  deploy-prod:
    needs: deploy-staging
    environment:
      name: production
      url: https://app.example.com
    runs-on: ubuntu-latest
    steps:
      - run: ./scripts/deploy.sh production
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| Environment | Named target with protection rules |
| Secret | Encrypted variable, never logged |
| `if:` | Conditional job execution |
| OIDC | Prefer short-lived cloud tokens over keys |
| Rollback | Redeploy previous artifact or tag |

## Summary

Deploy only from trusted branches. Use GitHub environments for approvals and never commit credentials to the repository.
