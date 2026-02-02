---
name: clay-tables-setup
description: Guides setup of 5 core Clay.com tables for GTM operations - Target Accounts, Contacts, Outreach, Competitive Intel, and ICP Analysis
argument-hint: "client name"
---

# Clay.com Tables Setup Skill

You are helping configure the core Clay tables for a GTM client's workspace. This creates a standardized structure for prospect research, data enrichment, and outreach automation.

## Prerequisites

Before running this skill, ensure:
- ✅ Clay account created (run `/clay-account-setup` first)
- ✅ Thrustlab team has admin access
- ✅ User is logged into Clay.com

## Required Information

Gather from the user:
1. **Client name** - The name of the client/company
2. **Target market** - Who they're selling to (industry, company size, roles)
3. **GTM focus** - What they're launching or expanding

## Overview

This skill guides creation of 5 core tables:

1. **Target Accounts** - Master list of companies to prospect
2. **Contacts** - Individual decision-makers at target accounts
3. **Outreach Campaigns** - Personalized messaging sequences
4. **Competitive Intel** - Market positioning and competitor tracking
5. **ICP Analysis** - Ideal customer profile validation

Each table takes 10-15 minutes to set up. Total time: ~1 hour.

---

## Table 1: Target Accounts Database

**Purpose**: Master list of target companies for this client's GTM motion

### Setup Instructions

```
📋 Create Target Accounts Table

1. In Clay dashboard, click "New Table" (top right)

2. Select "Start from scratch"

3. Name the table: "[Client Name] - Target Accounts"

4. Set up columns (click + to add new column):

   Column Setup:

   📌 Company Name (Title) - already exists

   🌐 Website
   - Type: URL
   - Description: "Company website"

   🏢 Industry
   - Type: Select
   - Options: SaaS, E-commerce, Financial Services, Healthcare,
             Manufacturing, Professional Services, Other
   - Description: "Primary industry"

   👥 Company Size
   - Type: Select
   - Options: 1-50, 51-200, 201-1000, 1001-5000, 5000+
   - Description: "Employee count"

   💰 Revenue Range
   - Type: Select
   - Options: <$1M, $1M-$10M, $10M-$50M, $50M-$200M, $200M+
   - Description: "Annual revenue estimate"

   📍 Location/HQ
   - Type: Text
   - Description: "Headquarters location"

   💻 Tech Stack
   - Type: Text
   - Description: "Key technologies used"

   💵 Funding Stage
   - Type: Select
   - Options: Bootstrapped, Seed, Series A, Series B, Series C+, Public
   - Description: "Funding status"

   ⭐ Account Score
   - Type: Number
   - Description: "ICP fit score (1-10)"

   📊 Status
   - Type: Select
   - Options: Prospecting, Contacted, Engaged, Qualified, Customer, Lost
   - Description: "Account stage"

   📝 Notes
   - Type: Long Text
   - Description: "Research notes and context"

5. Click "Create Table"

6. Add enrichment integrations:
   - Click "+ Add Column" → "Enrich"
   - Search for "Enrich Company" (Clearbit or Apollo)
   - Configure to auto-enrich on row creation

7. Add signals integration (optional but recommended):
   - Click "+ Add Column" → "Enrich"
   - Search for "Company News" or "Funding Rounds"
   - Set to refresh weekly

8. Create a view for high-priority accounts:
   - Click "Filter" → "Account Score" → "Greater than 7"
   - Click "Sort" → "Account Score" → "Descending"
   - Save view as "High Priority"
```

---

## Table 2: Contact Database

**Purpose**: Individual prospects and decision-makers at target accounts

### Setup Instructions

```
📋 Create Contacts Table

1. Click "New Table"

2. Select "Find people at companies"
   - This template is optimized for contact discovery

3. Name: "[Client Name] - Contacts"

4. Set up columns:

   📌 Full Name (Title) - already exists

   📧 Email
   - Type: Email
   - Enable "Verify email deliverability"

   📱 Phone
   - Type: Phone
   - Description: "Direct dial if available"

   💼 Job Title
   - Type: Text
   - Description: "Current role"

   🏢 Department
   - Type: Select
   - Options: Sales, Marketing, Product, Engineering, Operations,
             Finance, Executive, HR, Customer Success

   📊 Seniority
   - Type: Select
   - Options: IC, Manager, Director, VP, SVP, C-Level

   🏭 Company
   - Type: Relation to "Target Accounts" table
   - Link to the Target Accounts table you created

   🔗 LinkedIn URL
   - Type: URL
   - Description: "LinkedIn profile"

   📍 Location
   - Type: Text
   - Description: "Based in"

   🎯 Recent Activity
   - Type: Text
   - Description: "Job changes, posts, etc."

   ⭐ Engagement Score
   - Type: Number
   - Description: "Priority score (1-10)"

   📊 Status
   - Type: Select
   - Options: New, Research, Ready, Contacted, Replied, Qualified,
             Meeting Booked, Lost
   - Description: "Outreach stage"

   📅 Last Contacted
   - Type: Date
   - Description: "Last outreach date"

   📝 Notes
   - Type: Long Text
   - Description: "Personalization notes"

5. Add waterfall email enrichment:
   - Click "+ Add Column" → "Enrich"
   - Search for "Find Email"
   - Set up waterfall:
     a. Try Hunter.io first
     b. If empty, try Apollo
     c. If empty, try RocketReach
   - Enable "Verify email" at end of waterfall

6. Add LinkedIn enrichment:
   - Click "+ Add Column" → "Enrich"
   - Search for "LinkedIn Profile"
   - Auto-fill from LinkedIn URL when available

7. Create views:
   - "Ready to Contact": Status = "Ready" AND Email is not empty
   - "High Priority": Engagement Score > 7
   - "Needs Research": Status = "New"
```

