# Enrichment Sequence — Funding Trigger

## Data Flow

```
Crunchbase/PitchBook/News Alert → Clay Webhook/Import
  → Column 1-5: Funding Data Import
    → Column 6-9: Company Enrichment
      → Column 10-11: ICP Gate + Funding Relevance
        → Column 12-14: Growth Research
          → Column 15-19: Contact Finding
            → Column 20-26: Hook Type + Copy Generation
              → Column 27-30: Assembly + Export
```

## Trigger-Specific Columns

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `trigger_funding_amount` | Import | Crunchbase/news | Dollar amount raised |
| 2 | `trigger_funding_round` | Import | Crunchbase/news | Seed / A / B / C / D+ |
| 3 | `trigger_funding_date` | Import | Crunchbase/news | Announcement date |
| 4 | `trigger_funding_freshness_days` | Formula | `(TODAY() - funding_date)` | Recency — >60 days = stale |
| 5 | `trigger_funding_relevance` | AI Prompt | See clay-prompts.md | How funding connects to client's domain |
| 6 | `trigger_growth_priorities` | AI Prompt | Infer from round + industry | Where they'll invest |
| 7 | `research_hiring_surge` | Claygent | Check for post-funding hiring | Which departments growing |

## Gates

- `trigger_funding_freshness_days` > 60 → deprioritize (funding news is old)
- `trigger_funding_relevance` = "LOW" → route to general outbound
- Standard ICP gate on company fit

## Standard Layers

Company enrichment → ICP gate → Contact waterfall → Hook type + Copy → Meta tagging → Assembly.
See `workflow-website-trigger/references/enrichment-sequence.md` for standard layer details (columns 6-9 company enrichment, 15-19 contact finding, 26-28 assembly pattern).

## Credit Optimization

- Funding data comes from Crunchbase/PitchBook — no enrichment credits
- Freshness check (column 4) before ANY enrichment — stale funding saves credits
- Relevance check (column 5) before growth research — irrelevant funding stops early
