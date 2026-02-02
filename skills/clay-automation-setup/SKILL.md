---
name: clay-automation-setup
description: Guides setup of Clay.com automation workflows for auto-enrichment, email verification, trigger-based alerts, and data quality maintenance
argument-hint: "client name"
---

# Clay.com Automation Setup Skill

You are helping configure automation workflows in Clay.com to streamline data enrichment, maintain data quality, and trigger timely notifications.

## Prerequisites

Before running this skill, ensure:
- ✅ Clay account created
- ✅ Core tables set up (run `/clay-tables-setup` first)
- ✅ Enrichment integrations connected
- ✅ User logged into Clay.com

## Required Information

Gather from the user:
1. **Client name** - The name of the client/company
2. **Notification preferences** - Where to send alerts (Slack, email, etc.)
3. **Auto-enrichment rules** - Which fields to enrich automatically

## Overview

This skill sets up 6 key automation workflows:

1. **New Account Auto-Enrichment** - Enrich company data when accounts added
2. **Contact Email Waterfall** - Find and verify emails automatically
3. **Trigger-Based Alerts** - Notify team of high-value signals
4. **Data Quality Maintenance** - Clean and update stale records
5. **ICP Scoring Automation** - Calculate account fit scores
6. **Weekly Summary Reports** - Digest of new qualified leads

---

## Automation 1: New Account Auto-Enrichment

**Purpose**: Automatically enrich company data when new accounts are added to Target Accounts table

### Setup Instructions

```
📋 Create Account Enrichment Automation

1. Open "[Client Name] - Target Accounts" table

2. Click "Automations" tab (top right, thunder bolt icon)

3. Click "+ New Automation"

4. Configure trigger:
   - Trigger: "When row is created"
   - Condition: None (run on all new rows)

5. Add enrichment steps (click "+ Add Step" after each):

   Step 1: Enrich Company Data
   - Action: "Enrich" → "Company Overview"
   - Provider: Clearbit or Apollo
   - Map fields:
     * Website → Website column
     * Industry → Industry column
     * Employee Count → Company Size column
     * Revenue → Revenue Range column

   Step 2: Get Tech Stack
   - Action: "Enrich" → "Technology Stack"
   - Provider: Builtwith or Datanyze
   - Map field:
     * Technologies → Tech Stack column

   Step 3: Find Funding Info
   - Action: "Enrich" → "Funding Data"
   - Provider: Crunchbase
   - Map fields:
     * Funding Stage → Funding Stage column
     * Last Funding Date → (add to Notes)

   Step 4: Get Recent News
   - Action: "Enrich" → "Company News"
   - Provider: News API or similar
   - Timeframe: Last 30 days
   - Map field:
     * Top Headlines → Notes (append)

   Step 5: Calculate ICP Score (optional)
   - Action: "Run Formula"
   - Formula logic:
     * Start with base score: 5
     * +2 if Company Size matches target
     * +1 if Industry matches target
     * +1 if Funding Stage >= Series A
     * +1 if in target geography
   - Map to: Account Score column

   Step 6: Notify if High Priority
   - Action: "Send Notification"
   - Condition: IF Account Score >= 8
   - Channel: Slack (you'll configure in Step 4)
   - Message: "🎯 New high-priority account: {{Company Name}} (Score: {{Account Score}})"

6. Name the automation: "Auto-Enrich New Accounts"

7. Toggle "Active" to ON

8. Click "Save"

💡 Tip: Test on a sample row first before enabling!
```

---

## Automation 2: Contact Email Waterfall

**Purpose**: Automatically find and verify contact emails using waterfall enrichment

### Setup Instructions

