---
name: clay-client-workspace
description: Guides setup of a Clay.com workspace for a new GTM client with structured tables for prospect research, account intelligence, and outreach automation
argument-hint: "client name"
---

# Clay.com Client Workspace Setup Skill

You are helping guide the setup of a complete Clay.com workspace for a new GTM client. This skill provides step-by-step instructions for creating a standardized workspace structure for prospect research, data enrichment, and automated outreach.

**Important**: Clay.com does not have programmatic API access for workspace creation, so this skill guides the user through manual setup with best practices and templates.

## Required Information

Before proceeding, gather from the user:
1. **Client name** - The name of the client/company
2. **Target market** - Who they're selling to (ICP details: industry, company size, roles)
3. **GTM focus** - What they're launching or expanding
4. **Key team members** - Who from your team needs workspace access
5. **Workspace owner email** - The client's email for workspace ownership

## Step 1: Clay.com Account Setup

Guide the user through the following steps:

### A. Create Clay Account (with Referral)

Provide these instructions:

```
1. Visit Clay.com via Thrustlab's referral link:
   https://clay.com/?via=thrustlab

2. Sign up with the CLIENT's email address (they should own the workspace)

3. Choose the appropriate plan based on client needs:
   - Starter: $149/mo - Best for small teams testing Clay
   - Explorer: $349/mo - Most popular for growing GTM teams
   - Pro: $800/mo - For scaled outbound operations

4. Complete account setup and verify email

5. You'll receive 500 free Clay credits to get started
```

### B. Add Thrustlab Team to Workspace

After account creation, guide the user to add your team:

```
1. In Clay, click your workspace name (top left)

2. Go to "Settings" → "Members"

3. Add Thrustlab team members:
   - [ADD YOUR TEAM EMAIL HERE] (Admin role)
   - [ADD ADDITIONAL TEAM EMAILS] (Member role)

4. Set appropriate permissions:
   - Admin: Full workspace access (for your core team)
   - Member: Can create and edit tables (for collaborators)
   - Viewer: Read-only access (for stakeholders)
```

## Step 2: Workspace Structure Setup

Guide the user to create the following Clay tables:

### Table 1: Target Accounts Database

**Purpose**: Master list of target companies for this client's GTM motion

**Instructions**:
```
1. Click "New Table" → "Start from scratch"

2. Name: "[Client Name] - Target Accounts"

3. Set up these columns:
   - Company Name (default)
   - Website (URL)
   - Industry (Select)
   - Company Size (Select: 1-50, 51-200, 201-1000, 1000+)
   - Revenue Range (Select: <$1M, $1M-$10M, $10M-$50M, $50M+)
   - Location/HQ (Text)
   - Tech Stack (Text)
   - Funding Stage (Select: Bootstrapped, Seed, Series A, B, C+)
   - Account Score (Number: 1-10)
   - Status (Select: Prospecting, Engaged, Qualified, Customer)
   - Notes (Text)

4. Enable enrichment providers:
   - Add "Enrich Company" integration (Clearbit, Builtwith, etc.)
   - Add "Find Similar Companies" for lookalike expansion
   - Add "Company News & Signals" for trigger events
```

### Table 2: Contact Database

**Purpose**: Individual prospects and decision-makers at target accounts

**Instructions**:
```
1. Click "New Table" → "Find people at companies"

2. Name: "[Client Name] - Contacts"

3. Set up these columns:
   - Full Name (default)
   - Email (Email - with verification)
   - Phone (Phone)
   - Job Title (Text)
   - Department (Select: Sales, Marketing, Product, Engineering, Executive)
   - Seniority (Select: IC, Manager, Director, VP, C-Level)
   - Company (Relation to Target Accounts table)
   - LinkedIn URL (URL)
   - Location (Text)
   - Recent Activity (Text)
   - Engagement Score (Number: 1-10)
   - Status (Select: New, Contacted, Replied, Qualified, Lost)
   - Last Contacted (Date)
   - Notes (Text)

4. Enable enrichment:
   - Add "Waterfall Email Enrichment" for verified emails
   - Add "Find Mobile Phone" for direct dials
   - Add "LinkedIn Profile Enrichment" for background info
   - Add "Recent Job Changes" for timing signals
```

### Table 3: Outreach Sequences

**Purpose**: Personalized outreach campaigns with AI-generated messaging

