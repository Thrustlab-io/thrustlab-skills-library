# Enrichment Sequence — Tech Stack Change Trigger

## Data Flow

```
BuiltWith Alert / Clay Tech Signal → Clay Import
  → Column 1-4: Tech Change Import
    → Column 5-7: Change Classification + Gate
      → Column 8-11: Company Enrichment + ICP Gate
        → Column 12-13: Stack Context Research
          → Column 14-20: Hook Type + Copy Generation
            → Column 21-24: Assembly + Export
```

## Trigger-Specific Columns

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `trigger_tech_added` | Import | BuiltWith / Clay signal | Technology added |
| 2 | `trigger_tech_removed` | Import | BuiltWith / Clay signal | Technology removed (if any) |
| 3 | `trigger_tech_change_date` | Import | BuiltWith | When detected |
| 4 | `trigger_tech_change_type` | AI Prompt | See clay-prompts.md #1 | COMPLEMENTARY_ADD / COMPETITOR_REMOVED / COMPETITOR_ADDED / STACK_SHIFT / IRRELEVANT |
| 5 | `trigger_stack_context` | Claygent | See clay-prompts.md #2 | Broader stack research |

## Gates

- `trigger_tech_change_type` = "COMPETITOR_ADDED" → route to long-term nurture (6-12 month)
- `trigger_tech_change_type` = "IRRELEVANT" → stop
- Standard ICP gate on company fit
- Change detected >30 days ago → deprioritize

## Standard Layers

Company enrichment → ICP gate → Contact waterfall → Hook type + Copy → Meta tagging → Assembly.
See `workflow-website-trigger/references/enrichment-sequence.md` for standard layer details.

## Credit Optimization

- Classification (column 4) FIRST — filters COMPETITOR_ADDED and IRRELEVANT before spending
- BuiltWith is the primary data source — no extra enrichment credits for trigger data
- Stack context research only for confirmed COMPLEMENTARY/REMOVED/SHIFT
