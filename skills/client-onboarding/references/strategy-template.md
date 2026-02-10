# Strategy Template — 13 Sections + 2 Appendices

Every strategy follows this exact structure. Each section header includes guidance on what to include.

## Document Header

```markdown
# {Company Name} — GTM Outbound Strategy
**Prepared by:** Thrustlab
**Date:** {date}
**Version:** 1.0

## Quick Reference
- **Client:** {company_name} ({website})
- **Primary vertical:** {vertical_1}
- **Primary persona:** {persona_title}
- **Sales motion:** {PLG / sales-led / hybrid}
- **Recommended trigger plays:** {trigger_1}, {trigger_2}, {trigger_3}
```

---

## Phase 1: Intelligence Gathering (Sections 1-3)

### Section 1: Company Snapshot
One page max. Bullet format.
- What they sell, to whom, and why it matters
- Current traction: ARR, customer count, notable logos (if known)
- Market position: leader / challenger / niche player
- Key insight: What makes their outbound motion unique or challenging?
- Source all claims: `[Source: URL]`

### Section 2: ICP & Buying Personas
For each persona (primary, secondary, tertiary):
- Title variations (all titles that map to this persona)
- Department and reporting line
- Daily responsibilities and KPIs they're measured on
- Primary pain: the specific problem the client solves for them
- Secondary pains: adjacent challenges
- How they buy: research process, internal champions, blockers
- **Disqualifiers:** What makes a company/person NOT a fit (anti-patterns)
- **Buying triggers:** What events make them actively search for a solution

### Section 3: Marketing Ecosystem Integration
Bridge outbound insights to existing marketing:
- What content assets already exist that outbound can leverage?
- What messaging gaps exist between marketing and sales?
- How should outbound signals feed back into marketing (intent data, content topics)?
- Quick wins: existing content that can be repurposed for outbound touchpoints

---

## Phase 2: Market Mapping (Sections 4-5)

### Section 4: TAM → SAM Filtering
Provide EXACT Clay enrichment fields and filters:
- Total Addressable Market: size estimate with methodology
- Serviceable Available Market: specific filters that narrow TAM to SAM
- For each filter: which Clay enrichment provider to use, which field to check, what values to match
- Example Clay table filter configuration

### Section 5: List-Building Mechanics
Concrete, executable instructions:
- Boolean search strings for LinkedIn Sales Navigator (ready to paste)
- Boolean search strings for Apollo (ready to paste)
- Alternative data sources specific to this client's vertical
- Import strategy: which source first, expected volume per source
- Deduplication approach

---

## Phase 3: Messaging & Personalization (Sections 6-9)

### Section 6: Messaging Architecture
Map pains → outcomes → proof 1:1:
- For each persona pain: what outcome does the client deliver? What proof exists?
- Core narrative: the story arc of the outbound campaign
- Positioning against each competitor: what angle to take when prospect uses competitor X

### Section 7: Scalable Personalization Framework
3 effort tiers with time allocations:
- **Tier 1 (30 sec/prospect):** Automated — Clay enrichment + AI copy generation
- **Tier 2 (2 min/prospect):** Semi-manual — AI-generated draft + human review for A-tier accounts
- **Tier 3 (5+ min/prospect):** White-glove — Full manual research for A+ tier strategic accounts
- Which accounts go in which tier (based on ICP score)

### Section 8: Outbound Cadences
2 complete cadence variants, 10-14 business days max:

**Cadence A: Trigger-based** (for workflows with active signal)
- Step 1 (Day 1): Trigger-based email
- Step 2 (Day 3): LinkedIn connection request
- Step 3 (Day 5): Follow-up email (different angle)
- Step 4 (Day 8): LinkedIn message
- Step 5 (Day 12): Breakup email or new insight

**Cadence B: Research-based** (for general outbound, no trigger)
- Step 1 (Day 1): Company-observation email
- Step 2 (Day 3): LinkedIn connection request
- Step 3 (Day 6): Value-add email (case study / insight)
- Step 4 (Day 9): LinkedIn message
- Step 5 (Day 13): Breakup with new angle

All copy examples must follow `shared/references/copy-rules.md`.

### Section 9: Top 6 Observational Openers
Based on ACTUAL research about the client's market, not templates:
- 3 company-based openers (work for any prospect in ICP)
- 3 trigger-based openers (specific to the most common triggers)
- Each opener: the observation + why it matters + how to bridge to value prop
- Show exactly how each would read for a real example prospect

---

## Phase 4: Advanced Systems (Sections 10-13)

### Section 10: Trigger Playbook & Signal Stacking
Minimum 10 triggers across 3 tiers:

**Tier 1 — Direct Intent (highest conversion):**
- Tech migrations, compliance deadlines, vendor evaluations
- For each: signal source, detection method, response time window, copy angle

**Tier 2 — Behavioral (medium conversion):**
- Website visits, content consumption, social engagement, peer activity
- For each: signal source, detection method, response time window, copy angle

**Tier 3 — Environmental (lower conversion, higher volume):**
- Funding events, leadership changes, hiring surges, market dynamics
- For each: signal source, detection method, response time window, copy angle