```
📋 Create Email Enrichment Waterfall

1. Open "[Client Name] - Contacts" table

2. Click "Automations" tab

3. Click "+ New Automation"

4. Configure trigger:
   - Trigger: "When row is created"
   - Condition: IF Email is empty

5. Add waterfall enrichment steps:

   Step 1: Try Hunter.io
   - Action: "Enrich" → "Find Email" → "Hunter.io"
   - Input: Full Name + Company
   - Map to: Email column
   - Continue if: Result is empty

   Step 2: Try Apollo
   - Action: "Enrich" → "Find Email" → "Apollo"
   - Input: Full Name + Company
   - Map to: Email column
   - Continue if: Result is empty

   Step 3: Try RocketReach
   - Action: "Enrich" → "Find Email" → "RocketReach"
   - Input: Full Name + Company + Job Title
   - Map to: Email column
   - Continue if: Result is empty

   Step 4: Try Prospeo
   - Action: "Enrich" → "Find Email" → "Prospeo"
   - Input: Full Name + Company
   - Map to: Email column
   - Continue if: Result is empty

   Step 5: Verify Email Deliverability
   - Action: "Verify Email" → "Email Verification"
   - Provider: ZeroBounce or NeverBounce
   - Map result: (add verification status to Notes)
   - Condition: IF Email is not empty

   Step 6: Update Status
   - Action: "Update Field"
   - Condition: IF Email is valid
   - Field: Status
   - Value: "Ready"

   Step 7: Notify if Failed
   - Action: "Update Field"
   - Condition: IF Email still empty after all attempts
   - Field: Notes
   - Value: "❌ Email not found - requires manual research"

6. Name: "Email Waterfall Enrichment"

7. Toggle "Active" to ON

8. Click "Save"

⚡ Credit Saving Tip: This waterfall tries cheap providers first,
   only using expensive ones if needed!
```

---

## Automation 3: Trigger-Based Alerts

**Purpose**: Monitor for high-value buying signals and notify team immediately

### Setup Instructions

```
📋 Create Signal-Based Alerts

1. Open "[Client Name] - Target Accounts" table

2. Click "Automations" tab

3. Create multiple automations for different signals:

AUTOMATION A: Funding Alert
━━━━━━━━━━━━━━━━━━━━━━━━━
Trigger: "When field changes"
Field: Funding Stage OR Recent News contains "funding"

Steps:
1. Check if funding is recent (within 30 days)
2. IF yes, send Slack notification:
   "💰 {{Company Name}} just raised funding!
    Stage: {{Funding Stage}}
    This is a great time to reach out."
3. Update Account Score: +2
4. Add tag: "Hot - Recent Funding"

AUTOMATION B: Executive Hire Alert
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Trigger: "When row changes"
Condition: Recent News contains "hired" OR "appointed" OR "joins as"

Steps:
1. Parse news for role (VP, CTO, CMO, etc.)
2. IF role matches target buyer persona:
   - Send Slack: "👤 {{Company Name}} hired new {{Role}}
                 Perfect timing for outreach!"
   - Update Status: "High Priority"
   - Add to Notes: "New hire - {{parsed info}}"

AUTOMATION C: Product Launch Alert
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Trigger: "When field updates"
Field: Recent News

Condition: Contains "launch" OR "announced" OR "released"

Steps:
1. Extract product/feature info
2. Notify team:
   "🚀 {{Company Name}} just launched something new!
    {{Brief summary}}
    Could be expansion signal."
3. Add to priority outreach queue

AUTOMATION D: Company Growth Alert
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Trigger: "When field changes"
Field: Company Size (increased)

Steps:
1. Compare previous vs new size
2. IF increase > 20%:
   - Notify: "📈 {{Company Name}} is hiring fast!
            Grew from {{old size}} to {{new size}}"
   - Tag: "High Growth"
   - Increase Account Score

4. Enable all automation variations

5. Customize notification thresholds based on client's ICP
```

---

## Automation 4: Data Quality Maintenance

**Purpose**: Keep data fresh and clean automatically

### Setup Instructions

