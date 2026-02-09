# ICP Output Template

Generate `icp-mapping.md` following this structure. Every section is consumed by Phase 3 workflow skills.

```markdown
# {Company Name} — ICP Mapping

**Date:** {date}
**Based on:** profile.md + strategy.md + market-mapping.md

---

## 1. Account Scoring Formula

### Scoring Weights

| Factor | Weight | Max Points | Data Source | Logic |
|---|---|---|---|---|
| {factor} | {%} | {points} | {source} | {exact scoring rules} |
| ... | | | | |
| **Total** | **100%** | **{max}** | | |

### Tier Definitions

| Tier | Score Range | Action | Expected % of SAM |
|---|---|---|---|
| A+ | {range} | All trigger plays + general outbound + priority follow-up | ~{%} |
| A | {range} | Trigger plays + selective general outbound | ~{%} |
| B | {range} | Trigger-only (no proactive outbound) | ~{%} |
| DQ | Below {threshold} | Auto-remove | ~{%} |

### Instant Disqualifiers
- {DQ reason 1}
- {DQ reason 2}
- {DQ reason 3}

### Clay Formula

```
{Ready-to-paste Clay formula or pseudo-code that implements the scoring}
```

---

## 2. Persona Cards

{Repeat this block for each persona}

### Persona: {Title}

**Seniority:** {level}
**Department:** {dept}
**Reports to:** {who}
**Alternative titles:** {variant_1}, {variant_2}, {variant_3}

#### Priorities & KPIs
- {What they're measured on}
- {What keeps them up at night}
- {What gets them promoted}

#### Day-to-Day Pains
- {Specific pain 1}
- {Specific pain 2}
- {Specific pain 3}

#### Pain → Outcome Mapping (with {Client Product})
| Their Pain | Our Outcome | Proof Point |
|---|---|---|
| {pain} | {outcome} | {evidence} |
| {pain} | {outcome} | {evidence} |

#### Buying Role & Objections
- **Role:** {Decision maker / Influencer / Champion}
- **Top objection:** "{objection}" → **Counter:** {response}
- **Convincing evidence:** {what works — ROI, peer reference, pilot}

#### Messaging Rules
- **Lead with:** {what resonates}
- **Avoid:** {what doesn't work for this persona}
- **Tone:** {how formal/casual to be}

---

## 3. Industry Pain Mapping

{Repeat this block for each target vertical}

### {Industry Name}

#### Current Context
{2-3 sentences: what's happening in this industry NOW that creates urgency}

#### Pain Points (Ranked by Prevalence)

**1. {Pain Point Name}**
- What it looks like: {description}
- Business impact: {cost/risk/time impact}
- Detection signal: {how to spot in enrichment data — job posts, tech stack, news}
- Copy hook: "{one-line outbound angle}"

**2. {Pain Point Name}**
{Same structure}

**3. {Pain Point Name}**
{Same structure}

#### Industry Language Guide
| They Say | We Don't Say |
|---|---|
| {their term} | {generic B2B term to avoid} |
| {their jargon} | {our jargon that sounds outsider} |

#### Vertical-Specific Buying Triggers
- {Trigger + why it creates urgency}
- {Trigger + why it creates urgency}

#### Hook Type Mapping for This Vertical
*(See `shared/references/hook-types-guide.md` for hook type definitions)*

| Hook Type | Available? | Proof Point | Example Opener |
|---|---|---|---|
| Timeline | {yes/no} | {e.g., "Audit → full visibility in 6 weeks"} | "{industry} teams typically go from [X] to [Y] in [timeframe]" |
| Numbers | {yes/no} | {e.g., "40% reduction in review cycles"} | "{persona} teams see [metric] improvement when they [action]" |
| Social proof | {yes/no} | {e.g., "[Company] saw [result]"} | "[Named company] tackled [challenge] and saw [result]" |
| Hypothesis | {always yes — fallback} | N/A | "Curious whether [pain] is something your team is dealing with" |

**Default hook for this vertical:** {timeline / numbers / social_proof / hypothesis}
**Rationale:** {why this is the strongest hook for this specific vertical}

#### Signal Stacking Notes for This Vertical
- **Highest-converting signal combinations:**
  - {e.g., "Champion + Funding = 🔴 Hot — former customers at newly funded companies"}
  - {e.g., "Website (pricing) + Job posting = 🟠 Warm — active evaluation + hiring for function"}
- **Signals to prioritize:** {which triggers produce the strongest signals in this vertical}
- **Signals with limited value:** {any triggers that don't work well in this vertical and why}

#### Common Objections in This Vertical
| Objection | Counter | Evidence |
|---|---|---|
| "{objection}" | {response approach} | {supporting proof} |

#### Proof Points That Land
- {What type of evidence works here}
- {Specific client assets that apply}

---

## 4. Persona × Vertical Matrix

| | {Vertical 1} | {Vertical 2} | {Vertical 3} |
|---|---|---|---|
| **{Persona 1}** | Lead: {angle} / Pain: {pain} | Lead: {angle} / Pain: {pain} | Lead: {angle} / Pain: {pain} |
| **{Persona 2}** | Lead: {angle} / Pain: {pain} | Lead: {angle} / Pain: {pain} | Lead: {angle} / Pain: {pain} |
| **{Persona 3}** | Lead: {angle} / Pain: {pain} | Lead: {angle} / Pain: {pain} | Lead: {angle} / Pain: {pain} |

---

## 5. Fallback Copy Elements

When enrichment returns insufficient data for full personalization, use these industry-specific fallbacks instead of generic B2B language.

### By Vertical

| Vertical | Fallback Observation | Fallback Pain | Fallback CTA |
|---|---|---|---|
| {Vertical 1} | "{industry-specific observation}" | "{most common pain}" | "{stage-appropriate CTA}" |
| {Vertical 2} | "{observation}" | "{pain}" | "{CTA}" |

### By Persona

| Persona | Fallback Opening Angle | Fallback Value Prop |
|---|---|---|
| {Persona 1} | "{what to lead with when no trigger}" | "{persona-specific value}" |
| {Persona 2} | "{angle}" | "{value}" |

---

## 6. Anti-Patterns & Exclusions

### Companies to Avoid
- {Pattern to exclude — e.g., "consulting firms that look like target but aren't buyers"}
- {Pattern — e.g., "companies already using competitor X where switching cost is too high"}

### Persona Anti-Patterns
- {Title that looks right but isn't — e.g., "Operations Manager at <10 person company is actually the founder"}
- {Misleading signal — e.g., "IT Director at a bank is not our buyer, we need Operations"}

---

*This document feeds all Phase 3 workflow skills. Every Clay prompt for copy generation reads persona cards, pain mapping, and the persona × vertical matrix to produce client-specific outbound.*
```
