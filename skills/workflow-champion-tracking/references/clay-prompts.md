# Clay Prompts — Champion Tracking Workflow

## Prompt Architecture

```
Layer 1: Champion import + new company enrichment + ICP scoring + hook_type classification
Layer 2: Research on new company context + transition insight + overlap analysis
Layer 3: Copy generation (branched by relationship_type × hook_type)
```

---

## Layer 1 Prompts

### ICP Scoring — New Company Fit

**Column:** `score_icp_fit`
**Reads:** `company_name`, `company_industry`, `company_size`, enrichment data
**Purpose:** Does the CHAMPION'S NEW COMPANY fit client's ICP? A great relationship at a bad-fit company is a waste.

```
You are scoring whether {{company_name}} fits {client_name}'s ICP.

## ICP Criteria
{client_icp_criteria_from_strategy}

## Company Data
- Company: {{company_name}}
- Industry: {{company_industry}}
- Size: {{company_size}} employees
- HQ: {{company_hq}}
- Tech stack: {{company_tech_stack}}

## Scoring
Rate 1-10 where:
- 9-10: Perfect ICP fit (exact industry, size, tech stack match)
- 7-8: Strong fit (most criteria match)
- 5-6: Moderate fit (some criteria match)
- 3-4: Weak fit (few criteria match)
- 1-2: Not a fit

Output ONLY the number. If insufficient data to score, output "5" (neutral).
```

### Hook Type Classification

**Column:** `hook_type`
**Reads:** `company_industry`, `contact_title`, `champion_relationship_type`, `champion_result_achieved`

```
You are selecting the optimal email hook type for outreach to a former {{champion_relationship_type}}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Champion Context
- Relationship type: {{champion_relationship_type}}
- Result achieved (if customer): {{champion_result_achieved}}
- New company industry: {{company_industry}}
- New role: {{contact_title}}

## Classification Rules for Champions
1. If champion_relationship_type = "customer" AND champion_result_achieved is known → "social_proof" (they ARE the proof)
2. If we have a verified timeline for their new company's vertical → "timeline"
3. If we have a quantified metric relevant to their new role → "numbers"
4. Otherwise → "hypothesis"

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Layer 2 Prompts

### Research — New Company Context

**Column:** `research_new_company`
**Type:** Claygent (web research)

```
Research {{company_name}} to understand:
1. What the company does (core product/service, in 1-2 sentences)
2. What stage they're at (startup/growth/enterprise, any recent milestones)
3. Any public information about their {client_solution_area} approach

Context: We sell {client_product_description} to {target_persona_titles}. We want to understand if {{company_name}} has a relevant use case.

Rules:
- Only factual, verifiable information
- Maximum 80 words total
- If you can't find relevant info, output "NOT_FOUND"
- Do not fabricate
```

### Transition Insight

**Column:** `research_transition_insight`
**Reads:** `champion_previous_company`, `champion_previous_title`, `company_name`, `contact_title`
**Type:** AI prompt

```
You are a B2B sales researcher for {client_name}.

## Context
{{champion_first_name}} moved from {{champion_previous_title}} at {{champion_previous_company}} to {{contact_title}} at {{company_name}} ({{company_industry}}, {{company_size}} employees).

## Task
Write ONE insightful sentence about what this transition likely means for them professionally — what challenges or opportunities does the move create that relate to {client_problem_domain}?

## Rules
- Maximum 25 words
- Must be specific to THIS transition, not generic "new role challenges"
- No clichés ("hit the ground running", "make their mark")
- If insufficient data, output "SKIP"
```

### New Company Overlap

**Column:** `overlap_new_company`
**Reads:** `research_new_company`, `company_industry`, `company_size`, `company_tech_stack`
**Type:** AI prompt

```
You are analyzing how {client_name}'s product fits {{company_name}}'s context.

## {client_name} solves:
{client_value_prop}
Key pains: {pain_1}, {pain_2}, {pain_3}

## {{company_name}} context:
{{research_new_company}}
Industry: {{company_industry}}, Size: {{company_size}}, Tech: {{company_tech_stack}}

## Task
In ONE sentence (max 20 words), describe the most likely way {{company_name}} would use {client_product}. Be specific to their industry and size.

If no clear overlap exists, output "SKIP".
```

---

## Layer 3 Prompts — Copy Generation

### Opener — Branched by Relationship Type × Hook Type

**Column:** `copy_opener_trigger`
**Reads:** ALL upstream columns + `hook_type` + `champion_relationship_type`

```
You are writing an outbound email opener for {client_name}.

## Champion Context
- Name: {{champion_first_name}}
- Relationship: {{champion_relationship_type}}
- Previous company: {{champion_previous_company}}
- Previous title: {{champion_previous_title}}
- Result achieved (if known): {{champion_result_achieved}}
- Days in new role: {{champion_days_in_role}}

