# Day 18: Building and Testing in CI — Jobs, Caching, and Matrix

**Phase:** CI/CD — Continuous Integration & Deployment (Days 16-20)

## Phase Context

CI must run tests reliably and fast. Today you add test jobs, cache dependencies, and use matrix builds for multiple language versions.

## Tasks

### 1. Add test job with coverage

```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: "3.12"
      - run: pip install pytest pytest-cov
      - run: pytest --cov=app tests/
```

### 2. Cache pip dependencies

```yaml
      - uses: actions/cache@v4
        with:
          path: ~/.cache/pip
          key: pip-${{ runner.os }}-${{ hashFiles('requirements.txt') }}
```

### 3. Matrix strategy for multiple versions

```yaml
    strategy:
      matrix:
        python-version: ["3.11", "3.12"]
    steps:
      - uses: actions/setup-python@v5
        with:
          python-version: ${{ matrix.python-version }}
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| Matrix | Run same job across variants |
| Cache | Reuse downloaded dependencies |
| `needs:` | Job dependency chain |
| Artifact upload | Share build output between jobs |
| Fail fast | Stop matrix on first failure (default) |

## Summary

Fast CI uses caching and parallel matrix jobs. Tests run on every PR — a broken main branch is not acceptable in DevOps culture.
