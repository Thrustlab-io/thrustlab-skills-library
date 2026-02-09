# Clay Prompts — Competitor Customer Targeting Workflow

## Prompt Architecture

```
Layer 1: Import + competitor detection + ICP scoring + hook_type
Layer 2: Competitive intelligence research + switch angle + social proof matching
Layer 3: Copy generation (branched by competitor × hook_type)
```

---

## Layer 1 Prompts

### Competitor Detection

**Column:** `competitor_product_used`
**Type:** Claygent (web research) + BuiltWith (tech detection)
**Purpose:** Identify which competitor product (if any) this company uses.

```
Research whether {{company_name}} ({{company_website}}) uses any of these products:
{competitor_1}, {competitor_2}, {competitor_3}, {competitor_4}, {competitor_5}

Check:
1. Their website for technology signals (login pages, integrations mentioned, embedded widgets)
2. Job postings mentioning these tools
3. Case studies or testimonials they've published
4. G2/TrustRadius reviews they've written
5. LinkedIn posts from employees mentioning these tools

Output format: "{product_name}" if found, or "NOT_FOUND" if no evidence.
If multiple competitors detected, output the PRIMARY one (most evidence).
```

### Competitor Confidence Classification

**Column:** `competitor_confidence`
**Reads:** `competitor_product_used`, `competitor_detection_source`
**Type:** AI prompt

```
Based on the detection evidence, classify confidence level:

## Evidence
Competitor detected: {{competitor_product_used}}
Detection source: {{competitor_detection_source}}

## Classification
- "confirmed": Multiple sources confirm usage (BuiltWith + job posting, or public case study, or G2 review)
- "likely": Single strong signal (BuiltWith detection, or specific job posting requiring the tool)
- "possible": Indirect signal only (industry where this competitor is common, or vague mentions)
- "none": No evidence found

Output ONLY one word: confirmed, likely, possible, or none.
```

**CRITICAL GATE:** Only proceed to Layer 2 if `competitor_confidence` = "confirmed" or "likely". Route "possible" to general outbound. Drop "none" entirely.

### Hook Type Classification

**Column:** `hook_type`
**Reads:** `competitor_product_used`, `company_industry`, `contact_title`

```
Select the optimal hook type for competitive displacement outreach.

## Client Proof Points Against {{competitor_product_used}}
- Switch case studies: {switch_case_studies_for_this_competitor}
- Timeline from competitor to client: {typical_migration_timeline}
- Quantified differentiators: {metrics_where_client_beats_competitor}

## Classification for Competitor Targeting
1. If we have a named customer who switched from {{competitor_product_used}} → "social_proof"
2. If we have a verified migration timeline → "timeline"
3. If we have quantified metrics showing improvement over {{competitor_product_used}} → "numbers"
4. Otherwise → "hypothesis"

Note: For competitor targeting, social_proof is the STRONGEST hook (3x conversion with named switcher).

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Layer 2 Prompts

### Competitive Usage Research

**Column:** `research_competitor_usage`
**Type:** Claygent (web research)

```
Research how {{company_name}} uses {{competitor_product_used}}.

Find:
1. What specific workflows or processes they use it for
2. Scale of usage (team size, department, company-wide)
3. Any public complaints, limitations, or gaps they've mentioned
4. How long they've been using it (if detectable)

Context: We are {client_name}. We compete with {{competitor_product_used}} and win when {where_client_wins_vs_competitor}.

Rules:
- Only factual, verifiable information
- Maximum 80 words
- If insufficient information, output "NOT_FOUND"
- Do not fabricate
```

### Competitor Pain Analysis

**Column:** `research_competitor_pain`
**Reads:** `competitor_product_used`, `research_competitor_usage`, `company_industry`, `company_size`
**Type:** AI prompt

```
You are a competitive analyst for {client_name}.

## Known limitations of {{competitor_product_used}}:
{competitor_known_limitations}

## This company's context:
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- How they use it: {{research_competitor_usage}}

## Task
Identify the ONE most relevant limitation of {{competitor_product_used}} for THIS company's specific situation. Consider their industry, size, and usage pattern.

## Rules
- Maximum 25 words
- Must be a REAL limitation, not fabricated
- Frame as a challenge, not an attack: "scaling beyond X users" not "terrible at scale"
- If no specific limitation is relevant, output "SKIP"
```

### Switch Angle

**Column:** `research_switch_angle`
**Reads:** `research_competitor_pain`, `overlap_company`, `company_industry`
**Type:** AI prompt

```
You are crafting a competitive positioning angle for {client_name} vs {{competitor_product_used}}.

## Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Their competitor pain: {{research_competitor_pain}}
- Where {client_name} wins vs {{competitor_product_used}}: {where_client_wins}
- Switch case study (if available): {switch_case_study_summary}

