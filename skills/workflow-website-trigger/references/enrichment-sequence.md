# Enrichment Sequence — Website Visitor Trigger

## Data Flow

```
Intent Tool (RB2B/Clearbit) → Clay Webhook
  → Column 1-4: Import + Company Enrichment
    → Column 5-6: Trigger Analysis
      → Column 7-8: Research
        → Column 9: ICP Scoring + Gate
          → Column 10-15: Copy Generation
            → Column 16: Final Assembly + Export
```

## Column-by-Column Configuration

### Import Layer (Columns 1-4)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `company_name` | Import | Intent tool webhook | Company identified from visit |
| 2 | `company_website` | Import | Intent tool webhook | Domain of visiting company |
| 3 | `trigger_pages_visited` | Import | Intent tool webhook | Comma-separated page URLs |
| 4 | `trigger_visit_timestamp` | Import | Intent tool webhook | When the visit occurred |

### Company Enrichment Layer (Columns 5-8)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 5 | `company_industry` | Enrichment | Clearbit/Apollo (by domain) | Industry classification |
| 6 | `company_size` | Enrichment | Clearbit/Apollo | Employee count |
| 7 | `company_revenue` | Enrichment | Clearbit/Apollo | Revenue range if available |
| 8 | `company_location` | Enrichment | Clearbit/Apollo | HQ location |

### Trigger Analysis Layer (Columns 9-10)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 9 | `trigger_visit_recency_hours` | Formula | Calculate from `trigger_visit_timestamp` | `(NOW() - timestamp) / 3600` |
| 10 | `trigger_page_intent_score` | AI Prompt | See clay-prompts.md #1 | HIGH / MEDIUM / LOW / DISQUALIFY |

### Research Layer (Columns 11-13)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 11 | `trigger_visit_context` | Claygent | See clay-prompts.md #2 | Web research on company + visit context |
| 12 | `research_company_snapshot` | Claygent | Standard company research prompt | Recent news, initiatives |
| 13 | `research_pain_inference` | AI Prompt | See clay-prompts.md #3 | Inferred pain from visit + industry |

### ICP Scoring Gate (Column 14)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 14 | `score_icp_fit` | Formula/AI | ICP scoring formula from icp-mapping.md | A+ / A / B / DQ |

**GATE:** If `trigger_page_intent_score` = "DISQUALIFY" OR `score_icp_fit` = "DQ", stop here. Do not enrich contacts or generate copy. Saves Clay credits.

### Contact Finding Layer (Columns 15-19)

Only runs for accounts passing the gate.

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 15 | `contact_first_name` | Enrichment | Waterfall: Hunter → Apollo → RocketReach → Prospeo | Find target persona at company |
| 16 | `contact_last_name` | Enrichment | Same waterfall | |
| 17 | `contact_title` | Enrichment | Same waterfall | Must match target persona titles |
| 18 | `contact_email` | Enrichment | Same waterfall | |
| 19 | `contact_email_verified` | Enrichment | Email verification (ZeroBounce/NeverBounce) | valid / invalid / risky |

**Person search criteria:** Use target persona titles from icp-mapping.md. Search for the most senior matching title first.

### Copy Generation Layer (Columns 20-25)

Only runs for verified contacts at qualified accounts.

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 20 | `copy_opener_company` | AI Prompt | See clay-prompts.md #4 | Company-based observational opener |
| 21 | `copy_opener_trigger` | AI Prompt | See clay-prompts.md #5 | Intent-aware opener (no surveillance) |
| 22 | `copy_body` | AI Prompt | See clay-prompts.md #6 | Email body |
| 23 | `copy_linkedin` | AI Prompt | See clay-prompts.md #7 | LinkedIn connection request |
| 24 | `copy_cta` | AI Prompt | See clay-prompts.md #8 | Stage-appropriate CTA |
| 25 | `copy_email_subject` | AI Prompt | Subject line based on pain inference | ≤45 chars |

### Assembly & Export (Columns 26-28)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 26 | `export_email_1` | Formula | Concatenate: `copy_opener_trigger` + `copy_body` + `copy_cta` | Full Email 1 |
| 27 | `export_status` | Formula | "READY" if all copy columns populated + email verified, else "INCOMPLETE" | |
| 28 | `export_sequencer_pushed` | Checkbox | Manual or automation | Track if pushed to sequencer |

## Credit Optimization Notes

- Run company enrichment (columns 5-8) BEFORE trigger analysis — if company doesn't match basic SAM criteria, stop early
- The ICP gate at column 14 prevents wasting credits on copy generation for poor fits
- Contact waterfall tries cheapest provider first (Hunter) before paid options
- Email verification prevents wasting sequencer sends on bad addresses
- Batch test: run 20-30 records through the full sequence before activating automation

## Automation Trigger

When intent tool sends webhook to Clay:
1. New row created automatically
2. Columns 1-14 run in sequence
3. If passes gate → columns 15-28 run
4. If `export_status` = "READY" → push to sequencer (manual review recommended for first 2 weeks)
