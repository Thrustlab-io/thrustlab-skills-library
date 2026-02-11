---
name: strategy-refiner
description: Pulls the client-refined strategy from Notion, updates the local strategy.md, and generates detailed implementation documents (ICP mapping, market mapping, and top 3 workflow specifications). These documents serve as the basis for Clay table population and are pushed back to Notion. Use after initial client review meeting when the strategy has been refined with client feedback.
---

# Strategy Refiner

Bridges the gap between initial strategy and detailed implementation. Takes the client-refined strategy from Notion and produces actionable, detailed documents ready for execution.

**When to use:** After the initial client review meeting, when the strategy document in Notion has been updated with client feedback and clarifications.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- `Prospects/{client-slug}/strategy.md` exists (initial version)
- Notion GTM Client Hub exists with updated strategy content
- Notion MCP server is configured

**Produces:**
1. `Prospects/{client-slug}/strategy.md` — Updated with client refinements
2. `Prospects/{client-slug}/icp-mapping.md` — Detailed ICP with scoring, personas, and pain mapping
3. `Prospects/{client-slug}/market-mapping.md` — TAM→SAM mapping with search criteria and data sources
4. `Prospects/{client-slug}/workflows/` — Directory containing 3 workflow specifications
   - `workflow-1-{trigger-name}.md`
   - `workflow-2-{trigger-name}.md`
   - `workflow-3-{trigger-name}.md`
5. Updated Notion pages with all new content

---

## Workflow

### Step 1: Pull Refined Strategy from Notion

1. **Locate the strategy page** in the client's Notion workspace:
   - Navigate to `🎯 {Company Name} - GTM Client Hub`
   - Find the child page: `GTM Strategy - {Company Name}` 🎯

2. **Read the updated strategy content** from Notion:
   - Extract all sections (this has been refined with client input)
   - Note any new information, clarifications, or changes from the original
   - Pay special attention to:
     - Confirmed ICP details (firmographics, personas, pain points)
     - Top 3 recommended trigger plays (may have been reordered or modified)
     - Any new competitive intel or market insights
     - Clarified value propositions and positioning

3. **Update local strategy.md**:
   - Compare Notion version with `Prospects/{client-slug}/strategy.md`
   - Merge client refinements into the local file
   - Add a changelog note at the top:
     ```markdown
     *Last updated: {date} - Refined with client feedback*
     ```
   - Preserve the original structure and citations
   - Save the updated file

### Step 2: Generate ICP Mapping Document

Using the refined strategy and profile, create a detailed ICP mapping document.

**Reference:** See the `icp-mapping` skill for the full template and methodology.

**Key sections to include:**

#### Account-Level ICP
- **Firmographic criteria** with scoring weights:
  - Company size (employee count, revenue)
  - Industry/vertical (primary and secondary)
  - Geography (regions, countries)
  - Growth stage (funding, growth rate)
  - Technology stack (required and complementary tools)
- **Scoring formula**: Define how accounts are scored 0-100
- **Tier definitions**: Tier 1 (90-100), Tier 2 (70-89), Tier 3 (50-69)

#### Persona Cards
For each target persona (typically 2-4):
- **Title variations**: All relevant job titles
- **Seniority level**: IC vs. Manager vs. Director vs. VP vs. C-level
- **Department**: Primary team/function
- **Core responsibilities**: What they do day-to-day
- **Pain points**: Top 3-5 specific pains (not generic)
- **Success metrics**: How they're measured/evaluated
- **Buying triggers**: What makes them actively look for solutions
- **Objections**: Common pushback and how to address
- **Preferred channels**: Where to reach them (email, LinkedIn, events)

#### Industry-Specific Pain Mapping
For each target vertical:
- **Industry context**: Current trends, challenges, regulations
- **Pain point hierarchy**: Most urgent → least urgent
- **Proof points**: Customer stories, data, case studies specific to this vertical
- **Competitive landscape**: Who else serves this vertical, gaps in market

