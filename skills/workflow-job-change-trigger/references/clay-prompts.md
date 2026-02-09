# Clay Prompts — Job Role Change Trigger

---

## Prompt 1: Role Transition Type

**Column name:** `trigger_role_transition_type`
**Column type:** AI prompt

```
Classify this job change:

- Previous: {{trigger_previous_title}} at {{trigger_previous_company}}
- Current: {{contact_title}} at {{company_name}}

Categories:
- PROMOTION: Same company or moved to a more senior title
- LATERAL: Similar seniority, new company, same function
- CAREER_PIVOT: Different function or significant role change
- INDUSTRY_CHANGE: Moved to a different industry
- STEP_UP_COMPANY: Moved to a significantly larger or more prestigious company

Return ONLY the category.
If previous data is missing, return "UNKNOWN".
```

---

## Prompt 2: New Role Context Research

**Column name:** `trigger_new_role_context`
**Column type:** Claygent

```
Research the transition of {{contact_first_name}} {{contact_last_name}} who recently became {{contact_title}} at {{company_name}}.

Find:
1. What is {{company_name}} doing right now? (recent news, growth, challenges)
2. What would a new {{contact_title}} at this type of company typically prioritize in their first 90 days?
3. Any public posts or content from {{contact_first_name}} about their new role or priorities?

Check: {{company_website}}, LinkedIn, recent news.
Maximum 150 words. Facts only.
```

---

## Prompt 3: Previous Company Relevance

**Column name:** `research_previous_company_relevance`
**Column type:** AI prompt

```
You are researching a job change for {client_name}'s outbound.

## Context
{client_name} ({client_website}): {client_one_liner}
Our competitors: {competitor_1}, {competitor_2}, {competitor_3}

## The Transition
- Person: {{contact_first_name}} {{contact_last_name}}
- Previous: {{trigger_previous_title}} at {{trigger_previous_company}}
- Current: {{contact_title}} at {{company_name}}

## Task
Assess the previous company's relevance:
1. Did {{trigger_previous_company}} likely use {client_name}, a competitor, or a similar solution?
2. What experience might {{contact_first_name}} bring from {{trigger_previous_company}} that's relevant?
3. Is there a "I've seen this work before" angle?

## Output
Maximum 2 sentences. If no relevant connection, say "No direct relevance found — use general approach."
```

---

## Prompt 4: Company-Based Observational Opener

**Column name:** `copy_opener_company`
**Column type:** AI prompt

```
Write an observational opener about {{company_name}} for {client_name}'s outbound. Tone: {client_tone}.

## About {{company_name}}
- Industry: {{company_industry}}
- Size: {{company_size}} employees
- Research: {{trigger_new_role_context}}

## Rules
- ONE sentence, max 20 words
- Observation about the COMPANY (not the person's job change)
- No "I noticed" — declarative observations
- Use industry language: {industry_language_guide_from_icp}

Output ONLY the opener.
```

---

## Prompt 4b: Hook Type Classification

**Column name:** `hook_type`
**Column type:** AI prompt
**Purpose:** Select optimal hook pattern. See `shared/references/hook-types-guide.md`.

```
Select the optimal hook type for job-change triggered outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- New role: {{contact_title}} at {{company_name}} ({{company_industry}}, {{company_size}})
- Transition type: {{trigger_role_transition_type}}
- Days in role: {{trigger_days_in_role}}

## Classification Rules for Job Changes
1. If we have a verified ramp timeline for new-in-role personas → "timeline" (show speed to first win)
2. If we have quantified metrics for {{contact_title}}'s KPIs → "numbers"
3. If we have a named case study in {{company_industry}} → "social_proof"
4. Otherwise → "hypothesis"

Job change context: "timeline" is especially strong — new leaders want fast results in their first 90 days.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Prompt 5: Trigger-Based Opener (Transition-Aware, Hook Type Branched)

**Column name:** `copy_opener_trigger`
**Column type:** AI prompt

```
Write a transition-aware opener for {{contact_first_name}} who just became {{contact_title}} at {{company_name}}. Writing for {client_name}. Tone: {client_tone}.

