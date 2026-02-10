---
name: market-mapping
description: Builds the Total Addressable Market to Serviceable Available Market (TAM→SAM) mapping for a client. Use after the strategy is generated. Produces boolean search strings, Clay import criteria, enrichment recommendations, and discovers alternative/niche data sources specific to the client's vertical. Reads from profile.md and strategy.md.
---

# Market Mapping

Produces the complete market mapping saved to `Prospects/{client-slug}/market-mapping.md`.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- `Prospects/{client-slug}/strategy.md` exists (specifically: TAM→SAM filtering section, ICP criteria)

## Workflow

### Step 1: Load Client Data

Read from profile.md and strategy.md:
- Target verticals (ranked)
- Company size range
- Revenue range (if applicable)
- Target geographies
- Exclusions
- ICP criteria and scoring signals
- TAM→SAM filtering section from strategy

### Step 2: Build Standard Search Strings

For each target vertical, create ready-to-paste boolean searches:

**LinkedIn Sales Navigator format:**
```
(Title: "{persona_title_1}" OR "{persona_title_2}" OR "{persona_title_3}")
AND (Industry: "{industry_1}" OR "{industry_2}")
AND (Company headcount: {range})
AND (Geography: "{geo_1}" OR "{geo_2}")
```

**Apollo format:**
```
Person titles: {title_1}, {title_2}, {title_3}
Company industry: {industry_1}, {industry_2}
Company size: {min}-{max} employees
Location: {geography}
```

Create at minimum:
- 1 broad search per primary vertical (TAM-level)
- 1 filtered search per primary vertical (SAM-level, with additional qualifiers)
- 1 persona-specific search per target persona

### Step 3: Define Clay Import Criteria

For each search source, define:
- Which fields to import
- What enrichment to run immediately after import
- Deduplication rules (match on domain, not company name)
- Initial qualification filters (auto-disqualify companies that don't match SAM)

### Step 4: Discover Alternative & Niche Data Sources

This is critical — especially for SMBs and regulated industries where LinkedIn/Apollo coverage is weak.

**Research methodology:**
1. Web search: `"{vertical}" + database + directory + {geography}`
2. Web search: `"{vertical}" + licensed operators + registry + {geography}`
3. Web search: `"{vertical}" + association + member list + {geography}`
4. Web search: `"{vertical}" + industry events + conference + attendees + {geography}`
5. Web search: `"find {vertical} companies" + {geography}`
6. Check for government/regulatory databases relevant to the vertical
7. Check for industry-specific review platforms or directories

**For each discovered source, document:**

| Field | What to Capture |
|---|---|
| Source name | Name of the database/directory |
| URL | Direct link |
| What data it contains | Companies, contacts, certifications, etc. |
| Estimated volume | How many records approximately |
| Geography coverage | Which regions it covers |
| Data freshness | How often updated |
| Access method | Public/free, login required, API, scraping needed |
| Import feasibility | Can we get this into Clay? How? (CSV export, scraping, API, manual) |
| Relevance score | High/Medium/Low for this client's ICP |

**Examples of what to look for:**
- Government registries (food operators, licensed contractors, certified companies)
- Local business directories (5Bizzy for Belgian SMBs, Kompass, Europages)
- Industry association member lists
- Trade show exhibitor/attendee lists
- Certification body databases
- Industry-specific platforms (G2 for SaaS, Clutch for agencies, etc.)
- Regional business chambers of commerce

### Step 5: Market Size Estimates

For each data source (standard + alternative):
- Estimated number of companies matching SAM criteria
- Expected contact yield (% of companies where we can find target persona)
- Source quality ranking (which sources have most accurate/complete data)

### Step 6: Write market-mapping.md

Save to `Prospects/{client-slug}/market-mapping.md`. See `references/market-mapping-template.md` for the output structure.

### Step 7: Validate with User

Present findings and ask:
- "Do these search strings match how you think about your market?"
- "Are there any industry-specific databases or directories you already know about?"
- "Any verticals or company types missing from the mapping?"
