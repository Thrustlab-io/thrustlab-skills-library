# Thrustlab Skills - Usage Examples

This guide provides detailed examples of how to use Thrustlab GTM skills in real-world scenarios.

## Notion Client Onboarding

### Basic Usage

**Scenario:** You've just signed a new client called "Acme Corp" and need to set up their project workspace.

```
/notion-client-onboarding Acme Corp
```

**What happens:**
1. Claude asks for additional details:
   - Project start date
   - Primary contact
   - GTM focus (what they're launching)

2. Creates the complete workspace structure:
   - Main project hub page
   - 5 pre-configured databases
   - 3 ready-to-use templates
   - Resources page

3. Returns the Notion page URL for immediate access

### Step-by-Step Walkthrough

#### 1. Initial Invocation

**You:**
```
/notion-client-onboarding Acme Corp
```

**Claude:**
```
I'll help set up a complete Notion workspace for Acme Corp.

Let me gather a few details first:

1. When does the engagement start? (e.g., 2026-02-15)
2. Who is the primary client contact?
3. What is the GTM focus? (e.g., new product launch, market expansion, rebrand)
```

#### 2. Provide Details

**You:**
```
1. Start date: 2026-02-10
2. Contact: Sarah Johnson, VP Marketing
3. GTM focus: Launching their new SaaS platform for enterprise customers
```

**Claude:**
```
Perfect! Setting up the Notion workspace now...

✓ Created main project page: "Acme Corp - GTM Project"
✓ Created GTM Strategy & Planning database (15 pre-populated tasks)
✓ Created Competitive Intelligence database
✓ Created Customer Personas database
✓ Created Content & Deliverables database
✓ Created Meeting Notes database
✓ Created Weekly Status Report template
✓ Created GTM Strategy Document template
✓ Created Launch Checklist template
✓ Created Resources & Links page

Your workspace is ready! Access it here: [Link to Notion page]

Next steps:
- Review and customize the pre-populated tasks
- Begin competitive research
- Schedule kickoff meeting and use the meeting notes template
```

#### 3. What Gets Created

**Main Page Structure:**
```
📊 Acme Corp - GTM Project
├── Overview (with client details)
├── Quick Links
│   ├── Strategy & Planning →
│   ├── Competitive Intel →
│   ├── Personas →
│   ├── Content →
│   └── Meetings →
├── 🎯 GTM Strategy & Planning (Database)
├── 📊 Competitive Intelligence (Database)
├── 👥 Customer Personas (Database)
├── 📝 Content & Deliverables (Database)
├── 🗓️ Meeting Notes (Database)
├── Templates
│   ├── Weekly Status Report
│   ├── GTM Strategy Document
│   └── Launch Checklist
└── Resources & Links
```

### Real-World Scenarios

#### Scenario 1: Enterprise SaaS Launch

**Client:** TechStart Inc.
**Focus:** Launching new AI-powered analytics platform

```
/notion-client-onboarding TechStart Inc.
```

**Customize for this use case:**
- Add competitors: Tableau, Power BI, Looker
- Create personas: Data Analysts, CTO, Business Intelligence Manager
- Pre-populate content: Product demos, technical white papers, ROI calculators
- Channel focus: LinkedIn, industry events, analyst relations

#### Scenario 2: Market Expansion

**Client:** GlobalRetail Co.
**Focus:** Expanding into European markets

```
/notion-client-onboarding GlobalRetail Co.
```

**Customize for this use case:**
- Add regional competitors per market
- Create geo-specific personas
- Content needs: Localized marketing materials, regional case studies
- Channel focus: Local events, partnerships, PR

#### Scenario 3: Product Rebranding

**Client:** Legacy Software Ltd.
**Focus:** Complete rebrand and repositioning

```
/notion-client-onboarding Legacy Software Ltd.
```

**Customize for this use case:**
- Document current positioning vs. new
- Track all assets needing updates
- Manage stakeholder communications
- Timeline: Phased rollout plan

## Advanced Tips

### Customizing After Setup

After the initial setup, you can ask Claude to customize:

```
Add a "Partner Ecosystem" database to the Acme Corp project with fields for:
- Partner name, tier, region, status, contact info
```

```
Update the Content database to add a new content type: "Customer testimonial"
```

### Batch Operations

Setting up multiple clients:

```
I need to onboard 3 new clients this week. Let's start with:
1. /notion-client-onboarding ClientA
2. Then ClientB
3. Then ClientC

I'll give you details for each one at a time.
```

### Integration with Other Workflows

**After client onboarding, typical next steps:**

1. **Schedule kickoff meeting**
   ```
   Add a new meeting note for Acme Corp kickoff on Feb 10, 2026
   ```

2. **Start competitive research**
   ```
   Help me research top 5 competitors for enterprise SaaS analytics platforms
   and add them to the Acme Corp competitive database
   ```

3. **Create personas**
   ```
   Based on Acme Corp's target market (enterprise SaaS buyers),
   help me create 3 customer personas in their database
   ```

4. **Plan content**
   ```
   Generate a content plan for Acme Corp's launch with 10 key deliverables
   and add them to the Content database
   ```

## Troubleshooting

### Issue: Notion MCP not connected

**Error:**
```
I don't have access to Notion. Please ensure the Notion MCP server is installed.
```

**Solution:**
1. Check Claude Desktop config has Notion MCP configured
2. Verify your Notion API key is valid
3. Restart Claude Desktop
4. Try again

### Issue: Missing permissions

**Error:**
```
I cannot create pages in this workspace. Permission denied.
```

**Solution:**
1. Ensure your Notion integration has the right permissions
2. Share the parent page with your integration
3. Check that you're using the correct workspace

### Issue: Want to modify structure

**Question:** "Can I customize the databases before creation?"

**Answer:** Currently, the skill creates a standard structure. However, you can:
1. Let it create the full structure
2. Then ask Claude to modify specific parts
3. Or fork the skill and customize the SKILL.md file

## Best Practices

### 1. Consistent Naming
Always use the official client name exactly as it appears in contracts:
```
✓ /notion-client-onboarding Acme Corporation
✗ /notion-client-onboarding acme
```

### 2. Prepare Information
Have these ready before starting:
- Exact start date
- Primary contact name and role
- Clear GTM focus description

### 3. Review Before Customizing
- Let Claude create the full structure first
- Review what's been created
- Then customize based on specific client needs

### 4. Use Templates Immediately
- First meeting? Use the meeting notes template
- Weekly update? Use the status report template
- Strategy session? Use the strategy doc template

### 5. Keep It Updated
- Regular database maintenance
- Archive completed items
- Update competitor info monthly
- Refresh persona insights quarterly

---

Have more questions or examples to share? [Open an issue](https://github.com/your-org/thrustlab/issues) or contribute to this guide!
