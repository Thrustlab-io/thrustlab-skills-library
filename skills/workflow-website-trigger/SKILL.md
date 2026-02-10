---
name: workflow-website-trigger
description: Builds the complete website visitor trigger workflow — Clay table config, enrichment sequence, Claygent research prompts, and multi-step copy cadence. Use when the strategy recommends website intent as a trigger play. Copy angle is recency and intent-based — the prospect was JUST looking at the client's solution. Reads from all upstream client files.
---

# Website Visitor Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for website visitor intent signals.

**Copy DNA:** Recency + inferred intent. The prospect visited specific pages — we know what they're exploring. Copy is warm, timely, and connects their browsing behavior to a relevant pain point. Never creepy ("I saw you on our site"), always insightful ("teams exploring [area] usually...").

**Prerequisites:**
- `Prospects/{client-slug}/profile.md`
- `Prospects/{client-slug}/strategy.md`
- `Prospects/{client-slug}/icp-mapping.md` (persona cards, pain mapping, fallbacks)
- `Prospects/{client-slug}/market-mapping.md` (for enrichment field availability)
- `Prospects/{client-slug}/tooling-setup.md` (confirms which intent tool — RB2B, Clearbit, 6sense)

## Workflow

### Step 1: Load All Client Context

Read upstream files and extract:
- **profile.md:** Product, value prop, key differentiators, tone
- **strategy.md:** Website trigger section, messaging architecture, CTA rules
- **icp-mapping.md:** Persona cards, industry pain mapping, persona × vertical matrix, fallback copy
- **tooling-setup.md:** Which intent provider is configured

### Step 2: Build Clay Table Schema

See `references/enrichment-sequence.md` for the full column-by-column configuration.

**Unique columns for this trigger:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_pages_visited` | Import field | Pages the visitor viewed (from intent tool) |
| `trigger_visit_recency_hours` | Formula | Hours since last visit |
| `trigger_page_intent_score` | AI prompt | Score: Pricing/demo page = High, solution pages = Medium, blog = Low |
| `trigger_visit_context` | Claygent | Research what problem they're likely solving based on pages + industry |
| `research_company_snapshot` | Claygent | Recent company news, initiatives, challenges |
| `research_pain_inference` | AI prompt | Infer pain from pages visited + industry + company context |
| `copy_opener_company` | AI prompt | Observational opener based on company research (no trigger mention) |
| `copy_opener_trigger` | AI prompt | Intent-aware opener based on pages visited + pain inference |
| `copy_body` | AI prompt | Email body connecting inferred pain → client value prop |
| `copy_linkedin` | AI prompt | LinkedIn connection request (≤280 chars) |
| `copy_cta` | AI prompt | Stage-appropriate CTA |

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md` for every prompt, fully templated with client variables.

**Critical rules for this trigger's prompts:**
- NEVER say "I saw you visited our website" — that's surveillance, not insight
- DO say "Teams exploring [area] often..." or "When companies in [industry] start evaluating [solution category]..."
- Recency matters: <24 hours = more direct, 24-72 hours = softer, >72 hours = general approach
- Page intent shapes the angle: pricing page = buying mode, blog = education mode, case study = validation mode

### Step 4: Generate Copy Templates

See `references/copy-templates.md` for the full cadence.

**Cadence structure for website trigger:**
- **Email 1 (Day 0):** Intent-aware opener + pain inference + soft CTA
- **Email 2 (Day 3):** Value-add — share a relevant resource/insight related to pages visited
- **Email 3 (Day 7):** Social proof — case study or result from similar company in their vertical
- **LinkedIn request (Day 1):** Parallel touch, no email reference, company-observation based
- **LinkedIn follow-up (Day 4):** If connected, brief value message
- **Email 4 (Day 12):** Breakup — final touch with different angle

### Step 5: Generate Enrichment Sequence

See `references/enrichment-sequence.md` for the technical flow.

### Step 6: Quality Gate

Before outputting any copy or prompts, validate:
- [ ] Every `{client_*}` variable resolves to actual client data from profile.md
- [ ] Pain points reference pains from icp-mapping.md, not generic B2B pains
- [ ] Industry language matches the vertical's language guide from icp-mapping.md
- [ ] CTA matches the client's sales motion (PLG / sales-led / hybrid)
- [ ] Fallback values are industry-specific from icp-mapping.md fallback section
- [ ] No prompt says "I saw you visited" or similar surveillance language
- [ ] Copy ≤90 words per email, subject ≤45 chars, LinkedIn ≤280 chars

### Step 7: Output

Save all outputs to a structured document the user can use to configure Clay:
- Clay table schema with column types and order
- Every Clay prompt (copy-pasteable into Clay AI columns)
- Copy templates for the sequencer
- Enrichment sequence diagram

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. For website visitors, the hook_type classification should weight:
- If visitor hit pricing page + we have timeline data → "timeline" (they're evaluating, show speed)
- If visitor hit case study page + we have named case studies → "social_proof"
- If we have quantified metrics for their industry → "numbers"
- Fallback → "hypothesis"

All copy generation prompts in `references/clay-prompts.md` should branch on `hook_type`. See `shared/references/hook-types-guide.md`.

**Signal Stacking:** Website visit signals feed into the Signal Aggregation Table (if active):
- Pricing page visit: +25 points (decay: -5/week)
- Solution page visit: +15 points (decay: -5/week)
- Case study page: +10 points (decay: -3/week)
- Blog only: +5 points (decay: -3/week)

If composite score crosses tier thresholds, escalate treatment. See `shared/references/signal-stacking-guide.md`.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used` for attribution. See `shared/references/prompt-iteration-pipeline.md`.
