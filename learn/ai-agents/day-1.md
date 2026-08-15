# Day 1: Introduction to AI Agents and Autonomous Systems

**Track:** AI Agents · **Phase:** Introduction (Days 1–5)

## 1. What Is an AI Agent?

An **AI agent** is an autonomous entity that:
- **Perceives** its environment (via sensors, APIs, user input)
- **Decides** what to do next
- **Acts** to achieve a goal (via tools, code, APIs)

**Regular program vs agent:**

| Regular program | AI agent |
|---|---|
| Fixed if/else logic | Adapts based on context |
| Same input → same output | Can plan multi-step actions |
| No tools unless hardcoded | Uses LLM + tools dynamically |
| No memory across runs | Can remember prior context |

**Simple definition:** An AI agent = **LLM + goal + tools + loop** (observe → think → act → repeat).

## 2. History — Symbolic AI to LLM Agents

| Era | Approach | Example |
|---|---|---|
| 1950s–80s | Symbolic AI — rules & logic | Expert systems, chess engines |
| 1990s–2010s | ML agents — learn from data | Recommendation systems |
| 2020s+ | LLM agents — language + tools | ChatGPT with plugins, Cursor, AutoGPT |

Modern agents use **LLMs** as the reasoning engine and connect to **tools** (search, code, APIs) to act in the real world.

## 3. Types of Agents

| Type | Behavior | Example |
|---|---|---|
| **Simple reflex** | If condition → action | Thermostat: temp > 25 → turn on AC |
| **Model-based** | Keeps internal state | Robot vacuum maps the room |
| **Goal-based** | Plans steps to reach goal | Navigation app finds route to destination |
| **Utility-based** | Picks best option by score | Uber chooses driver by ETA + price |
| **Learning** | Improves from feedback | Spam filter learns from labels |

**Smart thermostat example:** A learning, model-based agent — it tracks your schedule (model), aims for comfort (goal), and adjusts over time (learning).

## 4. Key Components

```
┌─────────────┐     perceive      ┌─────────────┐
│ Environment │ ◄─────────────── │  AI Agent   │
│ (user, web, │                  │             │
│  APIs, DB)  │ ───────────────► │ LLM (brain) │
└─────────────┘     act (tools)  │ Tools       │
                                  │ Memory      │
                                  └─────────────┘
```

| Component | Role |
|---|---|
| **Perception** | Read user message, API response, file content |
| **Decision-making** | LLM decides next step (answer, call tool, ask clarifying question) |
| **Action** | Execute tool: search web, run code, send email |
| **Memory** | Short-term (conversation) + long-term (vector DB, files) |

## 5. Autonomous Systems

An **autonomous system** operates and makes decisions **without constant human intervention**.

Agents are building blocks of autonomous systems:
- **Single agent:** Research assistant that searches and summarizes
- **Multi-agent:** Planner agent + coder agent + reviewer agent working together
- **Human-in-the-loop:** Agent proposes, human approves (production best practice)

## 6. Vocabulary Drill

**AI Agent:** Autonomous entity that perceives, decides, and acts to achieve goals.

**Autonomous System:** System that operates without continuous human control.

**LLM (Large Language Model):** AI model that understands and generates human-like text — the "brain" of modern agents.

## 7. My Definition (in my own words)

<!-- Write 3–5 sentences in your own words -->

## 8. Scenario Practice

**Scenario:** A smart thermostat learns when you come home and pre-heats the room.

- **Agent type:** Model-based + learning (+ goal-based)
- **Perception:** Temperature sensor, time, motion
- **Action:** Adjust heating
- **Goal:** Comfort at lowest energy cost

## Notes

- Agents are not magic — they need clear goals, good tools, and guardrails.
- Production agents always include **human oversight** for high-risk actions.
- Next (Day 2): Agentic mindset — LLMs, tools, memory in depth.
