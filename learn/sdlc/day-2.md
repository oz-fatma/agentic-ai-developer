# Gün 2: SDLC Fazları Uçtan Uca

## 1. Phase Walkthrough — Input & Output

| Phase | Input | Output |
|---|---|---|
| Requirements | Business need, user feedback, stakeholder requests | Requirements doc, user stories, acceptance criteria |
| Design | Requirements, constraints, goals | Architecture, wireframes, API specs, technical design |
| Implementation | Design documents, specs | Working code, unit tests |
| Testing | Code, acceptance criteria | Test reports, bug list, QA sign-off |
| Deployment | Tested release candidate, release plan | Live software in production |
| Maintenance | Production system, user feedback, incidents | Patches, updates, bug fixes, new features |

## 2. Handoffs — "Done" Ne Demek?

### Design → Implementation

Design is **done** when: wireframes/specs are reviewed and approved, acceptance criteria are clear, and developers have no blocking open questions. Implementation **starts** when the team can build against a stable, agreed design.

### Testing → Release

Testing is **done** when: critical test cases pass, known bugs are triaged (fixed or accepted), and QA gives sign-off. Release **starts** when there is a tested build, release notes, and a rollback plan.

## 3. Roles Snapshot

| Role | Heavily influences which phase(s)? |
|---|---|
| PM / Product Owner | Requirements, prioritization; feedback loop in Maintenance |
| Designer | Design |
| Developer | Implementation (input to Design, output to Testing) |
| QA | Testing |
| DevOps | Deployment, Maintenance (monitoring, incidents, releases) |

## 4. Mini Timeline — Küçük Bir Özellik

**Feature:** Add a "dark mode" toggle to a mobile app

```
[Request]  →  [Requirements]  →  [Design]  →  [Build]  →  [Test]  →  [Production]
User asks   User story +      UI mockup +   Code the    QA verifies  Deploy to
for dark    acceptance        toggle        toggle +    toggle works live users
mode        criteria          placement     theme switch  on devices
```
