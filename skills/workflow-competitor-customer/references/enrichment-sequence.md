# Enrichment Sequence — Competitor Customer Trigger

## Data Flow

```
ICP List / Clay Import → Competitor Detection
  → Column 1-6: Company + Contact Import
    → Column 7-9: Competitor Detection + Confidence Gate
      → Column 10-11: ICP Scoring + Gate
        → Column 12-15: Competitive Intelligence Research
          → Column 16-22: Hook Type + Copy Generation
            → Column 23-26: Assembly + Signal Stacking + Export
```

## Column-by-Column Configuration

### Import Layer (Columns 1-6)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `company_name` | Import | ICP list / Clay import | Target company |
| 2 | `company_website` | Import | ICP list / enrichment | Company domain |
| 3 | `company_industry` | Enrichment | Clearbit/Apollo | Industry classification |
| 4 | `company_size` | Enrichment | Clearbit/Apollo | Employee count |
| 5 | `contact_first_name` | Import/Enrichment | Import or waterfall | Must be decision-maker for tool category |
| 6 | `contact_title` | Import/Enrichment | Import or waterfall | |

### Competitor Detection Layer (Columns 7-9)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 7 | `competitor_product_used` | Claygent + BuiltWith | See clay-prompts.md Layer 1 | Which competitor they use |
| 8 | `competitor_confidence` | AI Prompt | See clay-prompts.md | "confirmed" / "likely" / "possible" / "none" |
| 9 | `competitor_detection_source` | Formula | Which detection method succeeded | "builtwith" / "claygent_research" / "job_posting" / "g2_review" |

**GATE:** If `competitor_confidence` = "possible" or "none" → route to general outbound. Do NOT proceed with competitor-specific copy on unconfirmed data. Only "confirmed" and "likely" continue.

**Detection methods (run in order for credit efficiency):**
1. **BuiltWith** (cheapest, highest confidence for SaaS/tech products)
2. **Claygent web research** (searches company website, job postings, G2 reviews for competitor mentions)
3. **AI classification** (analyzes combined signals to assign confidence level)

### ICP Scoring Gate (Columns 10-11)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 10 | `score_icp_fit` | Formula/AI | Standard ICP scoring | A+ / A / B / DQ |
| 11 | `score_icp_tier` | Formula | ICP fit + competitor bonus | Competitor customers get +1 tier boost |

**Second gate:** DQ companies that don't fit ICP regardless of competitor usage.

### Competitive Intelligence Research (Columns 12-15)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 12 | `research_competitor_usage` | Claygent | See clay-prompts.md Layer 2 | How they use the competitor product |
| 13 | `research_competitor_pain` | AI Prompt | See clay-prompts.md | Known limitations relevant to this company |
| 14 | `research_switch_angle` | AI Prompt | See clay-prompts.md | Most compelling reason THIS company would switch |
| 15 | `competitor_customer_name` | Lookup | Client's CRM / case studies | Named customer who switched FROM same competitor |

**Column 15 is critical.** If the client has a customer who switched from the SAME competitor → copy uses social proof hook with 3x conversion rate. This lookup should search: client case studies, client CRM won deals with "switched from [competitor]" tag.

### Hook Type + Copy Generation (Columns 16-22)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 16 | `hook_type` | AI Prompt | See clay-prompts.md | social_proof (strongest if switcher exists) / timeline / numbers / hypothesis |
| 17 | `copy_opener_trigger` | AI Prompt | Branched by hook_type | Competitive seed opener |
| 18 | `copy_body` | AI Prompt | Competitive intelligence angle | Body with switch angle |
| 19 | `copy_cta` | AI Prompt | Signal-stacking aware | |
| 20 | `copy_linkedin` | AI Prompt | Category connection, NO competitor name | |
| 21 | `copy_email_subject` | AI Prompt | Category-focused, not competitor-focused | ≤45 chars |
| 22 | `copy_email_2_switch_story` | AI Prompt | Named switcher story for Day 5 | If available |

### Assembly + Export (Columns 23-26)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 23 | `signal_composite_score` | Formula/Webhook | From Signal Aggregation Table | +20 persistent relationship signal |
| 24 | `export_email_1` | Formula | Concatenate: opener + body + CTA | |
| 25 | `export_status` | Formula | "READY" if confirmed/likely + copy populated + email verified | |
| 26 | `meta_hook_type_used` + `meta_prompt_version` | Formula/Static | A/B tracking | |

## Credit Optimization

- BuiltWith check FIRST (cheapest, most reliable for tech detection)
- Claygent research ONLY if BuiltWith returns no result
- ICP gate BEFORE competitive intelligence research — save credits on DQ companies
- Competitor confidence gate at column 9 prevents wasting credits on uncertain data
- `competitor_customer_name` lookup is a table join, no credits

## Automation Trigger

This workflow can be triggered two ways:
1. **Batch:** Import ICP list → run competitor detection → filter confirmed/likely → continue
2. **Signal-based:** BuiltWith alert detects competitor tech added/removed → Clay webhook → enrich

For batch: run competitor detection on full ICP list first, then continue enrichment only on confirmed/likely matches.
