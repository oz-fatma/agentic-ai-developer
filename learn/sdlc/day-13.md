# Gün 13: Operate, Improve & Capstone — Deployment and Environments

**Product:** Study Buddy

## 1. Environment Map

| Environment | Purpose | Who uses it | Data |
|---|---|---|---|
| **Local** | Developer machine; fast iteration | Developers | Fake/seed data; disposable |
| **Staging** | Pre-prod validation; mirrors prod config | Dev, QA, PM | Anonymized or synthetic users |
| **Production** | Real users | Students | Real PII — protected |

**Responsibilities:**
- **Local:** break things freely; no uptime SLA
- **Staging:** must be stable enough for QA sign-off; deploy from `main` or release branch
- **Production:** changes only via promotion; monitored; rollback ready

## 2. Promotion Path

```
Developer laptop (local)
    → PR merge to main
    → CI builds artifact
    → Deploy to STAGING (auto or manual trigger)
    → QA smoke + acceptance tests
    → Manual approval gate
    → Deploy same artifact to PRODUCTION (no rebuild)
    → Post-deploy smoke test
```

**Rule:** Same binary/artifact promotes staging → prod — don't rebuild differently per env.

## 3. Config Awareness — Per Environment

| Config | Local | Staging | Production |
|---|---|---|---|
| `DATABASE_URL` | localhost Postgres | staging DB host | prod DB (secrets manager) |
| `API_BASE_URL` | http://localhost:3000 | https://api-staging.studybuddy.app | https://api.studybuddy.app |
| `JWT_SECRET` | dev-only secret | staging secret | strong rotatable secret |
| `LOG_LEVEL` | debug | info | warn |
| Feature flags | all on for dev | match prod or test flags | prod values |
| Analytics | disabled or test bucket | test bucket | prod bucket |

**Never:** copy prod secrets to laptop; commit `.env` to git.

## 4. Unsafe Shortcut vs Safer Alternative

**Unsafe:** Deploy untested hotfix straight to production because "it's just one line."

**Risks:**
- No CI run — might not compile on prod build
- No staging verification — breaks auth for all users
- No audit trail — hard to rollback
- "One line" often touches shared code paths

**Safer alternative:**
1. Fix on branch → CI passes → deploy to **staging** (5 min)
2. Smoke test critical path on staging
3. Deploy **same artifact** to prod
4. Monitor error rate for 15 minutes
5. Document in incident/changelog

**True emergency (P0 down):** Minimal fix + staging if time allows; else hotfix with extra reviewer, immediate rollback plan, post-incident review mandatory.
