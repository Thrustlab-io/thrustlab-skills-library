---
name: slack-channel-creator
description: Creates a Slack channel with the client, invites Thrustlab team, and posts a welcome message with Notion workspace link
argument-hint: "client name [client email]"
---

# Slack Channel Creator Skill

Creates a client collaboration Slack channel, invites the team and client, and posts a welcome message linking to the Notion Phase 1 page.

## Prerequisites

**REQUIRED**: The `notion-project-creator` skill must have been run first. This skill requires:
- The Notion workspace structure to exist
- Access to the "Phase 1: Infrastructure Setup - [Client Name]" page URL
- Slack MCP server configured with bot token

## Required Information

1. **Client name** - The name of the client/company (e.g., "Acme Corp")
2. **Client email** - Primary client contact email address (optional - can invite later)
3. **Notion Phase 1 URL** - Link to the "Phase 1: Infrastructure Setup - [Client Name]" page

## Execution Steps

### Step 1: Verify Prerequisites

Confirm that the Notion workspace was created and get the Phase 1 page URL.

If not available:
- Tell user to run `/notion-project-creator [client-name]` first
- Do not proceed without the Notion workspace

### Step 2: Look Up Users

Use `lookup_user` tool to find Slack user IDs for:
- `kwinten@thrustlab.io` (always invited)
- `jan@thrustlab.io` (always invited)
- Client contact email (if provided)

If a user is not found, ask user if they want to proceed and invite manually later.
If no client email provided, proceed with just Thrustlab team.

### Step 3: Create Channel

Use `create_channel` tool with:
- **Name format:** `[client-name]-thrustlab`
  - Lowercase, hyphens instead of spaces
  - Example: "Acme Corp" → `acme-corp-thrustlab`
  - Example: "DataMeister" → `datameister-thrustlab`
- **is_private:** true (always private for client channels)

**Save the channel ID** from the response.

### Step 4: Invite Users

Use `invite_users` tool to add (comma-separated user IDs):
- Kwinten (kwinten@thrustlab.io)
- Jan (jan@thrustlab.io)
- Client contact (if provided)

### Step 5: Post Welcome Message

Use `send_message` tool with this template:

```
Welcome to your Thrustlab collaboration channel! 👋

This is your direct line to our team for all GTM strategy and execution.

**Getting Started:**
📋 Phase 1 Infrastructure Setup: [Notion Phase 1 URL]

**Your Thrustlab Team:**
• Kwinten (kwinten@thrustlab.io)
• Jan (jan@thrustlab.io)

Feel free to ask questions anytime. Let's build something great together! 🚀
```

Replace `[Notion Phase 1 URL]` with the actual URL from the Notion workspace.

### Step 6: Confirm Completion

Output:
```
✅ Slack channel created!

📍 Channel: #[client-name]-thrustlab
🔗 Link: https://[workspace].slack.com/archives/[channel-id]
🔒 Private channel

👥 Invited:
• Kwinten (kwinten@thrustlab.io)
• Jan (jan@thrustlab.io)
• [Client Name] ([client-email]) - if provided

💬 Welcome message posted with Notion link
```

If no client was invited, remind user they can invite the client manually later.

## Channel Naming Examples

- "Acme Corp" → `acme-corp-thrustlab`
- "DataMeister" → `datameister-thrustlab`
- "O2O Bicycle Leasing" → `o2o-bicycle-leasing-thrustlab`

## Error Handling

**User not found:** Ask if they want to proceed without them (can invite manually)
**Channel exists:** Suggest adding `-2` or `-new` suffix
**No Notion workspace:** Direct user to run `/notion-project-creator` first

## Integration Workflow

Recommended order:
```bash
# 1. Generate strategy
/gtm-strategy-generator https://client.com

# 2. Create Notion workspace
/notion-project-creator "Client Name"

# 3. Create Slack channel (this skill)
/slack-channel-creator "Client Name" client@email.com
```

## Success Criteria

✅ Channel created as `[client-name]-thrustlab`
✅ Private visibility
✅ Kwinten and Jan invited
✅ Client invited
✅ Welcome message posted with Notion Phase 1 link
✅ User receives channel link