```
📋 Create Data Maintenance Automations

AUTOMATION A: Refresh Stale Data
━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Target Accounts table → Automations

2. Trigger: "On schedule"
   - Frequency: Weekly (every Monday)
   - Time: 9:00 AM

3. Condition: Last Updated > 30 days ago

4. Steps:
   - Re-run company enrichment
   - Refresh Recent News
   - Update Tech Stack
   - Mark Last Updated: Today

AUTOMATION B: Bounce Handling
━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Contacts table → Automations

2. Trigger: "When field changes"
   - Field: Status
   - Change to: "Bounced"

3. Steps:
   - Try to find new email (re-run waterfall)
   - IF still bounced:
     * Add to Notes: "Email bounced - needs update"
     * Search for LinkedIn profile
     * Try LinkedIn InMail as alternative
   - Remove from active outreach lists

AUTOMATION C: Unsubscribe Management
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Contacts table → Automations

2. Trigger: "When field changes"
   - Field: Status
   - Changes to: "Unsubscribed"

3. Steps:
   - Remove from all outreach campaigns
   - Add suppression tag
   - Update Notes: "Unsubscribed on {{date}}"
   - Archive (move to separate table or mark)

AUTOMATION D: Duplicate Detection
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Contacts table → Automations

2. Trigger: "When row created"

3. Steps:
   - Check if Email already exists in table
   - IF duplicate found:
     * Flag for review
     * Notify: "⚠️ Possible duplicate: {{Name}}"
     * Don't auto-enrich (saves credits)
```

---

## Automation 5: ICP Scoring Automation

**Purpose**: Automatically calculate and update account fit scores

### Setup Instructions

```
📋 Create ICP Scoring Logic

1. Target Accounts table → Automations

2. Trigger: "When row is created" OR "When enrichment completes"

3. Add scoring logic using Formula steps:

   Base ICP Scoring Formula:
   ━━━━━━━━━━━━━━━━━━━━━━

   Starting Score: 5

   Industry Match:
   +2 if perfect fit
   +1 if adjacent industry
   +0 if different

   Company Size:
   +2 if within target range
   +1 if slightly outside
   -1 if too small/large

   Revenue Range:
   +1 if meets minimum
   +1 if in sweet spot
   -1 if too low

   Tech Stack:
   +1 if using complementary tools
   +1 if using competitor tools (displacement opportunity)

   Funding Stage:
   +1 if funded (has budget)
   +1 if recent raise (actively spending)

   Growth Signals:
   +1 if hiring in target roles
   +1 if expanding geographically

   Location:
   +1 if in target geography
   +0 otherwise

4. Map result to Account Score column

5. Add conditional actions based on score:

   IF Score >= 9: "A+ Tier"
   - Tag: "Tier A+"
   - Notify sales team immediately
   - Route to senior AE

   IF Score 7-8: "A Tier"
   - Tag: "Tier A"
   - Add to priority outreach
   - Standard workflow

   IF Score 5-6: "B Tier"
   - Tag: "Tier B"
   - Monitor for signals
   - Nurture campaign

   IF Score < 5: "C Tier / Disqualify"
   - Tag: "Low Fit"
   - Consider removal or long-term nurture

6. Re-calculate scores monthly:
   - Schedule: First of each month
   - Re-run scoring on all active accounts
   - Adjust tiers based on new data
```

---

## Automation 6: Weekly Summary Reports

**Purpose**: Automated digest of activity and new opportunities

### Setup Instructions

```
📋 Create Weekly Summary Automation

1. Create a new automation (at workspace level if possible)

2. Trigger: "On schedule"
   - Frequency: Weekly
   - Day: Friday
   - Time: 4:00 PM

3. Aggregate data from multiple tables:

   From Target Accounts:
   - Count: New accounts added this week
   - Count: High-priority accounts (score >= 8)
   - List: Top 5 accounts by score

   From Contacts:
   - Count: New contacts added
   - Count: Ready for outreach
   - Count: Replied this week

   From Outreach Campaigns:
   - Count: Emails sent this week
   - Percentage: Reply rate
   - Count: Meetings booked
   - List: Top performing emails

4. Format summary message:

   "📊 Weekly Clay Summary - [Client Name]

   🎯 ACCOUNTS
   • {{new_accounts}} new target accounts added
   • {{high_priority}} high-priority opportunities
   • Top prospects: {{top_5_list}}

   👥 CONTACTS
   • {{new_contacts}} new contacts researched
   • {{ready_count}} ready for outreach
   • {{replies}} replies received ({{reply_rate}}% rate)

   📧 OUTREACH
   • {{sent}} emails sent
   • {{meetings}} meetings booked
   • Best performing subject: {{top_subject}}

   🚀 NEXT WEEK PRIORITIES
   {{auto_generated_priorities}}

   View in Clay: {{workspace_link}}"

5. Send to:
   - Slack channel
   - Email to key stakeholders
   - Internal dashboard (if available)

6. Enable and test
```

