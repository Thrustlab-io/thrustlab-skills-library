# Tooling Setup Document Template

Generate the client's `tooling-setup.md` following this structure. Only include sections relevant to their recommended trigger plays.

## Document Structure

```markdown
# {Company Name} — Tooling Setup Guide

**Prepared by:** Thrustlab
**Date:** {date}
**Based on strategy:** {link to strategy or strategy date}

## Overview

Your GTM strategy recommends 3 trigger-based outbound plays + 1 general outbound campaign:
1. **{Trigger Play 1}** — {one-line description of why}
2. **{Trigger Play 2}** — {one-line description of why}
3. **{Trigger Play 3}** — {one-line description of why}
4. **General Outbound** — Research-based outreach to ICP-fit accounts

This guide covers every tool you need, in priority order.

---

## 1. Clay — Your GTM Operating System

### Account Setup
1. Create your Clay account via our referral link: **https://clay.com?via=c75c72**
2. Sign up with your company email: {client_email}
3. Choose your plan: **{recommended_plan}** ({reasoning based on expected volume})
4. Complete payment and verify email

### Workspace Configuration
1. Rename workspace to: **"{Company Name} — GTM Operations"**
2. Invite Thrustlab team as Admins:
   - kwinten@thrustlab.io
   - jan@thrustlab.io
3. Invite your team members with appropriate roles

### Tables to Create
Based on your strategy, create these Clay tables:

| Table Name | Purpose | Priority |
|---|---|---|
| {Company Name} — {Trigger 1 Name} | {Purpose} | Week 1 |
| {Company Name} — {Trigger 2 Name} | {Purpose} | Week 2 |
| {Company Name} — {Trigger 3 Name} | {Purpose} | Week 2 |
| {Company Name} — General Outbound | Research-based ICP outreach | Week 1 |
| {Company Name} — Master Accounts | Deduplicated account database | Week 1 |

Table schemas and enrichment columns will be configured when we build each workflow.

### Enrichment Providers to Connect
Based on your ICP and trigger plays, connect these providers in Clay Settings → Integrations:

{List only providers needed for this client's specific plays}

---

## 2. Email Sequencer — {Recommended Tool}

### Why {Tool}
{1-2 sentences explaining why this tool fits the client's needs and budget}

### Setup Steps
1. Create account at {URL}
2. Connect {number} sending domains (NOT your primary domain)
   - Recommended: {domain suggestions, e.g., "getacme.io", "acme-team.com"}
3. Create {number} inboxes per domain (Google Workspace or Outlook)
4. Start inbox warmup — minimum 2-3 weeks before sending
5. Configure: SPF, DKIM, DMARC on each domain
6. Set daily sending limit: 30 emails/inbox/day (ramp slowly)

### Campaign Structure
Pre-create these campaigns (we'll populate copy in Phase 3):
{List campaigns matching the trigger plays + general outbound}

---

## 3. Signal & Intent Tools

{Only include sections for tools needed by the client's specific trigger plays. See `references/trigger-tooling-map.md` for the complete mapping.}

### {Signal Tool 1} — For {Trigger Play Name}
**What it does:** {description}
**Cost:** {price range}
**Setup:**
{Numbered steps}
**Clay integration:** {How data flows from this tool into Clay}

### {Signal Tool 2} — For {Trigger Play Name}
{Same structure}

### Champion Tracking Tools (if recommended)
{Include this section only if champion tracking trigger is recommended in strategy}

**UserGems** (or Clay Job Changes signal)
**What it does:** Monitors your past customers and prospects for job changes. When someone from your CRM starts a new role, you get alerted.
**Cost:** ~$500+/mo for UserGems, included with Clay for basic job change detection
**Setup:**
1. Export past customer/prospect contact list from {client_CRM}
   - Required fields: email, name, title, company, relationship_type (customer/prospect/meeting_had), product_used, result_achieved (if available)
2. Upload list to UserGems OR create Clay "Champion Watchlist" table
3. Configure alert: webhook → Clay when job change detected
4. Estimated contacts to monitor: {number from intake}
**Clay integration:** Job change webhook → new row in Champion Tracking table → auto-enrich new company → copy generation

### Competitor Detection Tools (if recommended)
{Include this section only if competitor customer trigger is recommended}

**BuiltWith** (for SaaS/tech products)
**What it does:** Detects which technologies companies use. Identifies your competitors' customers.
**Cost:** ~$295+/mo
**Setup:**
1. Create BuiltWith account
2. Set up technology alerts for competitors: {competitor_1}, {competitor_2}, {competitor_3}
3. Configure: alert when company adds/removes competitor technology
4. Connect to Clay via webhook or scheduled export
**Clay integration:** Tech alert → new row or enrichment column in Competitor Customer table

**Claygent Research** (alternative for non-tech products)
Built into Clay — no additional cost. Claygent researches company websites, job postings, and reviews to identify competitor usage. Lower confidence than BuiltWith but works for all product types.

### Dark Funnel Signal Tools (if recommended)
{Include this section based on client's primary market and strategy}

**Website Visitor Identification:**
{Choose based on primary market from intake}
- **Dealfront** (EU-preferred): Company-level website visitor ID, ~€199+/mo
  Setup: Install tracking script on {client_website}, configure webhook → Clay
- **RB2B** (US-preferred): Person-level website visitor ID, ~$149-399/mo
  Setup: Install tracking pixel on {client_website}, configure webhook → Clay

**LinkedIn Engagement Monitoring:**
- **Teamfluence**: Monitors LinkedIn engagement from ICP contacts, ~€99+/mo
  Setup: Connect LinkedIn, define ICP criteria, configure export → Clay
- **Trigify**: LinkedIn trigger events (job changes, posts, engagements), ~€149+/mo
  Setup: Define trigger criteria, configure webhook → Clay

**Community Monitoring:**
- **Common Room**: Tracks Slack, Discord, GitHub, forum activity, ~$500+/mo
  Setup: Connect community platforms, define ICP signals, configure API → Clay

{Only include tools relevant to client's market and budget}

---

## 3b. Signal Aggregation Table (if 2+ trigger plays active)
{Include this section when client has 2+ trigger plays — enables signal stacking}

**What it does:** Master Clay table that combines signals from all active trigger tables into composite scores. Routes accounts to treatment tiers (Hot/Warm/Active/Watching).

**Setup:**
1. Create Clay table: "{Company Name} — Signal Aggregation"
2. Configure webhooks/scheduled merge from each active trigger table
3. Set up composite scoring formula (see `shared/references/signal-stacking-guide.md`)
4. Configure Slack alerts for Hot tier (composite score ≥100):
   - Channel: #{client_slug}-alerts
   - Alert includes: company name, composite score, signal breakdown, assigned AE

**When to set up:** After at least 2 trigger tables are live and producing data.

---

## 4. LinkedIn Sales Navigator

### Setup
1. Upgrade to Sales Navigator (if not already)
2. Create saved searches for:
   - {Search 1 based on ICP — e.g., "VP Operations at food manufacturing companies in Belgium, 50-500 employees"}
   - {Search 2 based on trigger — e.g., "Job changes in past 90 days for target titles"}
3. Set up alerts for new results in each saved search
4. Connect to Clay for list exports (if on Team plan)

---

## 5. CRM Integration

### {Client's CRM} Setup
1. Connect {CRM} to Clay (Settings → Integrations)
2. Map these fields between Clay and {CRM}:
   {Field mapping table based on client's needs}
3. Configure two-way sync for:
   - New leads from Clay → {CRM}
   - Reply/meeting status from {CRM} → Clay
4. Set up Slack notifications for:
   - New A+ tier account detected
   - Reply received
   - Meeting booked

---

## 6. Integration Checklist

### Data Flow
```
{Signal source} → Clay (enrich + score + copy) → {Sequencer} (send) → {CRM} (track)
                                                 ↓
                                          Slack (notifications)
