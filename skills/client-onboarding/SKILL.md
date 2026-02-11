---
name: client-onboarding
description: End-to-end onboarding for a new Thrustlab client. Takes a LinkedIn Company URL and domain, auto-discovers the company profile, generates a full GTM strategy, and creates the Notion workspace — all in one skill. Outputs {client-slug}/profile.md, {client-slug}/strategy.md, and a fully populated Notion GTM Client Hub.
---

# Client Onboarding

The single entry point for every new Thrustlab client engagement. Combines profile discovery, strategy generation, and Notion workspace creation into one seamless workflow.

**Produces:**
1. `{client-slug}/profile.md` — the canonical client profile (read by ALL downstream skills)
2. `{client-slug}/strategy.md` — the full GTM strategy
3. Notion GTM Client Hub — fully populated workspace with strategy, competitor analysis, roadmap, and workflows

**Input required:** LinkedIn Company URL + company domain (e.g., `https://www.linkedin.com/company/acme-corp` + `acme.com`)

**Prerequisites:**
- Notion MCP server is configured

**No other prerequisites.** This skill auto-discovers everything it needs through web research.

---

## Phase 1: Profile & Strategy Generation

### Step 1: Generate Client Slug

Derive the slug from the company name (extracted from LinkedIn URL or domain).

Format: lowercase, hyphens, no spaces.
- "Acme Corp" → `acme-corp`
- "Quality Guard" → `quality-guard`
- "DataFlow.io" → `dataflow`

Create the directory: `{client-slug}/`

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
   - Recent posts for tone and messaging style

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

Save to `{client-slug}/profile.md` using the exact template structure from `references/profile-field-reference.md`.

**Important:** Strip the discoverability tier tags (`[Auto]`, `[Infer]`, `[Client]`) from the output — only include the confidence annotations where applicable (`[Inferred — verify with client]`, `[To be confirmed by client]`).

The profile must include:
- All 10 sections: Company Identity, Sales Motion, Target Market, Target Personas, Value Proposition, Competitive Landscape, Tone & Messaging, Tech Stack, Existing Assets, Engagement Parameters
- Footer with `*Profile created: {date}*` and `*Last updated: {date}*`

### Step 4: Research + Generate Strategy

Now use the auto-discovered profile as foundation to generate the full GTM strategy.

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

See `references/strategy-template.md` for the full 13-section + 2 appendix structure.

Every section must reference profile data. The strategy is not generic advice — it is a battle plan for THIS specific client, THIS specific market, THESE specific personas.

**4c. Quality validation:**

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

### Step 5: Save strategy.md

Save to `{client-slug}/strategy.md`

---

## Phase 2: Notion Workspace Creation

### Step 6: Create Main Hub Page

Create a parent page titled: `🎯 {Company Name} - GTM Client Hub`
Create an entry on the "Companies" page in the "Prospect" column: https://www.notion.so/quantascale/Companies-266fbebd816080549bfcccc0dee598b3

Include these sections on the main page:

#### Braindump & Quick Notes
A callout block with empty checkboxes — space for observations and ideas during execution.

#### To Discuss with Client
A callout block with checkboxes for topics to raise in next client meeting.

#### Documentation Hub
A collapsible callout with links to all sub-pages created below.

### Step 7: Create Sub-Pages

Create each as a child page of the main hub. For the strategy page, populate with actual strategy content from `{client-slug}/strategy.md`. All other pages get their template structure:

1. **GTM Strategy - {Company Name}** 🎯 — Populate with full strategy.md content
2. **Competitor Analysis - {Company Name}** 📊 — Pre-fill with competitors from profile.md
3. **Infrastructure** 📋 — Domains and tools used
4. **ICP Mapping** 🎯 — Empty, populated during ICP mapping phase
5. **Roadmap** 🗓️ — Pre-fill with 90-day blueprint from strategy
6. **Meeting Notes** 📅 — Empty template with date + attendees + notes structure
7. **Workflows** ⚡ — Pre-fill with trigger playbook from strategy
8. **Copy Repository** 💬 — Will store approved copy as cadences are built

---

## Phase 3: Confirm & Present Outputs

### Step 8: Present Everything

**Profile summary:**
- Company identity and positioning (1-2 sentences)
- Number of fields fully populated vs. needing confirmation
- List all fields marked `[To be confirmed by client]` grouped by section

**Strategy highlights:**
- Top 3 recommended trigger plays with rationale
- Key ICP refinements discovered during research
- Any gaps or assumptions that need client validation

**Notion workspace:**
- Main hub page URL
- Infrastructure page URL (needed for Slack channel creator)
- Confirmation that strategy content is populated
- Summary of what was pre-filled vs. what will be populated in later phases

**Next steps:**
- Review and confirm/correct fields marked `[To be confirmed by client]` in profile.md
- Once profile is confirmed, proceed with: `/slack-channel-creator` → `/tooling-setup-guide` → `/market-mapping` → `/icp-mapping` → `/workflow-*`
