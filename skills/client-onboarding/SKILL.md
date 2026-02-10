---
name: client-onboarding
description: Creates a structured client profile from intake data or discovery call notes. Use when onboarding a new Thrustlab GTM client, when a client provides new information that should update their profile, or when starting any new client engagement. Produces the canonical profile.md that ALL downstream skills read from and intake-questions.md that will further guide the discovery process.
---

# Client Onboarding

Creates the canonical client profile at `Prospects/{client-slug}/profile.md`.
Creates a list of follow-up questions to fill in by the client at `Prospects/{client-slug}/intake-questions.md`.

## Workflow

First check if the intake-questions.md already exists and if it contains answers and if the initial profile.md has already been created. If so, skip to step 6 to update the profile based on any new information. If not, follow the workflow below to create the initial profile and intake questions.
### Step 1: Gather Data from the web

If the user provides a client website URL research the company first:
1. Fetch the client website for primary intelligence
2. Run 2-3 targeted web searches for additional context


### Step 2: Generate Client Slug

Format: lowercase, hyphens, no spaces.
- "Acme Corp" → `acme-corp`
- "Quality Guard" → `quality-guard`
- "DataFlow.io" → `dataflow`

### Step 3: Create Profile Directory

```
Prospects/{client-slug}/
└── profile.md
```
### Step 4: Write first draft of profile.md

```markdown
# {Company Name} — Client Profile

## Company Identity
- **Name:** {company_name}
- **Website:** {website}
- **Slug:** {client-slug}
- **One-liner:** {1-sentence description of what they do}
- **Product/Service:** {2-3 sentence description}
- **Founded:** {year, if known}
- **HQ:** {location}
- **Company size:** {employee count/range}

### Step 5: Generate intake-questions.md
Write the intake question document. Use `references/intake-questions.md` as the source of truth for the questions.
### Step 6: Expand based on answers to intake questions
Update the profile based on the answers to the intake questions.
Use the following structure for the profile, and populate as much as possible based on the answers to the intake questions and any additional research you do. Mark any fields that are still missing information as "TBD" so we can easily identify gaps.
```markdown
## Sales Motion
- **Type:** {PLG / sales-led / hybrid / channel}
- **Average deal size:** {if known}
- **Sales cycle:** {if known}
- **Current outbound status:** {none / early / established}

## Target Market
- **Verticals (ranked):**
  1. {primary vertical}
  2. {secondary vertical}
  3. {tertiary vertical}
- **Company size range:** {e.g., 50-500 employees}
- **Revenue range:** {if relevant}
- **Geographies:** {target markets}
- **Exclusions:** {verticals, company types, or geos to avoid}

## Target Personas
### Primary Persona
- **Title(s):** {e.g., VP of Operations, Head of Quality}
- **Department:** {e.g., Operations, Engineering}
- **Seniority:** {IC / Manager / Director / VP / C-Level}
- **Key responsibilities:** {what they own}
- **Primary pain:** {the #1 problem they face that client solves}
- **Secondary pains:** {other relevant challenges}
- **How they buy:** {research process, stakeholders involved}

### Secondary Persona
{same structure}

### Tertiary Persona (if applicable)
{same structure}

## Value Proposition
- **Core value prop:** {1 sentence — what outcome do they deliver?}
- **Key differentiators:**
  1. {differentiator 1}
  2. {differentiator 2}
  3. {differentiator 3}
- **Proof points:** {case studies, metrics, logos — whatever exists}

## Competitive Landscape
- **Competitors:**
  1. {competitor 1} — {positioning difference}
  2. {competitor 2} — {positioning difference}
  3. {competitor 3} — {positioning difference}
- **Why clients choose {company} over competitors:** {key win themes}

## Tone & Messaging Preferences
- **Tone:** {formal / conversational / provocative}
- **Brand voice notes:** {any specific preferences or restrictions}
- **Words/phrases to avoid:** {client-specific banned terms}
- **Words/phrases to use:** {client-specific preferred language}

## Tech Stack & Infrastructure
- **CRM:** {Salesforce / HubSpot / Pipedrive / other}
- **CRM data quality:** {how clean is the data? Approximate contact count?}
- **Historical contacts in CRM:** {approximate number of past customers, past prospects, past meetings — whatever exists}
- **Email sequencer:** {Smartlead / Instantly / Outreach / other / none yet}
- **Existing signal/intent tools:** {list any tools already in use — e.g., Dealfront, RB2B, Teamfluence, BuiltWith, UserGems, Common Room, etc.}
- **Marketing automation:** {HubSpot Marketing / Marketo / none / other}
- **Other relevant tools:** {analytics, BI, project management, etc.}

## Existing Assets & Proof Points
- **Case studies:**
  {For each, capture: company name (can we reference them by name?), vertical, specific result, timeline to result if known}
  - {company_1} ({vertical}): "{result}" — {timeline if known} — {referenceable by name? yes/no}
  - {company_2} ({vertical}): "{result}" — {timeline if known} — {referenceable by name? yes/no}
- **Quantified metrics:** {any hard numbers — e.g., "40% reduction in X", "6 weeks to full deployment"}
- **Switch stories:** {customers who switched from a specific competitor — what they switched from, why, what result}
- **Customer logos available for use:** {list}
- **Content library:** {blog, resources, webinars, whitepapers — what exists, approximate volume}
- **Content publishing cadence:** {weekly / monthly / sporadic / none}
- **Content topics:** {main themes covered}
- **LinkedIn presence:** {company page URL, follower count if known, key team members who post}
- **Community presence:** {Slack communities, Discord, forums, events — anything where prospects or customers gather}

## Engagement Parameters
- **Clay plan tier:** {starter / explorer / pro / enterprise}
- **Expected sending volume:** {emails per week/month}
- **Budget considerations:** {any relevant constraints}
- **Timeline:** {launch urgency}
- **Primary market:** {EU / US / Global — affects tooling recommendations}

---
*Profile created: {date}*
*Last updated: {date}*
```
## Output

Final deliverable: `Prospects/{client-slug}/profile.md`

This file is now the input for `/gtm-strategy-generator`.
