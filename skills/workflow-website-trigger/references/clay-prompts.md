# Clay Prompts — Website Visitor Trigger

All prompts below use `{}` for client-level variables (resolved from profile.md/icp-mapping.md) and `{{}}` for Clay row-level variables (resolved per lead from enrichment columns).

---

## Prompt 1: Page Intent Scoring

**Column name:** `trigger_page_intent_score`
**Column type:** AI prompt (Claude/GPT)

```
You are scoring website visit intent for {client_name}.

## Pages Visited
{{trigger_pages_visited}}

## Scoring Rules
- HIGH: Pricing page, demo/trial page, comparison page, integration page, contact page
- MEDIUM: Solution/product pages, use case pages, case study pages, documentation
- LOW: Blog posts, about page, careers page, homepage only
- DISQUALIFY: Careers-only visits (they're job hunting, not buying)

## Output
Return ONLY one of: "HIGH", "MEDIUM", "LOW", or "DISQUALIFY"
If pages_visited is empty, return "UNKNOWN"
```

---

## Prompt 2: Visit Context Research

**Column name:** `trigger_visit_context`
**Column type:** Claygent (web research)

```
Research {{company_name}} ({{company_website}}).

Find:
1. What does this company do? (one sentence)
2. What recent news, initiatives, or changes are happening? (check their website, LinkedIn, press)
3. Given they visited these pages on {client_website}: {{trigger_pages_visited}} — what problem are they likely trying to solve?

Keep response under 150 words. Focus on facts, not speculation.
If you find nothing relevant, say "No recent context found."
```

---

## Prompt 3: Pain Inference

**Column name:** `research_pain_inference`
**Column type:** AI prompt (Claude/GPT)

```
You are a B2B sales research analyst for {client_name}.

## About {client_name}
{client_one_liner}
We help {target_persona_titles} at {target_company_types} with:
- {pain_1} → {outcome_1}
- {pain_2} → {outcome_2}
- {pain_3} → {outcome_3}

## This Prospect
- Company: {{company_name}} ({{company_industry}}, {{company_size}} employees)
- Pages visited: {{trigger_pages_visited}}
- Visit context: {{trigger_visit_context}}
- Intent score: {{trigger_page_intent_score}}

## Task
Infer the MOST LIKELY pain point this company is experiencing that led them to explore {client_name}'s website. Connect their industry, size, and the pages they visited to a specific problem.

## Rules
- Maximum 2 sentences
- Reference their specific industry, not generic B2B
- If intent score is LOW or data is thin, use this fallback pain: "{fallback_pain_for_industry}"
- If DISQUALIFY, output "SKIP"
- Never fabricate company details you don't have
```

---

## Prompt 3b: Hook Type Classification

**Column name:** `hook_type`
**Column type:** AI prompt (Claude/GPT)
**Purpose:** Select the optimal hook pattern for this prospect. See `shared/references/hook-types-guide.md`.

```
You are selecting the optimal email hook type for {client_name}'s outreach.

## Client Proof Points
- Timeline data by vertical: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Title: {{contact_title}}
- Intent score: {{trigger_page_intent_score}}
- Pain inference: {{research_pain_inference}}

## Classification Rules
1. If we have a verified achievement timeline for {{company_industry}} → "timeline"
2. If we have a quantified metric relevant to {{contact_title}}'s KPIs → "numbers"
3. If we have a named case study in {{company_industry}} or adjacent vertical → "social_proof"
4. Otherwise → "hypothesis"

Priority: timeline > numbers > social_proof > hypothesis
Website visitors with HIGH intent score: prefer "timeline" (they're evaluating — show speed to value).

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Prompt 4: Company-Based Observational Opener

**Column name:** `copy_opener_company`
**Column type:** AI prompt (Claude/GPT)

```
You are writing outbound email openers for {client_name}. Tone: {client_tone}.

## About the Prospect
- Company: {{company_name}}
- Industry: {{company_industry}}
- Size: {{company_size}} employees
- Company snapshot: {{research_company_snapshot}}

## Rules
- Write ONE opening sentence (max 20 words)
- Must be an observation about THEIR company — something you noticed, not something about us
- Reference a specific detail from their company snapshot
- Do NOT mention website visits, browsing, or any intent signal
- Do NOT start with "I" — start with their company, their work, or their industry
- Do NOT use "I noticed" or "I came across" — use declarative observations
- No flattery ("impressive", "amazing work")

## Industry Language
{industry_language_guide_from_icp}

## Examples of good openers for this vertical:
- "{example_opener_1_from_strategy}"
- "{example_opener_2_from_strategy}"

## Fallback
If insufficient data for a real observation, use: "{fallback_observation_from_icp}"

Output ONLY the opener sentence. Nothing else.
```

---

## Prompt 5: Trigger-Based Observational Opener (Hook Type Branched)

**Column name:** `copy_opener_trigger`
**Column type:** AI prompt (Claude/GPT)

```
You are writing intent-aware outbound email openers for {client_name}. Tone: {client_tone}.