---

## Testing & Validation

After setting up automations:

```
✅ Automation Testing Checklist

Test each automation with sample data:

□ Add test account to Target Accounts
   → Verify auto-enrichment runs
   → Check all fields populated
   → Confirm notification sent if high score

□ Add test contact without email
   → Verify waterfall enrichment runs
   → Check email found and verified
   → Confirm status updated to "Ready"

□ Update account with "funding" news
   → Verify trigger alert fires
   → Check Slack notification received
   → Confirm score increased

□ Check scheduled automations
   → Verify weekly refresh is queued
   → Confirm summary report scheduled
   → Test timing is appropriate

□ Monitor credit usage
   → Check enrichment costs per row
   → Adjust waterfall if too expensive
   → Disable unused enrichments

Troubleshooting:
- If automation doesn't run: Check trigger conditions
- If enrichment fails: Verify API connections
- If notifications missing: Check Slack integration
- If too many credits used: Add stricter conditions
```

## Best Practices

```
💡 Automation Best Practices

Start Small:
- Enable 1-2 automations first
- Test thoroughly before adding more
- Monitor credit consumption

Conditional Logic:
- Don't enrich every field on every row
- Use IF/THEN conditions to save credits
- Only notify on truly high-value signals

Error Handling:
- Add fallback steps if enrichment fails
- Log errors to Notes field for review
- Set up alerts if automation breaks

Performance:
- Batch operations when possible
- Schedule heavy tasks during off-hours
- Limit Claygent usage (expensive)

Maintenance:
- Review automation performance monthly
- Disable unused workflows
- Update scoring criteria as ICP evolves
- Clean up old notifications

Credit Management:
- Monitor spend per automation
- Adjust waterfall order (cheapest first)
- Disable auto-enrich on low-priority rows
- Use scheduled refreshes instead of real-time
```

## Troubleshooting

### Issue: Automation not triggering
**Solution**:
- Check trigger conditions are met
- Verify automation is toggled "Active"
- Ensure no conflicting automations
- Check workflow logs for errors

### Issue: Enrichment failing
**Solution**:
- Verify integration connections
- Check API keys are valid
- Ensure input data format is correct
- Try different enrichment provider

### Issue: Too many notifications
**Solution**:
- Add stricter conditions (score thresholds)
- Batch notifications (daily digest instead of real-time)
- Filter out low-priority signals
- Adjust alert rules

### Issue: Running out of credits
**Solution**:
- Review automation frequency
- Add conditions to limit enrichment runs
- Use waterfall with cheaper providers first
- Disable auto-enrichment on low-priority rows
- Upgrade plan if needed

## Expected Outcomes

By the end of this setup:
- ✅ Auto-enrichment running on new accounts
- ✅ Email waterfall finding & verifying contacts
- ✅ Trigger-based alerts for buying signals
- ✅ Data quality maintenance automated
- ✅ ICP scoring calculated automatically
- ✅ Weekly summary reports delivered
- ✅ Team notified of high-priority opportunities
- ✅ Credit usage optimized

## Next Steps

```
🎯 Automations are live! What's next?

1. Connect integrations for full workflow:
   Run: /clay-integrations-setup [Client Name]
   → CRM sync, email tools, Slack

2. Start populating data:
   - Import target account lists
   - Add contacts to enrich
   - Monitor automation runs

3. Optimize based on results:
   - Track credit consumption
   - Adjust scoring criteria
   - Refine notification rules
   - A/B test enrichment providers
```

---

**Success Metrics:**
- Automations running without errors
- Data enriched within 5 minutes of adding
- Team notified of high-value signals
- Credit usage optimized
- Weekly summaries delivered consistently
