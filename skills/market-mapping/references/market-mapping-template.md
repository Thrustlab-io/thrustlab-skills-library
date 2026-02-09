# Market Mapping Output Template

Generate the client's `market-mapping.md` following this structure.

```markdown
# {Company Name} — Market Mapping

**Date:** {date}
**Based on:** profile.md + strategy.md

---

## 1. Market Summary

**Total Addressable Market (TAM):** {estimate} companies
**Serviceable Available Market (SAM):** {estimate} companies
**Initial Target List Goal:** {number} accounts for first 90 days

### SAM Filtering Criteria Applied
- Geography: {geographies}
- Company size: {employee range}
- Revenue: {revenue range, if applicable}
- Verticals: {vertical_1}, {vertical_2}, {vertical_3}
- Exclusions: {what's filtered out and why}

---

## 2. Search Strings — LinkedIn Sales Navigator

### {Vertical 1} — Broad (TAM-level)
```
{Full boolean search string}
```
Estimated results: ~{number}

### {Vertical 1} — Filtered (SAM-level)
```
{Filtered boolean with additional qualifiers}
```
Estimated results: ~{number}

### {Persona 1} — Cross-Vertical
```
{Persona-specific search across verticals}
```
Estimated results: ~{number}

{Repeat for each vertical and persona}

---

## 3. Search Strings — Apollo

### {Vertical 1}
```
Person titles: {titles}
Company industry: {industries}
Company size: {range}
Location: {geo}
Additional filters: {any Apollo-specific filters}
```
Estimated results: ~{number}

{Repeat for each vertical}

---

## 4. Clay Import Configuration

### Import from LinkedIn Sales Navigator
- Export via: {method — manual CSV, PhantomBuster, Clay integration}
- Fields to capture: First name, Last name, Title, Company, Company URL, LinkedIn URL, Location
- Dedup rule: Match on company domain (not name — avoids "Inc" vs "Inc." mismatches)
- Post-import enrichment: Run Clearbit/Apollo company enrichment → fill industry, size, revenue, tech stack

### Import from Apollo
- Export via: {method}
- Fields to capture: {field list}
- Dedup rule: Same domain-based matching

### Import from Alternative Sources
{For each alternative source discovered in Step 4 of SKILL.md}
- Source: {name}
- Import method: {CSV upload / API / scraping with tool}
- Fields available: {what the source provides}
- Fields missing (need enrichment): {what Clay needs to fill}

---

## 5. Alternative & Niche Data Sources

{For each discovered source}

### {Source Name}
- **URL:** {link}
- **What it contains:** {description}
- **Volume:** ~{number} records
- **Geography:** {coverage}
- **Data freshness:** {how current}
- **Access:** {public/login/paid/API}
- **Import method:** {how to get into Clay}
- **Relevance:** {High/Medium/Low} — {why}

---

## 6. Enrichment Recommendations

### Priority Enrichment After Import

| Step | Provider | Purpose | Run On |
|---|---|---|---|
| 1 | Clearbit / Apollo | Company firmographics (size, revenue, industry) | All imported accounts |
| 2 | {Provider} | {Purpose specific to this client's ICP} | Accounts missing {field} |
| 3 | BuiltWith / Wappalyzer | Tech stack (if tech triggers are in play) | SAM-qualified accounts |
| 4 | Hunter / Apollo | Contact email finding for target persona | Accounts with ICP score ≥ {threshold} |
| 5 | Email verification | Verify deliverability | All found emails |

### SAM Qualification Formula (for Clay)

```
IF company_size >= {min} AND company_size <= {max}
AND company_industry IN ({industry_list})
AND company_geography IN ({geo_list})
AND NOT company_name IN ({exclusion_list})
THEN "SAM_QUALIFIED"
ELSE "DISQUALIFIED"
```

---

## 7. Source Priority Ranking

| Rank | Source | Est. Volume | Data Quality | Recommended Action |
|---|---|---|---|---|
| 1 | {best source} | {volume} | High | Import first, primary list |
| 2 | {second source} | {volume} | High | Supplement after primary |
| 3 | {third source} | {volume} | Medium | Cross-reference for coverage |
| ... | | | | |

---

## Next Step

This market mapping feeds into ICP Mapping (Phase 2.2), where we:
- Score and tier the imported accounts
- Build persona cards with role-specific pain points
- Create industry-specific messaging angles
```
