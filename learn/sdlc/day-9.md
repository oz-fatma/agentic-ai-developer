# Gün 9: Build, Quality & Release — Implementation Discipline

**Product:** Study Buddy  
**Sample story:** *As a student, I want to log a study session with start/end time, so that I can track how long I studied.*

## 1. Slice Work — Implementation Tasks

| # | Layer | Task | Output |
|---|---|---|---|
| 1 | **API** | `POST /sessions` — create session (start_time) | Endpoint + validation |
| 2 | **API** | `PATCH /sessions/:id` — end session (end_time, compute duration) | Endpoint + duration logic |
| 3 | **API** | `GET /sessions?week=current` — list sessions for aggregation | Query + tests |
| 4 | **DB** | Migration: `sessions` table (user_id, started_at, ended_at, duration_minutes) | Migration file |
| 5 | **UI** | "Start study" button → calls POST, shows active timer | Screen state |
| 6 | **UI** | "End study" button → calls PATCH, shows summary | Screen state |
| 7 | **Tests** | Unit tests for duration calculation (edge: midnight, timezone) | Test file |
| 8 | **Tests** | Integration test: start → end → session appears in GET | CI green |
| 9 | **Docs** | API doc for session endpoints + example payloads | README or OpenAPI snippet |

**Vertical slice order:** DB (4) → API (1, 2, 3) → UI (5, 6) → Tests (7, 8) → Docs (9)

Ship when 1–8 pass DoD — docs can follow in same PR or immediately after.

## 2. Branch Mentality — Safe Workflow

```
main (protected)
  └── feature/log-study-session
        ├── commit: add sessions migration
        ├── commit: POST /sessions endpoint
        ├── commit: PATCH /sessions/:id end session
        ├── commit: mobile start/end UI
        └── commit: tests for session flow
```

**Rules:**
1. Branch from latest `main`: `feature/log-study-session`
2. **Small commits** — one logical change per commit (easy to review, easy to revert)
3. Open **PR** when vertical slice works locally
4. **No direct push to main** — merge via PR after review
5. Rebase or merge `main` into branch if others merged meanwhile — resolve conflicts early
6. Delete branch after merge

**Commit message example:**
```
feat(sessions): add POST /sessions to start study session
```

## 3. Definition of Done — Coding

Implementation is **done** when ALL of the following are true:

- [ ] Code implements acceptance criteria for the story
- [ ] Unit tests added/updated for new logic (duration, validation)
- [ ] Integration or API test covers happy path
- [ ] Linter/formatter passes (`npm run lint` / `go vet` / equivalent)
- [ ] No secrets, debug logs, or commented-out junk left in diff
- [ ] PR opened with description linking to user story
- [ ] **At least one peer review** approved (or self-review checklist if solo)
- [ ] README or API doc updated if public contract changed
- [ ] Manually tested on target device (start → end → see session)

## 4. Spike vs Build

| Situation | Spike first? | Why |
|---|---|---|
| Standard CRUD session log with known stack | **No — build** | Pattern is well understood |
| Real-time sync across 2 devices while session active | **Yes — spike** | Conflict resolution unclear; 2–4 hour spike on sync strategy |
| Offline session log, sync when online | **Yes — spike** | Need to choose: local DB (SQLite) vs queue; test offline UX |
| Weekly progress bar UI | **No — build** | Design exists; standard UI work |
| Push reminder at 80% of week with no goal progress | **Maybe spike** | Platform permissions (iOS/Android) differ — 1-day spike on notification API |

**Spike rules:**
- **Timeboxed** (e.g. 2–4 hours, max 1 day)
- **Goal:** answer a specific question — output is notes or throwaway code, not production
- **Decision at end:** proceed, pivot, or cut scope

**Study Buddy decision for MVP session logging:** **Build directly** — no spike needed; defer offline sync to v2 unless user research demands it.
