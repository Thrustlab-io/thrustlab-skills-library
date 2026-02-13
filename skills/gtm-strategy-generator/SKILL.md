---
name: gtm-strategy-generator
description: The first step for any new Thrustlab client. Takes a LinkedIn Company URL and domain as input, auto-discovers the company profile through web research, produces the canonical profile.md, then generates a GTM strategy. Outputs both Prospects/{client-slug}/profile.md and Prospects/{client-slug}/strategy.md. Fields that cannot be auto-discovered are marked "[To be confirmed by client]" for follow-up.
---

# GTM Strategy Generator

The entry point for every new Thrustlab client engagement. Produces:
1. `Prospects/{client-slug}/profile.md` — the canonical client profile (read by ALL downstream skills)
2. `Prospects/{client-slug}/strategy.md` — the full GTM strategy
3. `Prospects/{client-slug}/onboarding-questions.md` — questions for the client onboarding meeting

**Input required:** LinkedIn Company URL + company domain (e.g., `https://www.linkedin.com/company/acme-corp` + `acme.com`)

**No other prerequisites.** This skill auto-discovers everything it needs through web research.

## Workflow

### Step 1: Generate Client Slug

Derive the slug from the company name (extracted from LinkedIn URL or domain).

Format: lowercase, hyphens, no spaces.
- "Acme Corp" → `acme-corp`
- "Quality Guard" → `quality-guard`
- "DataFlow.io" → `dataflow`

Create the directory: `Prospects/{client-slug}/`

### Step 2: Auto-Discover Company Profile

Build a complete client profile from scratch using web research. Reference `references/profile-field-reference.md` for the full field list and discoverability tiers.

**Research sequence:**

1. **Fetch company website** — analyze all key pages:
   - Homepage: positioning, value prop language, primary CTA style
   - Product/platform pages: what they actually sell, features, use cases
   - Pricing page: business model, target segment signals (enterprise vs SMB)
   - About page: founding story, team, mission, HQ, company size
   - Case studies/customers: named logos, verticals served, results with timelines
   - Blog/resources: content topics, publishing cadence, thought leadership positioning
   - Integrations/partners: tech ecosystem, complementary tools

2. **Parse LinkedIn company page** — extract:
   - Employee count, industry classification, HQ location
   - Company description and specialties
   - Founded year, follower count

3. **Run discovery web searches** (5-7 targeted searches):
   - `"{company_name}" product OR solution OR platform` — product positioning
   - `"{company_name}" customers OR "case study" OR testimonial` — proof points
   - `"{company_name}" vs OR alternative OR competitor` — competitive landscape
   - `"{company_name}" site:g2.com OR site:capterra.com` — reviews and market positioning (if SaaS)
   - `"{company_name}" funding OR raised site:crunchbase.com` — stage, funding, growth trajectory
   - `"{company_name}" {inferred_industry_terms}` — market context
   - `"{primary_vertical}" + trends + challenges + {year}` — industry context

4. **Synthesize into profile fields** using `references/profile-field-reference.md` as completeness checklist:
   - Populate all `[Auto]` fields directly from research findings
   - Populate `[Infer]` fields with best-guess values, annotate with `[Inferred — verify with client]`
   - Populate `[Client]` fields with reasonable defaults where possible, annotate with `[To be confirmed by client]`
   - Never leave a field completely blank

**Key inference rules:**
- Sales motion: "Try free" / "Start free trial" → PLG | "Book a demo" / "Talk to sales" → sales-led | Both → hybrid
- Target verticals: rank by frequency in case studies and website copy (most mentioned = primary)
- Personas: infer from "built for [role]" pages, case study contacts, job posting patterns
- Tone: analyze website copy — short/punchy sentences = conversational, formal language = formal, bold claims = provocative
- Geography: check language options, currency, office locations, compliance references (GDPR = EU focus)

### Step 3: Save Draft profile.md

Save to `Prospects/{client-slug}/profile.md` using the exact template structure from `references/profile-field-reference.md`.

**Important:** Strip the discoverability tier tags (`[Auto]`, `[Infer]`, `[Client]`) from the output — only include the confidence annotations where applicable (`[Inferred — verify with client]`, `[To be confirmed by client]`).

The profile must include:
- All 10 sections: Company Identity, Sales Motion, Target Market, Target Personas, Value Proposition, Competitive Landscape, Tone & Messaging, Tech Stack, Existing Assets, Engagement Parameters
- Footer with `*Profile created: {date}*` and `*Last updated: {date}*`

