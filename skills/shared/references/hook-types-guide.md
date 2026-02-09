# Hook Types Guide — Systematic A/B Testing Framework

## Why This Matters

The Digital Bloom's 2025 study across 4 hook types, 4 buyer profiles, and 4 industries shows massive performance differences between hook types. **Timeline hooks achieve 10.01% reply rate vs. 4.34% for problem hooks** — a 2.3x improvement. Your hook type is the single biggest copy lever available.

## The Four Hook Types

### 1. Timeline Hook (Default — Highest Performer)

**Reply rate: 10.01% | Meeting rate: 2.34%**

Shows the prospect a compressed achievement timeline — from start to measurable result. Triggers urgency without pressure, provides implicit social proof through specificity, and reduces perceived risk through intermediate milestones.

**Pattern:**
```
Week 1-2: [Discovery/assessment phase] →
Week 3-4: [Pilot/implementation phase] →
Week 5-6: [Rollout/scale phase] →
[Measurable result] by [timeframe]
```

**Example (inside an email opener):**
```
Teams like {{company_name}} typically go from initial data audit to full pipeline visibility
in about 6 weeks — most see a 15-20% lift in qualified pipeline by week 8.
```

**Why it works:**
- Creates urgency through concrete speed ("6 weeks, not 6 months")
- Implicit social proof ("teams like yours" = others have done this)
- Risk reduction through intermediate milestones (not "sign a 12-month contract")
- Outcome anchoring — prospect self-qualifies against the result

**When to use:** Default for most outbound. Especially effective for: new category solutions, complex B2B with long perceived implementation, prospects in evaluation mode.

**When NOT to use:** When client has no proven timeline data. When product is too new for credible milestones. When the prospect's urgency is already high (trigger-based plays with hot signals).

### 2. Numbers Hook (Second Highest Performer)

**Reply rate: 8.57% | Meeting rate: 1.72%**

Leads with a specific, quantified claim that challenges the prospect's current reality.

**Pattern:**
```
[Specific metric] for [comparable company type] in [timeframe]
```

**Example:**
```
{{company_industry}} teams using [client approach] are seeing 40% fewer manual reviews
per quarter — curious if that gap shows up in your workflow too.
```

**Why it works:**
- Numbers create instant credibility
- The specificity signals genuine data, not marketing
- Comparison frame ("40% fewer") creates a mental benchmark

**When to use:** When client has strong, specific proof points. When prospect is analytical/data-driven persona (RevOps, Finance, Engineering leads). When competitive positioning matters — numbers beat vague claims.

**When NOT to use:** When numbers aren't credible or verified. When the proof point doesn't resonate with the persona's KPIs.

### 3. Social Proof Hook (Third Performer)

**Reply rate: 6.53% | Meeting rate: 1.08%**

Leads with a named or recognizable reference — a peer company, an industry leader, or a shared connection.

**Pattern:**
```
[Named company/person] + [what they achieved/discovered] + [relevance to prospect]
```

**Example:**
```
After [similar company] switched from [old approach] to [client approach], their
[relevant metric] improved by [X%] in [timeframe]. Given {{company_name}}'s scale
in {{company_industry}}, figured this might resonate.
```

**Why it works:**
- Named references create trust through association
- FOMO — "if my competitor/peer is doing this..."
- Reduces perceived risk — "someone like me already validated this"

**When to use:** When client has named case studies the prospect would recognize. When targeting a vertical where peer proof matters (enterprise, regulated industries). For competitor-customer targeting (strongest use case).

**When NOT to use:** When client lacks recognizable logos. When the reference company is too different from the prospect. When NDA prevents naming customers.

### 4. Hypothesis Hook (Current Thrustlab Default)

**Reply rate: 4.34% | Meeting rate: 0.69%**

Frames a researched pain point as a question, inviting dialogue rather than asserting knowledge.

**Pattern:**
```
"Would I be right in thinking [specific pain hypothesis] is something your team deals with?"
```

**Example:**
```
Would I be right in thinking that scaling your annotation pipeline for real-time detection
is something your team is dealing with? If so, we wrote up how [similar company]
approached it: [blog_link]
```

**Why it works:**
- Disarming — asks, doesn't tell
- Shows research without being presumptuous
- Invites dialogue rather than a yes/no
- Feels consultative, not salesy

**When to use:** For deeply technical personas who reject overt selling. When the pain is nuanced and research-dependent. For general outbound where deep research IS the differentiator. When other hook types lack sufficient data (no timeline, no numbers, no proof).

