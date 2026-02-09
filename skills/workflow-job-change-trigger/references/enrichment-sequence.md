# Enrichment Sequence — Job Role Change Trigger

## Data Flow

```
Signal Source (UserGems/Clay Signal/SN Alert) → Clay Import
  → Column 1-5: Import + Transition Data
    → Column 6-9: Company Enrichment (NEW company)
      → Column 10-12: Transition Research
        → Column 13: ICP Scoring Gate
          → Column 14-18: Contact Verification
            → Column 19-25: Copy Generation
              → Column 26-28: Assembly + Export
```

## Column-by-Column Configuration

### Import Layer (Columns 1-5)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `contact_first_name` | Import | Signal source | Person who changed jobs |
| 2 | `contact_last_name` | Import | Signal source | |
| 3 | `contact_linkedin_url` | Import | Signal source | LinkedIn profile URL |
| 4 | `trigger_previous_company` | Import | Signal source / enrichment | Company they left |
| 5 | `trigger_previous_title` | Import | Signal source / enrichment | Role they had before |

### New Company Enrichment (Columns 6-9)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 6 | `company_name` | Enrichment | LinkedIn profile → current company | Their new employer |
| 7 | `company_website` | Enrichment | Clearbit/Apollo by company name | |
| 8 | `company_industry` | Enrichment | Clearbit/Apollo | Industry of NEW company |
| 9 | `company_size` | Enrichment | Clearbit/Apollo | |

### Current Role Enrichment (Column 10)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 10 | `contact_title` | Enrichment | LinkedIn enrichment | Current title at new company |

### Trigger Analysis (Columns 11-13)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 11 | `trigger_days_in_role` | Formula | Calculate from job start date | `(NOW() - start_date) / 86400` |
| 12 | `trigger_role_transition_type` | AI Prompt | See clay-prompts.md #1 | PROMOTION / LATERAL / CAREER_PIVOT / etc. |
| 13 | `trigger_new_role_context` | Claygent | See clay-prompts.md #2 | Research on new company + role |

### ICP Scoring Gate (Column 14)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 14 | `score_icp_fit` | Formula/AI | ICP scoring formula from icp-mapping.md | Score based on NEW company |

**GATE:** Stop if `score_icp_fit` = "DQ" OR `trigger_days_in_role` > 90 (trigger is stale). Saves credits.

### Research Layer (Columns 15-16)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 15 | `research_previous_company_relevance` | AI Prompt | See clay-prompts.md #3 | Did old company use similar solution? |
| 16 | `research_company_snapshot` | Claygent | Standard company research | Recent news about new company |

### Contact Verification (Columns 17-18)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 17 | `contact_email` | Enrichment | Waterfall: Hunter → Apollo → RocketReach → Prospeo | Email at NEW company |
| 18 | `contact_email_verified` | Enrichment | ZeroBounce/NeverBounce | Must verify — new role = new email |

**Note:** Job changers often don't have updated emails in databases yet. The waterfall is especially important here. If no email found, flag for LinkedIn-only outreach.

### Copy Generation (Columns 19-25)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 19 | `copy_opener_company` | AI Prompt | See clay-prompts.md #4 | |
| 20 | `copy_opener_trigger` | AI Prompt | See clay-prompts.md #5 | Transition-aware opener |
| 21 | `copy_body` | AI Prompt | See clay-prompts.md #6 | |
| 22 | `copy_linkedin` | AI Prompt | See clay-prompts.md #7 | |
| 23 | `copy_cta` | AI Prompt | See clay-prompts.md #8 | Low-commitment for new-in-role |
| 24 | `copy_email_subject` | AI Prompt | Subject line | |
| 25 | `copy_email_2_body` | AI Prompt | "First 90 days" value email | |

### Assembly & Export (Columns 26-28)

| # | Column Name | Type | Source | Notes |
|---|---|---|---|---|
| 26 | `export_email_1` | Formula | Concatenate opener + body + CTA | |
| 27 | `export_status` | Formula | READY / INCOMPLETE / LINKEDIN_ONLY | |
| 28 | `export_sequencer_pushed` | Checkbox | | |

## Credit Optimization

- ICP score the NEW company before researching the transition — most disqualifications happen here
- `trigger_days_in_role` > 90 auto-disqualifies — don't waste credits on stale job changes
- If no email found after waterfall, mark as LINKEDIN_ONLY instead of wasting more credits
- Previous company relevance is a bonus enrichment — skip if Clay credits are tight
