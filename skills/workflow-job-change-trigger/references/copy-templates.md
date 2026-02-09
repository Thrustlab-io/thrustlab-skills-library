# Copy Templates — Job Role Change Trigger

## Cadence Overview

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Congratulations + transition-specific observation + soft bridge |
| 2 | LinkedIn | Day 0 | Connection request: congrats + shared interest |
| 3 | Email | Day 5 | "First 90 days" value: insight for their new role |
| 4 | LinkedIn | Day 6 | Brief value share if connected |
| 5 | Email | Day 10 | Social proof: peer who made a similar transition |
| 6 | Email | Day 17 | Team-level angle: what their team needs, not just them |

---

## Email 1 — Transition Opener (Day 0)

**Subject line formula:** `{{contact_title}} at {{company_name}}` or `The move to {{company_name}}`
- ≤45 characters — keep it about them, not about us

**Structure:**
```
{{copy_opener_trigger}}

{{copy_body}}

{{copy_cta}}

{sender_signature}
```

**What makes this work:**
- Opener branches by `hook_type` (see `shared/references/hook-types-guide.md`):
  - **Timeline** (default): "30 days into a new VP role — teams at this stage see [result] within [timeframe] when they prioritize [area]."
  - **Numbers**: "VP-level leaders who [action] in their first quarter see [X% improvement] in [metric]."
  - **Social proof**: "When [named person] made a similar move to [company type], they [result]."
  - **Hypothesis** (fallback): "The move from [context A] to [context B] usually means [challenge] is top of mind."
- Body connects new-role challenges to client's outcome
- CTA calibrated by ICP tier AND signal composite score — low-commitment (they're overloaded in a new role)

---

## Email 2 — First 90 Days Value (Day 5)

**Subject line formula:** `{priority_topic} in your first quarter`

**Structure:**
```
{{copy_opener_company}}

{Value paragraph: "When [persona title]s join a new team, [specific challenge] usually surfaces in the first [timeframe]. Here's what we've seen work: [framework/insight/data point]."}

{Soft close: "Happy to share more detail if this is on your radar."}

{sender_signature}
```

**Prompt instructions:**
- Lead with company-based opener (vary from Email 1's transition opener)
- Value-add must be specific to their ROLE and INDUSTRY (from persona card in icp-mapping.md)
- Frame as helpful onboarding content, not product pitch
- Reference what new [title]s typically face — demonstrate you understand their world
- 90 words max, no meeting CTA

---

## Email 3 — Peer Transition Story (Day 10)

**Subject line formula:** `{peer_name_or_title} did this after joining {similar_company}`

**Structure:**
```
{Bridge: "A few weeks into a new [title] role is usually when [specific priority] starts moving up the list."}

{Peer story: "[Name/Title] joined [company] in a similar move — from [type of company] to [type of company]. They used {client_product} to [specific outcome with number]."}

{Connection: "Your transition from {{trigger_previous_company}} looks similar — [specific detail]."}

{{copy_cta}}

{sender_signature}
```

**Prompt instructions:**
- Peer story must use actual client case study (from profile.md)
- Emphasize the TRANSITION similarity, not just the company similarity
- Include a measurable result
- 90 words max

---

## LinkedIn Connection Request (Day 0)

```
{{copy_linkedin}}
```

**Rules:**
- ≤280 characters
- Brief congrats on the move + a shared industry connection
- Sent same day as Email 1 but must feel independent
- NO reference to email, NO pitch

---

## LinkedIn Follow-Up (Day 6, if connected)

**Template:**
```
Thanks for connecting, {{contact_first_name}}. How's the transition to {{company_name}} going? {Brief relevant resource or insight about their role's challenges.}
```

- ≤280 characters
- Casual, genuine curiosity about the transition
- Small value-add tied to new role

---

## Email 4 — Team-Level Angle (Day 17)

**Subject line formula:** `{team_function} at {{company_name}}`

**Structure:**
```
{Shift perspective: address what their TEAM needs, not just them personally}

{New angle: "New [title]s usually inherit [challenge]. Your [department] team at {{company_name}} is probably [dealing with X]. {client_name} helps teams like yours [specific outcome]."}

{Breakup energy: "If this isn't a priority right now, totally get it — wanted to make sure you had this on your radar before things get hectic."}

{sender_signature}
```

**Prompt instructions:**
- Different angle from Email 1-3: about their team, not their personal transition
- Use a different pain point from icp-mapping.md (not the same as Email 1)
- Breakup energy = give permission to not respond
- 70-80 words max (shorter than other emails)

---

## Timing Adjustments Based on Days in Role

| Days in Role | Cadence Adjustment |
|---|---|
| 0-14 days | Standard cadence — they're still excited about new role |
| 15-30 days | Delay Email 1 by 3 days — let them settle |
| 31-60 days | Skip congrats language — acknowledge they're "getting into the groove" |
| 61-90 days | Treat almost like general outbound — the transition is old news |
| 90+ days | Do NOT use this trigger play — switch to general outbound or another trigger |

---

## Variable Dependencies

| Variable | Source Column | Required For |
|---|---|---|
| `{{trigger_previous_company}}` | Import/Enrichment | Opener, body context |
| `{{trigger_previous_title}}` | Import/Enrichment | Opener, body context |
| `{{trigger_days_in_role}}` | Formula | Timing + tone adjustment |
| `{{trigger_role_transition_type}}` | AI prompt | Angle selection |
| `{{trigger_new_role_context}}` | Claygent | Research for body |
| `{{research_previous_company_relevance}}` | AI prompt | Experience angle |
| All standard variables | Enrichment + AI | Same as other triggers |