**When NOT to use as default:** Data shows this is the lowest-performing hook type when used as the primary opener. Use as a secondary element after a stronger lead.

---

## Hook Type Variable: `hook_type`

### Layer 2 Variable Definition

The `hook_type` variable is set per-prospect based on data availability and client proof points:

```
hook_type classification:

IF client has verified timeline data for prospect's vertical → "timeline"
ELSE IF client has specific metric/number for prospect's vertical → "numbers"
ELSE IF client has named case study prospect would recognize → "social_proof"
ELSE → "hypothesis" (fallback — always available through research)
```

### How hook_type Flows Through Clay

```
Layer 1: Foundation enrichment
  ↓
Layer 1.5: hook_type classification (AI prompt reads: company_industry, contact_title,
           client proof points from profile.md → outputs one of: timeline | numbers |
           social_proof | hypothesis)
  ↓
Layer 2: Research variables (prompts adjust based on hook_type)
  ↓
Layer 3: Copy generation (templates branch based on hook_type)
```

### Clay Prompt for hook_type Classification

```
You are selecting the optimal email hook type for a B2B outreach message.

## Client Proof Points Available
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}} employees)
- Contact: {{contact_first_name}}, {{contact_title}}
- ICP Angle: {{icp_angle}}

## Classification Rules
1. If we have a verified implementation timeline for their vertical or a closely adjacent one → output "timeline"
2. If we have a specific, quantified metric relevant to their role's KPIs → output "numbers"
3. If we have a named case study from a company they'd recognize (same vertical, similar size, or direct competitor) → output "social_proof"
4. If none of the above have strong data → output "hypothesis"

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Hook Type × Copy Templates

Every copy-generating skill should produce opener variants for each hook type. The copy generation prompt branches:

```
IF {{hook_type}} == "timeline":
  Lead with compressed achievement timeline from discovery to result.
  Pattern: "Teams in {{company_industry}} typically go from [phase 1] to [result] in [timeframe]."

IF {{hook_type}} == "numbers":
  Lead with specific quantified claim relevant to their role.
  Pattern: "{{company_industry}} teams using [approach] see [X% improvement] in [metric]."

IF {{hook_type}} == "social_proof":
  Lead with named reference and their result.
  Pattern: "After [named company] [did X], their [metric] [improved by Y]."

IF {{hook_type}} == "hypothesis":
  Lead with researched pain as a question.
  Pattern: "Would I be right in thinking [specific pain] is something your team deals with?"
```

---

## A/B Testing Protocol for Hook Types

### Phase 1: Establish Baseline (Weeks 1-2)
- Split traffic: 40% timeline, 20% numbers, 20% social_proof, 20% hypothesis
- Minimum 50 sends per hook type for statistical relevance
- Track: reply rate, positive reply rate, meeting book rate

### Phase 2: Optimize Winner (Weeks 3-4)
- 60% to winning hook type, 20% each to top 2 runners-up
- Begin testing WITHIN the winning type (which timeline narrative, which metric)

### Phase 3: Steady State (Week 5+)
- 70% to proven winner, 30% rotating test variants
- Re-test quarterly as client proof points evolve

### Reporting Variables
Add to Clay export columns:
- `hook_type_used` — which type was sent (for attribution)
- `hook_version` — A or B within type (for within-type testing)

---

## Client Onboarding: Hook Type Readiness Checklist

During client onboarding, populate these proof points for hook type classification:

```markdown
## Timeline Proof Points
- Vertical 1: [timeline from discovery to result]
- Vertical 2: [timeline from discovery to result]
- Average implementation time: [X weeks]
- First value milestone: [what happens at week Y]

## Numbers Proof Points
- Metric 1: [X% improvement in Y for Z company type]
- Metric 2: [X hours saved per week for Z role]
- Metric 3: [X% cost reduction in Y process]

## Named Case Studies
- Case Study 1: [Company name] ([vertical], [size]) — [result]
- Case Study 2: [Company name] ([vertical], [size]) — [result]
- Case Study 3: [Company name] ([vertical], [size]) — [result]

## Hook Type Availability Matrix
| Vertical | Timeline? | Numbers? | Social Proof? | Default Hook |
|---|---|---|---|---|
| {vertical_1} | ✅/❌ | ✅/❌ | ✅/❌ | {recommended} |
| {vertical_2} | ✅/❌ | ✅/❌ | ✅/❌ | {recommended} |
```

This matrix gets stored in `client-profiles/{slug}/profile.md` and is read by the `hook_type` classification prompt.