## Transition Details
- Previous: {{trigger_previous_title}} at {{trigger_previous_company}}
- Days in new role: {{trigger_days_in_role}}
- Transition type: {{trigger_role_transition_type}}
- Previous company relevance: {{research_previous_company_relevance}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed to first win in new role.
  Pattern: "[Days] into a new [role] — teams at this stage usually see [result] within [timeframe]."
  Use: {client_timeline_proof_by_vertical}

IF hook_type = "numbers":
  Lead with a metric relevant to new-role KPIs.
  Pattern: "[Persona] leaders who [action] in their first quarter see [X% improvement] in [metric]."
  Use: {client_number_proofs}

IF hook_type = "social_proof":
  Lead with a peer who made a similar transition.
  Pattern: "When [named person/company] made a similar move to [type of company], they [result]."
  Use: {client_named_case_studies}

IF hook_type = "hypothesis":
  Frame a question about transition-specific challenges.
  Pattern: "The move from [previous context] to [new context] usually means [specific challenge] is top of mind."

## Universal Rules
- ONE sentence, max 25 words
- If {{trigger_days_in_role}} < 30: lighter touch, acknowledge the move
- If {{trigger_days_in_role}} 30-60: more substantive
- If {{trigger_days_in_role}} > 60: skip congrats — they're established
- Do NOT use generic "Congrats on the new role!" without specific detail
- Do NOT start with "I"

Output ONLY the opener.
```

---

## Prompt 6: Email Body

**Column name:** `copy_body`
**Column type:** AI prompt

```
Write the email body for a job-change triggered outbound from {client_name}. Tone: {client_tone}.

## About {client_name}
{client_one_liner}
Key differentiator: {key_differentiator}

## The Prospect
- {{contact_first_name}} {{contact_last_name}}, {{contact_title}} at {{company_name}} ({{company_industry}})
- Transition context: {{trigger_new_role_context}}
- Relevant experience: {{research_previous_company_relevance}}

## Persona Rules
{persona_messaging_rules_from_icp}

## Task
Write the body connecting their new-role challenges to {client_name}'s outcome. NOT the opener (already written).

## Rules
- Maximum 60 words
- Frame around their first-90-day priorities, not our product features
- If they bring relevant experience from previous company, acknowledge it
- "You/your" > "we/our"
- Industry language: {industry_language_guide_from_icp}
- Natural transition to CTA (appended separately)

Output body only.
```

---

## Prompt 7: LinkedIn Connection Request

**Column name:** `copy_linkedin`
**Column type:** AI prompt

```
Write a LinkedIn connection request to {{contact_first_name}} ({{contact_title}} at {{company_name}}, previously {{trigger_previous_title}} at {{trigger_previous_company}}).

From: {sender_name} at {client_name}

## Rules
- Maximum 280 characters
- Congratulate briefly + reference the specific transition
- Connect on shared industry interest or relevant topic
- Do NOT pitch or mention product
- Warmer and more personal than email — LinkedIn is peer-to-peer

Output ONLY the message.
```

---

## Prompt 8: CTA (Signal-Stacking Aware)

**Column name:** `copy_cta`
**Column type:** AI prompt

```
Write a CTA for job-change triggered outbound from {client_name}.

## Prospect Context
- Title: {{contact_title}}
- ICP tier: {{score_icp_tier}}
- Days in role: {{trigger_days_in_role}}
- Signal composite score: {{signal_composite_score}} (if available)

## CTA Rules
{cta_rules_from_strategy}

## Job Change CTA Calibration
- Signal composite ≥100 (Hot) OR A+ tier: Direct but respectful ("worth a 15-min conversation about [specific challenge]?")
- Signal composite 70-99 (Warm): Resource share ("I put together a [framework/benchmark] for [role] leaders in their first 90 days — want me to send it?")
- Standard: Low-commitment value ("happy to share how [industry] teams approach [pain] — no strings")
- New-in-role people prefer insights over meetings — they're in 100 meetings already
- Never "let's hop on a call" — offer benchmarks, frameworks, intros instead

## Rules
- Maximum 20 words
- Question or soft offer, never a demand
- Match: {preferred_cta_style}

Output ONLY the CTA.
```

---

## Prompt 9: Meta Tagging

**Column name:** `meta_hook_type_used` + `meta_prompt_version`
**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "job-change-trigger"
```
