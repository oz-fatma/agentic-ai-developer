# Day 7: Requirements & Design — High-Level Design and Architecture Sketch

**Product:** Study Buddy (MVP scope: log sessions, weekly goal, home progress)

## 1. Context Diagram

```
                    ┌─────────────────┐
                    │  Student (user) │
                    └────────┬────────┘
                             │ mobile app
                             ▼
                    ┌─────────────────┐
                    │   Study Buddy   │
                    │     System      │
                    └────────┬────────┘
                             │
         ┌───────────────────┼───────────────────┐
         │                   │                   │
         ▼                   ▼                   ▼
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│  Push notify    │ │  Email service  │ │   Analytics     │
│  (optional P3)  │ │  (auth/invite)  │ │  (usage metrics)│
└─────────────────┘ └─────────────────┘ └─────────────────┘
```

**Actors:** Student (primary user)  
**External systems (later phases):** push notifications, email, analytics — not all required for MVP.

## 2. Component Sketch

```
┌──────────────────────────────────────────────────────────┐
│                     Mobile App (UI)                       │
│  Screens: Home, Log Session, Goal Settings, History      │
└────────────────────────────┬─────────────────────────────┘
                             │ HTTPS / REST (or GraphQL)
                             ▼
┌──────────────────────────────────────────────────────────┐
│                      API Server                           │
│  Auth · Sessions · Goals · Progress aggregation          │
└────────────────────────────┬─────────────────────────────┘
                             │
              ┌──────────────┴──────────────┐
              ▼                             ▼
     ┌─────────────────┐          ┌─────────────────┐
     │    Database     │          │  Background job │
     │  Users, Sessions│          │  (optional)     │
     │  WeeklyGoals    │          │  Weekly reset,  │
     └─────────────────┘          │  reminder cron  │
                                    └─────────────────┘
```

| Component | Responsibility |
|---|---|
| **Mobile UI** | Present data, capture session start/end, show progress |
| **API** | Validate input, enforce auth, business rules, aggregate weekly totals |
| **Database** | Persist users, sessions, goals |
| **Background jobs** | Weekly boundary reset, reminder notifications (post-MVP) |

## 3. Data Sketch

**Main entities:**

| Entity | Key fields | Relationships |
|---|---|---|
| **User** | id, email, created_at | 1 → many Session, 1 → many WeeklyGoal |
| **Session** | id, user_id, started_at, ended_at, duration_minutes | belongs to User |
| **WeeklyGoal** | id, user_id, week_start, target_hours | belongs to User; one active goal per week |

**Relationships (text):**
- User **has many** Sessions
- User **has many** WeeklyGoals (one per calendar week)
- Weekly progress = **sum(session.duration)** for sessions where `week_start <= session.started_at < week_start + 7 days`

**Example query logic:** Home screen reads `WeeklyGoal` for current week + aggregated session minutes for that user/week.

## 4. Trade-off Note

**Decision:** **Monolithic API + single PostgreSQL database** for MVP.

**Alternative rejected:** Microservices (separate auth service, session service, goal service) from day one.

**Why we chose monolith:**
- Small team, 6-week MVP — ops overhead of microservices is too high
- Traffic expected low at launch — no need to scale services independently yet
- Simpler debugging and deployment for first release

**Cost of choice:**
- Harder to split later if one area (e.g. analytics) needs different scaling
- Mitigation: keep clear module boundaries inside the monolith (auth, sessions, goals packages) so extraction is possible later

**When to revisit:** If DAU > threshold or team grows past ~8 engineers working on same deploy.
