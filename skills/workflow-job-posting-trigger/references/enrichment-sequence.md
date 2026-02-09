# Enrichment Sequence — Job Posting Trigger

## Data Flow

```
Signal (Clay Signal/Otta/LinkedIn Jobs) → Clay Import
  → Columns 1-4: Import (job title, URL, description, posting date)
    → Columns 5-8: Company Enrichment
      → Columns 9-12: JD Analysis (pain extraction, tools, intent score, team context)
        → Column 13: ICP Gate
          → Columns 14-18: Contact Finding (hiring manager, NOT the role being hired)
            → Columns 19-25: Copy Generation
              → Columns 26-28: Assembly + Export
```

## Key Differences from Other Trigger Sequences

### Import Layer
| # | Column Name | Type | Notes |
|---|---|---|---|
| 1 | `trigger_job_title_posted` | Import | The role being hired |
| 2 | `trigger_job_url` | Import | Link to the posting |
| 3 | `trigger_job_description_raw` | Import/Claygent | Full JD text — if not available from import, Claygent fetches from URL |
| 4 | `trigger_posting_date` | Import | When posted |

### JD Analysis Layer (unique to this trigger)
| # | Column Name | Type | Notes |
|---|---|---|---|
| 9 | `trigger_job_pain_indicators` | AI Prompt | 2-3 pain points from JD |
| 10 | `trigger_job_tools_mentioned` | AI Prompt | Tech stack from JD |
| 11 | `trigger_hiring_intent_score` | AI Prompt | HIGH/MEDIUM/LOW/NO_MATCH |
| 12 | `research_team_context` | Claygent | Research the hiring team |

### Contact Finding — CRITICAL DIFFERENCE
The target contact is NOT the role being hired. It's the **hiring manager** — typically one level above the posted role.

| Posted Role | Target Contact |
|---|---|
| Operations Analyst | VP Operations / Director Operations |
| Data Engineer | Head of Data / VP Engineering |
| Marketing Manager | CMO / VP Marketing |
| Sales Development Rep | VP Sales / Head of Revenue |

Use persona titles from icp-mapping.md to determine the correct target. Search for the most senior matching title at the company.

### Gates
- `trigger_hiring_intent_score` = "NO_MATCH" → stop
- `trigger_posting_age_days` > 30 → stop (stale)
- `score_icp_fit` = "DQ" → stop

### Credit Optimization
- Run intent scoring BEFORE company enrichment — NO_MATCH postings shouldn't cost enrichment credits
- JD text: import directly if available from signal source, only use Claygent to fetch if missing
- Team context research is valuable but optional if credits are tight
