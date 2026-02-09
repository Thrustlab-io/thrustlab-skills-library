# Copy Templates & Enrichment — General Outbound

## Scenario-Based Branching Architecture

General outbound isn't one linear cadence. It branches based on channel availability and prospect response. See `references/example-campaign-patterns.md` for the full real-campaign reference.

---

## Scenario Decision Tree

```
Day 0: LinkedIn Invite sent
  │
  ├── SCENARIO 1: Has email + LI accepted within 10 days
  │   Day 0:  LinkedIn Invite (peer curiosity)
  │   Day 2:  LI Message 1 (hypothesis question + blog 1)
  │   Day 5:  Email 1 (reinforce hypothesis + introduce client)
  │   Day 12: LI Message 2 (deep research + insight question + blog 2)
  │   Day 46: LI Message 3 (soft close + door open)
  │
  ├── SCENARIO 2: Has email + LI NOT accepted initially → Email sent → LI accepted later
  │   Day 0:  LinkedIn Invite
  │   Day 10: Email 1 (hypothesis + PS about LI invite + blog 1)
  │   Day 11: LI accepted → LI Message 1 (reference email)
  │   Day 16: Email 2 (deep research + insight)
  │   Day 48: LI Message 2 (soft close + blog 2)
  │
  ├── SCENARIO 3: Has email + LI never accepted
  │   Day 0:  LinkedIn Invite
  │   Day 10: Email 1 (hypothesis + PS about invite + blog 1)
  │   Day 17: Email 2 (deep research + insight)
  │   Day 51: Email 3 (soft close + blog 2)
  │
  └── SCENARIO 4: No email, LinkedIn only
      Day 0:  LinkedIn Invite
      Wait until accepted (no time limit)
      Day +2:  LI Message 1 (hypothesis question + blog 1)
      Day +12: LI Message 2 (deep research + blog 2)
      Day +44: LI Message 3 (soft close)
```

---

## Copy Escalation Pattern (All Scenarios Follow This)

### Phase 1: Opening — LinkedIn Invite (Day 0)

**Variables used:** `{{firstName}}`, `{{companyName}}`, `{{industry}}`, `{{overlap_company}}`

**A/B Versions:**

Version A (industry-research angle):
```
Hi {{firstName}}, came across {{companyName}} while researching {{industry}} teams working on [angle topic]. Would love to connect and share some research we've put together in this space.
```

Version B (company-relevance angle):
```
Hi {{firstName}}, we're [client one-liner]. Given {{companyName}}'s work in {{overlap_company}}, thought there might be some shared interests worth exploring.
```

**Rules:**
- ≤280 characters
- Peer curiosity — NOT selling, just connecting
- Each version tests a different approach: industry-research vs. company-specific

---

### Phase 2: Hook-Typed Question — LI Message 1 / Email 1

**Variables used:** `{{overlap_personal}}`, `{{overlap_company}}`, `{{blog_link}}`, `{{hook_type}}`

This is the key touch — the hook does ALL the heavy lifting since there's no trigger signal.

**A/B Versions by Hook Type:**

Version A — Timeline hook (DEFAULT, 10% reply rate):
```
Hi {{firstName}}, [warm opener].

Teams in {{industry}} tackling {{overlap_personal}} typically go from [phase 1] to [result] in [timeframe]. Curious whether that matches what you're seeing at {{companyName}}.

Thought this might be useful: {{blog_link}}
```

Version B — Numbers hook (8.6% reply rate):
```
Hi {{firstName}}, [warm opener].

[Persona] teams see [X% improvement] in [metric related to overlap_personal] — the math is especially interesting at {{companyName}}'s scale.

Related: {{blog_link}}
```

Version C — Social proof hook (6.5% reply rate):
```
Hi {{firstName}}, [warm opener].

[Named company] tackled a similar challenge with {{overlap_personal}} and saw [specific result]. {{companyName}}'s {{overlap_company}} setup looks like a similar fit.

Thought this could be relevant: {{blog_link}}
```

Version D — Hypothesis hook (4.3% reply rate, FALLBACK only):
```
Hi {{firstName}}, [warm opener].

Would I be right in thinking {{overlap_personal}} is something your team is dealing with?

If so, thought this could be useful: {{blog_link}}

Curious what you're seeing on your end.
```

**What makes this work:**
- Hook type is selected per-prospect in Clay based on available proof points (`hook_type` column)
- Timeline hooks show speed-to-value — the strongest opener for cold outbound (2.3x vs hypothesis)
- The blog link is value delivery, not a pitch
- The close is curiosity, not a CTA for a meeting
- A/B testing built in: track `meta_hook_type_used` for performance attribution

**Email version additions:**
- Subject line: `{{overlap_personal}} at {{companyName}}`
- If LinkedIn not accepted, add PS: "PS: Also sent you a LinkedIn invite — easier to stay in touch."

---

### Phase 3: Deep Research — LI Message 2 / Email 2

