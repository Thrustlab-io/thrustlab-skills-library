---
name: workflow-funding-trigger
description: Builds the complete funding trigger workflow — Clay table config, enrichment sequence, Claygent prompts, and copy cadence. Use when the strategy recommends funding events as a trigger play. Copy angle is growth momentum — the company just raised money and is investing in scaling. Timing is critical (first 2-4 weeks post-announcement). Reads from all upstream client files.
---

# Funding Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for funding event signals.

**Copy DNA:** Growth momentum + investment thesis alignment. When a company raises money, they've committed to a growth plan. The funding announcement often reveals WHERE they're investing (hiring, product, expansion). Copy connects the client's solution to the company's stated growth priorities, NOT just "congrats on the raise." The investor angle can add credibility — "companies backed by [investor] typically invest in [area]."

**Prerequisites:**
- `Prospects/{client-slug}/profile.md`
- `Prospects/{client-slug}/strategy.md`
- `Prospects/{client-slug}/icp-mapping.md`
- `Prospects/{client-slug}/tooling-setup.md`

## Workflow

### Step 1: Load Client Context

Special attention to:
- **icp-mapping.md:** Which company sizes / stages align post-funding? (Series A = different than Series D)
- **strategy.md:** Funding trigger section, how client's solution maps to growth-phase challenges

### Step 2: Build Clay Table Schema

**Unique columns for this trigger:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_funding_amount` | Import | Amount raised |
| `trigger_funding_round` | Import | Seed / A / B / C / D / Growth / Debt |
| `trigger_funding_date` | Import | Announcement date |
| `trigger_lead_investors` | Import | Key investors (credibility signal) |
| `trigger_funding_freshness_days` | Formula | Days since announcement |
| `trigger_growth_priorities` | Claygent | Research what they plan to do with the money |
| `trigger_funding_relevance` | AI prompt | How does their growth plan connect to {client_product}? |
| `research_hiring_surge` | Claygent | Are they hiring aggressively post-funding? What roles? |

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md`.

**Critical rules:**
- "Congrats on the raise" is the WORST opener — everyone sends it. Lead with what the funding MEANS for their business.
- Funding announcements often include quotes about priorities — use those as your angle
- Round stage matters: Seed/A = scrappy, need efficiency tools. B/C = scaling, need infrastructure. D+ = optimizing, need enterprise solutions.
- Freshness: <14 days = timely. 14-30 days = still relevant. >30 days = stale, switch angle to "as you scale."

### Step 4: Generate Copy Templates

See `references/copy-templates.md`.

**Cadence:**
- **Email 1 (Day 0):** Growth priority observation + how client helps scale [specific function]
- **LinkedIn (Day 1):** Growth-focused connection, shared industry interest
- **Email 2 (Day 5):** Scaling challenge value-add — resource/benchmark for their stage
- **Email 3 (Day 10):** Social proof — company at similar stage that scaled with client's help
- **Email 4 (Day 16):** Breakup — "before the team doubles" urgency

### Step 5: Quality Gate

- [ ] No "Congrats on the funding!" as the opener
- [ ] Funding amount/round referenced only when adding context (not bragging for them)
- [ ] Growth priorities come from actual research (press release, founder quotes), not assumptions
- [ ] Stage-appropriate messaging (don't pitch enterprise to a seed-stage company)
- [ ] Freshness < 30 days
- [ ] Standard copy rules

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