---

## Table 3: Outreach Campaigns

**Purpose**: Personalized outreach campaigns with AI-generated messaging

### Setup Instructions

```
📋 Create Outreach Table

1. Click "New Table"

2. Select "Import from" → "Contacts table"
   - This links to your contacts for easy campaign creation

3. Name: "[Client Name] - Outreach Campaigns"

4. Import contacts (you can do this now or later):
   - Select contacts with Status = "Ready"
   - Click "Import"

5. Set up columns:

   📌 Contact Name (Title) - from Contacts table

   📧 Email - from Contacts table

   🏢 Company - from Contacts table

   💼 Job Title - from Contacts table

   📢 Campaign Name
   - Type: Select
   - Options: Cold Outreach - Q1, Warm Intro, Product Launch,
             Event Follow-up, Re-engagement

   🔍 Research Notes
   - Type: Long Text
   - Description: "Claygent research findings"

   📧 Email Subject
   - Type: Text
   - Description: "Personalized subject line"

   📝 Email Body
   - Type: Long Text
   - Description: "Full email message"

   📊 Status
   - Type: Select
   - Options: Draft, Ready to Send, Sent, Opened, Replied,
             Bounced, Unsubscribed

   📅 Send Date
   - Type: Date
   - Description: "When email was sent"

   ✉️ Reply Received
   - Type: Checkbox
   - Description: "Did they respond?"

   💬 Reply Content
   - Type: Long Text
   - Description: "Their response"

   📈 Next Step
   - Type: Text
   - Description: "Follow-up action"

6. Add Claygent for AI research:
   - Click "+ Add Column" → "Claygent"
   - Name: "AI Research"
   - Prompt:
     "Research {{Company}} and {{Contact Name}}'s role as {{Job Title}}.
      Find:
      1. Recent company news or growth signals
      2. Their role responsibilities and likely pain points
      3. 3 specific ways our solution addresses their challenges

      Be specific and cite sources where possible."

   - Credit cost: ~10 credits per research

7. Add GPT-4 for email generation:
   - Click "+ Add Column" → "Enrich" → "OpenAI"
   - Name: "Generated Email"
   - Prompt:
     "Write a personalized cold email to {{Contact Name}} at {{Company}}.

      Context:
      - Their role: {{Job Title}}
      - Research: {{AI Research}}

      Requirements:
      - Under 100 words
      - Reference specific research findings
      - Focus on their pain points
      - Include clear CTA
      - Conversational tone
      - No generic buzzwords"

8. Create campaign views:
   - "Ready to Send": Status = "Ready to Send"
   - "Awaiting Reply": Status = "Sent" AND Reply Received = No
   - "Engaged": Reply Received = Yes
```

---

## Table 4: Competitive Intelligence

**Purpose**: Track competitors and market positioning

### Setup Instructions

```
📋 Create Competitive Intel Table

1. Click "New Table" → "Start from scratch"

2. Name: "[Client Name] - Competitive Intel"

3. Set up columns:

   📌 Competitor Name (Title)

   🌐 Website
   - Type: URL

   🎯 Positioning
   - Type: Text
   - Description: "How they position themselves"

   ⚡ Key Features
   - Type: Long Text
   - Description: "Main product features"

   💰 Pricing Model
   - Type: Text
   - Description: "Pricing structure and tiers"

   👥 Target Customers
   - Type: Text
   - Description: "Who they sell to"

   💪 Strengths
   - Type: Text
   - Description: "What they do well"

   ⚠️ Weaknesses
   - Type: Text
   - Description: "Gaps and vulnerabilities"

   📰 Recent News
   - Type: Text
   - Description: "Latest announcements"

   📊 Market Share
   - Type: Text
   - Description: "Estimated market position"

   ⭐ G2 Rating
   - Type: Number
   - Description: "G2 review score"

   📅 Last Updated
   - Type: Date
   - Description: "Last research date"

4. Add enrichment:
   - "+ Add Column" → "Enrich" → "Company Overview"
   - "+ Add Column" → "Enrich" → "G2 Reviews"
   - "+ Add Column" → "Enrich" → "Tech Stack"

5. Add Claygent for competitive analysis:
   - "+ Add Column" → "Claygent"
   - Name: "Competitive Analysis"
   - Prompt:
     "Analyze {{Competitor Name}} at {{Website}}.

      Provide:
      1. Their main positioning and messaging
      2. Key differentiators they emphasize
      3. Target customer segments
      4. Common customer complaints (check G2, Reddit, Twitter)
      5. Recent product launches or changes

      Be specific and cite sources."
```