**Top 3 Recommended Trigger Plays:**
Clearly identify the 3 triggers to build first, with rationale:
1. {trigger_1}: Why this is priority — expected volume, conversion potential, data availability
2. {trigger_2}: Why this is priority — ...
3. {trigger_3}: Why this is priority — ...

These 3 drives the Phase 1 tooling setup.

**Signal Stacking Configuration:**
Analyze which trigger combinations create the highest-converting stacks for THIS client:

```markdown
## Signal Stacking Matrix
Review all active trigger plays and identify compound signal stacks:

### Available Signal Sources (check which apply)
- [ ] Website visitor identification (Dealfront / RB2B)
- [ ] Job change monitoring (UserGems / Clay signal)
- [ ] Job posting monitoring (Clay signal / LinkedIn Jobs)
- [ ] Funding event monitoring (Crunchbase / Clay signal)
- [ ] Tech stack change monitoring (BuiltWith / Wappalyzer)
- [ ] Hiring surge monitoring (Clay signal / LinkedIn)
- [ ] Champion tracking (UserGems / CRM data)
- [ ] Competitor customer identification (Claygent / BuiltWith)
- [ ] Dark funnel - LinkedIn engagement (Teamfluence / Trigify)
- [ ] Dark funnel - community signals (Common Room)
- [ ] Content engagement (marketing automation / Phantom Buster)
- [ ] Compliance/regulatory events (industry sources)

### High-Converting Stacks (Client-Specific)
Based on {client_name}'s market, these signal combinations should trigger escalated treatment:

| Stack ID | Signals Combined | Expected Multiplier | Routing |
|---|---|---|---|
| stack_1 | {signal_a} + {signal_b} | {X}x baseline | {Hot/Warm} |
| stack_2 | {signal_a} + {signal_c} | {X}x baseline | {Hot/Warm} |
| stack_3 | {signal_b} + {signal_c} + {signal_d} | {X}x baseline | {Hot} |

### Composite Score Weights (Client-Specific)
Adjust default weights from signal-stacking-guide.md based on:
- Which signals are most predictive for {client_name}'s buyer journey
- Data availability and reliability of each signal source
- Client's sales capacity (don't create more Hot leads than AEs can handle)
```

See `shared/references/signal-stacking-guide.md` for the full composite scoring framework.

### Section 10b: Hook Type Recommendations
Based on {client_name}'s available proof points, recommend the hook type strategy:

```markdown
## Hook Type Availability Audit

### Timeline Proof Points
- Vertical 1 ({vertical}): {discovery → pilot → result in X weeks}
- Vertical 2 ({vertical}): {discovery → pilot → result in X weeks}
- Average implementation time: {X weeks}
- First value milestone: {what happens at week Y}

### Numbers Proof Points
- {X% improvement in Y for Z company type}
- {X hours/$ saved per week for Z role}
- {X% cost reduction / revenue increase}

### Named Case Studies (for Social Proof hooks)
- {Company name} ({vertical}, {size}) — {headline result}
- {Company name} ({vertical}, {size}) — {headline result}

### Recommended Hook Strategy
| Vertical | Default Hook | Fallback Hook | Rationale |
|---|---|---|---|
| {vertical_1} | {timeline/numbers/social_proof} | hypothesis | {why} |
| {vertical_2} | {timeline/numbers/social_proof} | hypothesis | {why} |

### A/B Testing Plan
- Weeks 1-2: 40% timeline, 20% numbers, 20% social proof, 20% hypothesis
- Weeks 3-4: Scale to winning type, continue testing within type
- Week 5+: 70% winner, 30% rotation
```

See `shared/references/hook-types-guide.md` for hook type definitions and performance benchmarks.

### Section 11: GTM Flywheel
Focus on compound effects:
- How outbound data feeds marketing intelligence
- How reply data refines ICP scoring
- How trigger data improves over time
- Feedback loops between Clay data and strategy refinement

### Section 12: Playbooks & Enablement
- 3-minute research checklist for SDRs
- Call script framework (opener → discovery question → bridge → CTA)
- Objection handling for top 5 objections per persona
- Voicemail script (≤20 seconds)

### Section 13: The Thrustlab 90-Day GTM Machine Blueprint
Week-by-week execution plan:
- Weeks 1-2: Infrastructure + tooling setup
- Weeks 3-4: Market mapping + ICP validation + first list build
- Weeks 5-6: First trigger play live + general outbound live
- Weeks 7-8: Second and third trigger plays live
- Weeks 9-10: Optimization (A/B testing copy, refining scoring)
- Weeks 11-12: Scale (increase volume, add new trigger plays, expand ICP)

---

## Appendices

### Appendix A: Competitive Battlecards
For each competitor:
- Their positioning vs. client's positioning
- When prospects mention competitor: what to say
- Where client wins vs. where competitor wins
- Proof points to counter competitor claims

### Appendix B: Metrics & KPIs
- Reply rate targets by workflow type
- Meeting book rate targets
- ICP accuracy targets (% of outreach that matches refined ICP)
- Trigger response time targets
- Weekly/monthly reporting cadence
