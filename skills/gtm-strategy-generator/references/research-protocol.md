# Research Protocol & Quality Checkpoints

## Research Execution Rules

### Web Research Sequence
1. **Primary:** Fetch client website — analyze product pages, pricing, case studies, about page, blog
2. **Competitive:** Search each competitor — positioning, pricing, target market, weaknesses
3. **Market:** Search target verticals — trends, challenges, regulations, market size
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
"{vertical}" + challenges + {year}
"{persona_title}" + "{vertical}" + priorities
"{vertical}" + compliance + regulations + {geography}
"{vertical}" + technology + trends + adoption
site:linkedin.com "{persona_title}" + "{vertical}" + {geography}
```

### Citation Standard
Every factual claim must have: `[Source: URL]`
If multiple sources confirm: `[Sources: URL1, URL2]`
If inferred from research: `[Inferred from: brief reasoning]`

## Quality Checkpoints — Run Before Delivering

### Completeness
- [ ] All 13 sections + 2 appendices are present
- [ ] No section is shorter than 3 paragraphs (except Company Snapshot)
- [ ] Top 3 recommended trigger plays are clearly identified with rationale
- [ ] At least 10 triggers in the playbook across 3 tiers
- [ ] Signal stacking section identifies high-converting combinations for this client
- [ ] Hook type recommendations populated with actual client proof points (not placeholders)
- [ ] Three new signal types evaluated: champion tracking, competitor customer, dark funnel

### Client Specificity
- [ ] Every recommendation ties back to profile.md data or web research
- [ ] Zero generic value props — all reference this client's specific product
- [ ] Pain points are for THIS client's personas, not generic B2B
- [ ] Competitor analysis uses real competitors from profile, not placeholders
- [ ] Copy examples use client's tone preference and banned word list

### Actionability
- [ ] Boolean search strings are ready to paste into Sales Navigator / Apollo
- [ ] Clay enrichment fields are named with specific providers
- [ ] Cadence steps have specific timing and copy examples
- [ ] 90-day blueprint has concrete weekly milestones
- [ ] Every trigger has a named signal source and detection method

### Copy Quality
- [ ] Email subjects ≤45 characters
- [ ] Email bodies ≤90 words
- [ ] LinkedIn messages ≤280 characters
- [ ] No banned words/phrases (see shared/references/copy-rules.md)
- [ ] All openers are observational, not generic
- [ ] CTAs are stage-appropriate

### Downstream Readiness
- [ ] Strategy clearly feeds into Phase 1 tooling setup (which tools to buy)
- [ ] ICP criteria are specific enough to create Clay scoring formula
- [ ] Trigger playbook is specific enough to build individual workflow skills
- [ ] Messaging architecture provides enough raw material for copy generation