## Task
Write ONE sentence that connects {{company_name}}'s likely pain with {{competitor_product_used}} to {client_name}'s specific advantage. This becomes the core messaging angle.

## Rules
- Maximum 25 words
- Frame as opportunity, not criticism: "teams who need X find that..." not "competitor can't do X"
- Must be credible — based on real differentiator
- If no compelling angle, output "SKIP"
```

---

## Layer 3 Prompts — Copy Generation

### Opener — Branched by Hook Type

**Column:** `copy_opener_trigger`
**Reads:** ALL upstream + `hook_type` + `competitor_product_used`

```
You are writing an outbound email opener for {client_name} to a {{competitor_product_used}} user.

## Prospect Context
- Name: {{contact_first_name}}, {{contact_title}} at {{company_name}}
- Uses: {{competitor_product_used}} (confidence: {{competitor_confidence}})
- Their usage: {{research_competitor_usage}}
- Their pain: {{research_competitor_pain}}
- Switch angle: {{research_switch_angle}}

## Available Social Proof
- Switch customer: {{competitor_customer_name}} (switched from {{competitor_product_used}})
- Their result: {{competitor_customer_result}}

## Hook Type: {{hook_type}}

IF hook_type = "social_proof":
  Lead with the named switcher and their result.
  Pattern: "When {{competitor_customer_name}} moved from {{competitor_product_used}} to {client_product}, they saw {{competitor_customer_result}}. Given {{company_name}}'s [context], that gap might be worth exploring."

IF hook_type = "timeline":
  Lead with migration speed.
  Pattern: "Teams switching from {{competitor_product_used}} typically are fully migrated in [X weeks] — most see [early win] before the first invoice."

IF hook_type = "numbers":
  Lead with quantified improvement over competitor.
  Pattern: "{{company_industry}} teams using {client_product} over {{competitor_product_used}} see [X% improvement] in [metric]."

IF hook_type = "hypothesis":
  Lead with a question about a known competitor limitation.
  Pattern: "Curious — as {{company_name}} scales, has {{competitor_product_used}}'s [known limitation] started showing up for your team?"

## Rules
- Maximum 2 sentences (35 words)
- NEVER trash {{competitor_product_used}} — they chose it for reasons
- Frame as "what if better were possible?" not "your choice is wrong"
- {client_tone} tone
- If insufficient data, output "SKIP"
```

### Email Body

**Column:** `copy_body`
**Reads:** opener + research + hook_type

```
You are writing the body of a competitive displacement email for {client_name}.

## Context (already written)
Opener: {{copy_opener_trigger}}

## Competitive Position
- Their competitor pain: {{research_competitor_pain}}
- Switch angle: {{research_switch_angle}}
- Client differentiator: {where_client_wins}

## Task
Write 2-3 sentences that:
1. Acknowledge they have a solution (never assume they're unhappy)
2. Introduce ONE specific capability where {client_name} differs from {{competitor_product_used}}
3. Connect that capability to {{company_name}}'s specific situation

## Rules
- Maximum 45 words
- Respectful of their current choice
- Specific differentiator, not "we're better overall"
- {client_tone} tone
```

### CTA

**Column:** `copy_cta`

```
Write a CTA for competitive displacement email from {client_name}.

## Calibration
- Competitor confidence: {{competitor_confidence}}
- Signal composite score: {{signal_composite_score}}

## Rules
- Non-threatening: they already have a solution — this is about exploring, not switching
- Pattern: "Worth a 15-min comparison to see if [specific differentiator] moves the needle for {{company_name}}?"
- If signal_composite_score >= 70: slightly more direct
- Maximum 20 words
- {client_tone} tone
```

### LinkedIn Message

**Column:** `copy_linkedin`

```
Write a LinkedIn message from {client_name} to a {{competitor_product_used}} user at {{company_name}}.

## Rules
- Maximum 280 characters
- Do NOT mention the competitor by name on LinkedIn (too aggressive for public-ish channel)
- Instead: focus on the category/capability where client differentiates
- Pattern: "Hi {{contact_first_name}}, we've been helping {{company_industry}} teams with [differentiated capability]. Given {{company_name}}'s approach to [area], thought there might be interesting overlap."
- {client_tone} tone
```

---

## Signal Stacking Integration

Competitor customer status is a persistent Layer 3 (relationship) signal worth +20 points (no decay — they're always using the competitor until proven otherwise).

Highest-converting stacks:
- Competitor customer + hiring for relevant role + website visit = 🔴 Hot
- Competitor customer + funding round = 🟠 Warm (minimum)
- Competitor customer + champion job change = 🔴 Hot (someone who knows your product now works at a competitor customer)

---

## Meta Tagging

**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "competitor-customer"
```

Tag every row for A/B testing attribution. See `shared/references/prompt-iteration-pipeline.md`.