## New Context
- New company: {{company_name}} ({{company_industry}}, {{company_size}})
- New title: {{contact_title}}
- Transition insight: {{research_transition_insight}}
- Company overlap: {{overlap_new_company}}

## Hook Type: {{hook_type}}

IF hook_type = "social_proof" AND relationship = "customer":
  Reference the specific result they achieved, then bridge to new company.
  Pattern: "At {previous_company}, your team [achieved result]. {New company}'s [context] seems like [bridge]."

IF hook_type = "timeline":
  Lead with how quickly teams like their new company see results.
  Pattern: "Teams making a move like {company_name}'s typically go from [phase] to [result] in [timeframe]."

IF hook_type = "numbers":
  Lead with a metric relevant to their new role's KPIs.
  Pattern: "{Industry} teams see [X% improvement] in [metric] — given your experience at {previous_company}, you'd know the difference that makes."

IF hook_type = "hypothesis":
  Frame a question about their new role's challenges.
  Pattern: "Moving from {previous_company} to {company_name} — would I be right that [specific challenge] is already on your radar?"

## Rules
- Maximum 2 sentences (35 words)
- ONE subtle reference to the past relationship — not heavy-handed
- Focus on THEIR FUTURE at the new company, not YOUR PAST together
- Congrats on the move is fine as a half-sentence, never the whole opener
- If champion_days_in_role > 100, skip the "congrats" — they're not new anymore
- {client_tone} tone
- If insufficient data, output "SKIP"
```

### Email Body

**Column:** `copy_body`
**Reads:** opener + research columns + hook_type

```
You are writing the body of a reconnection email for {client_name}.

## Context (already written)
Opener: {{copy_opener_trigger}}

## Champion Info
- Relationship: {{champion_relationship_type}}
- New company context: {{research_new_company}}
- Overlap: {{overlap_new_company}}

## Client Value
{client_value_prop}
Best proof point for {{company_industry}}: {relevant_case_study_or_metric}

## Task
Write 2-3 sentences that:
1. Bridge from the opener to why {client_product} is relevant at {{company_name}} specifically
2. Include ONE proof point (result, metric, or timeline) appropriate to the hook_type
3. Keep it forward-looking — about what's possible at the new company

## Rules
- Maximum 45 words
- Do NOT repeat the opener
- Do NOT over-pitch — this person already knows your product
- {client_tone} tone
- No banned phrases (see copy-rules.md)
```

### CTA

**Column:** `copy_cta`
**Reads:** `champion_relationship_type`, `champion_days_in_role`, `score_icp_fit`, `signal_composite_score`

```
Write a CTA for a reconnection email from {client_name} to a former {{champion_relationship_type}}.

## Calibration
- Days in new role: {{champion_days_in_role}}
- ICP fit of new company: {{score_icp_fit}}/10
- Signal composite score: {{signal_composite_score}}

## CTA Rules
IF former customer AND days_in_role < 60:
  Warm, personal: "Would love to catch up — grab 15 min?"

IF former customer AND days_in_role 60-100:
  Value-forward: "If [pain] comes up at {company_name}, happy to share what's changed since {previous}."

IF former prospect/engaged:
  Lighter touch: "If [topic] is on your radar at {company_name}, happy to pick the conversation back up."

IF signal_composite_score >= 70:
  More direct: "Worth a quick chat about how this could fit {company_name}'s setup?"

## Rules
- Maximum 20 words
- Always a question, never a command
- {client_tone} tone
```

### LinkedIn Message

**Column:** `copy_linkedin`
**Reads:** champion context + new company

```
Write a LinkedIn message for {client_name} reconnecting with {{champion_first_name}}.

## Context
- Moved from {{champion_previous_title}} at {{champion_previous_company}} to {{contact_title}} at {{company_name}}
- Relationship type: {{champion_relationship_type}}
- Days in role: {{champion_days_in_role}}

## Rules
- Maximum 280 characters
- Genuine professional reconnection, no pitch
- Pattern: "Congrats on {company_name}! [one observation]. Would love to stay connected."
- {client_tone} tone
```

---

## Signal Stacking Integration

Champion signals are the highest-weighted relationship signals. When detected:
1. Push to Signal Aggregation Table: `signal_type = "champion_job_change"`, points = 40 (customer) or 25 (prospect)
2. Check for compound signals on the new company (website visits, funding, hiring)
3. If composite score ≥ 100 or matches known high-converting stack → Slack alert to AE

---

## Meta Tagging

**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "champion-tracking"
```

Tag every row for A/B testing attribution. See `shared/references/prompt-iteration-pipeline.md`.
