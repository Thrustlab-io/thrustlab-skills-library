# Enrichment Sequence — General Outbound

## Data Flow

```
ICP List Import (Sales Navigator / Apollo / Clay)
  → Layer 1: Foundation (Columns 1-8) — ICP qualification + company enrichment
    → Layer 2a: Research Variables (Columns 9-12) — personal + company overlap
      → Layer 2b: Deep Research (Columns 13-18) — company workflow, tech, insights
        → Layer 2c: Content Matching (Column 19-20) — blog/resource lookup
          → Layer 3: Copy Generation (Columns 21-28) — hook type + all copy touches
            → Assembly + Export (Columns 29-33)
```

## Column-by-Column Configuration

### Layer 1: Foundation (Columns 1-8)

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 1 | `company_name` | Import | ICP list | |
| 2 | `company_website` | Import | ICP list | |
| 3 | `contact_first_name` | Import | ICP list | |
| 4 | `contact_title` | Import | ICP list | |
| 5 | `contact_email` | Import/Enrichment | List or waterfall | |
| 6 | `company_industry` | Enrichment | Clearbit/Apollo | |
| 7 | `company_size` | Enrichment | Clearbit/Apollo | |
| 8 | `icp_angle` | AI Prompt | See clay-prompts.md #1 | Which ICP angle fits this prospect |

**GATE:** `icp_angle` = "NO_FIT" → stop. Don't spend credits researching non-fits.

### Layer 2a: Research Variables (Columns 9-12)

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 9 | `overlap_personal` | AI Prompt | See clay-prompts.md #2 | Role-specific pain hypothesis |
| 10 | `overlap_company` | AI Prompt | See clay-prompts.md #3 | Company-level research overlap |
| 11 | `contact_email_verified` | Enrichment | ZeroBounce/NeverBounce | valid / invalid / risky |
| 12 | `language` | Formula/AI | Detect from profile/location | Output language for copy |

**GATE:** `contact_email_verified` = "invalid" → stop. `overlap_personal` AND `overlap_company` both = "SKIP" → stop.

### Layer 2b: Deep Research (Columns 13-18)

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 13 | `research_company_deep` | Claygent | See clay-prompts.md #4 | Thorough company web research |
| 14 | `company_workflow` | AI Prompt | See clay-prompts.md #5 | Specific relevant workflow |
| 15 | `tech_signals` | AI Prompt | See clay-prompts.md #6 | Technology signals |
| 16 | `company_insight` | AI Prompt | See clay-prompts.md #7 | Hypothesis question |
| 17 | `scale_trigger` | AI Prompt | See clay-prompts.md #8 | Scaling inflection point |
| 18 | `research_recent_activity` | Claygent | See clay-prompts.md #10 | Personal LinkedIn activity |

### Layer 2c: Content Matching (Columns 19-20)

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 19 | `blog_link` | Lookup | Content mapping table by `icp_angle` | Primary content URL |
| 20 | `blog_link_2` | Lookup | Content mapping table by `icp_angle` | Secondary content URL |

No AI credits — this is a table lookup.

### Layer 3: Copy Generation (Columns 21-28)

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 21 | `hook_type` | AI Prompt | See clay-prompts.md #10b | timeline (default) / numbers / social_proof / hypothesis |
| 22 | `copy_linkedin_invite` | AI Prompt | See clay-prompts.md #11 | LinkedIn connection request |
| 23 | `copy_hypothesis_touch` | AI Prompt | See clay-prompts.md #12 | Hook-typed email (Touch 2) |
| 24 | `copy_deep_research_touch` | AI Prompt | See clay-prompts.md #13 | Deep research email (Touch 3) |
| 25 | `copy_soft_close` | AI Prompt | See clay-prompts.md #14 | Final touch |
| 26 | `copy_email_subject_1` | AI Prompt | Subject for Touch 2 | ≤45 chars |
| 27 | `copy_email_subject_2` | AI Prompt | Subject for Touch 3 | ≤45 chars |
| 28 | `score_icp_tier` | Formula | ICP scoring formula | A+ / A / B |

### Assembly + Export (Columns 29-33)

| # | Column | Type | Source | Notes |
|---|---|---|---|---|
| 29 | `signal_composite_score` | Formula/Webhook | From Signal Aggregation Table (if exists) | May have signals from other triggers |
| 30 | `export_status` | Formula | "READY" if angle ≠ NO_FIT + overlaps exist + email verified + copy populated | |
| 31 | `export_sequencer_pushed` | Checkbox | Manual or automation | |
| 32 | `meta_hook_type_used` | Formula | = `hook_type` | A/B tracking |
| 33 | `meta_prompt_version` | Static | "v1.0" | |

## Credit Optimization

- ICP angle classification (column 8) runs BEFORE deep research — NO_FIT saves all downstream credits
- Email verification (column 11) runs early — invalid emails stop before copy generation
- Layer 2b is the most expensive layer (2 Claygent calls + 4 AI calls per row) — only runs on qualified, verified contacts
- Content matching (Layer 2c) uses lookups, not AI
- For large batches: run Layer 1 on full list, filter, then run Layer 2+ only on qualifiers
- Use Clay's Sandbox Mode for first 10 rows before scaling

## Sequencer Mapping

General outbound uses a 5-touch cadence:
1. LinkedIn invite (`copy_linkedin_invite`) — Day 0
2. Hook-typed email (`copy_hypothesis_touch`) — Day 2
3. LinkedIn message (if connected) — Day 4
4. Deep research email (`copy_deep_research_touch`) — Day 7
5. Soft close (`copy_soft_close`) — Day 14
