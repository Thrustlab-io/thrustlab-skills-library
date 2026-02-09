# Enrichment Sequence — Growth / Hiring Surge Trigger

## Data Flow

```
LinkedIn/Clay Headcount Signal → Clay Import
  → Column 1-5: Growth Data Import
    → Column 6-7: Growth Classification + Gate
      → Column 8-11: Company Enrichment + ICP Gate
        → Column 12-13: Scaling Challenge Research
          → Column 14-20: Hook Type + Copy Generation
            → Column 21-24: Assembly + Export
```

## Trigger-Specific Columns

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `trigger_headcount_current` | Import | LinkedIn / Clay signal | Current employee count |
| 2 | `trigger_headcount_6mo_ago` | Import | LinkedIn / Clay signal | 6-month-ago count |
| 3 | `trigger_growth_rate_pct` | Formula | `((current - 6mo) / 6mo) * 100` | Growth percentage |
| 4 | `trigger_growth_departments` | Import/Claygent | Which departments growing | LinkedIn analysis |
| 5 | `trigger_growth_type` | AI Prompt | See clay-prompts.md #1 | HEADCOUNT_SURGE / DEPARTMENT_BUILD / NEW_OFFICE / NEW_MARKET / SLOW_GROWTH |
| 6 | `research_scaling_challenges` | AI Prompt | See clay-prompts.md #2 | Inferred scaling pains |

## Gates

- `trigger_growth_type` = "SLOW_GROWTH" → stop (signal too weak)
- `trigger_growth_rate_pct` < 20 → deprioritize
- Standard ICP gate

## Standard Layers

Company enrichment → ICP gate → Contact waterfall → Hook type + Copy → Meta tagging → Assembly.
See `workflow-website-trigger/references/enrichment-sequence.md` for standard layer details.
