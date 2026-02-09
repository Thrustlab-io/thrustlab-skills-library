# Clay Prompts & Copy Templates — Growth Trigger

## Clay Prompts

### Prompt 1: Growth Type Classification

**Column:** `trigger_growth_type`

```
Classify the growth signal for {{company_name}}.

Data:
- Current headcount: {{trigger_headcount_current}}
- 6 months ago: {{trigger_headcount_6mo_ago}}
- Growth rate: {{trigger_growth_rate_pct}}%
- Growing departments: {{trigger_growth_departments}}
- Recent news: {{research_company_snapshot}}

Categories:
- HEADCOUNT_SURGE: >20% overall growth in 6 months
- DEPARTMENT_BUILD: One specific department growing disproportionately
- NEW_OFFICE: Opened new location(s)
- NEW_MARKET: Expanding to new geography or customer segment
- SLOW_GROWTH: <20% growth — not a strong enough signal

Return: CATEGORY — one sentence with the key detail.
```

### Prompt 2: Scaling Challenges Inference

**Column:** `research_scaling_challenges`

```
You are a B2B research analyst for {client_name}.

## {client_name}: {client_one_liner}
## Problems we solve: {pain_1}, {pain_2}, {pain_3}

## Company
{{company_name}} ({{company_industry}}, grew from {{trigger_headcount_6mo_ago}} to {{trigger_headcount_current}} employees).
Growing departments: {{trigger_growth_departments}}

## Task
Infer 2-3 scaling challenges this company is likely facing that {client_name} could address.

Consider:
- What breaks at the {{trigger_headcount_6mo_ago}} → {{trigger_headcount_current}} transition?
- Which processes can't be done manually anymore at this size?
- What does {{trigger_growth_departments}} growth specifically strain?

Maximum 3 bullet points, one sentence each.
If growth rate < 20%, output "WEAK_SIGNAL".
```

### Prompt 2b: Hook Type Classification

**Column:** `hook_type`

```
Select optimal hook type for growth-triggered outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Growth type: {{trigger_growth_type}}
- Growth rate: {{trigger_growth_rate_pct}}%

## Classification
1. If we have a timeline for teams scaling through this growth stage → "timeline"
2. If we have metrics on efficiency gains at this company size → "numbers"
3. If we have a named company that scaled through a similar transition → "social_proof"
4. Otherwise → "hypothesis"

Growth context: "timeline" is strong — fast-growing teams need fast results. "numbers" resonates when showing what breaks at scale.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

### Prompt 3: Trigger-Based Opener (Hook Type Branched)

**Column:** `copy_opener_trigger`

```
Write a growth-aware opener for {{company_name}}. Writing for {client_name}. Tone: {client_tone}.

## Growth Data
- Type: {{trigger_growth_type}}
- Rate: {{trigger_growth_rate_pct}}%
- Departments growing: {{trigger_growth_departments}}
- Scaling challenges: {{research_scaling_challenges}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed to solve the scaling pain.
  Pattern: "Teams going from {{trigger_headcount_6mo_ago}} to {{trigger_headcount_current}} typically fix [scaling challenge] within [timeframe]."

IF hook_type = "numbers":
  Lead with a metric that resonates at their scale.
  Pattern: "At {{trigger_headcount_current}} employees, {{company_industry}} teams spend [X hours/$ amount] on [manual process] — that number compounds fast."

IF hook_type = "social_proof":
  Lead with a company that scaled through the same transition.
  Pattern: "[Named company] was at the same inflection point — grew from [X] to [Y] and [result]."

IF hook_type = "hypothesis":
  Lead with the scaling challenge, not the growth.
  Pattern: "Going from [X] to [Y] in {{company_industry}} usually means [specific challenge] — that's where most teams feel it first."

## Rules
- ONE sentence, max 25 words
- Lead with SCALING CHALLENGE, not growth itself
- Do NOT say "Congrats on the growth"
- Do NOT start with "I"

Output ONLY the opener.
```

---

### Prompt 4: Email Body + CTA (Signal-Stacking Aware)

**Column:** `copy_body` + `copy_cta`

```
Write email body for growth-triggered outbound from {client_name}. Tone: {client_tone}.

## {client_name}: {client_one_liner}
## Prospect: {{contact_first_name}} ({{contact_title}}) at {{company_name}} ({{company_industry}})
## Scaling challenges: {{research_scaling_challenges}}
## Persona rules: {persona_messaging_rules_from_icp}

Connect scaling challenge → {client_name}'s outcome.
Maximum 60 words. Industry language: {industry_language_guide_from_icp}
Output body only.
```

**CTA prompt:**
```
Write CTA for growth-triggered outbound from {client_name}.

## Calibration
- ICP tier: {{score_icp_tier}}
- Growth rate: {{trigger_growth_rate_pct}}%
- Signal composite score: {{signal_composite_score}} (if available)

- Signal composite ≥100 OR A+ tier: Direct ("worth 15 min to see if [area] could scale with you?")
- Signal composite 70-99: Playbook offer ("we mapped the [X→Y employee] scaling playbook for {{company_industry}} — want it?")
- Standard: Value ("happy to share what works at your growth stage")

Maximum 20 words. Match: {preferred_cta_style}
```

---

### Prompt 5: Meta Tagging

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "growth-trigger"
```

---

## Copy Templates

### Cadence

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Scaling challenge specific to their growth rate |
| 2 | LinkedIn | Day 1 | Industry scaling interest |
| 3 | Email | Day 5 | Scaling playbook: "what works at [size range]" |
| 4 | Email | Day 10 | Peer story: company that scaled same transition |
| 5 | Email | Day 16 | Breakup: "before growing pains compound" |

### Email 2 — Scaling Playbook (Day 5)
- Subject: `{function} at {{trigger_headcount_current}}-person {{company_industry}} companies`
- Value: specific benchmark, framework, or data point for their size stage
- Frame as helpful content, not product pitch

### Key Difference from Funding Trigger
- Funding = cash in hand, investing. Copy is about growth PRIORITIES.
- Growth = operational strain happening NOW. Copy is about scaling PAIN.
- Don't conflate them — growth without funding is a very different dynamic.

### Gates
- `trigger_growth_rate_pct` < 20 → weak signal, consider deprioritizing
- `trigger_growth_type` = "SLOW_GROWTH" → stop
- Standard ICP gate