**Instructions**:
```
1. Click "New Table" → "Import from contacts"

2. Name: "[Client Name] - Outreach Campaigns"

3. Link to Contact Database (import relevant contacts)

4. Set up these columns:
   - Contact Name (from Contacts)
   - Email (from Contacts)
   - Campaign Name (Select: Cold Outreach, Warm Intro, Follow-up, Re-engagement)
   - Personalization Data (Text - research findings)
   - Email Subject Line (Text)
   - Email Body (Long text)
   - Send Status (Select: Draft, Ready, Sent, Replied)
   - Send Date (Date)
   - Reply Received (Checkbox)
   - Reply Content (Text)

5. Add Claygent AI Research:
   - Add "Claygent" column to research each prospect
   - Prompt: "Research {{Company}} and {{Full Name}}. Find recent company news,
     their role responsibilities, and 3 specific pain points our client's
     solution could address. Be specific and cite sources."

6. Add AI Personalization:
   - Add "Generate Email" column using GPT-4
   - Prompt: "Write a personalized cold email to {{Full Name}} at {{Company}}.
     Reference this research: {{Claygent Research}}. Keep it under 100 words,
     focus on their pain points, and include a specific call-to-action."
```

### Table 4: Competitive Intelligence

**Purpose**: Track competitors and market positioning

**Instructions**:
```
1. Click "New Table" → "Start from scratch"

2. Name: "[Client Name] - Competitive Intel"

3. Set up these columns:
   - Competitor Name (Title)
   - Website (URL)
   - Positioning (Text)
   - Key Features (Text)
   - Pricing Model (Text)
   - Target Customers (Text)
   - Strengths (Text)
   - Weaknesses (Text)
   - Recent News (Text)
   - Market Share Estimate (Text)
   - Last Updated (Date)

4. Add enrichment:
   - "Company Overview" for basic info
   - "G2 Reviews" for customer sentiment
   - "Tech Stack" to see their tools
   - "Recent Funding" for financial health
```

### Table 5: ICP Research & Validation

**Purpose**: Build and refine ideal customer profile

**Instructions**:
```
1. Click "New Table" → "Find companies"

2. Name: "[Client Name] - ICP Analysis"

3. Import 50-100 sample companies matching initial ICP hypothesis

4. Set up analysis columns:
   - Company Name (default)
   - Firmographic Data (size, industry, revenue)
   - Conversion Likelihood (Number: 1-10)
   - Fit Reason (Text - why they match ICP)
   - Disqualifiers (Text - red flags)
   - Win Pattern (Text - common traits of customers)

5. Use Claygent to analyze patterns:
   - "What common characteristics do top-scored companies share?"
   - "What distinguishes high-fit vs low-fit accounts?"
```

## Step 3: Automation & Workflows

Guide the user to set up these automated workflows:

### Workflow 1: New Account Enrichment

```
1. Go to any table → "Automations" tab

2. Create automation:
   - Trigger: "When row is added to Target Accounts"
   - Action: "Run enrichment sequence"
   - Steps:
     a. Enrich company data (Clearbit/Apollo)
     b. Find technology stack (Builtwith)
     c. Get recent news & funding (Crunchbase)
     d. Calculate account score based on ICP fit
     e. Notify team in Slack if score > 8

3. Save and enable automation
```

### Workflow 2: Contact Email Verification

```
1. In Contacts table → "Automations"

2. Create waterfall enrichment:
   - Trigger: "When contact is added"
   - Action: "Find email"
   - Sequence:
     a. Try Hunter.io
     b. If failed, try Apollo
     c. If failed, try RocketReach
     d. If failed, try Prospeo
     e. Verify email deliverability
     f. Mark status as "Ready" if valid email found
```

### Workflow 3: Trigger-Based Outreach

```
1. In Target Accounts table → "Automations"

2. Create signal-based workflow:
   - Trigger: "When company has trigger event"
   - Signals to monitor:
     * New funding announced
     * Executive hire (relevant role)
     * Product launch
     * Expansion to new market
   - Action: "Add to priority outreach queue"
   - Notify: Send Slack alert to sales team
```

## Step 4: Integration Setup

Guide the user to connect essential integrations:

### CRM Integration (Salesforce/HubSpot)

```
1. Go to "Integrations" in workspace settings

2. Connect your CRM:
   - Search for Salesforce or HubSpot
   - Authenticate with client's CRM credentials
   - Set up two-way sync:
     * Push qualified leads from Clay → CRM
     * Pull engagement data from CRM → Clay

3. Map fields between Clay and CRM

4. Set sync frequency (real-time recommended)
```

### Email Sequencing Tool (Outreach/SalesLoft/Lemlist)

```
1. In Integrations, connect email tool

2. Configure export settings:
   - Auto-export contacts when status = "Ready to Contact"
   - Include personalization variables
   - Set sequence assignment rules

3. Enable reply tracking to update Clay status
```

