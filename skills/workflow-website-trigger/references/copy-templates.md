# Copy Templates — Website Visitor Trigger

## Cadence Overview

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 (trigger fires) | Intent-aware: inferred pain + value prop |
| 2 | LinkedIn | Day 1 | Connection request: company observation, no pitch |
| 3 | Email | Day 3 | Value-add: relevant resource/insight tied to pages visited |
| 4 | LinkedIn | Day 4 | Follow-up if connected: brief value share |
| 5 | Email | Day 7 | Social proof: case study from same vertical |
| 6 | Email | Day 12 | Breakup: different angle, final touch |

---

## Email 1 — Intent-Aware (Day 0)

**Subject line formula:** `{industry_specific_topic} at {{company_name}}`
- Subject must be ≤45 characters
- No clickbait, no "Quick question", no "{first_name},"
- Reference the problem space, not {client_name}

**Structure:**
```
{{copy_opener_trigger}}

{{copy_body}}

{{copy_cta}}

{sender_signature}
```

**Total: ≤90 words**

**What makes this email work:**
- Opener branches by `hook_type` (see `shared/references/hook-types-guide.md`):
  - **Timeline** (default): "SaaS teams evaluating [category] typically go from first audit to full visibility in 6 weeks."
  - **Numbers**: "[Industry] teams see 40% fewer manual reviews after [approach] — that gap compounds at {{company_size}}+ scale."
  - **Social proof**: "[Named company] tackled the same evaluation and saw [result] within [timeframe]."
  - **Hypothesis** (fallback): "Teams in {{company_industry}} exploring [category] usually hit [specific challenge] first."
- Body links inferred pain → specific client outcome
- CTA calibrated by ICP tier AND signal composite score (if available)

---

## Email 2 — Value-Add Resource (Day 3)

**Subject line formula:** `{resource_topic} for {{company_industry}} teams`

**Structure:**
```
{{copy_opener_company}}

{Value-add paragraph: share a resource, data point, or insight relevant to the pages they visited. This should feel helpful, not salesy.}

{Soft bridge: "Thought this might be useful given [industry trend]. Happy to share more if relevant."}

{sender_signature}
```

**Prompt instructions for this email:**
- Lead with the company-based opener (not trigger-based — vary the approach)
- The value-add should connect to pages visited:
  - Pricing page → share ROI calculator, cost comparison, or benchmark
  - Solution page → share relevant use case or framework
  - Case study page → share another case study from their vertical
  - Blog → share additional content on the same topic
- Maximum 90 words total
- NO CTA asking for a meeting — this is pure value

---

## Email 3 — Social Proof (Day 7)

**Subject line formula:** `How {{similar_company}} handled {pain_topic}`

**Structure:**
```
{One-line bridge referencing previous email or their industry context}

{Social proof: specific result from a company in their vertical. Format: "[Company similar to them] was dealing with [pain]. After [using client product], they [specific measurable outcome]."}

{Tie back: "{{company_name}} seems to be in a similar position with [specific detail]."}

{CTA — slightly more direct than Email 1 since this is touch 3}

{sender_signature}
```

**Prompt instructions:**
- The similar company reference must come from {client_name}'s actual case studies (from profile.md)
- If no case study for their exact vertical, use the closest available + acknowledge it
- Result must include a number (%, $, time saved, etc.)
- Maximum 90 words

---

## LinkedIn Connection Request (Day 1)

**Template:**
```
{{copy_linkedin}}
```

**Rules:**
- ≤280 characters (hard LinkedIn limit)
- Company observation or shared industry interest
- NO pitch, NO mention of emails, NO "I also sent you an email"
- This runs parallel to email — they should feel like independent touches

---

## LinkedIn Follow-Up (Day 4, if connected)

**Template:**
```
Thanks for connecting, {{contact_first_name}}. {One sentence about something relevant from their profile or recent activity.} {One sentence offering a specific insight or resource — same theme as Email 2 but shorter.}
```

**Rules:**
- ≤280 characters
- Reference something from their LinkedIn profile (Claygent can research this)
- Offer value, don't pitch
- If not connected, skip this step entirely

---

## Email 4 — Breakup / Different Angle (Day 12)

**Subject line formula:** `{different_angle_topic}` (completely different from previous subjects)

**Structure:**
```
{Different opener angle — if previous emails led with pain, lead with opportunity. If they led with company observation, try industry trend.}

{Brief, different framing of the value prop — approach from a new persona's perspective or different pain point from icp-mapping.md}

{Low-pressure close: "Either way, no hard feelings — just wanted to make sure this was on your radar."}

{sender_signature}
```

**Prompt instructions:**
- This email must feel DIFFERENT from emails 1-3 — new angle, new energy
- Use the second-ranked pain point from icp-mapping.md (first was used in Email 1)
- Shorter than other emails — aim for 60-70 words
- The breakup framing gives permission to not respond, which paradoxically increases responses

---

## Variable Dependencies

Each copy template requires these Clay columns to be populated:

| Variable | Source Column | Required For |
|---|---|---|
| `{{copy_opener_trigger}}` | AI prompt column | Email 1 |
| `{{copy_opener_company}}` | AI prompt column | Email 2 |
| `{{copy_body}}` | AI prompt column | Email 1 |
| `{{copy_linkedin}}` | AI prompt column | LinkedIn request |
| `{{copy_cta}}` | AI prompt column | Email 1, 3 |
| `{{contact_first_name}}` | Enrichment | All |
| `{{contact_title}}` | Enrichment | Prompt context |
| `{{company_name}}` | Enrichment | All |
| `{{company_industry}}` | Enrichment | Prompt context |
| `{{research_company_snapshot}}` | Claygent | Opener generation |
| `{{research_pain_inference}}` | AI prompt | Body generation |
| `{{trigger_pages_visited}}` | Intent tool import | Trigger-specific prompts |
| `{{trigger_visit_recency_hours}}` | Formula | Urgency calibration |
| `{{trigger_page_intent_score}}` | AI prompt | Filtering + angle selection |