#### Data Source Recommendations
- Where to find these accounts (data sources, enrichment providers)
- Where to find these personas (LinkedIn filters, job boards)
- Signal sources for pain point validation

**Output:** Save to `Prospects/{client-slug}/icp-mapping.md`

### Step 3: Generate Market Mapping Document

Using the refined ICP, create the TAM→SAM mapping with specific search criteria.

**Reference:** See the `market-mapping` skill for the full template.

**Key sections to include:**

#### Total Addressable Market (TAM)
- **Market definition**: Who could theoretically use this product
- **Market size**: Number of companies, total revenue opportunity
- **Segments**: Break down by vertical, size, geography

#### Serviceable Available Market (SAM)
- **ICP filters applied**: How TAM narrows to SAM
- **Realistic targets**: Number of addressable accounts
- **Prioritization**: Which segments to target first and why

#### Boolean Search Strings
For each target segment, provide Clay/LinkedIn search criteria:
- **Company filters**: Industry, size, location, growth signals
- **Exclusion criteria**: Who to filter out
- **Technology filters**: Required stack, competitor tool usage

Example:
```
Segment: Mid-market SaaS companies in FinTech (US/UK)
- Industry: Financial Services, Fintech
- Employee count: 50-500
- Location: United States, United Kingdom
- Technology: Uses Stripe OR Plaid OR similar
- Exclude: Stealth mode, recently acquired
```

#### Enrichment Strategy
- **Required data points**: What to enrich for each account
- **Enrichment sequence**: Which providers in which order (Clearbit → Apollo → manual)
- **Signal enrichment**: Job postings, funding, tech stack, website visits

#### Alternative Data Sources
- **Niche directories**: Industry-specific lists, associations
- **Community sources**: Slack groups, forums, newsletters
- **Events**: Conference attendee lists, webinar participants
- **Partners**: Integration partner customer lists

**Output:** Save to `Prospects/{client-slug}/market-mapping.md`

### Step 4: Generate Top 3 Workflow Specifications

For each of the top 3 recommended trigger plays from the strategy, create a detailed workflow specification document.

**Create directory:** `Prospects/{client-slug}/workflows/`

**For each workflow, create:** `workflow-{1,2,3}-{trigger-name}.md`

Example filenames:
- `workflow-1-funding.md`
- `workflow-2-job-postings.md`
- `workflow-3-website-intent.md`

#### Workflow Document Structure

Each workflow specification should include:

**1. Overview**
- Trigger name and description
- Why this trigger is effective for this client
- Expected volume (accounts/month)
- Expected conversion rate (benchmark)

**2. Signal Detection**
- **Data sources**: Where to detect this trigger (e.g., Crunchbase for funding, Dealfront for website visits)
- **Signal criteria**: What qualifies as a valid trigger
  - Example (funding): "Series A+ rounds, $5M+, announced within last 14 days"
- **Filtering logic**: How to exclude false positives
- **Frequency**: How often to check for new signals

**3. Enrichment Sequence**
- **Required data points**: What to enrich after signal is detected
- **Enrichment order**: Step-by-step data gathering
  - Step 1: Basic firmographics (Clearbit)
  - Step 2: Contact discovery (Apollo, LinkedIn)
  - Step 3: Trigger-specific research (Claygent prompts)
- **Research prompts**: Specific Claygent/AI prompts for deep research
  - Example: "Find the company's growth plans from recent press releases and blog posts"

**4. Scoring & Prioritization**
- **ICP fit score**: How to score this account (reference ICP mapping)
- **Trigger urgency score**: How fresh/urgent is this signal
- **Combined prioritization**: Formula for final priority score
- **Tier assignment**: Which accounts go into high/medium/low priority

**5. Outreach Strategy**
- **Target personas**: Who to contact (reference ICP personas)
- **Contact discovery**: How to find the right person
- **Timing**: When to reach out after signal (e.g., "7 days after funding announcement")
- **Channel**: Primary channel (email) and secondary (LinkedIn)
- **Cadence structure**: Number of touchpoints, spacing, escalation
  - Example: Day 1 (email), Day 4 (email), Day 7 (LinkedIn), Day 11 (email)

