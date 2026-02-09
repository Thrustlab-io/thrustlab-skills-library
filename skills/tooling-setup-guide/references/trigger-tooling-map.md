# Trigger Play → Required Tooling Map

For each trigger play recommended in the strategy, these are the tools the client needs.

## Always Required (Every Client)

| Tool | Purpose | Setup Priority |
|---|---|---|
| **Clay** | Data enrichment, research, AI copy generation, workflow orchestration | #1 — everything runs through Clay |
| **Email sequencer** (Smartlead / Instantly) | Send email cadences, manage inbox rotation, track replies | #2 — needed before any outreach |
| **CRM** (HubSpot / Salesforce / Pipedrive) | Pipeline tracking, contact management, reporting | #3 — sync with Clay exports |
| **LinkedIn Sales Navigator** | Boolean searches, lead lists, InMail | #4 — primary prospecting source |

## Trigger-Specific Tooling

### Website Visitor Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **RB2B** | Identifies companies visiting client website | ~$149-399/mo |
| **Clearbit Reveal** (alternative) | Website visitor deanonymization | ~$99-499/mo |
| **6sense** (enterprise alternative) | Intent data + website identification | Enterprise pricing |

Clay integration: RB2B/Clearbit → Clay webhook → trigger enrichment + copy generation

### Job Role Change Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **UserGems** | Tracks job changes of target personas | ~$500+/mo |
| **LinkedIn Sales Navigator** (manual) | Job change alerts via saved searches | Included in SN subscription |
| **Apollo** (budget option) | Job change signals in contact database | ~$49-99/mo |
| **Clay signal** | Built-in job change detection | Included in Clay credits |

Clay integration: UserGems/Clay signal → Clay table → enrich + copy generation

### Job Posting Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **Clay signal (Job Postings)** | Monitors job boards for relevant postings | Included in Clay credits |
| **Otta API** | Job posting aggregator | Varies |
| **LinkedIn Jobs** (manual) | Boolean search for relevant job postings | Free / SN subscription |
| **Indeed API** | Job posting data | Varies |

Clay integration: Clay signal / manual import → Clay table → parse job description → enrich + copy

### Funding Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **Crunchbase** | Funding announcements, investor data | ~$29-49/mo |
| **PitchBook** (enterprise) | Comprehensive funding + deal data | Enterprise pricing |
| **Clay signal (Funding)** | Built-in funding detection | Included in Clay credits |
| **Google Alerts** (free backup) | News monitoring for funding announcements | Free |

### Tech Stack Change Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **BuiltWith** | Technology adoption/removal tracking | ~$295+/mo |
| **Wappalyzer** (via Clay) | Tech stack enrichment | Included in Clay |
| **HG Insights** (enterprise) | Technographic data at scale | Enterprise pricing |

### Growth / Hiring Surge Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **Clay signal (Headcount)** | Headcount growth monitoring | Included in Clay credits |
| **LinkedIn Sales Navigator** | Company headcount tracking via saved searches | Included in SN subscription |
| **Apollo** | Growth indicators in company data | ~$49-99/mo |

### Content Engagement Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **HubSpot / Marketo** | Track content downloads, webinar attendance | Depends on existing stack |
| **LinkedIn Analytics** | Track post engagement | Free |
| **Phantom Buster** (for LinkedIn scraping) | Extract post engagers | ~$56-128/mo |

### Compliance / Regulatory Trigger
| Tool | Purpose | Cost Range |
|---|---|---|
| **Google Alerts** | Monitor regulatory news | Free |
| **Industry-specific databases** | Varies by vertical (discovered in market mapping) | Varies |
| **Clay signal + Claygent** | Research regulatory events per company | Included in Clay credits |

