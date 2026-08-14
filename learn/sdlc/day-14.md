# Gün 14: Operate, Improve & Capstone — Maintenance, Monitoring, and Incidents

**Product:** Study Buddy API (post-launch)

## 1. Maintenance Types

| Type | Definition | Study Buddy example |
|---|---|---|
| **Corrective** | Fix defects | Session duration wrong when crossing midnight |
| **Adaptive** | Adjust to environment changes | iOS 18 breaks timer background behavior — update app |
| **Perfective** | Improve performance/UX | Optimize weekly aggregation query (was slow at 1000+ sessions) |
| **Preventive** | Reduce future failures | Add DB index on `(user_id, started_at)` before slow queries hurt users |

## 2. Signals — Simple Web API

**Metrics:**
- Request rate (req/s) per endpoint
- Error rate (5xx %, 4xx %)
- Latency p50, p95, p99 (`GET /sessions`, `POST /sessions`)
- DB connection pool usage

**Logs:**
- Structured JSON logs with `request_id`, `user_id` (hashed if needed), `endpoint`, `status`, `duration_ms`
- Error stack traces for 5xx

**Alerts (examples):**
- 5xx rate > 1% for 5 minutes → page on-call
- p95 latency > 2s for 10 minutes → warning
- DB connections > 80% → warning

**Business (optional):**
- Daily active users, sessions logged per day, weekly goal completion rate

## 3. Incident Outline — Mini Timeline

**Incident:** "Users cannot end study sessions — PATCH /sessions/:id returns 500"

| Phase | Time | Action |
|---|---|---|
| **Detect** | T+0 | Alert: 5xx spike on PATCH /sessions; user tweet report |
| **Mitigate** | T+10m | Identify bad deploy v0.1.1; rollback to v0.1.0 |
| **Verify** | T+15m | Smoke test end session on prod — works |
| **Fix** | T+2h | Root cause: null pointer when `ended_at` already set; PR with fix + test |
| **Review** | T+3d | Post-incident: add idempotent PATCH test; staging gate before prod |

## 4. On-call Empathy — What Helps

Information that makes incidents easier:

- **Runbook link** in alert — how to rollback Study Buddy
- **Recent deploys** — what changed in last 2 hours
- **Dashboards** — error rate, latency, one-click log search by `request_id`
- **Ownership** — who owns sessions module; escalation path
- **Rollback command** — documented one-liner or pipeline button
- **Status comms template** — "Investigating" / "Mitigated" / "Resolved"

**Avoid:** Alerts with no context ("something broke"); logs without correlation IDs.
