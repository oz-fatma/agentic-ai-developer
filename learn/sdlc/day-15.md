# Gün 15: Operate, Improve & Capstone — Retrospectives and Continuous Improvement

**Project:** Study Buddy MVP (Sprint 1–3 retrospective)

## 1. Retro Format — Personal Retro

| Went well | Improve | Actions |
|---|---|---|
| MVP shipped on time with core session + goal flows | Requirements for timezone/week boundary were unclear until mid-sprint | Document week boundary rule in spec before sprint 2 |
| PR reviews caught auth bug early | E2E tests added too late — found UI bug in staging | Add one E2E test per story in same PR |
| Team communication on Slack was clear | Too many scope additions mid-sprint (reminder feature creep) | Lock backlog during sprint; new items go to next sprint |
| CI pipeline saved us from broken migration | Changelog was written at last minute | Template changelog in release checklist |

## 2. Actionable Outcomes

| Complaint | Specific action | Owner | Date |
|---|---|---|---|
| "We always find bugs in staging" | Add integration test for every API endpoint before merge | Dev team | Next sprint start |
| "PM changes priorities daily" | Weekly backlog grooming; sprint goal locked after planning | PM + lead | Every Monday |

## 3. Measure Once — Process Metric

**Metric chosen:** **Escaped defects** — bugs found in production per release

**Target:** ≤ 1 P1/P2 escaped defect per MVP release

**How to track:** Label prod bugs in issue tracker; count at end of each release cycle

**Why:** Directly reflects testing + review quality in Build/Quality phase

## 4. Feedback Loops — User → Requirements

```
Production users
    → In-app feedback / app store reviews / support tickets
    → PM triages weekly (themes, frequency)
    → New user stories added to backlog
    → Prioritized in next sprint planning
    → Requirements phase (AC updated) → Design (if needed) → Build → Release
    → Measure impact (retention, NPS) → Retro
```

**Study Buddy example:**
- Users request "pause session" → story created → AC written → sprint 4
- Crash reports → corrective maintenance → hotfix or next sprint

**SDLC loop closed:** Operate/Maintenance feeds back into Requirements.
