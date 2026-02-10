# Thrustlab GTM — Claude Skills & MCP Servers

A collection of Claude skills and MCP servers for Go-To-Market operations, client onboarding, and outreach automation.

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/Thrustlab-io/thrustlab-skills-library/main/install.sh | bash
```

This installs 22 skills to `~/.claude/skills/` and 4 MCP server binaries (pre-built universal macOS binaries) to `~/.claude/bin/`.

### Configure MCP Credentials

After install, add your API credentials:

```bash
# Slack
claude mcp add -e SLACK_BOT_TOKEN="xoxb-..." -s user slack -- ~/.claude/bin/slack-mcp

# Clay
claude mcp add -e CLAY_WORKSPACE_ID="..." -e CLAY_SESSION_COOKIE="..." -s user clay -- ~/.claude/bin/clay-mcp-server

# Namecheap
claude mcp add -e NAMECHEAP_API_USER="..." -e NAMECHEAP_API_KEY="..." -e NAMECHEAP_USERNAME="..." -e NAMECHEAP_CLIENT_IP="..." -s user namecheap -- ~/.claude/bin/namecheap-mcp

# Premium Inboxes
claude mcp add -e PREMIUMINBOXES_API_TOKEN="..." -s user premiuminboxes -- ~/.claude/bin/premiuminboxes-mcp
```

Restart Claude Desktop after configuring.

## Skills

### Foundation
| Skill | Purpose |
|-------|---------|
| `/client-onboarding` | Create canonical client profile from intake data |
| `/gtm-strategy-generator` | Generate 13-section GTM strategy with Clay search plan |
| `/tooling-setup-guide` | Infrastructure checklist for signal sources |
| `/market-mapping` | Build market research and account lists |
| `/icp-mapping` | Create persona x vertical matrix |

### Admin & Integrations
| Skill | Purpose |
|-------|---------|
| `/notion-project-creator` | Create Notion workspace with 14 pages from strategy |
| `/slack-channel-creator` | Create private Slack channel with team invites |
| `/clay-campaign-generator` | Create Clay workbook and search tables from strategy |
| `/clay-account-setup` | Guide Clay.com account creation |

### Trigger-Based Workflows
| Skill | Signal |
|-------|--------|
| `/workflow-website-trigger` | Website visitor intent (RB2B/Dealfront) |
| `/workflow-job-change-trigger` | Person changes roles (UserGems/Clay) |
| `/workflow-job-posting-trigger` | Company posts relevant job |
| `/workflow-funding-trigger` | Company raises funding |
| `/workflow-tech-change-trigger` | Company adds/removes tech (BuiltWith) |
| `/workflow-growth-trigger` | Headcount surge or expansion |
| `/workflow-compliance-trigger` | Regulatory deadline approaching |
| `/workflow-content-trigger` | Prospect engages with content |
| `/workflow-champion-tracking` | Former customers change jobs |
| `/workflow-competitor-customer` | Companies using competitor products |
| `/workflow-dark-funnel` | Anonymous engagement signals |
| `/workflow-general-outbound` | Research-based outreach (no trigger) |

### Execution Order

```
/client-onboarding → /gtm-strategy-generator → /notion-project-creator → /slack-channel-creator → /tooling-setup-guide → /market-mapping → /icp-mapping → /workflow-*
```

## MCP Servers

| Server | Tools | Auth |
|--------|-------|------|
| Clay | create_workbook, search_companies_by_industry, search_businesses_by_geography | CLAY_WORKSPACE_ID + CLAY_SESSION_COOKIE |
| Slack | create_channel, lookup_user, invite_users, send_message | SLACK_BOT_TOKEN |
| Namecheap | check_domain, create_domain, list_domains | NAMECHEAP_API_KEY + user/ip |
| Premium Inboxes | get/create/cancel subscriptions and orders | PREMIUMINBOXES_API_TOKEN |

## Development

### Releasing

Tag a version to trigger the build:

```bash
git tag v0.1.0
git push origin v0.1.0
```

GitHub Actions will build universal macOS binaries for all MCP servers and attach them to the release.

### Adding a Skill

Create `skills/<skill-name>/SKILL.md`:

```yaml
---
name: skill-name
description: What the skill does
---

# Instructions for Claude...
```

## License

MIT
