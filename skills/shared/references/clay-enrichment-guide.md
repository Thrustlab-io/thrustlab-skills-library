# Clay Enrichment Guide — Universal Best Practices

Reference for every skill that produces Clay table configurations, enrichment sequences, or Claygent prompts.

## Clay Table Design Principles

### One Table Per Workflow
Each trigger play and the general outbound campaign gets its own Clay table. Don't mix workflows in one table — it makes automation triggers messy and reporting impossible.

Naming: `{client-slug} — {workflow-name}`
Examples: `acme — Website Visitors`, `acme — Job Changes`, `acme — General Outbound`

### Column Ordering
Follow the data flow — left to right mirrors the enrichment sequence:

```
[Source Data] → [Enrichment] → [Trigger Data] → [Research] → [Scoring] → [Copy] → [Export]
```

1. **Source columns**: company_name, company_website, contact_first_name, contact_email, etc.
2. **Enrichment columns**: company_industry, company_size, company_tech_stack, company_hq
3. **Trigger columns**: trigger_type, trigger_detail, trigger_date, trigger_source
4. **Research columns**: research_observation, research_pain_point, research_company_summary
5. **Scoring columns**: score_icp_fit, score_intent, score_engagement
6. **Copy columns**: copy_opener_company, copy_opener_trigger, copy_body, copy_cta, copy_linkedin
7. **Export columns**: export_status, export_date, sequencer_campaign_id

### Credit-Efficient Enrichment

#### Waterfall Enrichment (Contact Data)
Always try cheapest provider first. Stop when data is found.

**Email waterfall:**
1. Hunter.io (~1-2 credits) → check first
2. Apollo (~1-2 credits) → if Hunter empty
3. RocketReach (~2-3 credits) → if Apollo empty
4. Prospeo (~2-3 credits) → last resort
5. Email verification (always, regardless of source)

**Phone waterfall:**
1. Apollo → check first
2. RocketReach → if Apollo empty
3. Cognism/Lusha → premium fallback

#### Conditional Enrichment
Don't enrich every field on every row. Use IF/THEN conditions:
- Only run Claygent research if `score_icp_fit` ≥ 7
- Only generate copy if research columns are populated (not SKIP)
- Only run phone enrichment if email is verified and ICP score is A-tier

#### Batch Testing
- Always test enrichment on 5-10 rows before running full table
- Check credit consumption per row
- Verify data quality before scaling

## Claygent Prompt Standards

### Structure for Research Prompts
```
Research {specific_topic} for {company_name}.

Context: We are {client_name}, a {client_description}. We sell to {target_persona} at {target_company_type}.

Find:
1. [Specific datapoint 1]
2. [Specific datapoint 2]
3. [Specific datapoint 3]

Rules:
- Only include factual, verifiable information
- If you cannot find the information, output "NOT_FOUND"
- Maximum {X} words
- Do not fabricate or assume
```

### Structure for Copy Generation Prompts
```
You are writing outbound copy for {client_name}.

## Client Context
{client_name} ({client_website}): {client_one_liner}
Target buyer: {target_persona_title} at {target_company_type}
Value prop: {value_prop}
Pains we solve: {pain_1}, {pain_2}, {pain_3}
Tone: {client_tone}

## Prospect Context
Company: {{company_name}} ({{company_industry}}, {{company_size}} employees)
Contact: {{contact_first_name}} {{contact_last_name}}, {{contact_title}}
Research: {{research_observation}}
Trigger: {{trigger_detail}}

## Task
[Specific copy instruction]

## Rules
- Maximum {X} words
- {tone} tone
- Never use: [banned phrases from copy-rules.md]
- Reference specific details from prospect context
- If insufficient data, output "SKIP"
```

### Prompt Anti-Patterns in Clay

| Don't | Do |
|---|---|
| "Write a personalized email" (vague) | "Write a 2-sentence opener referencing {{research_observation}}" (specific) |
| Prompts without client context block | Every prompt starts with who the client is and what they sell |
| "Be creative and engaging" | "Use observational tone, reference {{trigger_detail}}, ≤90 words" |
| Generic system prompts | Role-specific: "You are a B2B sales researcher for {client_name}" |
| Prompts that can produce generic output | Include "If insufficient data to personalize, output SKIP" |