**Variables used:** `{{company_workflow}}`, `{{tech_signals}}`, `{{company_insight}}`, `{{scale_trigger}}`, `{{blog_link_2}}`

This is the CREDIBILITY touch — proves you've done genuine research.

**A/B Versions:**

Version A (insight question):
```
Hi {{firstName}},

Spent some time looking at how {{companyName}} handles {{company_workflow}}.

{{company_insight}}

PS: thought this blog might be useful as well: {{blog_link_2}}
```

Version B (tech signal + scale question):
```
Hi {{firstName}},

Spent some time looking at how {{companyName}} handles {{company_workflow}}.

From what I can see, it seems you're running {{tech_signals}} in-house.

Is {{scale_trigger}} something {{companyName}} is working on?

PS: thought this might be useful: {{blog_link_2}}
```

**Email version additions:**
- Subject A: `Re: {{overlap_personal}}?` (thread continuation)
- Subject B: `One more thought`
- PS for email: "We're happy to share our knowledge. Is there any specific topic where an extra pair of expert eyes would be useful?"

**What makes this work:**
- "Spent some time looking at..." demonstrates real effort
- The insight question proves understanding, not just research
- Tech signals show you understand their technical reality
- Blog link 2 is deeper/more specific content than blog 1

---

### Phase 4: Soft Close — LI Message 3 / Email 3

**Variables used:** `{{overlap_personal}}`, `{{companyName}}`

**Template:**
```
Hi {{firstName}},

If {{overlap_personal}} or any related [client domain] topic becomes a priority at {{companyName}}, we're here and happy to have a chat.

Sometimes an outside perspective helps. If you want to bounce ideas off someone in this space, our door is open.

All the best with what you're building.

PS: We share a lot of our research on our [LinkedIn page / newsletter]. Feel free to follow to stay up to date.
```

**What makes this work:**
- Zero pressure — gives explicit permission to not respond
- "Our door is always open" = standing invitation
- PS redirects to ongoing content = nurture even if they don't reply
- The cadence ENDS warm, not desperate

---

## A/B Testing Framework

Every touch has two versions testing different approaches:

| Touch | Version A Tests | Version B Tests |
|---|---|---|
| LI Invite | Industry research angle | Company-specific relevance |
| Hypothesis touch | Question about role pain | Observation about company work |
| Deep research | Insight question | Tech signal + scale question |

Track reply rates per version per scenario. After 50+ sends per variant, commit to winners.

---

## Multi-Language Support

Every message must be generated in `{{language}}` (EN, NL, FR, DE, etc. based on client's target geographies).

**Two approaches:**
1. **Generate natively** — prompt includes: "Write in {{language}}" (better for natural phrasing)
2. **Generate in EN, then translate** — more consistent but can feel translated

Recommend approach 1 for languages the client's team speaks. Approach 2 as fallback.

---

## Industry Proof Points (for Social Proof Touch)

Pre-map per vertical, used in Scenario variations that include a social proof touch:

| Industry | Key Metric | Source |
|---|---|---|
| {vertical_1} | {specific metric} | {case study / capability} |
| {vertical_2} | {specific metric} | {case study / capability} |
| {vertical_3} | {specific metric} | {case study / capability} |

These come from `profile.md` (existing case studies) and `icp-mapping.md` (proof points per vertical).

---

## Enrichment Sequence

```
Account List (from market-mapping sources) → Clay Import
  → Columns 1-4: Company identity (name, domain, industry, size)
    → Columns 5-8: Company enrichment (Clearbit/Apollo)
      → Column 9: ICP Scoring Gate (A+ and A tier ONLY for general outbound)
        → Column 10: ICP Angle classification
          → Columns 11-14: Contact finding (waterfall for target persona)
            → Columns 15-16: LAYER 2a — overlap_personal + overlap_company
              → Columns 17-20: LAYER 2b — Deep research (company_deep, company_workflow, tech_signals, company_insight, scale_trigger)
                → Columns 21-22: Content matching (blog_link, blog_link_2)
                  → Column 23: Language detection
                    → Columns 24-28: LAYER 3 — Copy generation (LI invite, hypothesis touch, deep research touch, soft close)
                      → Columns 29-30: Assembly + export status
```

### Gates
- `score_icp_fit` must be A+ or A — general outbound doesn't waste effort on B tier
- `overlap_personal` = "SKIP" → flag for manual research or remove
- `overlap_company` = "SKIP" → flag for manual research or remove
- Standard email verification gate

### Credit Optimization
- General outbound uses ~8-10 Clay columns per lead (most expensive workflow)
- Compensate with strict ICP gating — only A+ and A tier
- Layer 2b (deep research) only runs for leads that pass Layer 2a (overlap variables generated successfully)
- Batch test: 20-30 leads through full sequence, review variable quality before scaling
- If a vertical consistently produces "SKIP" variables, it may not be suitable for general outbound — switch to trigger-only for that vertical
