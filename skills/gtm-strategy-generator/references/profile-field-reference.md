# Profile Field Reference & Template

Complete field schema for `profile.md`. Used by the strategy generator during auto-discovery to ensure every field is populated.

## Discoverability Tiers

Each field is tagged with how it can be sourced:

- **[Auto]** — Directly discoverable from website, LinkedIn, or web search. Populate without annotation.
- **[Infer]** — Can be reasonably guessed from context. Populate and mark `[Inferred — verify with client]`.
- **[Client]** — Cannot be reliably discovered. Populate with best-guess default and mark `[To be confirmed by client]`.

---

## Profile Template

```markdown
# {Company Name} — Client Profile

## Company Identity
- **Name:** {company_name} `[Auto]`
- **Website:** {website} `[Auto]`
- **Slug:** {client-slug} `[Auto]`
- **One-liner:** {1-sentence description of what they do} `[Auto]`
- **Product/Service:** {2-3 sentence description} `[Auto]`
- **Founded:** {year} `[Auto]`
- **HQ:** {location} `[Auto]`
- **Company size:** {employee count/range} `[Auto]`

## Sales Motion
- **Type:** {PLG / sales-led / hybrid / channel} `[Infer]` — infer from website CTAs: "Try free" = PLG, "Book demo" = sales-led, "Find a partner" = channel
- **Average deal size:** {if known} `[Client]`
- **Sales cycle:** {if known} `[Client]`
- **Current outbound status:** {none / early / established} `[Client]`

## Target Market
- **Verticals (ranked):** `[Infer]` — infer from case studies, website copy, industry pages
  1. {primary vertical}
  2. {secondary vertical}
  3. {tertiary vertical}
- **Company size range:** {e.g., 50-500 employees} `[Infer]` — infer from pricing, case study companies, positioning language
- **Revenue range:** {if relevant} `[Infer]`
- **Geographies:** {target markets} `[Infer]` — infer from language, currencies, office locations, compliance references
- **Exclusions:** {verticals, company types, or geos to avoid} `[Client]`

## Target Personas
### Primary Persona
- **Title(s):** {e.g., VP of Operations, Head of Quality} `[Infer]` — infer from case studies, testimonials, website "for" pages
- **Department:** {e.g., Operations, Engineering} `[Infer]`
- **Seniority:** {IC / Manager / Director / VP / C-Level} `[Infer]`
- **Key responsibilities:** {what they own} `[Infer]`
- **Primary pain:** {the #1 problem they face that client solves} `[Infer]` — infer from product messaging, case study language
- **Secondary pains:** {other relevant challenges} `[Infer]`
- **How they buy:** {research process, stakeholders involved} `[Client]`

### Secondary Persona
{same structure}

### Tertiary Persona (if applicable)
{same structure}

## Value Proposition
- **Core value prop:** {1 sentence — what outcome do they deliver?} `[Auto]`
- **Key differentiators:** `[Auto]` — extract from website positioning, "why us" pages
  1. {differentiator 1}
  2. {differentiator 2}
  3. {differentiator 3}
- **Proof points:** {case studies, metrics, logos — whatever exists} `[Auto]`

## Competitive Landscape
- **Competitors:** `[Auto]` — from "vs" searches, G2 alternatives, review platforms
  1. {competitor 1} — {positioning difference}
  2. {competitor 2} — {positioning difference}
  3. {competitor 3} — {positioning difference}
- **Why clients choose {company} over competitors:** {key win themes} `[Infer]`

## Tone & Messaging Preferences
- **Tone:** {formal / conversational / provocative} `[Auto]` — analyze website and blog copy style
- **Brand voice notes:** {any specific preferences or restrictions} `[Auto]`
- **Words/phrases to avoid:** {client-specific banned terms} `[Client]`
- **Words/phrases to use:** {client-specific preferred language} `[Auto]` — extract from website copy patterns

## Tech Stack & Infrastructure
- **CRM:** {Salesforce / HubSpot / Pipedrive / other} `[Client]`
- **CRM data quality:** {how clean is the data? Approximate contact count?} `[Client]`
- **Historical contacts in CRM:** {approximate number of past customers, past prospects, past meetings} `[Client]`
- **Email sequencer:** {Smartlead / Instantly / Outreach / other / none yet} `[Client]`
- **Existing signal/intent tools:** {list any tools already in use} `[Client]`
- **Marketing automation:** {HubSpot Marketing / Marketo / none / other} `[Infer]` — sometimes detectable from website source/tracking scripts
- **Other relevant tools:** {analytics, BI, project management, etc.} `[Client]`

## Existing Assets & Proof Points
- **Case studies:** `[Auto]` — extract from website case study pages
  {For each: company name, vertical, result, timeline, referenceable by name?}
  - {company_1} ({vertical}): "{result}" — {timeline if known} — referenceable: `[Client]`
  - {company_2} ({vertical}): "{result}" — {timeline if known} — referenceable: `[Client]`
- **Quantified metrics:** {any hard numbers} `[Auto]`
- **Switch stories:** {competitor switches} `[Auto]` — if documented on website
- **Customer logos available for use:** {list} `[Auto]` — logos visible on website; permission to use in outbound: `[Client]`
- **Content library:** {blog, resources, webinars, whitepapers — what exists, approximate volume} `[Auto]`
- **Content publishing cadence:** {weekly / monthly / sporadic / none} `[Auto]`
- **Content topics:** {main themes covered} `[Auto]`
- **LinkedIn presence:** {company page URL, follower count, key team members who post} `[Auto]`
- **Community presence:** {Slack communities, Discord, forums, events} `[Infer]`

## Engagement Parameters
- **Clay plan tier:** {starter / explorer / pro / enterprise} `[Client]`
- **Expected sending volume:** {emails per week/month} `[Client]`
- **Budget considerations:** {any relevant constraints} `[Client]`
- **Timeline:** {launch urgency} `[Client]`
- **Primary market:** {EU / US / Global — affects tooling recommendations} `[Infer]`

---
*Profile created: {date}*
*Last updated: {date}*
```

---

## Discovery Research Sequence

Before populating the template, gather data in this order:

1. **Fetch company website** — product pages, pricing, about, case studies, blog, footer/integrations
2. **Parse LinkedIn company page** — size, industry, HQ, description, founded year, specialties
3. **Search: product positioning** — `"{company_name}" product OR solution OR platform`
4. **Search: customers & proof** — `"{company_name}" customers OR "case study" OR testimonial`
5. **Search: competitive landscape** — `"{company_name}" vs OR alternative OR competitor`
6. **Search: reviews** — `"{company_name}" site:g2.com OR site:capterra.com` (if SaaS)
7. **Search: funding & stage** — `"{company_name}" funding OR raised site:crunchbase.com`
8. **Search: market context** — `"{company_name}" {inferred_industry_terms}`

## Completeness Check

After populating, verify:
- [ ] Every `[Auto]` field has a value (no blanks — these are fully discoverable)
- [ ] Every `[Infer]` field has a best-guess value + annotation
- [ ] Every `[Client]` field has either a best-guess default + annotation, or is marked `[To be confirmed by client]`
- [ ] At least 1 persona is defined with title and primary pain
- [ ] At least 2 competitors identified
- [ ] Core value prop is specific to this company, not generic