## About the Prospect
- Company: {{company_name}} ({{company_industry}}, {{company_size}} employees)
- Pages visited: {{trigger_pages_visited}}
- Intent score: {{trigger_page_intent_score}}
- Pain inference: {{research_pain_inference}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with how quickly teams in their space see results.
  Pattern: "{{company_industry}} teams evaluating [category from pages visited] typically go from [phase 1] to [result] in [timeframe]."
  Use: {client_timeline_proof_by_vertical}

IF hook_type = "numbers":
  Lead with a quantified claim relevant to the problem they're researching.
  Pattern: "[Persona] teams see [X% improvement] in [metric] when they [action related to pages visited]."
  Use: {client_number_proofs}

IF hook_type = "social_proof":
  Lead with a named peer company's result.
  Pattern: "[Named company] tackled [problem inferred from pages] and saw [result] — {{company_name}}'s {{company_industry}} setup looks like a similar fit."
  Use: {client_named_case_studies}

IF hook_type = "hypothesis":
  Connect the INFERRED PROBLEM to a trend in their industry.
  Pattern: "Teams in {{company_industry}} evaluating [category]..." or "When {{company_industry}} companies start exploring [area]..."

## Universal Rules
- Write ONE opening sentence (max 25 words)
- NEVER mention "your team visited" or "I noticed browsing" or any website signal
- Frame as industry insight, not surveillance
- Do NOT start with "I"
- Recency: if {{trigger_visit_recency_hours}} < 24: more direct | 24-72: standard | > 72: lead with trend
- If {{trigger_page_intent_score}} is "DISQUALIFY", output "SKIP"

## Output
Return ONLY the opener. Nothing else.
```

---

## Prompt 6: Email Body

**Column name:** `copy_body`
**Column type:** AI prompt (Claude/GPT)

```
You are writing the body of an outbound sales email for {client_name}. Tone: {client_tone}.

## About {client_name}
{client_one_liner}
Key differentiator: {key_differentiator}

## About the Prospect
- Company: {{company_name}} ({{company_industry}})
- Pain inference: {{research_pain_inference}}
- Persona: {{contact_title}}

## Persona Messaging Rules
{persona_messaging_rules_from_icp}

## Task
Write the email body (NOT the opener — that's already written). Connect their inferred pain to {client_name}'s specific outcome.

## Rules
- Maximum 60 words (the opener adds ~20, total email must be ≤90)
- One clear idea, not a feature dump
- Use "you/your" more than "we/our"
- End with a natural transition to the CTA (which will be appended separately)
- Reference their industry specifically — use the language guide:
  {industry_language_guide_from_icp}
- If {{research_pain_inference}} is "SKIP", output "SKIP"
```

---

## Prompt 7: LinkedIn Connection Request

**Column name:** `copy_linkedin`
**Column type:** AI prompt (Claude/GPT)

```
Write a LinkedIn connection request from {sender_name} at {client_name} to {{contact_first_name}} ({{contact_title}} at {{company_name}}).

## Context
- Company snapshot: {{research_company_snapshot}}
- Industry: {{company_industry}}

## Rules
- Maximum 280 characters (hard limit — LinkedIn cuts off)
- Lead with an observation about their company or role, NOT about {client_name}
- Do NOT pitch or sell — this is a connection request, not a sales email
- Do NOT mention website visits or intent signals
- End with a reason to connect (shared interest, industry topic, mutual value)
- Tone: {client_tone} but slightly more casual than email
- No "I'd love to connect" — that's filler

## Output
Return ONLY the message. No quotes, no labels.
```

---

## Prompt 8: CTA (Signal-Stacking Aware)

**Column name:** `copy_cta`
**Column type:** AI prompt (Claude/GPT)

```
Write a call-to-action for an outbound sales email from {client_name}.

## Prospect Context
- Title: {{contact_title}}
- ICP tier: {{score_icp_tier}}
- Company size: {{company_size}}
- Intent score: {{trigger_page_intent_score}}
- Signal composite score: {{signal_composite_score}} (if available, otherwise ignore)

## CTA Rules from {client_name}'s sales motion:
{cta_rules_from_strategy}

## {client_name}'s Sales Motion: {sales_motion}

## Stage-Appropriate CTA Guide
- Signal composite ≥100 (Hot) OR A+ tier decision maker: Direct meeting ask ("worth a 15-minute conversation?")
- Signal composite 70-99 (Warm) OR A+ tier influencer: Resource share + soft ask ("happy to share how [peer company] approached this")
- Signal composite 40-69 (Active) OR A tier: Insight-led ("I put together a quick breakdown of how [industry] teams are handling [pain] — want me to send it over?")
- Signal composite <40 OR B tier: Pure value, no ask ("Thought this might be relevant — [resource]")
- HIGH intent score: Escalate one tier (they're actively evaluating)

## Rules
- Maximum 20 words
- Must be a question or soft offer — never a demand
- Never use: "schedule a demo", "book a call", "let's chat"
- Match the client's preferred CTA language: {preferred_cta_style}

## Output
Return ONLY the CTA sentence.
```

---

## Prompt 9: Meta Tagging

**Column name:** `meta_hook_type_used` + `meta_prompt_version`
**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "website-trigger"
```

These columns are NOT AI-generated — they're formula copies or static values for A/B tracking and prompt iteration attribution. See `shared/references/prompt-iteration-pipeline.md`.
