---
name: gtm-strategy-generator
description: Generates a comprehensive GTM strategy for a Thrustlab client. Use when a client profile exists at client-profiles/{client-slug}/profile.md and a full outbound strategy is needed. Produces a 13-section strategy document with trigger playbook, messaging architecture, and 90-day blueprint. Reads from the canonical client profile — never from just a website URL alone.
---

# GTM Strategy Generator

Generates a complete GTM strategy saved to `strategies/{client-slug}.md`.

**Prerequisite:** `client-profiles/{client-slug}/profile.md` must exist. Run `/client-onboarding` first if it doesn't.

## Workflow

### Step 1: Load Client Profile

Read `client-profiles/{client-slug}/profile.md` and extract:
- Company identity, product, value prop
- Target verticals, company size, geographies
- Target personas with pains
- Competitors and differentiators
- Sales motion and tone preference

If the profile has gaps marked `[To be determined in strategy phase]`, research and fill them during Step 2.

### Step 2: Research Protocol

1. **Fetch client website** for primary intelligence (product pages, pricing, case studies, blog)
2. **Run 3-5 targeted web searches** using profile data:
   - `"{company_name}" + {primary_vertical} + {geography}` — market positioning
   - `"{competitor_1}" vs "{competitor_2}" + {industry_terms}` — competitive landscape
   - `"{primary_vertical}" + trends + challenges + {year}` — industry context
   - `"{target_persona_title}" + challenges + {industry}` — persona pain validation
   - `{industry} + regulations + compliance + {geography}` — if regulated vertical
3. **Never use generic searches** — always include industry-specific terms from the profile
4. **Document everything** with `[Source: URL]` citations

### Step 3: Generate Strategy Document

See `references/strategy-template.md` for the full 13-section + 2 appendix structure.

Every section must reference profile data. The strategy is not generic advice — it is a battle plan for THIS specific client, THIS specific market, THESE specific personas.

### Step 4: Quality Validation

See `references/research-protocol.md` for the quality checkpoint list.

Before delivering, verify:
- Every recommendation is tied to source material or profile data
- Zero generic value props — all industry/role specific
- Trigger playbook has minimum 10 triggers across 3 tiers, each with specific signal sources
- First 3 recommended trigger plays are clearly identified with rationale
- **Signal stacking section** identifies high-converting signal combinations for THIS client (see template Section 10)
- **Hook type recommendations** are populated with client proof points (see template Section 10b)
- All copy examples use client-specific language, not templates
- Scoring formula references actual client ICP criteria
- **Three new signal types** are evaluated for inclusion:
  - Champion tracking: Does the client have past customers/prospects to monitor? If yes, ALWAYS recommend.
  - Competitor customer targeting: Are competitors identifiable via tech detection or research? If yes, recommend.
  - Dark funnel: Which tools are appropriate for the client's market? (Dealfront/RB2B + Teamfluence/Trigify)

### Step 5: Save & Confirm

Save to `strategies/{client-slug}.md`

Present summary to user highlighting:
- Top 3 recommended trigger plays (these drive Phase 1 tooling setup)
- Key ICP refinements discovered during research
- Any gaps or assumptions that need client validation
