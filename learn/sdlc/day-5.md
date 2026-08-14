# Gün 5: Requirements & Design — Gathering and Clarifying Needs

**Sample feature:** Study Buddy — *Weekly study goal tracking*

## 1. Stakeholder Questions (Discovery)

1. Who sets the goal — the student, a parent, or a tutor?
2. What time period matters most: daily, weekly, or per exam period?
3. How is "study time" defined — only logged sessions, or also reading offline?
4. What happens when a student misses their weekly goal — reminder, streak loss, or nothing?
5. Should goals be private by default or visible to friends?
6. Do we need different goal types (hours vs sessions vs topics completed)?
7. What devices will students use most — phone, tablet, web?
8. Is offline logging required (study without internet)?
9. What does success look like for the product owner — retention, daily active users, completed goals?
10. Are there school/university policies about tracking student activity we must respect?

## 2. Problem vs Solution

| User problem (keep) | Premature solution (avoid too early) |
|---|---|
| "I don't know if I'm studying enough each week." | "Build a React Native app with Firebase." |
| "I lose motivation when I can't see progress." | "Add a Redis cache for goal calculations." |
| "I forget to log sessions unless reminded." | "Use microservices for the goal module." |
| "I want a simple weekly target, not a complex planner." | "Integrate with Google Calendar on day one." |

**Rule:** Stay in the **problem space** during discovery. Tech choices come after requirements are clear.

## 3. Constraints

| Type | Constraints for Study Buddy |
|---|---|
| **Business** | MVP in ~6 weeks; small team (2 devs, 1 designer); free tier for students |
| **Time** | Launch before exam season; weekly goal feature in sprint 2, not sprint 5 |
| **Legal / privacy** | GDPR/KVKK-aware; minimal personal data; clear consent for friend sharing |
| **Technical** | Must work on mid-range Android phones; backend budget limited — prefer simple stack |

## 4. Ambiguity Hunt

**Vague request:** *"Make it faster."*

**Clarifying questions:**
1. Faster for **whom** — all users or a specific segment (e.g. slow networks)?
2. Faster **what** — app launch, screen load, saving a session, syncing goals?
3. How slow is it **today** — do we have metrics (p95 load time, crash rate)?
4. What is the **target** — under 2 seconds to open the home screen?
5. On which **devices/OS versions** was slowness reported?
6. Is the issue **network**, **server**, or **client UI**?
7. Does "faster" mean fewer taps, or literally lower latency?
8. Is this a **requirement** for MVP or a post-launch optimization?

**Rewritten requirement (example):**
> "Home screen should load and show the current weekly progress within **2 seconds** on a mid-range Android device over **4G**, for **95%** of sessions (measured in analytics)."
