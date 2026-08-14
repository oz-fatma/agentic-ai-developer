# Gün 10: Build, Quality & Release — Code Review and Collaboration

**Product:** Study Buddy

## 1. Review Checklist (10 items)

- [ ] **Correctness** — Does the code meet acceptance criteria? Edge cases handled (empty session, double end)?
- [ ] **Tests** — New logic has unit/integration tests; CI passes
- [ ] **Naming** — Variables, functions, endpoints are clear (`endSession` not `doThing`)
- [ ] **Security** — Auth on all endpoints; user can only access own sessions; no secrets in diff
- [ ] **Error handling** — API returns proper status codes (400, 401, 404); no silent failures
- [ ] **Scope** — PR does one story; no unrelated refactors
- [ ] **Performance** — No obvious N+1 queries; duration calc is O(1)
- [ ] **Readability** — A junior dev can follow the flow without asking
- [ ] **Docs** — API changes documented; PR description links to story
- [ ] **Rollback safety** — Migration reversible or forward-only with plan noted

## 2. Give Feedback — Sample Diff Practice

**Sample diff:** Developer adds `POST /sessions` but skips auth check — any user can create sessions for any `user_id` in body.

**Kind comment:**
> Nice clean handler structure — the validation on `started_at` is clear. One suggestion: let's derive `user_id` from the JWT instead of the request body so users can't spoof another account.

**Critical comment (blocking):**
> **Blocking:** Missing auth middleware on `POST /sessions`. Unauthenticated requests can create sessions. Please add auth and a test that returns 401 without token.

## 3. Receive Feedback — Disagree Professionally

**Review says:** "Rename `SessionService` to `StudySessionManager` for clarity."

**Professional response:**
> Thanks for the suggestion. I kept `SessionService` because it matches our existing `GoalService` naming pattern in this codebase. Happy to rename if we're standardizing across the module — could we decide team-wide in a follow-up PR rather than blocking this story?

**Principles:** Acknowledge → explain reasoning → offer compromise → don't take it personally.

## 4. PR Hygiene — Good PR Description

```markdown
## Summary
Implements session logging (start/end) for Study Buddy MVP — Story #1.

## User story
As a student, I want to log a study session with start/end time, so that I can track how long I studied.

## Changes
- Add `sessions` table migration
- POST /sessions (start), PATCH /sessions/:id (end)
- Mobile: Start/End study buttons + active timer
- Unit + integration tests

## How to test
1. Log in as test user
2. Tap "Start study" → timer runs
3. Tap "End study" → session appears in history
4. Run `npm test` — all green

## Screenshots
[optional: mobile flow]

## Checklist
- [x] Tests added
- [x] No secrets in code
- [x] Linked to story SB-1
```
