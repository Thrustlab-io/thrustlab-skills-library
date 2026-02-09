---
name: workflow-compliance-trigger
description: Builds the complete compliance/regulatory trigger workflow. Use when the strategy recommends regulatory deadlines, new laws, or policy changes as a trigger play. Copy angle is urgency from external mandate — the company MUST act due to a regulation, not just should. This creates the strongest urgency of any trigger type. Highly vertical-specific. Reads from all upstream client files.
---

# Compliance / Regulatory Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for regulatory and compliance signals.

**Copy DNA:** External mandate urgency. Unlike other triggers where the company chooses to act, compliance triggers mean they MUST act — there's a deadline, a fine, or a legal requirement. This is the highest-urgency trigger type. Copy combines empathy (compliance is stressful) with authority (demonstrating deep understanding of the regulation) and positions the client's product as the compliance path of least resistance.

**Prerequisites:**
- All standard upstream files
- **CRITICAL:** This trigger requires deep vertical knowledge. The icp-mapping.md industry pain section must include specific regulatory context.

## Workflow

### Step 1: Load Client Context

Special attention to:
- **icp-mapping.md:** Vertical-specific regulatory landscape, compliance pain points
- **strategy.md:** Compliance trigger section, which regulations map to client's product
- **profile.md:** How specifically does the client's product help with compliance?

### Step 2: Define Regulatory Signal Criteria

This trigger only works when you can map a specific regulation → specific client capability:

**Regulation inventory for this client:**

| Regulation/Standard | Affects | Deadline | How {client_product} Helps |
|---|---|---|---|
| {regulation_1} | {who} | {when} | {specific capability} |
| {regulation_2} | {who} | {when} | {specific capability} |

### Step 3: Build Clay Table Schema

**Unique columns:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_regulation_name` | Import/Manual | The specific regulation or standard |
| `trigger_regulation_deadline` | Import/Manual | Compliance deadline date |
| `trigger_days_until_deadline` | Formula | Urgency countdown |
| `trigger_regulation_impact` | AI prompt | How does this regulation specifically affect {{company_name}}? |
| `trigger_penalty_risk` | Claygent | What are the penalties for non-compliance? |
| `research_compliance_readiness` | Claygent | Signs of whether {{company_name}} is prepared or scrambling |

### Step 4: Generate Prompts + Copy

**Critical rules:**
- ACCURACY IS NON-NEGOTIABLE — never misstate a regulation, deadline, or requirement
- Demonstrate genuine regulatory knowledge (use correct terminology, reference specific articles/sections)
- Urgency is built-in — don't manufacture fake urgency on top of a real deadline
- Two personas: compliance officer (detail-oriented, wants specifics) vs. executive (wants the risk summary)
- Deadline proximity matters: >6 months = educational, 3-6 months = planning, <3 months = urgent, past = remediation

### Step 5: Cadence

- **Email 1 (Day 0):** Regulation observation + how it affects their specific business + bridge to help
- **Email 2 (Day 5):** Compliance checklist or readiness framework (pure value)
- **LinkedIn (Day 2):** Connect on regulatory topic in their industry
- **Email 3 (Day 10):** Social proof: how a peer company achieved compliance using client's product
- **Email 4 (Day 15):** Deadline reminder + offer to help (if deadline < 3 months)

### Step 6: Quality Gate

- [ ] Regulation name, deadline, and requirements are FACTUALLY ACCURATE (verify with web search)
- [ ] Client's product genuinely helps with this specific regulation (not a stretch)
- [ ] Deadline proximity reflected in urgency level
- [ ] Industry language is correct — regulators have specific terminology
- [ ] No fearmongering — factual urgency only
- [ ] Standard copy rules

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
