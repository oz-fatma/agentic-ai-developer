# Day 3: Waterfall, Agile, and Iterative Models

## 1. Waterfall Sketch

**Linear sequence:**

```
Requirements → Design → Implementation → Testing → Deployment → Maintenance
     ↓              ↓            ↓             ↓           ↓
  (complete)   (complete)   (complete)    (complete)  (complete)
```

Each phase largely finishes before the next one starts.

**2 strengths:**
1. Clear documentation and predictable milestones — easy to plan budget and timeline upfront.
2. Works well when requirements are stable and regulation/audit trails matter (e.g. banking, healthcare).

**2 weaknesses:**
1. Late feedback — users see working software only near the end; wrong assumptions are costly to fix.
2. Hard to adapt to change — new requirements mid-project often mean rework or long change requests.

## 2. Agile Sketch

**Short iterations:** Work happens in timeboxed cycles (sprints, e.g. 1–2 weeks). Each iteration delivers a small, usable increment of software.

**Feedback:** Stakeholders and users review working software often — not just documents. Teams adjust based on what they learn.

**Changing priorities:** Backlog can be reordered each iteration. The most valuable work ships first; less important items can wait or drop.

```
[ Sprint 1 ] → demo → feedback
[ Sprint 2 ] → demo → feedback
[ Sprint 3 ] → demo → feedback
      ...
```

**Core idea:** Working software + collaboration + responding to change over heavy upfront planning.

## 3. Compare for a Project

| Project | Better model | Why |
|---|---|---|
| **Bank core system** (accounts, transactions, compliance) | **Waterfall** (or strict hybrid) | Requirements and regulation are heavy; mistakes are high-risk; audit and documentation matter; change is controlled. |
| **Marketing landing page** (campaign site, A/B tests) | **Agile** | Requirements change often; fast feedback wins; low risk to ship small iterations; design/copy can pivot weekly. |

**Sample choice:** For a marketing landing page → **Agile**. Ship a minimal page in week 1, measure clicks, iterate copy and layout in week 2–3.

## 4. Hybrid Reality

Most real teams **do not** pick pure Waterfall or pure Agile. They mix:

- **Upfront planning** for architecture, compliance, and major milestones (Waterfall discipline)
- **Iterative delivery** for features in sprints with demos and retros (Agile rhythm)
- **Continuous deployment** for small, tested changes (DevOps culture)

**Example hybrid:** A bank builds a new mobile app with a fixed regulatory design phase (Waterfall), then delivers features in 2-week sprints (Agile), with automated tests and staged releases (Hybrid + DevOps).

**Takeaway:** Match the model to **risk**, **regulation**, and **rate of change** — not ideology.
