# Enrichment Sequence — Champion Tracking Trigger

## Data Flow

```
CRM Export / UserGems Alert → Clay Import
  → Column 1-6: Champion Data Import + New Company Enrichment
    → Column 7-8: ICP Scoring (new company) + Gate
      → Column 9-12: Research (new company context + transition insight)
        → Column 13-14: Hook Type + Copy Generation
          → Column 15-17: Assembly + Signal Stacking + Export
```

## Column-by-Column Configuration

### Import Layer — Champion Data (Columns 1-6)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `champion_first_name` | Import | CRM export / UserGems webhook | The person who changed jobs |
| 2 | `champion_previous_company` | Import | CRM / UserGems | Where they worked before |
| 3 | `champion_previous_title` | Import | CRM / UserGems | Previous role |
| 4 | `champion_relationship_type` | Import | CRM | "customer" / "prospect" / "meeting_had" / "engaged" |
| 5 | `champion_product_used` | Import | CRM | Which {client_product} feature/tier they used (if customer) |
| 6 | `champion_result_achieved` | Import | CRM / CS notes | Specific result if known (e.g., "cut review time by 35%") |

**CRITICAL:** Columns 4-6 come from the CRM export. The client MUST provide this data. Without `champion_relationship_type`, copy can't be properly tiered. Without `champion_result_achieved`, social proof hooks are weaker.

### New Company Enrichment (Columns 7-12)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 7 | `company_name` | Import | UserGems / LinkedIn | Their NEW company |
| 8 | `contact_title` | Import | UserGems / LinkedIn | Their NEW title |
| 9 | `company_industry` | Enrichment | Clearbit/Apollo (by domain) | New company industry |
| 10 | `company_size` | Enrichment | Clearbit/Apollo | New company employee count |
| 11 | `company_website` | Enrichment | Clearbit/Apollo | New company domain |
| 12 | `champion_days_in_role` | Formula | `(TODAY() - job_change_date)` | Urgency calibration |

### ICP Scoring Gate (Columns 13-14)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 13 | `score_icp_fit` | Formula/AI | ICP scoring on the NEW company | A+ / A / B / DQ |
| 14 | `score_icp_tier` | Formula | Combined ICP + relationship bonus | Champions get +1 tier boost |

**GATE:** If `score_icp_fit` = "DQ" on the new company, mark SKIP — even though they're a champion, their new company doesn't fit. Saves credits and prevents irrelevant outreach.

**Champion tier boost:** A company that would normally be "A tier" gets promoted to "A+ tier" if the champion was a customer. This reflects the higher conversion probability.

### Research Layer (Columns 15-18)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 15 | `research_new_company` | Claygent | See clay-prompts.md Layer 2 | What the new company does, relevant to client |
| 16 | `research_role_context` | AI Prompt | See clay-prompts.md | How new role relates to client's solution |
| 17 | `research_transition_insight` | AI Prompt | See clay-prompts.md | What moving from old→new means |
| 18 | `overlap_new_company` | AI Prompt | See clay-prompts.md | Where client product fits at new company |

### Hook Type + Copy Generation (Columns 19-25)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 19 | `hook_type` | AI Prompt | See clay-prompts.md | timeline/numbers/social_proof/hypothesis |
| 20 | `copy_opener_trigger` | AI Prompt | Branched by hook_type × relationship_type | Reconnection opener |
| 21 | `copy_body` | AI Prompt | Branched by relationship_type | Body connecting old relationship to new context |
| 22 | `copy_cta` | AI Prompt | Signal-stacking aware | Stage-appropriate CTA |
| 23 | `copy_linkedin` | AI Prompt | Genuine reconnection, no pitch | LinkedIn message |
| 24 | `copy_email_subject` | AI Prompt | Personal — `{first_name} + {new_company}` | ≤45 chars |
| 25 | `copy_email_2_value` | AI Prompt | Value bridge for Day 7 email | Resource share for new role |

### Assembly + Signal Stacking + Export (Columns 26-30)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 26 | `signal_composite_score` | Formula/Webhook | From Signal Aggregation Table (if exists) | Check for other signals on new company |
| 27 | `export_email_1` | Formula | Concatenate: opener + body + CTA | Full reconnection email |
| 28 | `export_status` | Formula | "READY" if all populated + email exists | |
| 29 | `meta_hook_type_used` | Formula | = `hook_type` | A/B tracking |
| 30 | `meta_prompt_version` | Static | "v1.0" | Iteration tracking |

## Credit Optimization

- Company enrichment on NEW company BEFORE research — DQ early if bad fit
- Champion data (columns 1-6) comes from CRM — no enrichment credits needed
- Contact email usually available from UserGems — skip Hunter/Apollo waterfall
- `champion_days_in_role` > 100 doesn't auto-DQ, but should lower priority
- Batch test: run 10 champion records through full sequence before activating

## Automation Trigger

When UserGems detects a job change OR when CRM export identifies a match:
1. New row with champion data (columns 1-6) + new company (7-8)
2. Columns 9-14 run (enrich + ICP gate)
3. If passes → columns 15-30 run
4. If `signal_composite_score` ≥ 100 → Slack alert to AE for manual follow-up
5. Otherwise, if `export_status` = "READY" → push to champion-specific sequencer
