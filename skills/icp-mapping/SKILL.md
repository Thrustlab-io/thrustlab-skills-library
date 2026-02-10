---
name: icp-mapping
description: Refines the Ideal Customer Profile with account scoring, persona cards, and industry-specific pain mapping. Use after market mapping is complete. Produces the detailed ICP document that all Phase 3 workflow skills read from for copy generation. This is the critical bridge between strategy and outbound — it translates market data into messaging ammunition.
---

# ICP Mapping

Produces the refined ICP saved to `Prospects/{client-slug}/icp-mapping.md`.

This is the most copy-critical Phase 2 output. Every workflow skill in Phase 3 reads this file to generate personalized Clay prompts and outbound copy. Invest time here — weak ICP mapping = generic outbound.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md` exists
- `Prospects/{client-slug}/strategy.md` exists (ICP section, messaging architecture)
- `Prospects/{client-slug}/market-mapping.md` exists (data sources, enrichment fields)

## Workflow

### Step 1: Load All Upstream Data

Read and cross-reference:
- **profile.md** — Product, value prop, differentiators, competitors, tone
- **strategy.md** — ICP criteria, persona definitions, messaging architecture, pain→outcome mapping, trigger playbook
- **market-mapping.md** — Available data sources, enrichment fields, SAM criteria, alternative sources discovered

### Step 2: Define Account Scoring Formula

Build a weighted scoring formula for Clay. Every factor must map to an enrichable field.

**Framework:**

| Factor | Weight | Data Source | Scoring Logic |
|---|---|---|---|
| Company size | {weight} | Clearbit/Apollo | {exact range = max points, adjacent = partial} |
| Industry match | {weight} | Clearbit/Apollo | {primary vertical = max, adjacent = partial} |
| Geography | {weight} | Clearbit/Apollo | {target geo = max, adjacent = partial} |
| Tech stack signals | {weight} | BuiltWith/Wappalyzer | {uses complementary tech = bonus, uses competitor = flag} |
| Growth signals | {weight} | Clay signals | {hiring surge, funding, expansion = bonus points} |
| Revenue range | {weight} | Clearbit/Apollo | {sweet spot = max, too small/large = penalty} |
| Content engagement | {weight} | HubSpot/website | {any engagement = bonus} |

**Tier definitions:**
- **A+ (Score ≥ X):** Perfect ICP fit. Prioritize for all trigger plays + general outbound.
- **A (Score X-Y):** Strong fit. Include in trigger plays, selective in general outbound.
- **B (Score X-Y):** Partial fit. Only include if triggered (don't proactively outbound).
- **Disqualified (Score < X):** Auto-remove from all workflows.

**Disqualifiers (instant DQ regardless of score):**
- {List specific disqualifiers — e.g., "government entities", "companies < 10 employees", "already a customer", "competitor"}

### Step 3: Build Persona Cards

For each target persona defined in the strategy, create a detailed card.

**Per persona, document:**

```
### Persona: {Persona Title}
**Seniority:** {C-level / VP / Director / Manager}
**Department:** {Operations / IT / Finance / etc.}
**Reports to:** {Who they report to — relevant for multi-threading}
**Common alternative titles:** {Title variants to search for}

#### What They Care About
- {Priority 1 — tied to their KPIs}
- {Priority 2}
- {Priority 3}

#### Day-to-Day Pains
- {Specific operational pain 1 — not generic "efficiency"}
- {Specific pain 2 — grounded in what this role actually deals with}
- {Specific pain 3}

#### How {Client Product} Helps Them
- Pain: {pain} → Outcome: {specific outcome with client's product}
- Pain: {pain} → Outcome: {specific outcome}

#### Buying Role
- {Decision maker / Influencer / Champion / End user}
- {What objections they typically raise}
- {What evidence convinces them — ROI calc, peer reference, pilot, etc.}

#### Messaging Angle
- **Opening approach:** {What to lead with for this persona}
- **Avoid:** {What turns this persona off — e.g., "don't lead with cost savings for a CTO"}
```

### Step 4: Industry-Specific Pain Mapping

This is what makes outbound copy feel "insider" rather than generic. For each target vertical from market-mapping.md:

**Research methodology:**
1. Web search: `"{vertical}" challenges {year}` and `"{vertical}" trends {year}`
2. Web search: `"{vertical}" + "{persona_title}" + frustrations OR pain points`
3. Web search: `"{vertical}" + technology adoption OR digital transformation`
4. Check alternative sources from market-mapping.md for industry context
5. Review client's competitor content for pain language they use

**Per vertical, document:**

```
### Vertical: {Industry Name}

#### Industry Context
{2-3 sentences on current state — what's happening in this industry right now that creates urgency for the client's solution}

#### Top Pain Points (Ranked)
1. **{Pain Point 1}** — {Why it's painful, how common, what it costs them}
   - Signal in data: {How you'd detect this in enrichment — e.g., "job posting for data analyst", "uses legacy ERP"}
   - Copy hook: {One-line angle for outbound copy}
2. **{Pain Point 2}** — {Same structure}
3. **{Pain Point 3}** — {Same structure}

#### Industry-Specific Language
- They say: "{term they use}" (not "{generic B2B term}")
- They say: "{their jargon}" (not "{our jargon}")
- Avoid: "{terms that mark you as an outsider}"

#### Buying Triggers Specific to This Vertical
- {Trigger 1 — e.g., "regulatory deadline approaching", "seasonal peak coming"}
- {Trigger 2 — e.g., "competitor just adopted similar tech"}
- {Trigger 3 — e.g., "new compliance requirement announced"}

#### Objection Patterns
- Common objection: "{objection}" → Counter: "{how to address}"
- Common objection: "{objection}" → Counter: "{how to address}"

#### Proof Points That Resonate
- {Type of social proof that works — e.g., "ROI calculations", "peer company case study", "industry benchmark"}
- {Specific proof from client's assets — e.g., "case study with [similar company]"}
```

### Step 5: Cross-Reference Personas × Verticals

Create a matrix showing how pains and messaging shift when the same persona appears in different verticals:

```
| Persona | Vertical A | Vertical B | Vertical C |
|---|---|---|---|
| {Title 1} | Lead with: {angle} | Lead with: {angle} | Lead with: {angle} |
| {Title 2} | Lead with: {angle} | Lead with: {angle} | Lead with: {angle} |
```

This matrix directly feeds the Clay prompts in Phase 3 — the AI opener prompt can select the right angle based on `{{contact_title}}` + `{{company_industry}}`.

### Step 6: Generate Fallback Content for Copy

For each persona × vertical combination, pre-generate fallback copy elements that workflow skills use when enrichment data is thin:

- **Fallback observation:** Industry-specific (not "your recent growth")
- **Fallback pain assumption:** Most common pain for this persona in this vertical
- **Fallback CTA:** Stage-appropriate for this persona's buying role

### Step 7: Write icp-mapping.md

Save to `Prospects/{client-slug}/icp-mapping.md` using the structure in `references/icp-output-template.md`.

### Step 8: Validate

Present to user:
- "Do these persona cards match the buyers you actually sell to?"
- "Are the industry pain points accurate from your experience?"
- "Any vertical-specific language I'm missing?"
- "Do the scoring weights reflect your actual deal patterns?"
