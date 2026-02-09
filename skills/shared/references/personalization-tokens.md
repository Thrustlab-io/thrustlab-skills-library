# Personalization Tokens — Variable Standards

Standard format for all Clay column variables across every workflow skill.

## Variable Format

```
{{category_specific_field|"fallback_value"}}
```

## Category Prefixes

| Prefix | Scope | Examples |
|---|---|---|
| `company_` | Company-level data | `company_name`, `company_industry`, `company_size`, `company_hq`, `company_tech_stack` |
| `contact_` | Person-level data | `contact_first_name`, `contact_last_name`, `contact_title`, `contact_linkedin`, `contact_email` |
| `trigger_` | Trigger-specific data | `trigger_type`, `trigger_date`, `trigger_detail`, `trigger_source` |
| `research_` | AI-researched data | `research_pain_point`, `research_observation`, `research_company_summary` |
| `copy_` | AI-generated copy | `copy_opener_company`, `copy_opener_trigger`, `copy_body`, `copy_cta`, `copy_linkedin` |
| `score_` | Scoring/qualification | `score_icp_fit`, `score_intent`, `score_engagement`, `score_composite`, `score_fit`, `score_relationship` |
| `signal_` | Signal stacking data | `signal_composite_score`, `signal_count`, `signal_strongest`, `signal_stack_id`, `signal_list` |
| `hook_` | Hook type system | `hook_type`, `hook_version`, `hook_timeline_data`, `hook_numbers_data`, `hook_proof_data` |
| `meta_` | Workflow metadata | `meta_workflow_name`, `meta_cadence_step`, `meta_send_date`, `meta_prompt_version`, `meta_hook_type_used` |
| `champion_` | Champion tracking | `champion_previous_company`, `champion_previous_title`, `champion_days_in_role`, `champion_relationship_type` |
| `competitor_` | Competitor customer | `competitor_product_used`, `competitor_customer_name`, `competitor_win_angle` |
| `darkfunnel_` | Dark funnel signals | `darkfunnel_source`, `darkfunnel_engagement_type`, `darkfunnel_content_topic`, `darkfunnel_recency` |

## Fallback Rules

### The #1 Rule: No Generic Fallbacks

Fallback values must still feel relevant to the client's domain and the prospect's industry. A fallback should read as if the SDR simply didn't find a hyper-specific detail but still knows the space.

```
# ❌ BAD — fully generic, could be anyone
{{research_observation|"your recent growth"}}
{{trigger_detail|"recent changes at your company"}}
{{research_pain_point|"common challenges in your industry"}}

# ✅ GOOD — still references client domain + prospect context
{{research_observation|"the shift toward {client_industry_trend} in {{company_industry}}"}}
{{trigger_detail|"the momentum in {{company_industry}} right now"}}
{{research_pain_point|"the pressure to {client_core_pain_verb} without adding headcount"}}
```

### Fallback Construction Pattern

Every fallback should follow: `[industry/role reference] + [client-relevant challenge/trend]`

The `{client_*}` variables in fallbacks are resolved at skill execution time from the client profile — they are NOT Clay variables. They are baked into the prompt when the skill generates the Clay configuration.

### When to Use "SKIP" Instead of Fallback

Output `SKIP` (not a fallback) when:
- The entire personalization angle depends on data that's missing (e.g., trigger-based opener but no trigger data)
- Using a fallback would make the message feel obviously templated
- The missing data is the core differentiator of the message

The sequencer should be configured to skip rows with `SKIP` values in critical fields.

## Clay Column Naming Convention

Clay column names should match the variable prefix system:

```
Column name in Clay          → Variable in prompts/copy
─────────────────────────────────────────────────────
Company Name                 → {{company_name}}
Company Industry             → {{company_industry}}
Contact First Name           → {{contact_first_name}}
Research: Pain Point         → {{research_pain_point}}
Research: Company Observation→ {{research_observation}}
Copy: Opener (Company)       → {{copy_opener_company}}
Copy: Opener (Trigger)       → {{copy_opener_trigger}}
Copy: Email Body             → {{copy_body}}
Copy: CTA                    → {{copy_cta}}
Copy: LinkedIn Message       → {{copy_linkedin}}
Trigger: Type                → {{trigger_type}}
Trigger: Detail              → {{trigger_detail}}
Score: ICP Fit               → {{score_icp_fit}}
```

## Variable Resolution Order in Clay

Prompts that generate copy should explicitly reference upstream columns. This creates a compounding personalization chain:

```
Column 1: Enrichment         → company_size, company_industry, company_tech_stack
Column 2: Trigger Detection  → trigger_type, trigger_detail, trigger_date
Column 3: Claygent Research  → research_observation, research_pain_point
Column 4: ICP Scoring        → score_icp_fit (qualification gate — only proceed if fit)
Column 5: Hook Type          → hook_type (timeline/numbers/social_proof/hypothesis — see hook-types-guide.md)
Column 6: Signal Stacking    → signal_composite_score, signal_count (if aggregation table in use)
Column 7: AI Opener (Co.)    → copy_opener_company (reads columns 1, 3, 5)
Column 8: AI Opener (Trig.)  → copy_opener_trigger (reads columns 1-5)
Column 9: AI Email Body      → copy_body (reads columns 1-8)
Column 10: AI LinkedIn       → copy_linkedin (reads columns 1-3, 5)
Column 11: AI CTA            → copy_cta (reads strategy CTA rules + columns 1-6)
Column 12: Final Assembly    → Full email = best opener + copy_body + copy_cta
Column 13: Meta Tagging      → meta_prompt_version, meta_hook_type_used, meta_workflow_name
```

Each AI column prompt must explicitly state which upstream columns it reads. Never assume implicit context.

## Clay Column Naming Convention — Extended

```
Column name in Clay          → Variable in prompts/copy
─────────────────────────────────────────────────────
Company Name                 → {{company_name}}
Company Industry             → {{company_industry}}
Contact First Name           → {{contact_first_name}}
Research: Pain Point         → {{research_pain_point}}
Research: Company Observation→ {{research_observation}}
Hook: Type                   → {{hook_type}}
Hook: Timeline Data          → {{hook_timeline_data}}
Hook: Numbers Data           → {{hook_numbers_data}}
Hook: Proof Data             → {{hook_proof_data}}
Signal: Composite Score      → {{signal_composite_score}}
Signal: Active List          → {{signal_list}}
Signal: Count                → {{signal_count}}
Signal: Strongest            → {{signal_strongest}}
Copy: Opener (Company)       → {{copy_opener_company}}
Copy: Opener (Trigger)       → {{copy_opener_trigger}}
Copy: Email Body             → {{copy_body}}
Copy: CTA                    → {{copy_cta}}
Copy: LinkedIn Message       → {{copy_linkedin}}
Trigger: Type                → {{trigger_type}}
Trigger: Detail              → {{trigger_detail}}
Score: ICP Fit               → {{score_icp_fit}}
Meta: Prompt Version         → {{meta_prompt_version}}
Meta: Hook Type Used         → {{meta_hook_type_used}}
```
