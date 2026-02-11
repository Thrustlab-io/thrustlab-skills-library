# Research Protocol & Quality Checkpoints

## Phase 1: Profile Discovery Research

### Objective
Build a comprehensive company profile from scratch using only a LinkedIn Company URL and domain. This runs during Step 2 of the workflow.

### Research Sequence
1. **Website analysis:** Fetch all key pages (/, /about, /product or /platform, /pricing, /customers, /case-studies, /blog). Extract: product description, value proposition language, customer logos, case study details, pricing model, team page, integrations/partners.
2. **LinkedIn company page:** Extract: employee count, industry classification, HQ location, founding year, description, specialties, follower count.
3. **Competitive intelligence:** Search `"{company_name}" vs` and `"{company_name}" alternative` to identify competitors. For each competitor found, note positioning difference.
4. **Review platforms:** Search G2, Capterra, TrustRadius for customer sentiment, strengths/weaknesses, and competitive positioning.
5. **Funding & stage:** Search Crunchbase, press releases for funding rounds, investor info, employee growth trajectory.
6. **Tone analysis:** From website copy and blog posts, characterize the brand voice (formal/conversational/provocative), identify preferred terminology, note words they use frequently.

### Profile Completeness Checklist
After discovery research, check every field in `references/profile-field-reference.md`. For any field still empty:
- If a reasonable default exists for this company type/industry → fill it and mark `[Inferred — verify with client]`
- If no reasonable default exists → mark `[To be confirmed by client]`
- Never leave a field completely blank

### Transition to Strategy Research
Once profile.md is saved, continue to Phase 2 (Strategy Research) below. The strategy research uses the profile data as the foundation for deeper market and persona analysis.

---

## Phase 2: Strategy Research

### Web Research Sequence
1. **Primary:** Fetch client website — analyze product pages, pricing, case studies, about page, blog
2. **Competitive:** Search each competitor — positioning, pricing, target market, weaknesses
3. **Market (per vertical):** For EACH ranked vertical from the profile, perform a separate research pass:
   - Search `"{vertical}" trends {year}` and `"{vertical}" challenges {year}` — identify 2-3 macro trends (regulatory changes, technology shifts, market pressures)
   - Search `"{vertical}" {geography} outlook OR forecast` — understand the market direction
   - Connect trends to client's product — articulate why prospects in this vertical need to act now
   - Identify vertical-specific pain points beyond the general persona pains
   - Research buying behavior: typical procurement process, budget cycles, compliance requirements
   - Do this for each vertical separately — trends in healthcare are completely different from trends in manufacturing
4. **Persona:** Search target titles — what they care about, how they buy, communities they're in
5. **Triggers:** Search for signal sources — which events are detectable and actionable
6. **Hook Type Data:** Extract from case studies — timeline milestones, specific metrics, named results. For each vertical, identify: (a) fastest time-to-value story, (b) strongest quantified outcome, (c) most referenceable named customer. See `shared/references/hook-types-guide.md`.
7. **Champion Potential:** Check if client has CRM data on past customers/prospects. Assess volume and data quality for champion tracking viability.
8. **Competitor Detection:** For each competitor, check if detectable via BuiltWith/tech scanning, job posting keywords, or G2/review platforms. Note detection method and confidence.
9. **Dark Funnel Readiness:** Identify existing signal tools (website visitor ID, LinkedIn monitoring, community presence). Note client's primary market (EU/US/Global) for tool recommendations. See `shared/references/clay-enrichment-guide.md` dark funnel section.

### Search Query Patterns (Always Industry-Specific)
```
"{company_name}" + "{vertical}" + {geography}
"{competitor}" + vs + "{company_name}"
"{persona_title}" + "{vertical}" + priorities
site:linkedin.com "{persona_title}" + "{vertical}" + {geography}
```

### Per-Vertical Trend Research Queries (run for EACH vertical)
```
"{vertical}" + trends + {year}
"{vertical}" + challenges + {year} + {geography}
"{vertical}" + compliance + regulations + {geography}
"{vertical}" + technology + trends + adoption
"{vertical}" + market + outlook + forecast + {year}
"{vertical}" + buying + process OR procurement
"{vertical}" + budget + priorities + {year}
```

### Citation Standard
Every factual claim must have: `[Source: URL]`
If multiple sources confirm: `[Sources: URL1, URL2]`
If inferred from research: `[Inferred from: brief reasoning]`

## Quality Checkpoints — Run Before Delivering

### Completeness
- [ ] profile.md was saved before strategy generation began
- [ ] No section is shorter than 3 paragraphs (except Company Snapshot)
- [ ] Top 3 recommended trigger plays are clearly identified with rationale
- [ ] At least 10 triggers in the playbook across 3 tiers
- [ ] Signal stacking section identifies high-converting combinations for this client
- [ ] Hook type recommendations populated with actual client proof points (not placeholders)

### Client Specificity
- [ ] Every recommendation ties back to profile.md data or web research
- [ ] Zero generic value props — all reference this client's specific product
- [ ] Pain points are for THIS client's personas, not generic B2B
- [ ] Competitor analysis uses real competitors from profile, not placeholders

### Actionability
- [ ] Boolean search strings are ready to paste into Sales Navigator / Apollo
- [ ] Clay enrichment fields are named with specific providers
- [ ] 90-day blueprint has concrete weekly milestones
- [ ] Every trigger has a named signal source and detection method


### Downstream Readiness
- [ ] Strategy clearly feeds into Phase 1 tooling setup (which tools to buy)
- [ ] ICP criteria are specific enough to create Clay scoring formula
- [ ] Trigger playbook is specific enough to build individual workflow skills

