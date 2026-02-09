# Enrichment Sequence — Content Engagement Trigger

## Data Flow

```
HubSpot/Marketo/LinkedIn Analytics → Clay Webhook/Import
  → Column 1-5: Engagement Data Import
    → Column 6-7: Engagement Depth + Intensity Gate
      → Column 8-11: Company Enrichment + ICP Gate
        → Column 12-14: Topic Mapping + Content Match
          → Column 15-18: Contact Verification
            → Column 19-25: Hook Type + Copy Generation
              → Column 26-29: Assembly + Export
```

## Trigger-Specific Columns

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `trigger_content_type` | Import | Marketing tool webhook | DOWNLOAD / WEBINAR / LINKEDIN_ENGAGEMENT / MULTI_PAGE_VISIT |
| 2 | `trigger_content_title` | Import | Marketing tool webhook | Specific content they engaged with |
| 3 | `trigger_engagement_timestamp` | Import | Marketing tool webhook | When engagement occurred |
| 4 | `trigger_engagement_recency_hours` | Formula | `(NOW() - timestamp) / 3600` | Freshness |
| 5 | `trigger_content_topic` | AI Prompt | See clay-prompts.md #1 | Map content to pain category |
| 6 | `trigger_engagement_depth` | AI Prompt | See clay-prompts.md #2 | DEEP / MODERATE / LIGHT |
| 7 | `content_match` | Lookup | Client content library | Next content to share, matched to topic |

## Gates

- `trigger_engagement_depth` = "LIGHT" → lower priority cadence (3 touches) or route to general outbound
- `trigger_engagement_recency_hours` > 168 (7 days) → reference topic generally, not specific asset
- Standard ICP gate
- Contact often already known (from form fill or marketing tool) — skip waterfall if email exists

## Unique Characteristics

This is the WARMEST trigger — contact data often comes pre-populated from the marketing tool. Focus enrichment budget on company context and topic matching rather than contact finding.

## Standard Layers

Company enrichment → ICP gate → Contact waterfall (often skippable) → Hook type + Copy → Meta tagging → Assembly.
See `workflow-website-trigger/references/enrichment-sequence.md` for standard layer details.
