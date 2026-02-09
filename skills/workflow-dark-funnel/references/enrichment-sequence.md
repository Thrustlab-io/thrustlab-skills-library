# Enrichment Sequence — Dark Funnel Trigger

## Data Flow

```
Signal Tool (Dealfront/RB2B/Teamfluence/Trigify/Common Room) → Clay Webhook
  → Column 1-5: Signal Import + Classification
    → Column 6-9: Company Enrichment
      → Column 10-11: Intensity Gate + ICP Gate
        → Column 12-15: Intent Interpretation + Research
          → Column 16-22: Hook Type + Copy Generation
            → Column 23-27: Assembly + Signal Stacking + Export
```

## Column-by-Column Configuration

### Signal Import Layer (Columns 1-5)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `company_name` | Import | Signal tool webhook | Company identified |
| 2 | `company_website` | Import | Signal tool webhook | Company domain |
| 3 | `darkfunnel_source` | Import | Webhook metadata | "dealfront" / "rb2b" / "teamfluence" / "trigify" / "common_room" |
| 4 | `darkfunnel_engagement_type` | Import | Signal tool webhook | What they did: "website_visit" / "linkedin_like" / "linkedin_comment" / "community_post" / "webinar_signup" |
| 5 | `darkfunnel_raw_detail` | Import | Signal tool webhook | Page URL, post URL, comment text, etc. |

**Source-specific webhook fields:**
- **Dealfront/RB2B:** `pages_visited`, `visit_duration`, `visit_timestamp`, `contact_email` (RB2B person-level only)
- **Teamfluence:** `linkedin_post_url`, `engagement_type`, `engagement_timestamp`
- **Trigify:** `trigger_event`, `linkedin_profile_url`, `event_timestamp`
- **Common Room:** `platform`, `activity_type`, `content_url`, `activity_timestamp`

### Signal Classification (Column 6)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 6 | `darkfunnel_intensity` | AI Prompt | See clay-prompts.md Layer 1 | "high" / "medium" / "low" |

**INTENSITY GATE:** If `darkfunnel_intensity` = "low" → route to general outbound with topic tag. Do NOT trigger dedicated dark funnel cadence. Only "high" and "medium" continue.

### Company Enrichment (Columns 7-10)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 7 | `company_industry` | Enrichment | Clearbit/Apollo | Industry classification |
| 8 | `company_size` | Enrichment | Clearbit/Apollo | Employee count |
| 9 | `darkfunnel_content_topic` | AI Prompt | Map engagement to pain/topic | What the signal means thematically |
| 10 | `darkfunnel_recency` | Formula | `(NOW() - signal_timestamp)` in hours | Freshness check |

### ICP Gate (Columns 11-12)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 11 | `score_icp_fit` | Formula/AI | Standard ICP scoring | A+ / A / B / DQ |
| 12 | `score_icp_tier` | Formula | ICP fit score | |

**GATE:** DQ if ICP doesn't fit. Also check `darkfunnel_recency`:
- Website visits: act within 48 hours, discard after 7 days
- LinkedIn engagement: act within 72 hours, discard after 10 days
- Community signals: act within 7 days, discard after 21 days

### Intent Interpretation + Research (Columns 13-16)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 13 | `research_engagement_context` | AI Prompt | See clay-prompts.md | What this engagement MEANS |
| 14 | `research_inferred_pain` | AI Prompt | See clay-prompts.md | What pain they're likely experiencing |
| 15 | `research_company_context` | Claygent | See clay-prompts.md | What's happening at their company |
| 16 | `content_match` | Lookup/AI | Client's content library | Best content to share, matched to engagement topic |

### Contact Finding (Columns 17-20)

Only for company-level signals (Dealfront). RB2B provides person-level data directly.

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 17 | `contact_first_name` | Import/Enrichment | RB2B direct OR waterfall: Hunter → Apollo → RocketReach | |
| 18 | `contact_title` | Import/Enrichment | Same | Must match target persona |
| 19 | `contact_email` | Import/Enrichment | Same | |
| 20 | `contact_email_verified` | Enrichment | ZeroBounce/NeverBounce | valid / invalid / risky |

### Hook Type + Copy Generation (Columns 21-26)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 21 | `hook_type` | AI Prompt | See clay-prompts.md | hypothesis tends to be default for dark funnel |
| 22 | `copy_opener_trigger` | AI Prompt | THE GOLDEN RULE: must feel like coincidence | Value-forward opener |
| 23 | `copy_body` | AI Prompt | Content-forward, not surveillance | |
| 24 | `copy_cta` | AI Prompt | Signal-stacking aware | |
| 25 | `copy_linkedin` | AI Prompt | Content share, thought leadership | |
| 26 | `copy_email_subject` | AI Prompt | Topic-focused | ≤45 chars |

### Assembly + Signal Stacking + Export (Columns 27-31)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 27 | `signal_composite_score` | Formula/Webhook | Push to Signal Aggregation Table | Points per signal-stacking-guide.md |
| 28 | `export_email_1` | Formula | Concatenate: opener + body + CTA | |
| 29 | `export_status` | Formula | "READY" if intensity med/high + ICP pass + copy + email verified | |
| 30 | `meta_hook_type_used` | Formula | = `hook_type` | |
| 31 | `meta_prompt_version` | Static | "v1.0" | |

## Credit Optimization

- Signal classification (column 6) is the FIRST AI call — filter "low" intensity before any enrichment spend
- ICP gate + recency check before research — save Claygent credits
- For RB2B (person-level): skip contact finding waterfall — person data included
- For Dealfront (company-level): contact finding required — use the waterfall
- Content match (column 16) is a lookup/formula, not an AI call
- Batch: consolidate multiple signals from the same company (e.g., 3 page visits in 24h = 1 enrichment row with highest intensity)

## Automation Triggers

Each signal source sends webhooks to Clay:
1. **Dealfront webhook** → new row with company-level data
2. **RB2B webhook** → new row with person-level data (skip contact finding)
3. **Teamfluence export** → scheduled Clay import (daily/weekly)
4. **Trigify webhook** → new row with LinkedIn trigger data
5. **Common Room webhook** → new row with community signal data

**Deduplication:** If same company appears from multiple sources within 48 hours, merge into one row with the highest intensity signal. Track all sources in `darkfunnel_source` as comma-separated.
