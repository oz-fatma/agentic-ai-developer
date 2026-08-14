# Day 11: Build, Quality & Release — Testing in the Life Cycle

**Product:** Study Buddy — Story: log study session

## 1. Testing Levels

| Level | What it tests | Example (Study Buddy) |
|---|---|---|
| **Unit** | One function/class in isolation | `calculateDuration(start, end)` returns correct minutes; handles same-day vs midnight edge |
| **Integration** | Components together (API + DB) | `POST /sessions` then `PATCH /sessions/:id` → row in DB with correct duration |
| **E2E** | Full user flow through UI + backend | Playwright: login → start study → wait 2s → end → home shows updated weekly total |

## 2. Map Tests to Story

**Story:** Log a study session with start/end time

| Level | Tests to write |
|---|---|
| Unit | Duration calc; validate `started_at < ended_at`; reject negative duration |
| Integration | POST creates session; PATCH sets end; GET returns session; 401 without auth |
| E2E | Happy path on mobile; error when ending session twice |

**Priority:** Unit + integration before merge; E2E for MVP release candidate.

## 3. Bug Life Cycle

```
1. Report     → User/QA logs bug (title, steps, expected vs actual, env)
2. Triage     → PM/lead sets severity (P0–P3) and priority
3. Assign     → Developer picks up; links to story/bug ticket
4. Fix        → Branch, fix, tests, PR
5. Verify     → QA or reporter confirms on staging
6. Close      → Merge to main; deploy; mark closed; optional regression test added
```

**Example:** "Ending session twice shows 500" → P2 → fix PATCH idempotency → verify on staging → close.

## 4. Quality != Only QA

**Developer responsibilities before handoff:**
- Write unit tests for new logic — don't rely on QA to find basic breaks
- Run linter and test suite locally before opening PR
- Self-test happy path + one edge case manually
- Document known limitations in PR description
- Fix flaky tests — don't disable them without ticket

**QA adds:** exploratory testing, device matrix, regression suites, acceptance sign-off.

**Rule:** Quality is built in during Implementation and Testing phases — QA deepens it; developers don't outsource it entirely.