```

### Verification Checklist
- [ ] Clay account created with Thrustlab referral
- [ ] Thrustlab team has admin access to Clay
- [ ] {number} Clay tables created (schemas configured later)
- [ ] {Sequencer} account created
- [ ] {number} sending domains purchased and configured
- [ ] {number} inboxes created and warming
- [ ] SPF/DKIM/DMARC configured on all domains
- [ ] {Signal tool 1} connected to Clay
- [ ] {Signal tool 2} connected to Clay (if applicable)
- [ ] Champion tracking: CRM contact list exported and uploaded (if applicable)
- [ ] Champion tracking: UserGems / Clay job change monitoring configured (if applicable)
- [ ] Competitor detection: BuiltWith alerts configured for competitors (if applicable)
- [ ] Dark funnel: Website visitor ID tool installed and sending to Clay (if applicable)
- [ ] Dark funnel: LinkedIn monitoring configured (if applicable)
- [ ] Signal Aggregation Table created with composite scoring (if 2+ triggers active)
- [ ] LinkedIn Sales Navigator active with saved searches
- [ ] {CRM} connected to Clay
- [ ] Slack notifications configured (including Hot tier alerts if signal stacking active)

---

## Timeline

| Week | Action | Status |
|---|---|---|
| Week 1 | Clay account + sequencer + domains + inbox warmup starts | |
| Week 1 | Signal tools subscribed and configured | |
| Week 2 | LinkedIn Sales Navigator searches created | |
| Week 2 | CRM integration configured | |
| Week 3 | Inbox warmup complete — ready to send | |
| Week 3 | First workflow goes live | |

---

*Questions? Reach out in your Slack channel or email kwinten@thrustlab.io*
```
