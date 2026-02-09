# Copy Templates — Competitor Customer Targeting Workflow

## Cadence Overview

Competitor customers need a specific escalation: plant doubt → show proof → differentiate → open door.

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Competitive seed: result that implies competitor can't match |
| 2 | LinkedIn | Day 2 | Category connection: talk about the space, not the switch |
| 3 | Email | Day 5 | Switch story: named customer who moved + their result |
| 4 | Email | Day 10 | Differentiation: specific capability comparison for THEIR use case |
| 5 | Email | Day 16 | Soft close: "even if happy, worth a comparison" |

---

## Email 1 — Competitive Seed (Day 0)

**Subject line formula:** `{differentiating_metric} for {{company_industry}} teams`
- Focus on the RESULT, not the switch — they don't know you know about their competitor yet

**Structure:**
```
{{copy_opener_trigger}}

{{copy_body}}

{{copy_cta}}

{sender_signature}
```

**Total: ≤90 words**

**What makes this email work:**
- Doesn't mention the competitor by name (yet) — just shares a result that creates doubt
- The prospect self-compares: "Are WE getting this kind of result?"
- Sets up Email 2's switch story naturally

---

## Email 2 — Switch Story (Day 5)

**Subject line formula:** `how {{competitor_customer_name}} approached {topic}`

**Structure:**
```
{One sentence connecting to Email 1's topic — but stands alone}

{Switch story: "{{competitor_customer_name}} was in a similar position — using [competitor approach]
for [workflow]. After switching to {client_product}, they saw {{competitor_customer_result}}
within [timeline]."}

{Bridge to prospect: "Given {{company_name}}'s scale in {{company_industry}}, the math
could be even more compelling."}

{Soft CTA: "Worth seeing if the numbers hold for your setup?"}

{sender_signature}
```

**Prompt instructions:**
- The switch customer MUST be real — never fabricate
- If no switch case study exists for this specific competitor, use a general case study + numbers hook
- The timeline makes it tangible — "6 weeks" not "quickly"
- Maximum 90 words
- Never disparage the competitor: "similar position" not "stuck with an inferior product"

---

## Email 3 — Differentiation Deep Dive (Day 10)

**Subject line formula:** `{specific_capability} at scale`

**Structure:**
```
{Open with their specific use case — reference {{research_competitor_usage}} if available}

{Differentiation: "Where teams like {{company_name}} typically hit a wall with [competitor approach]
is [specific limitation]. {client_name} handles this differently: [specific capability]."}

{Proof: "That's what drove [X%] improvement in [metric] for {{competitor_customer_name}}."}

{CTA: "Happy to walk through a side-by-side comparison specific to your workflow."}

{sender_signature}
```

**Prompt instructions:**
- This is the most TECHNICAL email in the cadence — specific capability, not generic positioning
- The limitation must be REAL and relevant to THIS company's scale/industry
- Always pair the limitation with client's specific alternative approach
- Maximum 90 words

---

## LinkedIn Connection (Day 2)

**Template:**
```
{{copy_linkedin}}
```

**Rules:**
- ≤280 characters
- Do NOT mention the competitor on LinkedIn — too aggressive for a semi-public channel
- Focus on the category/capability, not the switch
- Pattern: "Hi {{contact_first_name}}, we work with {{company_industry}} teams on [capability area]. Given {{company_name}}'s approach, thought there might be interesting overlap."

---

## Email 4 — Soft Close (Day 16)

**Subject line formula:** `either way` (or similar low-pressure)

**Structure:**
```
{Quick, different angle — maybe a new development in the category or a relevant trend}

{"Even if {{company_name}} is happy with your current setup, it's worth benchmarking
every [6 months / year] — the category is moving fast."}

{Open door: "If a side-by-side comparison is ever useful, happy to set one up. No pressure either way."}

{sender_signature}
```

**Prompt instructions:**
- This is a PERMISSION email — giving them a reason to evaluate without feeling disloyal
- "Even if you're happy" is important — it respects their choice
- 60-70 words max
- Do NOT reference previous emails

---

## Variable Dependencies

| Variable | Source Column | Required For |
|---|---|---|
| `{{copy_opener_trigger}}` | AI prompt (branched by hook_type) | Email 1 |
| `{{copy_body}}` | AI prompt | Email 1 |
| `{{copy_cta}}` | AI prompt | Email 1, 2, 3 |
| `{{copy_linkedin}}` | AI prompt | LinkedIn |
| `{{contact_first_name}}` | Enrichment | All |
| `{{contact_title}}` | Enrichment | Context |
| `{{company_name}}` | Enrichment | All |
| `{{company_industry}}` | Enrichment | Context |
| `{{competitor_product_used}}` | Claygent/BuiltWith | Core targeting |
| `{{competitor_confidence}}` | AI classification | Quality gate |
| `{{research_competitor_usage}}` | Claygent | Emails 1, 3 |
| `{{research_competitor_pain}}` | AI prompt | Email 3 |
| `{{research_switch_angle}}` | AI prompt | Emails 1, 2 |
| `{{competitor_customer_name}}` | Lookup/Static | Email 2 |
| `{{competitor_customer_result}}` | Lookup/Static | Email 2, 3 |
| `{{hook_type}}` | AI classification | Copy branching |
| `{{signal_composite_score}}` | Signal aggregation | CTA calibration |
