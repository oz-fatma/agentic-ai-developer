# Gün 8: Requirements & Design — Practice Project & Review

**Product:** Study Buddy — MVP one-pager spec

---

## 1. Mini Spec (One-Pager)

### Problem
Students struggle to know if they study enough each week. They need a simple way to log time, set a weekly target, and see progress without a complex planner.

### Users
- **Primary:** University / high-school students (16–25), smartphone users
- **Secondary (later):** Study groups / friends (post-MVP)

### Top 5 User Stories (MVP)

| # | Story |
|---|---|
| 1 | As a student, I want to log a study session with start/end time, so that I can track how long I studied. |
| 2 | As a student, I want to set a weekly study hour goal, so that I know my target. |
| 3 | As a student, I want to see my progress toward this week's goal on the home screen, so that I stay motivated. |
| 4 | As a student, I want to view my study history for the last 7 days, so that I can reflect on habits. |
| 5 | As a student, I want to sign up and log in securely, so that my data is mine. |

### Acceptance Criteria (MVP gate — Story 1 + 3)

**Log session:** Given logged in, when user starts and ends a session, then duration is saved and appears in weekly total within 1 second.

**Home progress:** Given a weekly goal and ≥1 session this week, when user opens home, then they see `hours_logged / target_hours` and a progress bar.

### Constraints
- **Time:** MVP in ~6 weeks
- **Team:** 2 devs, 1 designer
- **Tech:** Mobile app + monolith API + PostgreSQL
- **Legal:** KVKK/GDPR — minimal PII, clear privacy policy
- **Out of scope for MVP:** friend sharing, push reminders, calendar sync

---

## 2. Design Attachments

### Context

```
Student → Study Buddy App → API → PostgreSQL
                ↓
         (future: push, analytics)
```

### Components

| Layer | Responsibility |
|---|---|
| Mobile UI | Home, Log Session, Goal Settings, History, Auth |
| API | Auth, sessions CRUD, goals CRUD, weekly aggregation |
| DB | users, sessions, weekly_goals |
| Jobs (post-MVP) | reminders, weekly reset notifications |

### Data (core)

- **User** 1—* *→ **Session** (started_at, ended_at, duration)
- **User** 1—* *→ **WeeklyGoal** (week_start, target_hours)

---

## 3. Review Pass

**Reviewer:** Self-review (simulate peer review)

| Section | Unclear / issue | Revision |
|---|---|---|
| Stories | Story 5 light on AC | Added: email+password signup, session token, logout |
| Constraints | "6 weeks" — which platforms? | Clarified: iOS + Android **or** responsive web for MVP — pick one in sprint 0 |
| Design | Weekly boundary timezone | Added: use user's device timezone for week_start |
| Out of scope | Friends mentioned in Day 6 | Explicitly deferred to v2 in spec |

**Revised AC for Story 5 (Auth):**
- User can register with email + password (min 8 chars)
- User can log in and receive a session token
- Invalid credentials show generic error (no email enumeration)
- User can log out and token is invalidated client-side

---

## 4. Ready-to-Build Check

| Question | Ready? | Notes |
|---|---|---|
| Is the problem statement clear? | ✅ | Yes — weekly study visibility |
| Are MVP stories small and ordered? | ✅ | P1: log + goal + home; auth parallel sprint 0 |
| Are acceptance criteria testable? | ✅ | Given/When/Then for core flows |
| Is architecture sketched? | ✅ | Monolith + mobile + Postgres |
| Are constraints and out-of-scope listed? | ✅ | Friends, push deferred |
| Could a dev start without guessing intent? | ✅ | Entities, screens, and API responsibilities defined |

**Definition of Done for Requirements & Design phase:**
- [x] One-pager spec complete
- [x] Design sketches attached
- [x] Review pass done with revisions
- [x] Ready-to-build sign-off

**Next phase (Day 9+):** Build, Quality & Release — implementation, code review, testing, CI.

---

## Phase 2 recap (Days 5–8)

| Day | Focus |
|---|---|
| 5 | Discovery questions, constraints, ambiguity |
| 6 | User stories, AC, epic split, priority |
| 7 | Context, components, data, trade-offs |
| 8 | Mini spec + review + ready-to-build |

**Faz 2 tamamlandı.**
