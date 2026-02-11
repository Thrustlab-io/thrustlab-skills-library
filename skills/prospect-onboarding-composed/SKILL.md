---
name: prospect-onboarding-composed
description: End-to-end onboarding for a new Thrustlab prospect. Takes a LinkedIn Company URL and domain as input, generates the GTM strategy (profile.md + strategy.md), creates the Notion workspace, and generates a beautiful HTML presentation to share with the prospect. Composed skill that orchestrates gtm-strategy-generator, notion-project-creator, and html-presentation-generator.
---

# Prospect Onboarding (Composed)

The single entry point for every new Thrustlab prospect engagement. Combines profile discovery, strategy generation, Notion workspace creation, and presentation generation into one seamless workflow.

**Note:** This is a "composed" skill — it orchestrates other specialized skills rather than implementing logic directly.

**Input required:** LinkedIn Company URL + company domain (e.g., `https://www.linkedin.com/company/acme-corp` + `acme.com`)

**Prerequisites:**
- Notion MCP server is configured

**Produces:**
1. `Prospects/{prospect-slug}/profile.md` — the canonical prospect profile
2. `Prospects/{prospect-slug}/strategy.md` — the full GTM strategy
3. `Prospects/{prospect-slug}/presentation.html` — beautiful HTML presentation of the strategy
4. Notion GTM Client Hub — fully populated workspace

---

## Workflow

### Step 1: Generate GTM Strategy

Use the `gtm-strategy-generator` skill to:
- Auto-discover the company profile through web research
- Generate the canonical profile.md
- Create the full GTM strategy document

This will output both `Prospects/{prospect-slug}/profile.md` and `Prospects/{prospect-slug}/strategy.md`.

### Step 2: Create Notion Workspace

Use the `notion-project-creator` skill to:
- Create the GTM Client Hub in Notion
- Populate it with strategy content
- Set up all sub-pages and structure

This will create the complete Notion workspace with the strategy, competitor analysis, roadmap, workflows, and other pages.

### Step 3: Generate Presentation

Use the `html-presentation-generator` skill to:
- Read the strategy.md document
- Extract key insights and recommendations
- Generate a beautiful, modern HTML presentation
- Save it as `Prospects/{prospect-slug}/presentation.html`

This creates a shareable presentation that summarizes the strategy for the prospect.

### Step 4: Present Outputs

Provide the user with:
- Profile summary (company identity, fields needing confirmation)
- Strategy highlights (top 3 recommended trigger plays)
- Notion workspace URL and confirmation of what was populated
- Presentation file path and preview

**Next steps:**
- Review and confirm/correct fields marked `[To be confirmed by client]` in profile.md
- Share the presentation.html with the prospect for review
- Once profile is confirmed, proceed with: `/slack-channel-creator` → `/tooling-setup-guide` → `/market-mapping` → `/icp-mapping` → `/workflow-*`
