# Copy Templates — Job Posting Trigger

## Cadence Overview

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Hiring observation + pain from JD + bridge to solution |
| 2 | LinkedIn | Day 1 | Team growth observation, no pitch |
| 3 | Email | Day 4 | Value-add: how peers handle this need (build vs buy framework) |
| 4 | Email | Day 9 | Social proof: team that achieved more with fewer hires |
| 5 | Email | Day 14 | Breakup: different angle — team productivity, not headcount |

---

## Email 1 — Hiring Observation (Day 0)

**Subject line formula:** `{function} at {{company_name}}` or `Building {{department}} teams in {{company_industry}}`
- ≤45 chars, about their challenge, not about us

**Structure:**
```
{{copy_opener_trigger}}

{{copy_body}}

{{copy_cta}}

{sender_signature}
```

**What makes this work:**
- Opener branches by `hook_type` (see `shared/references/hook-types-guide.md`):
  - **Timeline** (default): "Teams building out [function] go from [state] to [result] in [timeframe] — often before the new hire starts."
  - **Numbers**: "{{company_industry}} teams see [X% improvement] in [JD area] — usually before the req even closes."
  - **Social proof**: "[Named company] was in the same position — growing [dept]. They [result]."
  - **Hypothesis** (fallback): "Building out [function] at {{company_name}} — the [pain from JD] is usually what drives that."
- Shows you understand WHY they're hiring (pain from JD), not just that they are
- Positions solution as helping the team, not replacing the hire
- CTA calibrated by ICP tier AND signal composite score

---

## Email 2 — Build vs Buy Value (Day 4)

**Subject line formula:** `{function_challenge} — build or buy?`

**Structure:**
```
{{copy_opener_company}}

{Value: "Most {{company_industry}} teams hiring for [function] face a choice: build the capability from scratch or layer in tools that get the team productive faster. We've seen {{similar_companies}} take the second approach and [specific result]."}

{"Happy to share what we've seen work — might save your new hire's first few months."}

{sender_signature}
```

**Prompt instructions:**
- Company-based opener (different from Email 1)
- Frame as strategic insight about building capabilities, not product pitch
- Reference what the JD says they're building
- 90 words max

---

## Email 3 — Social Proof (Day 9)

**Subject formula:** `How {{peer_company}} scaled {function} without adding headcount`

**Structure:**
```
{Bridge to team scaling challenge}

{Case study: specific results — "[Company] was hiring for the same function. After adopting {client_product}, their [metric] improved by [%] and they [specific outcome]."}

{Connection to their situation}

{{copy_cta}}

{sender_signature}
```

---

## Email 4 — Breakup (Day 14)

**Subject formula:** `Last thought on {{department}} at {{company_name}}`

**Structure:**
```
{Completely different angle from Emails 1-3 — different pain point from icp-mapping.md}

{Brief reframe: approach from team productivity angle rather than hiring angle}

{"If the timing's off, no worries — wanted to make sure this was on your radar."}

{sender_signature}
```

---

## Key Differences from Other Triggers

- **Target contact:** The HIRING MANAGER (one level above posted role), not the person being hired
- **Posting freshness matters:** <7 days = high urgency, 7-21 days = standard, >21 days = lower priority
- **Never imply "don't hire"** — always position as "help the team / new hire succeed"
- **JD is your research source** — extract specific pain language and mirror it back