### Champion Tracking Trigger (NEW — Highest Converting Signal)
| Tool | Purpose | Cost Range |
|---|---|---|
| **UserGems** (preferred) | Tracks job changes of former customers/prospects | ~$500+/mo |
| **Clay signal (Job Changes)** | Built-in job change detection for uploaded contact lists | Included in Clay credits |
| **LinkedIn Sales Navigator** | Manual tracking via saved searches + alerts | Included in SN subscription |

Clay integration: Upload past customer/prospect contact list → UserGems/Clay monitors for job changes → webhook fires to Clay → enrich new company → champion-aware copy generation

**Setup requirement:** Client must provide a list of past customer contacts and/or past engaged prospects from their CRM.

### Customer Competitor Targeting Trigger (NEW — 2.5x Conversion)
| Tool | Purpose | Cost Range |
|---|---|---|
| **Claygent** | Research whether target companies use a competitor product | Included in Clay credits |
| **BuiltWith** (for SaaS/tech products) | Detect competitor technology on prospect websites | ~$295+/mo |
| **G2 / TrustRadius data** | Identify companies reviewing competitor products | Via intent data providers |
| **Clay AI column** | Classify competitor status from research + tech stack data | Included in Clay credits |

Clay integration: ICP list → Claygent researches competitor usage → AI column classifies competitor_status (confirmed/likely/none) → competitor-specific copy generation

### Dark Funnel / Social Engagement Trigger (NEW)
| Tool | Purpose | Cost Range |
|---|---|---|
| **Dealfront** (EU-preferred) | Company-level website visitor ID, strong EU coverage | ~€199+/mo |
| **RB2B** (US-preferred) | Person-level website visitor ID, US-focused | ~$149-399/mo |
| **Teamfluence** (LinkedIn engagement) | Monitors when ICP contacts engage with LinkedIn content (likes, comments, shares) | ~€99+/mo |
| **Trigify** (LinkedIn triggers) | Monitors LinkedIn events — job changes, posts, engagement patterns, company updates | ~€149+/mo |
| **Common Room** (community) | Aggregates signals from Slack, Discord, GitHub, forums, social media | ~$500+/mo |
| **Phantom Buster** (budget LinkedIn) | Scrape LinkedIn post engagers, commenters | ~$56-128/mo |

Clay integration: Dealfront/RB2B webhook → Clay (website signals) | Teamfluence/Trigify export → Clay import (LinkedIn signals) | Common Room API → Clay (community signals) → enrich + score + copy generation

**Preferred stack by market:**
- **EU clients:** Dealfront + Teamfluence + Trigify
- **US clients:** RB2B + Trigify + Common Room
- **Global:** Dealfront + RB2B + Teamfluence + Trigify

## Signal Stacking Infrastructure (When 2+ Trigger Plays Active)

| Tool | Purpose | Cost Range |
|---|---|---|
| **Clay (Signal Aggregation Table)** | Master table combining signals across all trigger plays | Included in Clay |
| **Clay webhooks** | Push signals from individual trigger tables to aggregation table | Included in Clay |
| **Slack integration** | Alert team for Hot-tier accounts (composite score ≥ 100) | Free (Slack API) |
| **CRM integration** | Pull relationship signals (past meetings, past replies) | Depends on CRM |

See `shared/references/signal-stacking-guide.md` for the complete scoring framework.

## Email Sequencer Recommendations

| Budget | Tool | Why |
|---|---|---|
| Budget-friendly | **Instantly** ($30-77/mo) | Good deliverability, simple UI, Clay integration |
| Mid-range | **Smartlead** ($39-94/mo) | More advanced sequencing, native Clay integration, inbox rotation |
| Enterprise | **Outreach / Salesloft** | Full-featured, CRM integration, call + email + LinkedIn |

## Email Infrastructure (Always Required)
- 2-3 domains per client (not their primary domain)
- 2-3 inboxes per domain (Google Workspace or Outlook)
- 2-3 week warmup period before sending
- Max 30 emails/inbox/day during ramp
- Domain authentication: SPF, DKIM, DMARC configured
