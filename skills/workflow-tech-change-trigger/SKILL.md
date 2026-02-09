---
name: workflow-tech-change-trigger
description: Builds the complete tech stack change trigger workflow. Use when the strategy recommends technology adoption or removal as a trigger play. Copy angle is stack evolution — the company just adopted a complementary tool (integration opportunity) or dropped a competing tool (replacement window). Reads from all upstream client files.
---

# Tech Stack Change Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for tech stack change signals.

**Copy DNA:** Stack evolution. When a company adds or removes technology, it reveals a strategic shift. Adding a complementary tool = they're investing in the ecosystem your product fits into. Removing a competitor = active evaluation window. Copy demonstrates ecosystem awareness and positions the client's product as the natural next step in their stack evolution.

**Prerequisites:**
- All standard upstream files
- `client-profiles/{client-slug}/tooling-setup.md` (confirms BuiltWith/Wappalyzer setup)

## Workflow

### Step 1: Load Client Context

Special attention to:
- **profile.md:** Client's tech integrations, complementary tools, competitor products
- **icp-mapping.md:** Tech stack signals in ICP scoring, ideal tech stack patterns

### Step 2: Define Tech Stack Signals

Before building Clay columns, define what tech changes matter:

**Complementary additions (positive signal):**
- {Tool that integrates with client's product} — "They just added X, which works great with {client_product}"
- {Tool that indicates a strategic shift toward client's category}

**Competitor removals (high-intent signal):**
- {Competitor_1} removed — active replacement window
- {Competitor_2} removed — evaluation in progress

**Competitor additions (negative signal — possibly disqualify):**
- Just added {competitor} — they may have already decided

### Step 3: Build Clay Table Schema

**Unique columns:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_tech_added` | Import/Enrichment | Technology recently added |
| `trigger_tech_removed` | Import/Enrichment | Technology recently removed |
| `trigger_tech_change_type` | AI prompt | COMPLEMENTARY_ADD / COMPETITOR_REMOVED / COMPETITOR_ADDED / STACK_SHIFT |
| `trigger_tech_change_date` | Import | When the change was detected |
| `trigger_stack_context` | Claygent | What else is in their stack? How does this change fit? |
| `research_tech_decision_context` | Claygent | WHY did they make this change? (migration, consolidation, growth) |

### Step 4: Generate Prompts + Copy

See `references/clay-prompts.md`.

**Critical rules:**
- Two distinct copy paths:
  1. **Complementary add:** "Now that you've invested in [tool], the natural next step is [area {client_product} covers]"
  2. **Competitor removed:** "Evaluating alternatives to [competitor]? Here's what [industry] teams are looking for"
- Never bash the competitor directly
- Show you understand their tech ecosystem, not just one tool
- If `trigger_tech_change_type` = "COMPETITOR_ADDED" → auto-disqualify or drastically different approach

### Step 5: Cadence

- **Email 1 (Day 0):** Stack observation + how client complements/replaces
- **LinkedIn (Day 1):** Tech ecosystem interest, shared challenges
- **Email 2 (Day 5):** Integration value / migration guide (depending on signal type)
- **Email 3 (Day 10):** Social proof: company with similar stack that added client's product
- **Email 4 (Day 15):** Breakup with technical angle

### Step 6: Quality Gate

- [ ] Tech change type correctly identified (complementary vs competitor)
- [ ] Copy path matches the signal type
- [ ] No competitor bashing
- [ ] Technical accuracy — don't reference integrations that don't exist
- [ ] Standard copy rules

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
