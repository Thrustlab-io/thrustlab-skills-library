# Copy Templates & Enrichment — Funding Trigger

## Cadence

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Growth priorities + scaling challenge → solution bridge |
| 2 | LinkedIn | Day 1 | Growth-focused connection |
| 3 | Email | Day 5 | Stage-specific value: scaling framework or benchmark |
| 4 | Email | Day 10 | Social proof: similar-stage company that scaled with client |
| 5 | Email | Day 16 | Breakup: "before the team doubles" urgency |

## Email 1 — Growth Priority (Day 0)
**Subject:** `Scaling {function} at {{company_name}}`
- Opener branches by `hook_type` (see `shared/references/hook-types-guide.md`):
  - **Timeline**: "Post-[round] teams in {{company_industry}} see [result] within [timeframe] when they prioritize [area]."
  - **Numbers**: "[Industry] teams investing in [area] see [X% return] — that math gets better at {{company_name}}'s stage."
  - **Social proof**: "[Named company] was at a similar stage after their [round] — they [result]."
  - **Hypothesis** (fallback): "Scaling [function] after a [round] is one of the hardest transitions in {{company_industry}}."
- Body: connect growth plan → client outcome
- CTA: calibrated by ICP tier + signal composite score + growth stage

## Email 2 — Scaling Framework (Day 5)
**Subject:** `{function} at {{trigger_funding_round}}-stage {{company_industry}} companies`
- Company-based opener
- Value: share benchmark, framework, or data point relevant to their growth stage
- No CTA for meeting — pure value

## Email 3 — Social Proof (Day 10)
**Subject:** `How {{peer_company}} scaled after their {{similar_round}}`
- Peer story from similar funding stage + industry
- Specific metric/outcome
- Slightly more direct CTA

## Email 4 — Breakup (Day 16)
**Subject:** `{different_topic} at {{company_name}}`
- Different pain point from icp-mapping.md
- Urgency: growth-phase problems compound — "before things get hectic"
- Permission to not respond

---

## Enrichment Sequence

```
Signal (Crunchbase/Clay Signal/News) → Clay Import
  → Columns 1-5: Funding data (amount, round, date, investors, URL)
    → Columns 6-9: Company enrichment
      → Columns 10-13: Funding analysis (growth priorities, relevance, hiring surge, company snapshot)
        → Column 14: ICP Gate (score + relevance + freshness check)
          → Columns 15-18: Contact finding (target persona at company)
            → Columns 19-25: Copy generation
              → Columns 26-28: Assembly + export
```

### Gates
- `trigger_funding_relevance` starts with "LOW" → stop
- `trigger_funding_freshness_days` > 30 → stop
- `score_icp_fit` = "DQ" → stop

### Contact Finding Note
Post-funding companies are often chaotic. C-suite may be harder to reach. Consider targeting VP-level (one below C-suite) as they're the ones executing the growth plan.
