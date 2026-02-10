---
name: slack-channel-creator
description: Creates a private Slack channel for a Thrustlab client, invites the team, and posts a welcome message with Notion workspace links. Use after the Notion workspace has been created. Requires Slack MCP server.
---

# Slack Channel Creator

Creates a private Slack channel and posts a welcome message linking to the Notion workspace.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- Notion workspace has been created (need Phase 1 page URL)
- Slack MCP server is configured with bot token

## Workflow

### Step 1: Load Client Data

Read `Prospects/{client-slug}/profile.md` for:
- Company name (for channel naming)
- Client contact email (if available, for Slack invite)

Get from user if not already provided:
- Notion hub page URL
- Notion Phase 1: Infrastructure Setup page URL
- Client contact email for Slack invite (optional)

### Step 2: Look Up Slack Users

Use Slack `lookup_user` for:
- `kwinten@thrustlab.io` (always invited)
- `jan@thrustlab.io` (always invited)
- Client contact email (if provided)

### Step 3: Create Channel

Use Slack `create_channel`:
- **Name format:** `{client-slug}-thrustlab` (lowercase, hyphens, no spaces)
  - "Acme Corp" → `acme-corp-thrustlab`
  - "Quality Guard" → `quality-guard-thrustlab`
- **is_private:** `true` (always private for client channels)

### Step 4: Invite Users

Add all looked-up users to the channel.

### Step 5: Post Welcome Message

```
Welcome to your Thrustlab collaboration channel! 👋

This is your direct line to our team for all GTM strategy and execution.

📋 **Your GTM Hub:** {Notion Hub URL}
🏗️ **Phase 1 Setup:** {Notion Phase 1 URL}

**Your Thrustlab Team:**
• Kwinten (kwinten@thrustlab.io)
• Jan (jan@thrustlab.io)

Next step: We'll share your tooling setup guide here shortly. Feel free to ask questions anytime — let's build pipeline! 🚀
```

### Step 6: Confirm & Output

Provide to user:
- Slack channel link
- Confirmation of who was invited
- Reminder that tooling-setup-guide output should be shared in this channel
