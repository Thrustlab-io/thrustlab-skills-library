# Clay Prompts & Copy Templates — Tech Stack Change Trigger

## Clay Prompts

### Prompt 1: Tech Change Classification

**Column:** `trigger_tech_change_type`

```
Classify this tech stack change for {client_name}'s outbound.

## {client_name}'s Ecosystem
- Our product: {client_product_category}
- Complementary tools: {complementary_tools}
- Competitors: {competitor_1}, {competitor_2}, {competitor_3}

## The Change
- Company: {{company_name}}
- Tech added: {{trigger_tech_added}}
- Tech removed: {{trigger_tech_removed}}

Categories:
- COMPLEMENTARY_ADD: They added a tool that works well with {client_product}. Integration opportunity.
- COMPETITOR_REMOVED: They dropped a competitor. Active replacement/evaluation window.
- COMPETITOR_ADDED: They just adopted a competitor. May be disqualified.
- STACK_SHIFT: They're restructuring their tech stack. Window to be part of the new architecture.
- IRRELEVANT: Change has no meaningful connection.

Return: CATEGORY — one sentence explaining why.
```

### Prompt 2: Stack Context

**Column:** `trigger_stack_context` (Claygent)

```
Research {{company_name}}'s technology stack ({{company_website}}).

Check BuiltWith, Wappalyzer, their integrations page, and job postings.

1. What major tools/platforms do they use?
2. How does the recent change (added: {{trigger_tech_added}}, removed: {{trigger_tech_removed}}) fit their broader stack?
3. Any signs of tech consolidation or migration?

Maximum 100 words.
```

### Prompt 2b: Hook Type Classification

**Column:** `hook_type`

```
Select optimal hook type for tech-stack-change outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Change type: {{trigger_tech_change_type}}
- Added: {{trigger_tech_added}}, Removed: {{trigger_tech_removed}}

## Classification
1. If we have a verified integration/migration timeline → "timeline"
2. If we have quantified improvements from stack changes → "numbers"
3. If we have a named company that made a similar stack transition → "social_proof"
4. Otherwise → "hypothesis"

Tech change context: "numbers" strong for COMPLEMENTARY_ADD (show combined value); "social_proof" strong for COMPETITOR_REMOVED (show switcher results).

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

### Prompt 3: Trigger-Based Opener (Hook Type Branched)

**Column:** `copy_opener_trigger`

```
Write a tech-stack-aware opener for {{company_name}}. Writing for {client_name}. Tone: {client_tone}.

## Change Type: {{trigger_tech_change_type}}
- Added: {{trigger_tech_added}}
- Removed: {{trigger_tech_removed}}
- Stack context: {{trigger_stack_context}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed of integration/migration.
  Pattern: "Teams adding [tool] to their stack typically see [result] within [timeframe] when they [complement with client area]."

IF hook_type = "numbers":
  Lead with quantified improvement from the stack combination.
  Pattern: "{{company_industry}} teams running [tool] + [client product] see [X% improvement] in [metric]."

IF hook_type = "social_proof":
  Lead with a named company that made a similar stack change.
  Pattern: "[Named company] made the same move with [tool] — paired it with [client approach] and saw [result]."

IF hook_type = "hypothesis":
  Use the original change-type branching:
  IF COMPLEMENTARY_ADD: "Teams that invest in [added tool] usually find [area client covers] is the next bottleneck."
  IF COMPETITOR_REMOVED: "When [industry] teams move on from [category], the priorities usually shift to [client's strength]."
  IF STACK_SHIFT: "Modernizing the stack usually means [specific challenge for client's domain]."

## Rules
- ONE sentence, max 25 words
- Do NOT start with "I"
- COMPETITOR_ADDED or IRRELEVANT → output "SKIP"

Output ONLY the opener.
```

### Prompt 4: Email Body

**Column:** `copy_body`

```
Write email body for tech-stack-triggered outbound from {client_name}. Tone: {client_tone}.

## {client_name}: {client_one_liner}
## Prospect: {{contact_first_name}} ({{contact_title}}) at {{company_name}} ({{company_industry}})
## Change: {{trigger_tech_change_type}} — added: {{trigger_tech_added}}, removed: {{trigger_tech_removed}}
## Persona rules: {persona_messaging_rules_from_icp}

## Task
Connect the stack change to {client_name}'s value. Maximum 60 words.

IF COMPLEMENTARY_ADD: position as "completes the stack" or "unlocks more value from [tool they added]"
IF COMPETITOR_REMOVED: position on merits, not vs competitor. Focus on what matters when evaluating alternatives.
IF STACK_SHIFT: position as part of the modernization journey.

Industry language: {industry_language_guide_from_icp}
Output body only.
```

---

### Prompt 5: CTA (Signal-Stacking Aware)

**Column:** `copy_cta`

```
Write a CTA for tech-stack triggered outbound from {client_name}.

## Prospect Context
- Title: {{contact_title}}, ICP tier: {{score_icp_tier}}
- Change type: {{trigger_tech_change_type}}
- Signal composite score: {{signal_composite_score}} (if available)

## CTA Calibration
- Signal composite ≥100 OR COMPETITOR_REMOVED + A+ tier: Direct ("worth a quick comparison of how [client product] fits your new stack?")
- Signal composite 70-99 OR COMPLEMENTARY_ADD: Integration-forward ("we put together an integration guide for [added tool] + [client area] — want it?")
- Standard: Value offer ("happy to share how {{company_industry}} teams are building their [function] stack")

## Rules
- Maximum 20 words, question or soft offer
- Match: {preferred_cta_style}

Output ONLY the CTA.
```

---

### Prompt 6: Meta Tagging

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "tech-change-trigger"
```

---

## Copy Templates

### Cadence

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Stack observation + ecosystem positioning |
| 2 | LinkedIn | Day 1 | Tech/industry interest connection |
| 3 | Email | Day 5 | Integration guide / migration framework (type-dependent) |
| 4 | Email | Day 10 | Social proof: company with same stack that added client |
| 5 | Email | Day 15 | Breakup: technical angle with different value prop |

### Type-Specific Email 2 (Day 5)

**IF COMPLEMENTARY_ADD:**
- Subject: `Getting more from {{trigger_tech_added}} + {client_product_category}`
- Value: integration benefits, workflow that connects both tools, efficiency gain

**IF COMPETITOR_REMOVED:**
- Subject: `What to look for in a {product_category} in {year}`
- Value: evaluation framework (positions client's strengths without naming competitors)

**IF STACK_SHIFT:**
- Subject: `Rebuilding {function} infrastructure in {{company_industry}}`
- Value: best practices for tech consolidation in their vertical

### Gates
- `trigger_tech_change_type` = "COMPETITOR_ADDED" → auto-DQ or route to separate nurture
- `trigger_tech_change_type` = "IRRELEVANT" → stop
- Standard ICP + freshness gates
