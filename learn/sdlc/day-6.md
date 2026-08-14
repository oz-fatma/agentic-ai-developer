# Day 6: Requirements & Design — User Stories and Acceptance Criteria

**Product:** Study Buddy

## 1. User Stories (5)

1. **As a** student, **I want to** log a study session with start and end time, **so that** I can track how long I studied.
2. **As a** student, **I want to** set a weekly study hour goal, **so that** I know what target I am working toward.
3. **As a** student, **I want to** see my progress toward this week's goal on the home screen, **so that** I stay motivated.
4. **As a** student, **I want to** receive a reminder if I am far behind my weekly goal, **so that** I don't forget to study.
5. **As a** student, **I want to** view my study history for the last 7 days, **so that** I can reflect on my habits.

## 2. Acceptance Criteria (2 stories)

### Story 2 — Set a weekly study hour goal

**Given** I am logged in  
**When** I open goal settings and enter a valid number of hours (1–40) and save  
**Then** my weekly goal is stored and shown on the home screen  

**Also:**
- Invalid input (0, negative, non-number) shows a clear error message
- Default goal for new users is empty until they set one
- Goal applies to the current calendar week (Mon–Sun or locale-appropriate)

### Story 3 — See progress on home screen

**Given** I have a weekly goal and at least one logged session this week  
**When** I open the home screen  
**Then** I see hours logged vs goal (e.g. "12 / 20 hours") and a visual progress indicator  

**Also:**
- If no goal is set, show a prompt to set one (not a broken empty state)
- Progress updates within 1 second after saving a new session
- Progress resets at the start of a new week

## 3. Split Large Stories (Epic → smaller stories)

**Epic:** *As a student, I want to share my study progress with friends, so that we motivate each other.*

| # | Smaller shippable story |
|---|---|
| 1 | As a student, I want to invite a friend by email/link, so that we can connect in the app. |
| 2 | As a student, I want to see my friend's weekly hours (if they opt in), so that I can compare progress. |
| 3 | As a student, I want to control who can see my stats (private / friends only), so that I control my privacy. |
| 4 | As a student, I want to send a one-tap "nudge" to a friend, so that we remind each other to study. |

**Ship order:** 1 → 3 → 2 → 4 (privacy before visibility; nudge last)

## 4. Priority Pass

| Priority | Story | Why (value / risk) |
|:---:|---|---|
| **P1** | Log a study session | Core value — without this, nothing else matters |
| **P1** | Set weekly goal | Defines the main loop with logging |
| **P2** | See progress on home screen | High motivation value; depends on P1 stories |
| **P2** | View 7-day history | Useful feedback; lower risk than reminders |
| **P3** | Weekly reminder notification | Nice-to-have; needs permission handling + edge cases |
| **P4** | Friend sharing epic | Social layer — defer until MVP proves retention |

**Rule:** Ship **P1** first as MVP; validate retention before **P4**.