---

## Table 5: ICP Research & Validation

**Purpose**: Build and refine ideal customer profile

### Setup Instructions

```
📋 Create ICP Analysis Table

1. Click "New Table" → "Find companies"
   - This template helps discover similar companies

2. Name: "[Client Name] - ICP Analysis"

3. Import 50-100 sample companies matching initial ICP hypothesis
   - Use filters for industry, size, location
   - Include both best-fit and edge cases

4. Set up analysis columns:

   📌 Company Name (Title)

   🏢 Industry - from enrichment

   👥 Company Size - from enrichment

   💰 Revenue - from enrichment

   📍 Location - from enrichment

   💻 Tech Stack - from enrichment

   ⭐ ICP Fit Score
   - Type: Number
   - Description: "How well they match ICP (1-10)"

   ✅ Fit Reasons
   - Type: Text
   - Description: "Why they're a good fit"

   ❌ Disqualifiers
   - Type: Text
   - Description: "Red flags or concerns"

   📊 Conversion Likelihood
   - Type: Select
   - Options: High, Medium, Low

   🎯 Win Patterns
   - Type: Text
   - Description: "Common traits of best customers"

5. Add Claygent for pattern analysis:
   - "+ Add Column" → "Claygent"
   - Name: "ICP Pattern Analysis"
   - Prompt (run on high-scoring accounts):
     "Analyze {{Company Name}} to understand why they're a strong ICP fit.

      Identify:
      1. Specific characteristics that make them ideal
      2. Business model or growth stage alignment
      3. Tech stack or process indicators
      4. Likely budget and decision-making process

      Look for patterns we can use to find similar companies."

6. Create views:
   - "Best Fit": ICP Fit Score >= 8
   - "Need Research": ICP Fit Score is empty
   - "Low Fit": ICP Fit Score <= 4
```

---

## Post-Setup Actions

After creating all 5 tables:

```
✅ Table Setup Verification

Confirm with user that all tables are created:

□ Target Accounts database configured
□ Contacts database with email enrichment
□ Outreach campaigns with Claygent and GPT-4
□ Competitive intelligence table
□ ICP analysis table

Next steps:
1. Populate initial data (import CSVs or manual entry)
2. Test enrichment on 5-10 sample rows
3. Run automation setup: /clay-automation-setup [Client Name]
```

## Best Practices

Share these tips with the user:

```
💡 Clay Table Best Practices

Data Quality:
- Test enrichments on small batches first (5-10 rows)
- Verify email deliverability before sending
- Keep notes updated with research context
- Archive old/stale contacts regularly

Credit Management:
- Email finding: ~2-5 credits per contact
- Claygent research: ~10-20 credits per query
- Basic enrichment: ~1-2 credits per field
- Use waterfall to try cheaper providers first

Table Hygiene:
- Update target account scores monthly
- Mark lost deals with loss reason
- Track outreach results to improve messaging
- Refresh competitive intel quarterly

Workflow:
1. Add accounts to Target Accounts
2. Find contacts at those accounts
3. Research with Claygent
4. Generate personalized outreach
5. Track engagement in Outreach table
6. Refine ICP based on results
```

## Troubleshooting

### Issue: Enrichment not working
**Solution**:
- Check integration connections in Settings → Integrations
- Verify API keys are valid
- Try different enrichment provider
- Check credit balance

### Issue: Claygent responses too generic
**Solution**:
- Make prompts more specific
- Include more context variables
- Ask for cited sources
- Limit to 2-3 specific questions

### Issue: Relations between tables not working
**Solution**:
- Ensure both tables exist first
- Use "Relation" column type
- Select correct table to link to
- May need to refresh page

### Issue: Running out of credits
**Solution**:
- Use waterfall enrichment (cheaper providers first)
- Don't auto-enrich every row
- Batch research with Claygent
- Upgrade plan if doing high-volume

## Expected Outcomes

By the end of this setup:
- ✅ 5 core tables created with proper schema
- ✅ Enrichment integrations configured
- ✅ Claygent prompts customized for use case
- ✅ Relations between tables established
- ✅ Sample views created for common workflows
- ✅ Ready for automation and integration setup

## Next Steps

```
🎯 Tables are ready! What's next?

1. Set up automation workflows:
   Run: /clay-automation-setup [Client Name]
   → Auto-enrichment, notifications, triggers

2. Connect your CRM and email tools:
   Run: /clay-integrations-setup [Client Name]
   → Salesforce, HubSpot, Outreach, Slack

3. Populate initial data:
   - Import existing prospect lists
   - Add known target accounts
   - Begin enrichment on sample batch
```

---

**Success Metrics:**
- All 5 tables operational
- Enrichment providers connected
- Ready for data population and automation
