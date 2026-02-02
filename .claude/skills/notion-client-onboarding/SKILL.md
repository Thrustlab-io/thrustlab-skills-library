---
name: notion-client-onboarding
description: Sets up a complete Notion workspace for a new GTM client with standardized project structure, databases, and templates
argument-hint: "client name"
---

# Notion Client Onboarding Skill

You are helping set up a complete Notion workspace for a new GTM (Go-To-Market) client. This skill creates a standardized project structure to manage the client's GTM strategy, execution, and deliverables.

## Required Information

Before proceeding, gather from the user:
1. **Client name** - The name of the client/company
2. **Project start date** - When the engagement begins
3. **Primary contact** - Main client stakeholder (optional)
4. **GTM focus** - What they're launching (product, feature, market expansion, etc.)

## Project Structure to Create

Create the following hierarchy in Notion:

### 1. Main Client Page
Create a parent page titled: `[Client Name] - GTM Project`

This page should include:
- Client overview section
- Key contacts
- Project timeline
- Quick links to all databases and sub-pages

### 2. Core Databases

Create these databases as sub-pages:

#### A. **GTM Strategy & Planning Database**
Properties:
- Name (title)
- Status (select: Not Started, In Progress, In Review, Completed)
- Priority (select: High, Medium, Low)
- Category (select: Market Research, Positioning, Messaging, Competitive Analysis, Personas, Channel Strategy)
- Owner (person)
- Due Date (date)
- Notes (text)

Pre-populate with standard GTM tasks:
- Market analysis and sizing
- Competitive landscape research
- Target persona development
- Value proposition definition
- Messaging framework
- Channel strategy
- Launch timeline
- Success metrics definition

#### B. **Competitive Intelligence Database**
Properties:
- Competitor Name (title)
- Category (select: Direct, Indirect, Alternative)
- Strength (select: Strong, Moderate, Weak)
- Market Position (text)
- Key Differentiators (text)
- Pricing Model (text)
- Target Audience (text)
- Last Updated (date)
- Source Links (url)

#### C. **Customer Personas Database**
Properties:
- Persona Name (title)
- Role/Title (text)
- Industry (select: customize based on client)
- Company Size (select: Enterprise, Mid-Market, SMB, Startup)
- Pain Points (text)
- Goals (text)
- Decision Criteria (text)
- Preferred Channels (multi-select: Email, LinkedIn, Events, Search, Content, Partner)
- Priority (select: Primary, Secondary, Tertiary)

#### D. **Content & Deliverables Database**
Properties:
- Title (title)
- Type (select: Blog Post, Case Study, White Paper, Sales Deck, One-Pager, Email Template, Landing Page, Social Post, Video Script)
- Status (select: Planned, In Progress, In Review, Approved, Published)
- Owner (person)
- Due Date (date)
- Target Persona (relation to Personas DB)
- Channel (multi-select)
- URL/Link (url)

#### E. **Meeting Notes Database**
Properties:
- Meeting Title (title)
- Date (date)
- Attendees (text)
- Type (select: Kickoff, Strategy Session, Review, Check-in, Workshop)
- Action Items (text)
- Key Decisions (text)
- Next Steps (text)

### 3. Template Pages

Create these template pages as sub-pages:

#### **Weekly Status Report Template**
Structure:
- Summary/Highlights
- Completed This Week
- In Progress
- Blockers/Risks
- Next Week Priorities
- Questions for Client

#### **GTM Strategy Document Template**
Structure:
- Executive Summary
- Market Analysis
- Target Personas
- Value Proposition
- Positioning & Messaging
- Channel Strategy
- Success Metrics
- Timeline & Milestones

#### **Launch Checklist Template**
Pre-populated checklist with:
- Pre-launch tasks (30/60/90 days out)
- Launch week tasks
- Post-launch tasks
- Stakeholder communication plan
- Content publishing schedule
- Monitoring & optimization plan

### 4. Resources Page

Create a "Resources & Links" page with sections for:
- Brand guidelines
- Product documentation
- Research reports
- Useful tools
- Internal team contacts

## Execution Steps

1. **Confirm details** - Verify all required information with the user
2. **Create parent page** - Set up the main client project page with overview content
3. **Create databases** - Build each database with the specified properties
4. **Add sample data** - Pre-populate the GTM Strategy database with standard tasks
5. **Create templates** - Set up the template pages
6. **Create resources page** - Add the resources hub
7. **Link everything** - Add navigation links on the main page to all databases and templates
8. **Confirm completion** - Provide the user with the link to the main client page

## Best Practices

- Use consistent naming conventions: `[Client Name] - [Section Name]`
- Add helpful descriptions to each database explaining its purpose
- Use emoji icons to make pages visually distinct (🎯 for Strategy, 👥 for Personas, 📊 for Competitive, 📝 for Content, 🗓️ for Meetings)
- Set up linked database views on the main page for quick access to active tasks
- Add helpful inline comments or callouts with instructions for using each section

## After Setup

Ask the user if they want to:
- Customize any database properties for this specific client
- Add any additional sections or databases
- Pre-populate any content based on initial discovery
- Set up automations or reminders
- Share the workspace with specific team members

## Notes

- All databases should be created as full-page databases (not inline) for maximum flexibility
- Use relations between databases where it makes sense (e.g., Content → Personas, Tasks → Owner)
- Keep the structure consistent across clients for easy team onboarding
- The goal is to have everything ready for Day 1 of client work
