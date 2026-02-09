---
name: workflow-job-change-trigger
description: Builds the complete job role change trigger workflow — Clay table config, enrichment sequence, Claygent research prompts, and multi-step copy cadence. Use when the strategy recommends job changes as a trigger play. Copy angle is the "new broom" window — a person just started a new role and is evaluating tools, processes, and vendors. Reads from all upstream client files.
---

# Job Role Change Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for job change signals.

**Copy DNA:** Transition momentum. A target persona just changed roles — they're in "new broom" mode, evaluating everything, eager to make an impact in the first 90 days. Copy acknowledges the move, references their background, and positions the client's solution as part of their success at the new company. Warm, congratulatory, never presumptuous.

**Prerequisites:**
- `client-profiles/{client-slug}/profile.md`
- `strategies/{client-slug}.md`
- `client-profiles/{client-slug}/icp-mapping.md`
- `client-profiles/{client-slug}/market-mapping.md`
- `client-profiles/{client-slug}/tooling-setup.md` (confirms signal source — UserGems, Clay signal, or manual SN alerts)

## Workflow

### Step 1: Load All Client Context

Same upstream file loading as other triggers. Pay special attention to:
- **icp-mapping.md:** Persona cards (what does this role care about in their first 90 days?)
- **strategy.md:** Job change trigger section, messaging for "new in role" angle

### Step 2: Build Clay Table Schema

See `references/enrichment-sequence.md` for full column config.

**Unique columns for this trigger:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_previous_company` | Import/Enrichment | Where they came from |
| `trigger_previous_title` | Import/Enrichment | What they did before |
| `trigger_days_in_role` | Formula | How new they are (days since start date) |
| `trigger_role_transition_type` | AI prompt | Promotion / lateral / new company / industry change |
| `trigger_new_role_context` | Claygent | Research what challenges come with this transition |
| `research_previous_company_relevance` | AI prompt | Did previous company use similar solutions? Any learnings they'd bring? |

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md`.

**Critical rules for this trigger's prompts:**
- Congratulate naturally but briefly — don't gush
- Reference the transition specifically (previous role → new role), not just "congrats on the new job"
- The "new broom" window has phases: 0-30 days = learning, 30-60 days = planning, 60-90 days = executing
- Tailor urgency to `trigger_days_in_role`
- If they came from a company that used the client's product (or a competitor), that's gold — reference it

### Step 4: Generate Copy Templates

See `references/copy-templates.md`.

**Cadence structure for job change trigger:**
- **Email 1 (Day 0):** Congratulations + observation about the transition + soft bridge to relevance
- **LinkedIn request (Day 0):** Congratulations + shared industry interest (parallel, NOT duplicate)
- **Email 2 (Day 5):** "First 90 days" value — share an insight relevant to their new role's challenges
- **LinkedIn follow-up (Day 6):** Brief value share if connected
- **Email 3 (Day 10):** Social proof — how a peer in same role used client's solution after a similar transition
- **Email 4 (Day 17):** Different angle — approach from their team's perspective, not just theirs

### Step 5: Quality Gate

Before outputting:
- [ ] Transition-specific details (previous company, previous title) are used, not ignored
- [ ] Copy adjusts tone based on days_in_role (warmer/softer if very new, more direct if 60+ days)
- [ ] Persona messaging rules from icp-mapping.md applied to this specific title
- [ ] No generic "congrats on the new role" without a specific observation
- [ ] Industry language matches their NEW company's vertical (not the old one)
- [ ] Standard copy rules (≤90 words, ≤45 char subject, etc.)

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
