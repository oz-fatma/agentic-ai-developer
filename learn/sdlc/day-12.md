# Day 12: Build, Quality & Release — Release Readiness and CI Basics

**Product:** Study Buddy — MVP v0.1.0 release

## 1. Release Checklist

- [ ] All MVP stories merged; acceptance criteria signed off
- [ ] DB migrations tested on staging (forward + rollback plan documented)
- [ ] Feature flags: N/A for MVP (or flag off for incomplete features)
- [ ] Secrets in env vars / vault — not in repo
- [ ] CI green on `main` (build, lint, unit, integration)
- [ ] Smoke test on staging: login → log session → see progress
- [ ] Rollback plan written (see below)
- [ ] Changelog / release notes drafted
- [ ] On-call or owner identified for first 24h post-launch
- [ ] Monitoring/alerts configured (error rate, 5xx)

## 2. CI Purpose

**Continuous Integration** runs automated checks on every push/PR merge:

| Gate | What it catches |
|---|---|
| Build | Code compiles; dependencies resolve |
| Lint | Style issues, common bugs |
| Unit tests | Broken logic in isolation |
| Integration tests | Broken API/DB contracts |
| (Optional) E2E | Broken user flows |

**SDLC role:** CI enforces quality gates so humans focus on judgment (design, edge cases), not "did someone run tests?" Forgotten steps become automatic blockers before merge.

**Study Buddy CI pipeline (example):**
```
push/PR → install deps → lint → unit tests → integration tests (with test DB) → build artifact
```

## 3. Rollback Plan

**Scenario:** v0.1.0 deployed; session POST returns 500 for all users.

| Step | Action |
|---|---|
| 1 | **Detect** — alert fires on 5xx spike / user reports |
| 2 | **Decide** — incident lead confirms rollback vs hotfix (< 5 min) |
| 3 | **Rollback** — redeploy previous artifact (v0.0.9) OR revert migration if safe |
| 4 | **Verify** — smoke test: login + log session works |
| 5 | **Communicate** — status page / team channel update |
| 6 | **Follow-up** — root cause PR; add regression test before re-release |

**Mitigation if rollback hard:** Feature flag disable session endpoint; show maintenance message.

## 4. Version Note — Changelog Entry

```markdown
## [0.1.0] — 2026-08-14

### Added
- User registration and login
- Log study session (start/end with duration)
- Set weekly study hour goal
- Home screen weekly progress bar
- 7-day study history view

### Fixed
- N/A (initial release)

### Known issues
- Offline session logging not supported (planned v0.2.0)
```
