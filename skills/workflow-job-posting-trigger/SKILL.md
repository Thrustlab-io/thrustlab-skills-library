---
name: workflow-job-posting-trigger
description: Builds the complete job posting trigger workflow — Clay table config, enrichment sequence, Claygent prompts, and copy cadence. Use when the strategy recommends job postings as a trigger play. Copy angle is hiring-as-pain-signal — the company is hiring for a role that reveals a gap the client's product fills. Reads from all upstream client files.
---

# Job Posting Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for job posting signals.

**Copy DNA:** Hiring reveals pain. When a company posts a job, they're admitting they have a gap. The job description is a goldmine — it lists the exact responsibilities, tools, and challenges the company faces. Copy connects the hiring need to the client's solution, positioning it as a force multiplier for the new hire (or a reason they might not need to hire at all).

**Prerequisites:**
- `client-profiles/{client-slug}/profile.md`
- `strategies/{client-slug}.md`
- `client-profiles/{client-slug}/icp-mapping.md`
- `client-profiles/{client-slug}/tooling-setup.md` (confirms signal source — Clay signal, Otta, LinkedIn Jobs)

## Workflow

### Step 1: Load Client Context

Pay special attention to:
- **profile.md:** What the product does — is it a tool that augments the hired role, or replaces the need?
- **icp-mapping.md:** Which personas does the client sell to? The job posting should match a RELATED role, not necessarily the buyer persona. E.g., if client sells to VP Operations, a posting for "Operations Analyst" is the signal.
- **strategy.md:** Job posting trigger section, relevant job titles to monitor

### Step 2: Define Job Posting Detection Criteria

Before building Clay columns, define which job postings are relevant:

**Signal job titles:** Titles that, when posted, indicate a need the client solves
- Direct match: {titles from icp-mapping.md personas}
- Adjacent match: {titles one level below the buyer persona — the team they're building}
- Inverse match: {titles that indicate the company is building the capability in-house instead of buying — competitive signal}

**Keyword filters in job descriptions:**
- Include terms: {terms related to client's problem domain}
- Exclude terms: {terms that indicate irrelevant context — e.g., "intern" if targeting enterprise}

### Step 3: Build Clay Table Schema

See `references/enrichment-sequence.md`.

**Unique columns for this trigger:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_job_title_posted` | Import | The specific job title |
| `trigger_job_url` | Import | Link to the posting |
| `trigger_job_description_raw` | Import/Claygent | Full or summarized job description |
| `trigger_job_pain_indicators` | AI prompt | Extract 2-3 pain points from the JD that {client_product} addresses |
| `trigger_job_tools_mentioned` | AI prompt | Technologies/tools mentioned in JD (tech stack signal) |
| `trigger_posting_age_days` | Formula | Days since posted (freshness) |
| `trigger_hiring_intent_score` | AI prompt | How relevant is this posting to {client_product}? HIGH/MEDIUM/LOW |
| `research_team_context` | Claygent | Research the team/department that's hiring |

### Step 4: Generate Clay Prompts

See `references/clay-prompts.md`.

**Critical rules for this trigger's prompts:**
- Parse the job description for PAIN INDICATORS, not just responsibilities
- "Must have experience with X" = they currently lack X
- "Responsible for building/creating Y" = Y doesn't exist yet
- "Managing Z processes" = Z is a pain point they're throwing headcount at
- Connect the hiring gap to how the client's product fills or augments it
- Two valid angles: "This tool helps your new hire succeed faster" OR "This tool handles part of what you're hiring for"

### Step 5: Generate Copy Templates

See `references/copy-templates.md`.

**Cadence structure:**
- **Email 1 (Day 0):** Job posting observation + pain inference + bridge to solution
- **Email 2 (Day 4):** Value-add: how similar companies handle this need (tool vs hire trade-off)
- **LinkedIn request (Day 1):** About their team growth, no pitch
- **Email 3 (Day 9):** Social proof from same industry — results with fewer headcount
- **Email 4 (Day 14):** Breakup — different angle (team productivity, not hiring)

### Step 6: Quality Gate

- [ ] Job description pain indicators are actually extracted from the JD, not assumed
- [ ] Copy doesn't imply "you don't need to hire" (offensive to hiring manager) — instead positions as "help the new hire / existing team do more"
- [ ] Target contact is the HIRING MANAGER or their boss, not the role being hired for
- [ ] Posting age < 30 days (stale postings = stale outreach)
- [ ] Industry language from icp-mapping.md applied
- [ ] Standard copy rules

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
