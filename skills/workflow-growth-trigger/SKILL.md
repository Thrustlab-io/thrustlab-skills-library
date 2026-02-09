---
name: workflow-growth-trigger
description: Builds the complete growth/hiring surge trigger workflow. Use when the strategy recommends headcount growth, new office openings, or market expansion as a trigger play. Copy angle is scaling pain — rapid growth creates operational strain that the client's product alleviates. Reads from all upstream client files.
---

# Growth / Hiring Surge Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for growth signals.

**Copy DNA:** Scaling strain. When a company grows >20% headcount in a quarter, opens a new office, or enters a new market, existing processes break. What worked for 50 people doesn't work for 100. Copy connects this scaling pain to the client's solution, framing it as infrastructure for the next stage. Tone is empathetic (scaling is hard) not congratulatory (everyone says "congrats on the growth").

**Prerequisites:**
- All standard upstream files

## Workflow

### Step 1: Load Client Context

Special attention to:
- **icp-mapping.md:** Which company size transitions create the most pain? (50→100, 100→250, 250→500)
- **strategy.md:** Growth trigger section, which scaling challenges map to client's product

### Step 2: Build Clay Table Schema

**Unique columns:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_growth_type` | AI prompt | HEADCOUNT_SURGE / NEW_OFFICE / NEW_MARKET / DEPARTMENT_BUILD |
| `trigger_headcount_current` | Enrichment | Current employee count |
| `trigger_headcount_6mo_ago` | Enrichment | Employee count 6 months ago (for growth rate) |
| `trigger_growth_rate_pct` | Formula | ((current - 6mo_ago) / 6mo_ago) × 100 |
| `trigger_growth_departments` | Claygent | Which departments are growing fastest? |
| `research_scaling_challenges` | AI prompt | Infer scaling challenges from growth pattern + industry |

### Step 3: Generate Prompts + Copy

**Critical rules:**
- Don't lead with "Congrats on the growth" — lead with the CHALLENGE growth creates
- Growth rate matters: 20-50% = measured scaling. 50-100% = hyper-growth. >100% = chaos mode.
- Department-level growth tells you where the pain is (engineering growing = product scaling, sales growing = GTM scaling, ops growing = process strain)
- Frame client product as "infrastructure for scale," not a nice-to-have

### Step 4: Cadence

- **Email 1 (Day 0):** Scaling observation + specific challenge at their growth rate
- **LinkedIn (Day 1):** Industry scaling interest, peer connection
- **Email 2 (Day 5):** Scaling playbook / benchmark for their size transition
- **Email 3 (Day 10):** Social proof: company that scaled through same transition
- **Email 4 (Day 16):** Breakup: "before the growing pains compound"

### Step 5: Quality Gate

- [ ] Growth data is real (enriched, not assumed)
- [ ] Scaling challenge is specific to their growth rate + industry, not generic
- [ ] Not confusing growth with funding (different trigger, different angle)
- [ ] Department-level insight used when available
- [ ] Standard copy rules

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