### Step 4: Research + Generate Strategy

Now use the auto-discovered profile as foundation to generate the full GTM strategy. Do **NOT** add messaging copy. This is a strategic battle plan, not a messaging document.

**4a. Deepen research using the research protocol:**

See `references/research-protocol.md` for detailed execution rules.

Using profile data as context, run targeted research:
- `"{company_name}" + {primary_vertical} + {geography}` — market positioning
- `"{competitor_1}" vs "{competitor_2}" + {industry_terms}` — competitive landscape
- `"{primary_vertical}" + trends + challenges + {year}` — industry context
- `"{target_persona_title}" + challenges + {industry}` — persona pain validation
- `{industry} + regulations + compliance + {geography}` — if regulated vertical

Never use generic searches — always include industry-specific terms from the profile.
Document everything with `[Source: URL]` citations.

**4b. Generate strategy document:**

See `references/strategy-template.md` for the structure.

Every section must reference profile data. The strategy is not generic advice — it is a battle plan for THIS specific client, THIS specific market, THESE specific personas.

**4c. Quality validation:**

See `references/research-protocol.md` for the quality checkpoint list.

Before delivering, verify:
- Every recommendation is tied to source material or profile data
- ** NO MESSAGING COPY** - this is an initial strategy and this will make the document too big and unfocused. Messaging will come in later steps once we have the strategy locked down.
- Zero generic value props — all industry/role specific
- 3 recommended plays (INCLUDING general outbound) are clearly identified with rationale

### Step 5: Save strategy.md

Save to `Prospects/{client-slug}/strategy.md`

### Step 6: Generate Onboarding Questions Document

Generate `Prospects/{client-slug}/onboarding-questions.md` — a structured document of questions for the client onboarding meeting.

**Purpose:** This document consolidates everything we need to ask the client to confirm/correct from the auto-discovered profile and validate the strategy. It ensures the onboarding meeting is efficient and nothing is missed.

**The document must include these sections:**

1. **Profile Confirmation** — Questions for every field marked `[To be confirmed by client]` or `[Inferred — verify with client]` in profile.md, organized by section (Sales Motion, Target Market, Personas, Engagement Parameters)

2. **Strategy Validation** — Questions to validate the recommended trigger plays, targeting approach, messaging tone, and any strategy-specific assumptions

3. **Product & Proof Points** — Questions about customer count, success stories, testimonials available for outreach, product capabilities, pricing confirmation, trial/guarantee offers

4. **Tech Stack & Infrastructure** — Questions about CRM, email sequencer, LinkedIn access, existing tools, databases, and any GTM tooling already in use

5. **Existing Content & Assets** — Questions about content library, lead magnets, social presence, reviews, sales collateral, customer logos available for use

6. **Operations & Capacity** — Questions about who handles meetings, booking tools, follow-up process, budget for tooling, and whether Thrustlab manages operations end-to-end

7. **Quick Wins & Priorities** — Questions about existing warm leads, past trial users, referral programs, and seasonal priorities

8. **Success Metrics & Expectations** — Questions about primary KPIs, reporting cadence, reporting format, and decision-maker for strategy changes

9. **Action Items After Meeting** — A table template with Owner and Due Date columns for post-meeting follow-ups

Use `references/onboarding-questions-template.md` as the structure template.

**Format:** Each question should be a checkbox `- [ ]` for easy tracking during the meeting. Group questions logically. Include brief context for why each question matters where non-obvious.

### Step 7: Present Outputs

Present all three outputs to the user:

**Profile summary:**
- Company identity and positioning (1-2 sentences)
- Number of fields fully populated vs. needing confirmation
- List all fields marked `[To be confirmed by client]` grouped by section

**Strategy highlights:**
- Top 3 recommended trigger plays with rationale
- Key ICP refinements discovered during research

**Onboarding questions:**
- Confirm the onboarding-questions.md was generated
- Highlight the most critical questions that MUST be answered before implementation can begin

**Next steps:**
- Review and confirm/correct fields marked `[To be confirmed by client]` in profile.md
- Use onboarding-questions.md as the agenda for the client onboarding meeting
- Once profile is confirmed, proceed with: `/notion-project-creator` → `/slack-channel-creator` → `/tooling-setup-guide` → `/market-mapping` → `/icp-mapping` → `/workflow-*`
