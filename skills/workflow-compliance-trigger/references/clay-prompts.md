# Clay Prompts & Copy Templates — Compliance Trigger

## Clay Prompts

### Prompt 1: Regulation Impact Assessment

**Column:** `trigger_regulation_impact`

```
You are a compliance-aware B2B analyst for {client_name}.

## Regulation
- Name: {{trigger_regulation_name}}
- Deadline: {{trigger_regulation_deadline}}
- Days until deadline: {{trigger_days_until_deadline}}

## Company
- {{company_name}} ({{company_industry}}, {{company_size}} employees, {{company_location}})

## Task
Assess how {{trigger_regulation_name}} specifically affects {{company_name}}.
- What must they comply with?
- What changes are required?
- What's their likely exposure/risk level given their industry and size?

## Rules
- Maximum 3 sentences
- Be factually precise — never misstate regulatory requirements
- If the regulation doesn't clearly affect this company's industry, output "LOW_RELEVANCE"
- If you don't have enough information about the regulation, output "VERIFY_MANUALLY"
```

### Prompt 2: Compliance Readiness Research

**Column:** `research_compliance_readiness` (Claygent)

```
Research whether {{company_name}} ({{company_website}}) appears to be prepared for {{trigger_regulation_name}} (deadline: {{trigger_regulation_deadline}}).

Look for:
1. Any mention of the regulation on their website, blog, or press
2. Job postings related to compliance, legal, or the regulation's domain
3. Recent hires in compliance or related functions
4. Public statements about their compliance posture

Maximum 80 words. If nothing found, say "No public compliance signals found."
```

### Prompt 2b: Hook Type Classification

**Column:** `hook_type`

```
Select optimal hook type for compliance-triggered outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Regulation: {{trigger_regulation_name}}
- Days until deadline: {{trigger_days_until_deadline}}

## Classification
1. If we have a verified compliance-to-production timeline → "timeline" (show they can comply in time)
2. If we have quantified compliance cost/effort reduction → "numbers"
3. If we have a named company that achieved compliance using client → "social_proof"
4. Otherwise → "hypothesis"

Compliance context: "timeline" is especially powerful when deadline is <6 months. "social_proof" builds credibility.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

### Prompt 3: Trigger-Based Opener (Hook Type Branched)

**Column:** `copy_opener_trigger`

```
Write a compliance-aware opener for {{contact_first_name}} at {{company_name}}. For {client_name}. Tone: {client_tone}.

## Context
- Regulation: {{trigger_regulation_name}}
- Deadline: {{trigger_regulation_deadline}} ({{trigger_days_until_deadline}} days away)
- Impact: {{trigger_regulation_impact}}
- Readiness: {{research_compliance_readiness}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed to compliance.
  Pattern: "{{company_industry}} teams typically achieve {{trigger_regulation_name}} compliance in [X weeks] — even starting from scratch."

IF hook_type = "numbers":
  Lead with compliance cost/effort reduction.
  Pattern: "Teams using [approach] reduce {{trigger_regulation_name}} compliance effort by [X%] — that's [Y hours/dollars] back."

IF hook_type = "social_proof":
  Lead with a peer company's compliance success.
  Pattern: "[Named company] achieved full {{trigger_regulation_name}} compliance [X weeks] before deadline using [approach]."

IF hook_type = "hypothesis":
  Use deadline-proximity framing:
  - >6 months: "{{trigger_regulation_name}} is reshaping how {{company_industry}} companies handle [area]."
  - 3-6 months: "With {{trigger_regulation_name}} taking effect in [month], {{company_industry}} teams are rethinking [area]."
  - <3 months: "{{trigger_regulation_name}}'s [deadline] is creating real pressure on [function]."
  - Past: "Most {{company_industry}} companies are still catching up with {{trigger_regulation_name}}."

## Rules
- ONE sentence, max 25 words
- Reference regulation by name — show expertise
- Use correct regulatory terminology
- Do NOT fearmongering — factual urgency only
- Do NOT start with "I"

Output ONLY the opener.
```

### Prompt 4: Email Body

```
Write email body for compliance-triggered outbound from {client_name}. Tone: {client_tone}.

## {client_name}: {client_one_liner}
## How we help with {{trigger_regulation_name}}: {specific_compliance_capability}
## Prospect: {{contact_first_name}} ({{contact_title}}) at {{company_name}} ({{company_industry}})
## Regulation impact: {{trigger_regulation_impact}}

## Rules
- Maximum 60 words
- Connect the specific regulatory requirement → specific {client_name} capability
- Don't explain the regulation (they know it) — explain how to comply efficiently
- Industry language: {industry_language_guide_from_icp}
- If {{trigger_days_until_deadline}} < 90: include timeline awareness
- Natural transition to CTA

Output body only.
```

---

## Copy Templates

### Cadence

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Regulation-specific observation + compliance path |
| 2 | LinkedIn | Day 2 | Regulatory topic connection |
| 3 | Email | Day 5 | Compliance checklist / readiness framework (pure value) |
| 4 | Email | Day 10 | Peer compliance success story |
| 5 | Email | Day 15 | Deadline proximity reminder + offer |

### Email 2 — Compliance Checklist (Day 5)
- Subject: `{{trigger_regulation_name}} readiness checklist for {{company_industry}}`
- Pure value: actionable checklist or framework
- Position {client_name} as ONE item on the checklist, not the whole thing
- Shows authority — you know the regulation well enough to create a checklist

### Key Uniqueness
- Highest external urgency of any trigger — real deadlines, real penalties
- Authority and accuracy matter more than personalization
- Compliance officers respond to specifics, not generalities
- This trigger has natural follow-up cadence tied to deadline proximity

### Gates
- `trigger_regulation_impact` = "LOW_RELEVANCE" or "VERIFY_MANUALLY" → stop or manual review
- Regulation factual accuracy must be verified before sending
- Standard ICP gate

---

### Prompt 5: CTA (Signal-Stacking Aware)

**Column:** `copy_cta`

```
Write CTA for compliance-triggered outbound from {client_name}.

## Calibration
- Title: {{contact_title}}, ICP tier: {{score_icp_tier}}
- Days until deadline: {{trigger_days_until_deadline}}
- Signal composite score: {{signal_composite_score}} (if available)

- Signal composite ≥100 OR deadline <90 days + A+ tier: Direct ("worth a 30-min compliance walkthrough specific to {{company_name}}'s setup?")
- Signal composite 70-99 OR deadline 90-180 days: Framework offer ("we built a {{trigger_regulation_name}} readiness checklist for {{company_industry}} — want me to send it?")
- Standard OR deadline >180 days: Educational value ("happy to share how {{company_industry}} teams are approaching {{trigger_regulation_name}}")

Maximum 20 words. Match: {preferred_cta_style}
```

---

### Prompt 6: Meta Tagging

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "compliance-trigger"
```
