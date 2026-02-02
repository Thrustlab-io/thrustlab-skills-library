# Thrustlab GTM Claude Skills Library

A curated collection of Claude Desktop skills designed to streamline Go-To-Market (GTM) operations, client onboarding, and strategic planning workflows.

## 📚 Available Skills

### Notion Client Onboarding
Automates the setup of a complete Notion workspace for new GTM clients, including:
- 5 pre-configured databases (Strategy, Competitive Intel, Personas, Content, Meeting Notes)
- 3 ready-to-use templates (Status Reports, Strategy Docs, Launch Checklists)
- Standardized project structure for consistent client delivery

**Usage:** `/notion-client-onboarding [client name]`

**Requirements:** [Notion MCP Server](https://github.com/modelcontextprotocol/servers/tree/main/src/notion)

### Clay Client Workspace
Guides setup of a Clay.com workspace for new GTM clients with structured tables for prospect research, data enrichment, and outreach automation:
- 5 pre-configured Clay tables (Target Accounts, Contacts, Outreach, Competitive Intel, ICP Analysis)
- Automated enrichment workflows and trigger-based alerts
- AI-powered research with Claygent prompts
- CRM and email tool integrations
- Thrustlab team admin access for ongoing support

**Usage:** `/clay-client-workspace [client name]`

**Requirements:** Clay.com account (guided setup via Thrustlab referral link)

## 🚀 Installation

### Quick Install (Recommended)

```bash
curl -fsSL https://raw.githubusercontent.com/your-org/thrustlab/main/install.sh | bash
```

### Manual Installation

**For Claude Desktop:**
```bash
# Clone the repository
git clone https://github.com/your-org/thrustlab.git

# Copy skills to your Claude skills directory
cp -r thrustlab/skills/* ~/.claude/skills/
```

**For Claude Code (CLI):**
```bash
# Clone the repository
git clone https://github.com/your-org/thrustlab.git

# Copy skills to your project
cp -r thrustlab/skills/* .claude/skills/
```

### Development Setup

If you want to contribute or develop locally:

```bash
# Clone and symlink for live updates
git clone https://github.com/your-org/thrustlab.git
ln -s "$(pwd)/thrustlab/skills"/* ~/.claude/skills/
```

## 📋 Prerequisites

### Required MCP Servers

Some skills require specific MCP (Model Context Protocol) servers to be installed:

#### Notion MCP Server
Required for: `notion-client-onboarding`

**Installation:**
1. Add to your Claude Desktop config at `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "notion": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-notion"
      ],
      "env": {
        "NOTION_API_KEY": "your-notion-integration-token"
      }
    }
  }
}
```

2. Get your Notion API key:
   - Go to https://www.notion.so/my-integrations
   - Create a new integration
   - Copy the "Internal Integration Token"
   - Share relevant Notion pages with your integration

3. Restart Claude Desktop

#### Clay.com Account
Required for: `clay-client-workspace`

**Setup:**
The skill guides you through setting up a Clay.com workspace via their web interface:

1. Create account using Thrustlab referral link (included in skill)
2. Add Thrustlab team members as workspace admins
3. Configure tables, enrichments, and automations
4. Connect CRM and email tool integrations

**Note:** Clay.com doesn't have programmatic API access, so the skill provides step-by-step guided setup instructions

## 💡 Usage Examples

### Setting Up a New Client Project

```
/notion-client-onboarding Acme Corp
```

Claude will guide you through:
1. Gathering client details (start date, contact, GTM focus)
2. Creating the complete Notion workspace structure
3. Pre-populating databases with standard tasks
4. Setting up templates and resources

### Typical Workflow

**Complete Client Onboarding:**
1. New client signed → Run `/notion-client-onboarding [Client Name]`
2. Notion workspace created with all databases and templates
3. Run `/clay-client-workspace [Client Name]`
4. Follow guided setup to create Clay.com workspace with Thrustlab referral
5. Add Thrustlab team as admins to Clay workspace
6. Configure 5 core tables with enrichment and automation
7. Begin prospect research and ICP validation in Clay
8. Use Notion for strategy, deliverables, and client reporting
9. Leverage Clay for data enrichment and automated outreach

## 🛠️ Skill Development

Want to add your own GTM skills? Here's the structure:

```
skills/
└── your-skill-name/
    └── SKILL.md
```

**SKILL.md format:**
```yaml
---
name: your-skill-name
description: What your skill does
argument-hint: "expected arguments"
---

# Your Skill Instructions

Detailed instructions for Claude on how to execute this skill...
```

See [skills/notion-client-onboarding/SKILL.md](skills/notion-client-onboarding/SKILL.md) for a complete example.

## 🤝 Contributing

We welcome contributions! To add a new skill:

1. Fork this repository
2. Create a new skill directory under `skills/`
3. Add your `SKILL.md` file with proper frontmatter
4. Test the skill locally
5. Submit a pull request with:
   - Skill description
   - Usage examples
   - Any prerequisites (MCP servers, etc.)

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

## 🔗 Resources

- [Claude Code Documentation](https://code.claude.com)
- [Model Context Protocol (MCP)](https://modelcontextprotocol.io)
- [Notion API Documentation](https://developers.notion.com)

## 🆘 Support

Found a bug or have a feature request? [Open an issue](https://github.com/your-org/thrustlab/issues)

---

Built with ❤️ for GTM teams using Claude Desktop
