---
name: client-activation
description: Activates a new signed client by setting up all communication and collaboration channels. Creates a shared Google Drive folder and private Slack channel with team access. Use when a prospect signs the contract and transitions to active client status. Requires Google Drive and Slack MCP servers.
---

# Client Activation

Transitions a prospect to active client by setting up all necessary communication channels and collaboration infrastructure.

**When to use:** Immediately after a prospect signs the contract and becomes an active client.

**Prerequisites:**
- `Prospects/{client-slug}/` directory exists with profile.md and strategy.md
- Notion GTM Client Hub exists
- Google Drive MCP server is configured
- Slack MCP server is configured

**Produces:**
1. Shared Google Drive folder with initial document structure
2. Private Slack channel with team members and welcome message
3. Updated Notion workspace with collaboration links

---

## Workflow

### Step 1: Load Client Data

Read client information from existing files:
1. `Prospects/{client-slug}/profile.md` — Company name, domain, team contacts
2. `Prospects/{client-slug}/strategy.md` — Top 3 trigger plays, timeline, milestones
3. Notion workspace URL (prompt user if not stored locally)

Get from user:
- Client primary contact email (for Slack invite)
- Client primary contact name
- Any additional team members to invite (emails)
- Contract start date (defaults to today)

### Step 2: Create Shared Google Drive Folder

Use Google Drive MCP server to create the folder structure:

**Main folder name:** `[Client] {Company Name} - Thrustlab GTM`

**Sub-folders to create:**
1. **01 - Strategy & Planning**
   - Purpose: Strategy documents, ICP mapping, market analysis
   - Initial files: Export and upload strategy.md, profile.md

2. **02 - Workflows & Copy**
   - Purpose: Workflow specifications, email copy, messaging frameworks
   - Initial files: Empty (populated during workflow creation)

3. **03 - Data & Lists**
   - Purpose: Account lists, enrichment exports, Clay table exports
   - Initial files: Empty (populated during execution)

4. **04 - Reporting**
   - Purpose: Weekly reports, performance dashboards, optimization notes
   - Initial files: Create a README with reporting schedule

5. **05 - Meeting Notes**
   - Purpose: Agendas, meeting recordings, action items
   - Initial files: Create first meeting agenda template

**Permissions:**
- Share folder with:
  - `kwinten@thrustlab.io` (Editor)
  - `jan@thrustlab.io` (Editor)
  - Client primary contact email (Editor)
  - Any additional client team members (Commenter or Editor based on role)

**Output:** Drive folder URL

### Step 3: Create Private Slack Channel

Use the Slack MCP server to set up the collaboration channel:

**Channel name format:** `client-{client-slug}`
- Example: "Acme Corp" → `client-acme-corp`
- Note: Different from prospect channels which use `{client-slug}-thrustlab`

**Channel settings:**
- **Type:** Private channel
- **Purpose:** "Thrustlab GTM execution for {Company Name} - Strategy, workflows, and performance"

**Invite team members:**
- `kwinten@thrustlab.io` (always)
- `jan@thrustlab.io` (always)
- Client primary contact (from user input)
- Additional client team members (if provided)

**Welcome message:**

```
🎉 Welcome to your Thrustlab GTM workspace!

Excited to partner with you on building predictable pipeline. This channel is your direct line to our team for strategy, execution, and optimization.

**📋 Quick Links:**
• GTM Strategy Hub: {Notion Hub URL}
• Shared Drive: {Google Drive URL}

**👥 Your Thrustlab Team:**
• Kwinten (kwinten@thrustlab.io) - Strategy & Execution
• Jan (jan@thrustlab.io) - Operations & Support

**🗓️ What's Next:**
Week 1: Tooling setup (Clay, enrichment tools, sequencer)
Week 2-3: Market mapping & ICP refinement
Week 4: First workflow launch

We'll share detailed setup guides and schedule our kickoff call here. Let's get started! 🚀
```

**Output:** Slack channel URL

### Step 4: Update Notion with Collaboration Links

Update the main GTM Client Hub page:

1. **Add a "Collaboration" section** (callout or section at top):
   ```
   🤝 Collaboration Hub
   • Slack Channel: {Slack channel URL}
   • Shared Drive: {Google Drive URL}
   • Team: Kwinten, Jan @ Thrustlab
   ```

2. **Update the Infrastructure page** with:
   - Slack channel name and URL
   - Google Drive folder URL
   - Access instructions for client team

### Step 5: Send Setup Summary to Slack

Post a summary message in the new Slack channel:

```
✅ **Client Activation Complete!**

Your Thrustlab workspace is ready. Here's what we've set up:

**📁 Shared Drive:** {Drive URL}
└─ Strategy docs, workflows, data, and reporting folders

**📋 Notion Workspace:** {Notion URL}
└─ GTM strategy, workflows, and ICP mapping

**Next Steps:**
1. Schedule kickoff call (Calendly link or propose times)
2. Begin tooling setup (setup guide coming shortly)
3. Review strategy and confirm priorities

Questions? Drop them here anytime. Let's build pipeline! 🚀
```

### Step 6: Confirm & Output

Provide the user with a complete activation summary:

**Created:**
- ✅ Google Drive folder: {URL}
  - 5 sub-folders with initial structure
  - Shared with {list of people}
- ✅ Slack channel: {URL}
  - Private channel: `client-{client-slug}`
  - Members: {list of people}

**Updated:**
- ✅ Notion GTM Client Hub with collaboration links
- ✅ Infrastructure page with access details

**Client status:** Prospect → Active Client ✨

**Next actions for Thrustlab:**
1. Schedule kickoff call with client
2. Run `/tooling-setup-guide` and share in Slack
3. Begin tooling and infrastructure setup
4. Update internal tracker with client start date

---

## Notes

- This skill marks the official transition from prospect to active client
- All collaboration infrastructure is centralized (Slack, Drive, Notion)
- Client can now access all resources and collaborate with the team directly
- This skill can be run immediately after contract signature
- The Drive folder structure provides a clear home for all project artifacts
- Roadmap and detailed planning can be handled separately as needed