### Slack Notifications

```
1. Connect Slack workspace

2. Set up notification rules:
   - New high-priority account added (score > 8)
   - Contact replied to outreach
   - Trigger event detected
   - Weekly summary of new qualified leads
```

## Step 5: Template & Best Practices

Provide these best practices:

### Data Quality Rules

```
- Run email verification on all contacts before outreach
- Update account scores monthly as you learn ICP patterns
- Archive contacts who bounce or unsubscribe
- Tag accounts with loss reasons when disqualified
- Keep notes updated with recent interactions
```

### Claygent Prompts Library

Share these effective prompts:

**For Company Research**:
```
"Research {{Company}} and summarize: 1) Their core business model,
2) Recent growth indicators (hiring, funding, expansion),
3) Three specific pain points related to [your client's solution category].
Cite recent sources."
```

**For Personalization**:
```
"Based on {{Full Name}}'s role as {{Job Title}} at {{Company}},
identify 3 specific challenges they likely face that relate to [problem area].
Reference their company's recent {{Recent News}} if relevant."
```

**For Competitive Intel**:
```
"Compare {{Competitor}} to our positioning. What do they emphasize in their
messaging? What customer segments do they target? What are common complaints
in their G2 reviews?"
```

### Table Maintenance Schedule

```
Weekly:
- Review and score new accounts added
- Update contact statuses based on outreach results
- Refresh company news and signals

Monthly:
- Analyze ICP patterns and refine scoring criteria
- Update competitive intelligence
- Clean up stale/bounced contacts
- Review automation performance

Quarterly:
- Comprehensive ICP review and pivot if needed
- Major enrichment provider evaluation
- Workflow optimization based on conversion data
```

## Step 6: Thrustlab Handoff

After setup is complete, ensure proper handoff:

### Checklist for Thrustlab Team Access

```
✅ Thrustlab team added to workspace with Admin access
✅ All 5 core tables created and configured
✅ Enrichment providers connected and tested
✅ Automations enabled and running
✅ CRM integration active (if applicable)
✅ Initial ICP research data populated (50+ accounts)
✅ Claygent prompts customized for client's use case
✅ Team trained on table usage and workflows
✅ Documentation shared with client team
```

### Knowledge Transfer Session

Schedule a walkthrough covering:
```
1. Clay workspace tour (15 min)
2. How to add accounts and contacts (10 min)
3. Reading enrichment data and scores (10 min)
4. Using Claygent for research (15 min)
5. Exporting to CRM/email tools (10 min)
6. Monitoring automation and alerts (10 min)
7. Q&A (10 min)
```

## Expected Outcomes

By the end of this setup:
- ✅ Clay.com workspace created with Thrustlab referral
- ✅ Thrustlab team has admin access to workspace
- ✅ 5 core tables configured with proper schema
- ✅ Enrichment providers connected (email, company data, signals)
- ✅ Automations running for enrichment and notifications
- ✅ CRM and email tool integrations active
- ✅ Initial ICP data populated and scored
- ✅ Client team trained on Clay usage
- ✅ Ongoing maintenance plan established

## After Setup

Ask the user if they want to:
1. **Populate initial data** - Import existing prospect lists
2. **Refine ICP criteria** - Based on existing customer data
3. **Customize Claygent prompts** - For specific research needs
4. **Set up additional integrations** - Other tools in their stack
5. **Create custom views** - Filtered tables for different team members
6. **Schedule training** - For client's sales/marketing team

## Notes

- Clay.com does not have a public API for programmatic table creation
- All setup must be done through the web interface
- Thrustlab referral link provides 500 free credits for new accounts
- Admin access allows you to monitor usage and optimize configurations
- Clay credits are consumed by enrichment operations (plan accordingly)
- Email verification consumes credits but significantly improves deliverability
- Claygent research is powerful but credit-intensive (use strategically)

## Troubleshooting

### Issue: Can't add team members
**Solution**: Ensure workspace owner has completed account setup and email verification first

### Issue: Enrichment not working
**Solution**: Check that enrichment provider integrations are connected in Settings → Integrations

### Issue: CRM sync failing
**Solution**: Verify field mappings and authentication. Check CRM API limits.

### Issue: Running out of credits quickly
**Solution**: Review automation triggers and use waterfall enrichment to try cheaper providers first

---

**Success Metrics:**
- Workspace accessible by both client and Thrustlab team
- All core tables operational with live data
- Enrichment automations running successfully
- CRM integration syncing bidirectionally
- Client team trained and actively using Clay
- ICP becoming clearer through data analysis
