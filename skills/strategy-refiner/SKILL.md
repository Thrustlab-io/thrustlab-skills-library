---
name: strategy-refiner
description: Orchestrates the refinement pipeline after client feedback. Pulls the client-refined strategy from Notion, updates local strategy.md, then delegates to specialized skills (market-mapping, icp-mapping, and 3 workflow-* skills) to build implementation documents. All outputs pushed back to Notion. Use after initial client review meeting.
---

# Strategy Refiner

Orchestration skill that bridges initial strategy and detailed implementation. Takes client-refined strategy from Notion and coordinates specialized skills to produce execution-ready documents.

**When to use:** After the initial client review meeting, when the strategy document in Notion has been updated with client feedback and clarifications.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- `Prospects/{client-slug}/strategy.md` exists (initial version)
- Notion GTM Client Hub exists with updated strategy content
- Notion MCP server is configured

**Produces (via delegation):**
1. `Prospects/{client-slug}/strategy.md` — Updated with client refinements *(this skill)*
2. `Prospects/{client-slug}/market-mapping.md` — Via `market-mapping` skill
3. `Prospects/{client-slug}/icp-mapping.md` — Via `icp-mapping` skill
4. `Prospects/{client-slug}/workflows/` — Via 3 `workflow-*-trigger` skills
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
   - Note any changes from the original strategy:
     - Confirmed ICP details (firmographics, personas, pain points)
     - Top 3 recommended trigger plays (may have been reordered or modified)
     - New competitive intel or market insights
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

4. **Identify the top 3 trigger plays** from the refined strategy:
   - Extract play names from the strategy's trigger playbook section
   - Map each to its corresponding workflow skill:
     - Funding Events → `workflow-funding-trigger`
     - Job Postings → `workflow-job-posting-trigger`
     - Job Role Changes → `workflow-job-change-trigger`
     - Website Intent → `workflow-website-trigger`
     - Tech Stack Changes → `workflow-tech-change-trigger`
     - Content Engagement → `workflow-content-trigger`
     - Compliance/Regulatory → `workflow-compliance-trigger`
     - Growth/Hiring Surge → `workflow-growth-trigger`
     - Competitor Customer → `workflow-competitor-customer`
     - Champion Tracking → `workflow-champion-tracking`
     - Dark Funnel Signals → `workflow-dark-funnel`
     - General Outbound → `workflow-general-outbound`

### Step 2: Generate Market Mapping

Delegate to the specialized skill:

```
Use the /market-mapping skill
```

This skill will:
- Read from `profile.md` and updated `strategy.md`
- Generate TAM→SAM mapping with boolean search strings
- Create enrichment recommendations and data source discovery
- Output to `Prospects/{client-slug}/market-mapping.md`

**Wait for completion** before proceeding to Step 3.

### Step 3: Generate ICP Mapping

Delegate to the specialized skill:

```
Use the /icp-mapping skill
```

This skill will:
- Read from `profile.md`, `strategy.md`, and `market-mapping.md`
- Define account scoring formula and tier definitions
- Create detailed persona cards with pain mapping
- Build industry-specific pain hierarchies
- Output to `Prospects/{client-slug}/icp-mapping.md`

**Wait for completion** before proceeding to Step 4.

### Step 4: Generate Top 3 Workflow Specifications

For each of the 3 identified trigger plays, delegate to the corresponding workflow skill.

**Create directory first:** `Prospects/{client-slug}/workflows/` if it doesn't exist.

**For each play in order of priority (1, 2, 3):**

```
Use the /workflow-{trigger-type}-trigger skill
```

Example sequence:
1. `/workflow-funding-trigger` → produces `workflows/workflow-1-funding.md`
2. `/workflow-job-posting-trigger` → produces `workflows/workflow-2-job-postings.md`
3. `/workflow-website-trigger` → produces `workflows/workflow-3-website-intent.md`

Each workflow skill will:
- Read from `profile.md`, `strategy.md`, `icp-mapping.md`, `tooling-setup.md`
- Generate Clay table schema with trigger-specific columns
- Create Claygent research prompts and enrichment sequence
- Build copy templates and outreach cadence
- Define scoring, prioritization logic, and success metrics
- Output to `Prospects/{client-slug}/workflows/workflow-{n}-{trigger-name}.md`

**Wait for all 3 workflows to complete** before proceeding to Step 5.

### Step 5: Push All Documents to Notion

Update the client's Notion workspace with all new content:

1. **Update GTM Strategy page** (if local version has additional refinements):
   - Navigate to existing `GTM Strategy - {Company Name}` page
   - Replace content with updated `strategy.md`
   - Preserve any Notion-specific formatting
   - Return the Notion page URL

2. **Update Market Mapping page**:
   - Navigate to or create `Market Mapping` page in the hub
   - Populate with content from `market-mapping.md`
   - Format boolean search strings in code blocks for easy copy-paste
   - Format TAM→SAM tables clearly
   - Return the Notion page URL

3. **Update ICP Mapping page** 🎯:
   - Navigate to the existing `ICP Mapping` page
   - Populate with content from `icp-mapping.md`
   - Format persona cards cleanly (use callout blocks or tables)
   - Format scoring formulas and tier definitions clearly
   - Return the Notion page URL

4. **Update Workflows page** ⚡:
   - Navigate to the existing `Workflows` page
   - Create 3 child pages (or update if they exist):
     - `Workflow 1: {Trigger Name}`
     - `Workflow 2: {Trigger Name}`
     - `Workflow 3: {Trigger Name}`
   - Populate each with content from respective workflow files
   - Use clear headings and formatting for readability
   - Return the Notion page URL

5. **Update Documentation Hub links** on main hub page if needed

### Step 6: Confirm & Output

Provide the user with a complete summary:

**Local files created/updated:**
- ✅ `strategy.md` - Updated with client refinements
- ✅ `market-mapping.md` - TAM→SAM with search criteria (via `/market-mapping`)
- ✅ `icp-mapping.md` - Account scoring + personas (via `/icp-mapping`)
- ✅ `workflows/workflow-1-{name}.md` (via `/workflow-*-trigger`)
- ✅ `workflows/workflow-2-{name}.md` (via `/workflow-*-trigger`)
- ✅ `workflows/workflow-3-{name}.md` (via `/workflow-*-trigger`)

**Notion pages updated:**
- 🎯 GTM Strategy: [URL]
- 📊 Market Mapping: [URL]
- 🎯 ICP Mapping: [URL]
- ⚡ Workflows (with 3 child pages): [URL]

**Key changes from original strategy:**
- [Summary bullet point 1]
- [Summary bullet point 2]
- [Summary bullet point 3]

**Top 3 trigger plays confirmed:**
1. {Play 1 name} - {One-sentence rationale}
2. {Play 2 name} - {One-sentence rationale}
3. {Play 3 name} - {One-sentence rationale}

**Next steps:**
- Review the detailed workflow specifications in Notion or locally
- Proceed with Clay table setup using these specifications
- Begin implementation with the highest-priority trigger play first
- Use `/clay-campaign-generator` to start populating Clay tables

---

## Notes

- **This is an orchestration skill** — it coordinates other skills rather than doing the work inline
- The execution order matters: strategy → market → ICP → workflows (each builds on previous outputs)
- All delegated skills read from the updated `strategy.md`, so pulling from Notion first is critical
- These documents are the "blueprint" for all downstream execution work
- Keep this skill focused on coordination and Notion sync — the specialized skills handle the details
- If any delegated skill fails, stop and report the error before proceeding
