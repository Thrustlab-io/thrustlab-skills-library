---
name: client-onboarding
description: Full client onboarding pipeline. Generates GTM strategy, creates Notion workspace, sets up Slack channel, and builds Clay campaign — all in sequence.
argument-hint: "client website [client email]"
---
# Client Onboarding

Runs the complete onboarding pipeline for a new GTM client. Executes these skills in order, passing outputs between them.

## Required Input

- **Client website** (required) — e.g., `https://acme.com`
- **Client email** (optional) — primary contact email for Slack invite

## Pipeline

### Step 1: Generate GTM Strategy
Run `/gtm-strategy-generator [client website]`

This produces `strategies/[CLIENT].md` with ICP, market mapping, messaging, and the Clay Search Plan.

**Wait for completion before proceeding.** Extract the client name from the generated file.

### Step 2: Create Notion Workspace
Run `/notion-project-creator [client name]`

Uses the strategy file to create the full Notion GTM Client Hub. **Save the Notion Phase 1 URL** from the output — it's needed for the Slack welcome message.

### Step 3: Create Slack Channel
Run `/slack-channel-creator [client name] [client email]`

Creates `[client-name]-thrustlab` private channel, invites the Thrustlab team (and client if email provided), and posts the welcome message with the Notion Phase 1 link.

### Step 4: Create Clay Campaign
Run `/clay-campaign-generator strategies/[CLIENT].md`

Creates a Clay workbook and executes all searches from the Clay Search Plan section of the strategy.

## Completion

After all steps, summarize what was created:

```
Client onboarding complete for [Client Name]!

1. GTM Strategy: strategies/[CLIENT].md
2. Notion Hub: [notion URL]
3. Slack Channel: #[client-name]-thrustlab
4. Clay Workbook: [clay workbook URL] ([N] search tables created)
```

## Error Handling

- If any step fails, report the error and ask the user whether to continue with the remaining steps or stop
- Each step depends on the previous one, so do not skip ahead if a step fails
- If Clay auth is not configured, suggest running clay-account-setup first and continue with the other steps
