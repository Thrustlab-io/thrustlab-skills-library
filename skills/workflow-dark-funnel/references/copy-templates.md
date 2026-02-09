# Copy Templates — Dark Funnel Engagement Workflow

## Cadence Overview

Dark funnel outreach is SOFTER and more VALUE-FORWARD than triggered outreach. The prospect hasn't raised their hand — you're meeting them where they are in their research journey.

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 (signal detected) | Value-forward: share relevant content |
| 2 | LinkedIn | Day 1 | Content share: thought leadership, not connection pitch |
| 3 | Email | Day 4 | Insight share: proprietary data or framework on the topic |
| 4 | Email | Day 9 | Soft engagement: thought-provoking question |

**Note:** 4 touches max. Dark funnel prospects are early in their journey — don't overwhelm.

**GOLDEN RULE:** No email in this cadence should make the prospect think "how did they know I was looking at this?" Every touch must feel like natural thought leadership that happens to be perfectly timed.

---

## Email 1 — Value-Forward Content Share (Day 0)

**Subject line formula:** `{content_topic} in {{company_industry}}`
- Lead with the topic, not {client_name}
- Feel like a newsletter from a knowledgeable peer, not a sales email

**Structure:**
```
{{copy_opener_trigger}}

{{copy_body}}

{{copy_cta}}

{sender_signature}
```

**Total: ≤90 words**

**Framing examples by hook type:**

**Timeline hook:**
```
{{company_industry}} teams tackling {topic} typically go from initial assessment
to measurable improvement in 6-8 weeks. We just mapped the timeline for companies
at {{company_name}}'s stage — thought you'd find the milestones useful: {content_match}
```

**Hypothesis hook (most common for dark funnel):**
```
Curious — as {{company_name}} scales, is {inferred_pain} becoming a priority?
We've been researching how {{company_industry}} teams approach this and put together
a framework: {content_match}. Thought it might be relevant given {{company_name}}'s
trajectory.
```

---

## LinkedIn Content Share (Day 1)

**Template:**
```
{{copy_linkedin}}
```

**Rules:**
- ≤280 characters
- Share the SAME content piece as Email 1 — it's a parallel touch
- Frame as "thought you'd find this interesting" not "I'm selling to you"
- If they engage with the LinkedIn share, that's a SECOND dark funnel signal → escalate in signal stacking

---

## Email 2 — Proprietary Insight (Day 4)

**Subject line formula:** `{X}% of {{company_industry}} teams {surprising_finding}`

**Structure:**
```
{Brief opener — new angle on the same topic, not reference to Email 1}

{Proprietary insight: share a data point, benchmark, or counterintuitive finding
from {client_name}'s experience working with similar companies. This positions
{client_name} as an expert, not a vendor.}

{One sentence connecting the insight to {{company_name}}'s specific situation.}

{No CTA in this email — pure value. End with "More research coming next week"
or similar to create anticipation.}

{sender_signature}
```

**Prompt instructions:**
- The insight must be REAL — based on actual client data or industry research
- Position {client_name} as a thought leader with unique perspective
- No CTA — this is a pure value touch. The goal is to be remembered, not to book a meeting.
- Maximum 90 words

---

## Email 3 — Soft Engagement Question (Day 9)

**Subject line formula:** `quick question about {topic} at {{company_name}}`

**Structure:**
```
{Open with a thoughtful question about how {{company_name}} approaches the topic}

{Brief context: "We're mapping how {{company_industry}} teams handle [specific aspect
of the topic]. Your perspective would be valuable."}

{Soft CTA: "Open to a brief exchange on this? No pitch — genuinely interested
in how teams at {{company_name}}'s scale approach it."}

{sender_signature}
```

**Prompt instructions:**
- This email asks for THEIR input — flipping the dynamic from seller to researcher
- The question must be genuinely interesting and specific to their role
- "No pitch" framing is important — and must be HONEST (if they respond, engage on the topic first)
- Maximum 70 words

---

## Response Handling (Post-Cadence)

If they respond to any touch:
1. Engage on the TOPIC first — answer their question, share more insight
2. After 1-2 value exchanges, bridge: "This actually connects to what we do at {client_name}..."
3. Only THEN ask for a meeting

If they don't respond:
- Add to general outbound nurture with the topic tag
- If a new dark funnel signal fires later → start a fresh dark funnel play (new topic)

---

## Variable Dependencies

| Variable | Source Column | Required For |
|---|---|---|
| `{{copy_opener_trigger}}` | AI prompt (hook_type branched) | Email 1 |
| `{{copy_body}}` | AI prompt | Email 1 |
| `{{copy_cta}}` | AI prompt (calibrated by intensity) | Email 1, 3 |
| `{{copy_linkedin}}` | AI prompt | LinkedIn |
| `{{contact_first_name}}` | Enrichment | All |
| `{{contact_title}}` | Enrichment | Context |
| `{{company_name}}` | Signal source | All |
| `{{company_industry}}` | Enrichment | Context |
| `{{darkfunnel_source}}` | Import | Internal routing |
| `{{darkfunnel_engagement_type}}` | Import | Intensity classification |
| `{{darkfunnel_content_topic}}` | Import/AI | Content matching + copy angles |
| `{{darkfunnel_intensity}}` | AI classification | Quality gate + CTA calibration |
| `{{research_inferred_pain}}` | AI prompt | Opener + body |
| `{{research_company_context}}` | Claygent | Body generation |
| `{{content_match}}` | Lookup | Email 1, LinkedIn |
| `{{hook_type}}` | AI classification | Copy branching |
| `{{signal_composite_score}}` | Signal aggregation | Routing + CTA calibration |
