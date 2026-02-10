---
name: tooling-setup-guide
description: Generates a client-specific tooling setup guide based on the GTM strategy. Use after the strategy has been generated. Reads the strategy to determine which tools the client needs based on their recommended trigger plays. Produces step-by-step instructions for Clay, email sequencer, intent/signal tools, and integrations — the client sets these up themselves with our guidance.
---

# Tooling Setup Guide

Generates a comprehensive, client-specific tooling setup document saved to `Prospects/{client-slug}/tooling-setup.md`.

The client pays for their own tooling. This skill produces the instructions — not automated creation.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- `Prospects/{client-slug}/strategy.md` exists (specifically: top 3 trigger plays, ICP scoring criteria, cadence structure)

## Workflow

### Step 1: Load Strategy & Profile

Read both files. Extract:
- **From strategy.md:** Top 3 recommended trigger plays, ICP scoring formula, cadence structure, enrichment needs
- **From profile.md:** Existing tech stack, CRM, budget/plan tier, team members

### Step 2: Map Trigger Plays → Required Tooling

For each of the 3 recommended trigger plays, determine what tools are needed.
See `references/trigger-tooling-map.md` for the complete mapping.

### Step 3: Generate Setup Document

Write `Prospects/{client-slug}/tooling-setup.md` using the structure in `references/setup-document-template.md`.

The document must be:
- Specific to THIS client's trigger plays (don't include tools they don't need)
- Ordered by priority (Clay first, then signal tools for trigger play #1, etc.)
- Include exact configuration steps, not just "set up Clay"
- Reference the strategy for WHY each tool is needed

### Step 4: Output & Share

Save to `Prospects/{client-slug}/tooling-setup.md`.

Suggest sharing in the client's Slack channel as the next action item.