## ICP Scoring in Clay

### Default Scoring Formula (Customize Per Client)
Starting score: 0 (build up, not subtract down)

| Signal | Points | Source |
|---|---|---|
| Industry match (exact) | +3 | Enrichment |
| Industry match (adjacent) | +1 | Enrichment |
| Company size in target range | +2 | Enrichment |
| Revenue in target range | +1 | Enrichment |
| Tech stack match (complementary) | +1 | Enrichment |
| Geography match | +1 | Enrichment |
| Trigger signal detected | +2 | Signal source |
| Persona title match | +2 | Enrichment |
| Funding stage match | +1 | Enrichment |

**Tiers:**
- Score ≥ 9: **A+ Tier** → Priority outreach, full personalization
- Score 7-8: **A Tier** → Standard outreach with trigger-based personalization
- Score 5-6: **B Tier** → Lighter touch, general outbound only
- Score < 5: **C Tier** → Do not contact, review ICP criteria

### Scoring Must Be Client-Specific
The default formula above is a starting point. Every client gets a customized scoring formula based on:
- Their actual ICP criteria from `icp-mapping.md`
- Weighted by what matters most for their sales motion
- Validated against their existing customer base (if available)

## Signal Stacking — Composite Scoring Layer

Beyond basic ICP scoring, signal stacking adds a dynamic layer that combines multiple buying signals. See `shared/references/signal-stacking-guide.md` for the complete framework.

### How It Interacts With ICP Scoring

```
ICP Score (static fit)     → Qualifies the account (gate: must pass before signal processing)
Signal Composite Score     → Prioritizes and routes qualified accounts based on active buying signals
```

ICP scoring answers: "Is this a good account for us?"
Signal stacking answers: "Is this account ready to buy NOW, and how should we engage?"

### Signal Aggregation Table
For clients with 2+ active trigger plays, create a master Signal Aggregation Table that:
1. Receives rows from each trigger play table (via webhook or scheduled merge)
2. Calculates time-decayed composite scores across all signals
3. Routes accounts to appropriate treatment tier (Hot/Warm/Active/Watching)
4. Alerts the team via Slack for Hot tier accounts

### Dark Funnel Signal Sources

Dark funnel signals illuminate the 70%+ of the B2B buyer journey that happens anonymously:

| Tool | Signal Type | Clay Integration | Cost |
|---|---|---|---|
| **Dealfront** (EU-strong) | Website visitor ID, company-level | Webhook → Clay | ~€199+/mo |
| **RB2B** | Person-level website visitor ID (US-focused) | Webhook → Clay | ~$149-399/mo |
| **Teamfluence** | LinkedIn engagement monitoring (likes, comments, shares by ICP) | Export → Clay import | ~€99+/mo |
| **Trigify** | LinkedIn trigger event monitoring (job changes, posts, engagement) | Webhook/API → Clay | ~€149+/mo |
| **Common Room** | Community signals (Slack, Discord, GitHub, forums) | API → Clay | ~$500+/mo |

**Preferred tooling by market:**
- EU clients: Dealfront (website) + Teamfluence (LinkedIn engagement) + Trigify (LinkedIn triggers)
- US clients: RB2B (website) + Trigify (LinkedIn triggers) + Common Room (community)
- Global: Dealfront + RB2B + Teamfluence + Trigify (full coverage)

## Export & Sequencer Integration

### Fields to Export
Every Clay table should have these export-ready columns:
- `contact_email` (verified)
- `contact_first_name`
- `contact_last_name`
- `company_name`
- `copy_opener_company` or `copy_opener_trigger` (best available)
- `copy_body`
- `copy_cta`
- `copy_linkedin`
- `meta_workflow_name` (for tracking which play generated the lead)

### Export Rules
- Only export rows where `score_icp_fit` ≥ threshold (default 7)
- Only export rows where email is verified
- Only export rows where copy columns are populated (not SKIP)
- Tag with workflow name for sequencer-level reporting