**6. Messaging Guidelines** *(Not full copy - just guidelines)*
- **Hook angle**: The core narrative (e.g., "Growth momentum hook - they just raised money and are scaling")
- **Key elements to reference**:
  - The specific trigger (funding amount, job posting, etc.)
  - Relevant pain point for this trigger
  - Social proof from similar companies
  - Clear CTA
- **Personalization requirements**: What must be customized per account
- **Tone**: Conversational, formal, or provocative (from profile)

**7. Clay Table Configuration**
- **Table purpose**: What this table tracks
- **Input sources**: Where accounts enter (CSV import, API, webhook)
- **Key columns**: List all necessary columns
  - Account fields (name, domain, industry, etc.)
  - Signal fields (trigger date, trigger details, signal source)
  - Enrichment fields (employee count, tech stack, etc.)
  - Contact fields (name, title, email, LinkedIn)
  - Workflow fields (status, priority, last touch)
- **Automations**: What happens automatically
  - Auto-enrichment rules
  - Auto-assignment to sequences
  - Alerts/notifications

**8. Success Metrics**
- **Activity metrics**: Emails sent, open rate, reply rate
- **Outcome metrics**: Meetings booked, opportunities created, pipeline value
- **Benchmarks**: Expected performance (from strategy)
- **Optimization triggers**: When to adjust (e.g., "If reply rate < 2% after 50 sends, revise copy")

**Output:** Save 3 workflow files to `Prospects/{client-slug}/workflows/` directory

### Step 5: Push All Documents to Notion

Update the client's Notion workspace with all new content:

1. **Update GTM Strategy page** (if local version has additional refinements)

2. **Update ICP Mapping page** 🎯:
   - Navigate to the existing empty ICP Mapping page
   - Populate with content from `icp-mapping.md`
   - Ensure persona cards are formatted cleanly (use callout blocks or tables)

3. **Create Market Mapping page** (if it doesn't exist) or update existing:
   - Add to the main hub as a child page
   - Populate with content from `market-mapping.md`
   - Include search strings in code blocks for easy copy-paste

4. **Update Workflows page** ⚡:
   - Navigate to the existing Workflows page
   - Create 3 sub-pages (or toggle blocks) for each workflow:
     - `Workflow 1: {Trigger Name}`
     - `Workflow 2: {Trigger Name}`
     - `Workflow 3: {Trigger Name}`
   - Populate each with content from respective workflow files
   - Use clear headings and formatting for readability

5. **Update Documentation Hub section** on main hub page with any new page links

### Step 6: Confirm & Output

Provide the user with:

**Local files created/updated:**
- `strategy.md` - Updated with client refinements
- `icp-mapping.md` - Detailed ICP with scoring and personas
- `market-mapping.md` - TAM→SAM with search criteria
- `workflows/workflow-1-{name}.md`
- `workflows/workflow-2-{name}.md`
- `workflows/workflow-3-{name}.md`

**Notion pages updated:**
- GTM Strategy page URL
- ICP Mapping page URL
- Market Mapping page URL (if new)
- Workflows page URL with 3 sub-pages

**Key changes from original strategy:**
- Summary of client refinements (2-3 bullet points)
- Any new ICP insights
- Confirmed top 3 trigger plays

**Next steps:**
- Review the detailed workflow specifications
- Proceed with Clay table setup using these specifications
- Begin with: `/clay-campaign-generator` or `/workflow-{trigger}-trigger` for detailed copy generation

---

## Notes

- This skill is the critical bridge between high-level strategy and hands-on execution
- The documents produced here are the "blueprint" that all downstream work references
- Be thorough in this phase — clarity here saves time in execution
- These documents should be specific enough that someone else could execute from them
- Keep messaging guidelines at the strategic level — full copy comes in the next phase with individual workflow skills
