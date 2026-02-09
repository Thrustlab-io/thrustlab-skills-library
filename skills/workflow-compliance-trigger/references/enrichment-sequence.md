# Enrichment Sequence — Compliance / Regulatory Trigger

## Data Flow

```
Manual/News Import (regulation identified) → Clay Import (ICP list filtered by affected industry)
  → Column 1-4: Regulation Data Import
    → Column 5-7: Company Enrichment + ICP Gate
      → Column 8-9: Regulation Impact Assessment + Gate
        → Column 10-11: Compliance Readiness Research
          → Column 12-18: Hook Type + Copy Generation
            → Column 19-22: Assembly + Export
```

## Trigger-Specific Columns

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `trigger_regulation_name` | Manual/Import | Strategy document | Specific regulation name |
| 2 | `trigger_regulation_deadline` | Manual/Import | Official source | Compliance deadline date |
| 3 | `trigger_days_until_deadline` | Formula | `(deadline - TODAY())` | Urgency countdown |
| 4 | `trigger_regulation_impact` | AI Prompt | See clay-prompts.md #1 | How regulation affects THIS company |
| 5 | `research_compliance_readiness` | Claygent | See clay-prompts.md #2 | Public signs of compliance prep |
| 6 | `trigger_penalty_risk` | Claygent | Web research | Non-compliance penalties |

## Gates

- `trigger_regulation_impact` = "LOW_RELEVANCE" → stop
- `trigger_regulation_impact` = "VERIFY_MANUALLY" → manual review queue
- Standard ICP gate
- **CRITICAL:** All regulation facts (name, deadline, requirements) must be VERIFIED before sending

## Unique Characteristics

This trigger is often BATCH-based, not real-time:
1. New regulation identified → filter ICP list by affected industries
2. Import filtered list into Clay → enrich + assess impact
3. Copy cadence timing tied to deadline proximity, not trigger freshness

## Standard Layers

Company enrichment → ICP gate → Contact waterfall → Hook type + Copy → Meta tagging → Assembly.
See `workflow-website-trigger/references/enrichment-sequence.md` for standard layer details.
