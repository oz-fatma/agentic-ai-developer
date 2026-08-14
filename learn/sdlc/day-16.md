# Gün 16: Operate, Improve & Capstone — Full-Cycle Capstone Plan

**Feature:** Study Buddy — *Weekly reminder when behind on goal* (post-MVP v0.2.0)

---

## 1. Pick a Feature

**Feature:** Send a push notification on Thursday if the student has logged < 50% of their weekly study goal.

**Why this feature:** Motivation loop; builds on existing session + goal data; small vertical slice.

---

## 2. Full Packet

### Requirements

**Problem:** Students forget to study mid-week and miss weekly goals.

**User story:** As a student, I want a reminder when I'm behind on my weekly goal, so that I can catch up before the week ends.

**Acceptance criteria:**
- Given a weekly goal and < 50% progress by Thursday 18:00 (user timezone), when the job runs, then user receives one push notification
- Given user opted out of notifications, then no push sent
- Given user already at ≥ 50%, then no push sent
- Given user already notified this week, then no duplicate push

**Constraints:** iOS + Android push APIs; user consent required; job runs once daily

### Design sketch

```
Cron job (daily) → Query users with goal + progress < 50% + Thursday + not notified
    → Push service (FCM/APNs) → Mobile app
```

**Components:** `ReminderJob`, `NotificationService`, `user_notification_prefs` table

### Build / test plan

| Task | Layer |
|---|---|
| Migration: notification prefs | DB |
| PATCH /users/me/notification-settings | API |
| ReminderJob with unit tests (mock clock) | Backend |
| Integration: job → mock push provider | API |
| Mobile: permission prompt + settings toggle | UI |
| E2E: opt-in → simulate Thursday → receive notification | E2E |

### Release checklist

- [ ] Staging: job dry-run mode (log only, no push)
- [ ] Prod: feature flag `weekly_reminder` — enable 10% users first
- [ ] Rollback: disable flag; no migration rollback needed
- [ ] Changelog entry for v0.2.0

### Ops notes

- Monitor: job success rate, push delivery failures, opt-out rate
- Alert: job fails 2 days in a row
- Runbook: disable flag in admin; link in on-call doc

---

## 3. Model Choice

**Choice:** **Hybrid (Agile delivery)**

**Why:**
- Small feature fits one 2-week sprint — Agile iteration
- Push permissions need a **spike** (half day) on iOS vs Android — research before build
- Release uses **flagged rollout** — controlled release discipline from Waterfall/hybrid ops

Not pure Waterfall (requirements will tune threshold 50% → maybe 40% after data).  
Not pure Agile chaos (staging + flag + rollback required for user-facing notifications).

---

## 4. Reflection

**Phase I underestimate most:** **Requirements / discovery**

**Why:** I jump to implementation ("just add a cron job") before clarifying timezone, duplicate prevention, and opt-out — causing rework in testing.

**How I'll practice next:**
1. Write acceptance criteria **before** opening IDE
2. Run "ambiguity hunt" on every vague word ("behind", "reminder", "week")
3. Do 5-minute stakeholder question pass even for solo projects

---

## SDLC Track Complete — 16/16

| Phase | Days | Status |
|---|---|---|
| SDLC Fundamentals | 1–4 | ✅ |
| Requirements & Design | 5–8 | ✅ |
| Build, Quality & Release | 9–12 | ✅ |
| Operate, Improve & Capstone | 13–16 | ✅ |

**Capstone deliverable:** Full life-cycle packet for *weekly reminder* feature — ready to guide a real sprint.
