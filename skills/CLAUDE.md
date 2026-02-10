# Thrustlab GTM System

## Team
- Kwinten (kwinten@thrustlab.io)
- Jan (jan@thrustlab.io)

## Client Data Location
- All client data: `Prospects/{client-slug}/`
  - Profile: `Prospects/{client-slug}/profile.md`
  - Strategy: `Prospects/{client-slug}/strategy.md`
  - Market mapping: `Prospects/{client-slug}/market-mapping.md`
  - ICP mapping: `Prospects/{client-slug}/icp-mapping.md`
  - Tooling setup: `Prospects/{client-slug}/tooling-setup.md`
- Clay referral: https://clay.com?via=c75c72

## Universal Copy Rules
- Email subject: ≤45 characters
- Email body: ≤90 words
- LinkedIn message: ≤280 characters
- Opening line: Always observational, never generic. MUST use the hook type system (see below).
- CTA: Stage-appropriate ("sanity check" not "demo")
- Never use: "innovative", "cutting-edge", "game-changing", "synergy", "leverage", "disruptive"
- Two opener types per workflow: company-based (evergreen) + trigger-based (timely hook)

## Hook Type System (2025 Performance Data)
Every email opener must use one of four hook types. See `shared/references/hook-types-guide.md`.
- **Timeline** (DEFAULT — 10.0% reply rate): Compressed achievement milestones
- **Numbers** (8.6% reply rate): Specific quantified claims
- **Social proof** (6.5% reply rate): Named reference + result
- **Hypothesis** (4.3% reply rate): Pain framed as question — FALLBACK only
The `hook_type` variable is classified per-prospect in Clay and drives copy generation branching.

## Signal Stacking
When a client has 2+ active trigger plays, deploy the Signal Aggregation Table.
See `shared/references/signal-stacking-guide.md`.
- Layer 1 (Fit): Static ICP scoring — gate before signal processing
- Layer 2 (Intent): Time-decayed intent signals (website, content, dark funnel)
- Layer 3 (Relationship): Champion tracking, competitor customer status, past engagement
- Composite score drives routing: Hot (≥100) → Slack alert + manual | Warm (70-99) → priority play | Active (40-69) → standard play

## Prompt Iteration Pipeline
Every Clay prompt should be tracked with `meta_prompt_version` for systematic improvement.
See `shared/references/prompt-iteration-pipeline.md`.
- Baseline → Sample → Categorize failures → Diagnose → Iterate ONE change → Test → Validate → Log
- Target: 70%+ pass rate, 0% fail rate, 30-40+ iterations per campaign over time

## Skill Execution Order
1. `/client-onboarding` → `Prospects/{slug}/profile.md`
2. `/gtm-strategy-generator` → `Prospects/{slug}/strategy.md`
3. `/notion-project-creator` → Notion workspace
4. `/slack-channel-creator` → Slack channel
5. `/tooling-setup-guide` → `Prospects/{slug}/tooling-setup.md`
6. `/market-mapping` → `Prospects/{slug}/market-mapping.md`
7. `/icp-mapping` → `Prospects/{slug}/icp-mapping.md`
8. `/workflow-{trigger-name}` → Clay prompts + copy + enrichment config

## Available Workflow Skills
### Standard Trigger Plays
- `/workflow-website-trigger` — Website visitor intent (RB2B/Dealfront/Clearbit)
- `/workflow-job-change-trigger` — Target persona changes jobs (UserGems/Clay)
- `/workflow-job-posting-trigger` — Company posts relevant job (Clay signal/LinkedIn)
- `/workflow-funding-trigger` — Company raises funding (Crunchbase/Clay)
- `/workflow-tech-change-trigger` — Company adds/removes tech (BuiltWith)
- `/workflow-growth-trigger` — Headcount surge or expansion (Clay signal)
- `/workflow-compliance-trigger` — Regulatory/compliance deadline
- `/workflow-content-trigger` — Prospect engages with client content

### High-Converting Signal Plays (NEW)
- `/workflow-champion-tracking` — Former customers/prospects change jobs (#1 converting signal)
- `/workflow-competitor-customer` — Companies using competitor products (2.5x conversion, 3x with named switcher)
- `/workflow-dark-funnel` — Anonymous engagement signals from LinkedIn/community/website (Dealfront/RB2B + Teamfluence/Trigify)

### Non-Trigger
- `/workflow-general-outbound` — Proactive ICP outreach, research-based personalization

## Phase Boundaries
Do NOT run all phases in one conversation. Between phases:
- Save all outputs to files
- Compact/clear context
- Next phase reads from files, not conversation history

## Quality Standards
- Every output must be client-specific. Generic = failure.
- Every Clay prompt must include the client context block (name, product, value prop, pains).
- Every copy must pass the "would a human SDR write this?" test.
- Research prompts must include ICP qualification datapoints.
- Fallback values must be industry-specific, never fully generic.
- Every copy-producing column must tag `hook_type` used and `prompt_version`.
- For shared rules on copy, tokens, hook types, signal stacking, and Clay: see `shared/references/`.

## Shared Reference Files
- `shared/references/copy-rules.md` — Universal copy standards
- `shared/references/personalization-tokens.md` — Variable naming and resolution order
- `shared/references/clay-enrichment-guide.md` — Clay best practices, ICP scoring, signal stacking, dark funnel tools
- `shared/references/hook-types-guide.md` — Hook type definitions, performance data, A/B testing protocol
- `shared/references/signal-stacking-guide.md` — Composite scoring framework, tier routing, known stacks
- `shared/references/prompt-iteration-pipeline.md` — Systematic prompt improvement process
