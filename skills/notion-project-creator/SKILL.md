---
name: notion-project-creator
description: Creates a Notion client workspace with the complete GTM Client Hub structure. Use when a client profile and strategy exist and the Notion workspace needs to be set up. Reads from profile.md and strategy.md to populate the workspace with client-specific content. Requires Notion MCP server.
---

# Notion Project Creator

Creates a fully populated Notion workspace for a new GTM client.
Make sure to create this page under the "Thrustlab" teamspace!

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- Notion MCP server is configured

## Workflow

### Step 1: Load Client Data

Read these files and extract key information:
1. `Prospects/{client-slug}/profile.md` — company name, website, product description, personas, value prop

### Step 2: Create Main Hub Page

Create a parent page titled: `🎯 {Company Name} - GTM Client Hub`
Create an entry on the "Companies" page in the "Prospect" column: https://www.notion.so/quantascale/Companies-266fbebd816080549bfcccc0dee598b3

Include these sections on the main page:

#### Braindump & Quick Notes
A callout block with empty checkboxes — an empty list

#### To Discuss with Client:
A callout block with empty checkboxes — an empty list

#### Documentation Hub
A collapsible callout with links to all sub-pages created below.
Make sure to update these links as you create the sub-pages in Step 3.

### Step 3: Create Sub-Pages 

Create each as a child page of the main hub. For the strategy page, populate with actual strategy content from `Prospects/{client-slug}/strategy.md`. All other pages get their template structure:

1. **GTM Strategy - {Company Name}** 🎯 — Populate with full strategy.md content
2. **Competitor Analysis - {Company Name}** 📊 — Pre-fill with competitors from profile.md
3. **Infrastructure** 📋 — Domains and tools used
4. **ICP Mapping** 🎯 — Empty, populated during ICP mapping phase
5. **Roadmap** 🗓️ — Pre-fill with 90-day blueprint from strategy
6. **Meeting Notes** 📅 — Empty template with date + attendees + notes structure
7. **Workflows** ⚡ — Pre-fill with trigger playbook from strategy
8. **Copy Repository** 💬 — Will store approved copy as cadences are built

Then update the Documentation Hub section on the main page with links to each of these sub-pages.

### Step 4: Confirm & Output

Provide to the user:
- Main hub page URL
- Infrastructure page URL (needed for Slack channel creator)
- Confirmation that strategy content is populated
- Summary of what was pre-filled vs. what will be populated in later phases
