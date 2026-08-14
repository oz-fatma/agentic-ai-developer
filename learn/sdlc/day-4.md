# Gün 4: SDLC Fundamentals — Practice

## 1. Case Study

**Product idea:** **Study Buddy** — a mobile app where students log study sessions, set weekly goals, and optionally share progress with a small friend group.

**Chosen model:** **Hybrid** (Agile delivery + upfront design spike)

**Justification (5 bullets):**
1. Requirements will evolve after real student feedback — pure Waterfall is too rigid.
2. No heavy regulation (unlike banking) — we can iterate quickly without audit gates on every change.
3. We still need a short **design spike** upfront (auth, data model, core screens) so the team doesn't rewrite architecture every sprint.
4. **2-week sprints** fit a small team: ship "log a session" in sprint 1, "weekly goals" in sprint 2, "friend sharing" in sprint 3.
5. **Hybrid** balances speed with enough planning to avoid chaos in a student-facing product.

## 2. Phase Checklist — Artifacts per Phase

| Phase | Expected artifacts |
|---|---|
| **Requirements** | Problem statement, 5–10 user stories, acceptance criteria for MVP |
| **Design** | Wireframes (3–5 screens), simple data model, API sketch |
| **Implementation** | Source code, README, unit tests for core logic |
| **Testing** | Test plan, manual test checklist, bug list, QA sign-off for MVP |
| **Deployment** | Release notes, app store / deploy config, rollback steps |
| **Maintenance** | Feedback backlog, crash/error monitoring, patch release plan |

**One-page MVP checklist (Study Buddy):**
- [ ] User can sign up / log in
- [ ] User can start and end a study session with duration
- [ ] User can set a weekly hour goal
- [ ] App shows progress toward weekly goal
- [ ] Critical paths tested on iOS + Android (or web)
- [ ] Privacy note for friend-sharing feature documented

## 3. Risk Flag

**Riskiest phase:** **Requirements** (building the wrong MVP)

**Why:** Students might want social features more than logging — or the opposite. If we guess wrong, we waste sprints on low-value work.

**Mitigation:**
1. Interview 5–10 students before sprint 1; validate the core problem.
2. Ship the smallest usable version (session logging only) in week 2–3 and measure retention.
3. Keep the backlog ordered by value — defer "friend sharing" until logging proves useful.
4. Define explicit **acceptance criteria** per story so "done" is not vague.

## 4. Teach-back — 2-Minute Script (Non-Developer)

> "SDLC is basically the journey software takes from idea to something people actually use.
>
> First you figure out **what** people need — that's requirements. Then you **design** how it will look and work. Then developers **build** it. **Testers** check it works. Then you **release** it so users can download or use it online. After launch, you **maintain** it — fix bugs, add features, keep it running.
>
> Some teams plan everything upfront like a waterfall — one phase after another. Others work in **short cycles** called sprints: build a small piece, show it, get feedback, improve. Most real teams mix both.
>
> The point isn't bureaucracy — it's shipping something **useful** and **reliable** without building the wrong thing or breaking it for users."

**~90 seconds when read at a calm pace.**
