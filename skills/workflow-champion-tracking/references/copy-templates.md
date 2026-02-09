# Copy Templates — Champion Tracking Workflow

## Cadence Overview

Champion tracking uses a SHORTER, WARMER cadence than cold outreach. These people know you — respect that.

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 (trigger fires) | Reconnection: subtle past reference + new company context |
| 2 | LinkedIn | Day 2 | Connection/message: genuine congrats, no pitch |
| 3 | Email | Day 7 | Value bridge: share something useful for their new role |
| 4 | Email | Day 14 | Soft close: open door, no pressure |

**Note:** 4 touches max, not 5-6. Champions deserve fewer, higher-quality touches.

---

## Email 1 — Reconnection (Day 0)

**Subject line formula:** `{first_name} + {new_company}`
- Keep casual — this is a person you know, not a cold prospect

**Structure:**
```
{{copy_opener_trigger}}

{{copy_body}}

{{copy_cta}}

{sender_signature}
```

**Total: ≤90 words**

**Relationship-tier examples:**

**Former customer (social_proof hook):**
```
At {previous_company}, your team cut review cycles by 35% in the first quarter.
{company_name}'s {industry} operation looks like it could benefit from a similar
approach, especially at your current scale. Worth a quick catch-up to see if
the playbook translates?
```

**Former prospect (timeline hook):**
```
Congrats on {company_name}. Teams in {{company_industry}} typically go from
first data audit to full pipeline visibility in 6 weeks — and given what
you explored at {previous_company}, the ramp would be even faster.
If {topic} comes up in your first quarter, happy to pick up where we left off.
```

**Engaged contact (hypothesis hook):**
```
Saw you made the move to {{company_name}} — congrats. Would I be right that
scaling {pain_area} is high on the priority list in a new {{contact_title}} role?
We put together a playbook on how {{company_industry}} teams approach this:
{blog_link}
```

---

## LinkedIn Connection/Message (Day 2)

**Rules:**
- ≤280 characters
- If already connected: message, not connection request
- Pure reconnection — no product mentions

---

## Email 2 — Value Bridge (Day 7)

**Subject line formula:** `{resource_topic} for new {role/team} leads`

**Structure:**
```
{Brief opener — reference new context, not previous email}

{Value-add: something genuinely useful for their new role — industry report,
framework, intro, content piece. Must be relevant to NEW role.}

{Soft bridge: "Thought this might be useful given [new company context]."}

{sender_signature}
```

- Pure VALUE — no pitch, no ask
- Maximum 90 words
- Do NOT reference Email 1

---

## Email 3 — Soft Close (Day 14)

**Subject line formula:** `{topic} at {{company_name}}`

**Structure:**
```
{One-line context about their industry/role challenge}

{One sentence: what {client_name} could do for {company_name} specifically}

{Open-door CTA: "If [pain area] comes up as you settle in, our door is open.
Either way, rooting for you in the new role."}

{sender_signature}
```

- LAST touch — 60-70 words max
- Do NOT say "just following up"

---

## Variable Dependencies

| Variable | Source Column | Required For |
|---|---|---|
| `{{copy_opener_trigger}}` | AI prompt (branched by relationship × hook_type) | Email 1 |
| `{{copy_body}}` | AI prompt | Email 1 |
| `{{copy_cta}}` | AI prompt (calibrated by relationship + timing) | Email 1, 3 |
| `{{copy_linkedin}}` | AI prompt | LinkedIn |
| `{{champion_first_name}}` | Import | All |
| `{{champion_previous_company}}` | Import | Opener context |
| `{{champion_relationship_type}}` | Import/CRM | Copy branching |
| `{{champion_result_achieved}}` | Import/CRM | Social proof hook |
| `{{champion_days_in_role}}` | Formula | Urgency calibration |
| `{{company_name}}` | Import | All |
| `{{contact_title}}` | Import | Context |
| `{{company_industry}}` | Enrichment | Context |
| `{{research_new_company}}` | Claygent | Body generation |
| `{{research_transition_insight}}` | AI prompt | Opener context |
| `{{hook_type}}` | AI classification | Copy branching |
| `{{signal_composite_score}}` | Signal aggregation | CTA calibration |
