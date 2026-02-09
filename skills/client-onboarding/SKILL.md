---
name: client-onboarding
description: Creates a structured client profile from intake data or discovery call notes. Use when onboarding a new Thrustlab GTM client, when a client provides new information that should update their profile, or when starting any new client engagement. Produces the canonical profile.md that ALL downstream skills read from.
---

# Client Onboarding

Creates the canonical client profile at `client-profiles/{client-slug}/profile.md`.

This file is the single source of truth for every downstream skill. Nothing ships without it.

## Workflow

### Step 1: Gather Intake Data

If the user provides a client website URL or unstructured notes, research the company first:
1. Fetch the client website for primary intelligence
2. Run 2-3 targeted web searches for additional context
3. Extract as much intake data as possible before asking follow-up questions

For any intake fields not answerable from research, ask the user. See `references/intake-questions.md` for the full field list and guidance on each field.

Prioritize getting these critical fields before proceeding:
- Company name & website
- Product/service description
- Target verticals
- Target personas
- Key differentiators

### Step 2: Generate Client Slug

Format: lowercase, hyphens, no spaces.
- "Acme Corp" → `acme-corp`
- "Quality Guard" → `quality-guard`
- "DataFlow.io" → `dataflow`

### Step 3: Create Profile Directory

```
client-profiles/{client-slug}/
└── profile.md
```

### Step 4: Write profile.md

Structure the profile using this exact format:

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

## Tech Stack
- **CRM:** {Salesforce / HubSpot / Pipedrive / other}
- **Email sequencer:** {Smartlead / Instantly / Outreach / other / none yet}
- **Intent/signal tools:** {existing tools, if any}
- **Other relevant tools:** {marketing automation, analytics, etc.}

## Existing Assets
- **Case studies:** {list with URLs if available}
- **Social proof:** {notable logos, metrics, quotes}
- **Content library:** {blog, resources, webinars — what exists}
- **LinkedIn presence:** {company page URL, key team members}

## Hook Type Readiness
*(Determines which hook patterns are available — directly impacts reply rates)*

- **Timeline proof points by vertical:**
  - {vertical_1}: "{discovery_to_result} timeline — e.g., 'From first audit to full visibility in 6 weeks'"
  - {vertical_2}: "{timeline}"
- **Numbers proof points:**
  - "{specific metric — e.g., '40% reduction in review cycles'}"
  - "{specific metric}"
- **Named case studies (referenceable in copy):**
  - {company_1} ({vertical}): "{result}"
  - {company_2} ({vertical}): "{result}"
- **Switch stories (customers who switched FROM a competitor):**
  - Switched from {competitor}: {company} — "{result}"

## Champion Tracking Readiness
*(Required for the highest-converting signal type)*

- **CRM contact export available:** {yes/no}
- **Estimated past customer contacts:** {number}
- **Estimated past prospect contacts (meeting stage+):** {number}
- **UserGems or job change tracking in place:** {yes/no/planned}
- **CRM fields available:** {relationship_type, product_used, result_achieved — which exist?}

## Competitor Detection
- **Competitors detectable via BuiltWith/tech scanning:** {yes for which competitors / no}
- **Competitors detectable via job postings:** {yes — which keywords / no}
- **Strongest switch stories:**
  - From {competitor_1}: {company} saw {result}
  - From {competitor_2}: {company} saw {result}

## Dark Funnel & Signal Preferences
- **Website visitor ID tool in place:** {Dealfront / RB2B / none / planned}
- **LinkedIn engagement monitoring:** {Teamfluence / Trigify / none / planned}
- **Community monitoring:** {Common Room / none / planned}
- **Primary market for dark funnel tooling:** {EU / US / Global}
- **Content publishing frequency:** {weekly / monthly / sporadic / none}
- **Content topics:** {list main content themes}

## Engagement Parameters
- **Clay plan tier:** {starter / explorer / pro / enterprise}
- **Expected sending volume:** {emails per week/month}
- **Budget considerations:** {any relevant constraints}
- **Timeline:** {launch urgency}

---
*Profile created: {date}*
*Last updated: {date}*
```

### Step 5: Validate Completeness

Before saving, check:
- Every section has client-specific data (no placeholders like "TBD" or "N/A" for critical fields)
- Personas have actual pain points, not generic ones
- Competitors are real companies, not placeholders
- Value prop is specific to THIS client, not generic B2B language

For non-critical fields where data isn't available yet, mark as `[To be determined in strategy phase]` — the strategy generator will research and fill these.

### Step 6: Confirm with User

Present a summary of the profile and ask:
- "Does this accurately capture your product and positioning?"
- "Are the personas and pain points correct?"
- "Anything missing or wrong?"

Update based on feedback before proceeding.

## Output

Final deliverable: `client-profiles/{client-slug}/profile.md`

This file is now the input for `/gtm-strategy-generator`.
