# Clay Prompts — Dark Funnel Engagement Workflow

## Prompt Architecture

```
Layer 1: Signal capture + enrichment + intensity classification + ICP scoring + hook_type
Layer 2: Intent interpretation + inferred pain + company research + content matching
Layer 3: Copy generation (branched by signal_source × intensity × hook_type)
```

---

## Layer 1 Prompts

### Signal Intensity Classification

**Column:** `darkfunnel_intensity`
**Reads:** `darkfunnel_engagement_type`, `darkfunnel_content_topic`, `darkfunnel_source`
**Purpose:** Determine if this signal warrants dedicated outreach or just general outbound nurture.

```
Classify the buying intent intensity of this dark funnel signal.

## Signal Data
- Source: {{darkfunnel_source}}
- Engagement type: {{darkfunnel_engagement_type}}
- Content topic: {{darkfunnel_content_topic}}
- Company: {{company_name}} ({{company_industry}}, {{company_size}})

## Classification Rules

HIGH intensity (active buying behavior):
- Website: visited pricing page, comparison page, or 3+ pages in one session
- LinkedIn: commented on competitor content, or engaged with {client_name}'s content 2+ times
- Community: asked a question about the problem {client_name} solves, or compared solutions

MEDIUM intensity (research behavior):
- Website: visited solution/product page or case study page (single visit)
- LinkedIn: liked or shared content related to {client_problem_domain}
- Community: participated in discussion about the problem space

LOW intensity (passive awareness):
- Website: blog visit only, homepage only
- LinkedIn: single like on general content
- Community: lurking, no active participation detected

Output ONLY one word: high, medium, or low
```

**GATE:** Only proceed if "high" or "medium". Route "low" to general outbound queue with topic tag.

### Hook Type Classification

**Column:** `hook_type`
**Reads:** `darkfunnel_intensity`, `darkfunnel_content_topic`, `company_industry`

```
Select the optimal hook type for dark funnel outreach.

## Context
- Signal intensity: {{darkfunnel_intensity}}
- Content they engaged with: {{darkfunnel_content_topic}}
- Company industry: {{company_industry}}

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Classification Rules for Dark Funnel
1. If HIGH intensity + we have timeline data for their vertical → "timeline" (they're evaluating — show speed)
2. If we have a quantified metric on the topic they engaged with → "numbers"
3. If we have a case study in their vertical → "social_proof"
4. Otherwise → "hypothesis" (research the topic, ask an insightful question)

Dark funnel default tends toward "hypothesis" because the signal gives TOPIC but not enough for direct proof.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Layer 2 Prompts

### Engagement Context Interpretation

**Column:** `research_engagement_context`
**Reads:** `darkfunnel_engagement_type`, `darkfunnel_content_topic`, `contact_title`, `company_industry`
**Type:** AI prompt
**Purpose:** Interpret WHAT the engagement means — why would this person look at this content?

```
You are a B2B sales analyst for {client_name}.

## Signal
A {{contact_title}} at {{company_name}} ({{company_industry}}, {{company_size}}) just {{darkfunnel_engagement_type}} content about "{{darkfunnel_content_topic}}".

## Task
In ONE sentence, interpret what this engagement likely means for someone in their role. What business need or challenge would drive a {{contact_title}} to engage with this topic?

## Rules
- Maximum 25 words
- Be specific to their role + industry combination
- Don't reference the signal itself — focus on the BUSINESS CONTEXT
- If the engagement doesn't have meaningful business interpretation, output "SKIP"
```

### Inferred Pain

**Column:** `research_inferred_pain`
**Reads:** `research_engagement_context`, `company_industry`, `contact_title`
**Type:** AI prompt

```
You are inferring a business pain for {client_name}'s outreach.

## Context
- Person: {{contact_first_name}}, {{contact_title}} at {{company_name}} ({{company_industry}})
- Engagement interpretation: {{research_engagement_context}}
- {client_name} solves: {client_value_prop}
- Pains we address: {pain_1}, {pain_2}, {pain_3}

## Task
Which of {client_name}'s pain points is MOST likely relevant to this person given their engagement context? Write the pain in THEIR language (how they'd describe it), not {client_name}'s marketing language.

## Rules
- Maximum 20 words
- Frame as THEIR problem, not your solution
- If no clear pain connection, output "SKIP"
```

### Company Context Research

**Column:** `research_company_context`
**Type:** Claygent (web research)

```
Research {{company_name}} for recent news, announcements, or changes related to {{darkfunnel_content_topic}}.

Find:
1. Any recent company news that connects to this topic
2. Job postings that suggest they're investing in this area
3. Public statements about their approach to {{darkfunnel_content_topic}}

Context: We want to understand why {{company_name}} might be interested in {{darkfunnel_content_topic}} right now.

Rules:
- Only factual, verifiable information
- Maximum 60 words
- If nothing found, output "NOT_FOUND"
```

### Content Match

**Column:** `content_match`
**Reads:** `darkfunnel_content_topic`, `company_industry`, `icp_angle`
**Type:** Lookup/Formula

```
Match the best {client_name} content piece to share based on:
- Topic: {{darkfunnel_content_topic}}
- Industry: {{company_industry}}

Content library:
{client_content_library_with_topics_and_urls}

Output the URL of the best-matching content piece.
If no relevant match, output the most relevant general resource URL.
```

---

## Layer 3 Prompts — Copy Generation

### Opener — The Art of Serendipity

**Column:** `copy_opener_trigger`
**Reads:** ALL upstream + `hook_type`
**CRITICAL:** This opener must NEVER reference the dark funnel signal. It must feel like natural outreach.

```
You are writing an outbound email opener for {client_name}.

## THE GOLDEN RULE: This email must feel like coincidence, not surveillance.
You know the prospect is interested in {{darkfunnel_content_topic}} but you CANNOT reveal how you know.
The email must be framed as: "We're actively working in this space and thought you'd find this relevant."

## Prospect Context
- Name: {{contact_first_name}}, {{contact_title}} at {{company_name}}
- Industry: {{company_industry}}, Size: {{company_size}}
- Inferred pain: {{research_inferred_pain}}
- Company context: {{research_company_context}}
- Content to share: {{content_match}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed of result in their area of interest.
  Pattern: "{{company_industry}} teams tackling [topic] typically see [result] within [timeframe]."

IF hook_type = "numbers":
  Lead with a metric on the topic they're researching.
  Pattern: "We just published data showing [X% metric] for {{company_industry}} teams working on [topic]."

IF hook_type = "social_proof":
  Lead with a peer company's result in this area.
  Pattern: "[Similar company] recently tackled [topic] and saw [result]."

IF hook_type = "hypothesis":
  Lead with an insightful question about the topic.
  Pattern: "Curious — as {{company_name}} scales, is [inferred pain] something your team is navigating?"

## FRAMING OPTIONS (choose the most natural)
- "We just published research on [topic]..." (content-forward)
- "We're seeing a trend in {{company_industry}} around [topic]..." (industry insight)
- "Came across {{company_name}} while researching [topic] in {{company_industry}}..." (research context)

## Rules
- Maximum 2 sentences (35 words)
- ZERO references to their browsing, engagement, or activity
- Feel like helpful coincidence, not targeted surveillance
- {client_tone} tone
- If insufficient data, output "SKIP"
```

### Email Body

**Column:** `copy_body`

```
You are writing the body of a value-forward email for {client_name}.

## Context (already written)
Opener: {{copy_opener_trigger}}

## Prospect Context
- Inferred pain: {{research_inferred_pain}}
- Company context: {{research_company_context}}
- Content to share: {{content_match}}

## Task
Write 2-3 sentences that:
1. Connect the topic to {{company_name}}'s specific situation
2. Offer genuine value — a resource, insight, or framework
3. Position {client_name} as a thought leader, not a vendor (yet)

## Rules
- Maximum 45 words
- Lead with VALUE, not pitch
- Reference {{content_match}} naturally, not as a hard sell
- {client_tone} tone
```

### CTA

**Column:** `copy_cta`

```
Write a CTA for a dark funnel outreach email from {client_name}.

## Calibration
- Signal intensity: {{darkfunnel_intensity}}
- Signal composite score: {{signal_composite_score}}

## Rules
- Dark funnel CTAs should be SOFTER than trigger-based CTAs
- These prospects haven't raised their hand — meet them where they are

IF darkfunnel_intensity = "high":
  Slightly warmer: "Worth a quick chat about how teams in {{company_industry}} are approaching this?"

IF darkfunnel_intensity = "medium":
  Content-forward: "Happy to share more of our research on [topic] if useful."

- Maximum 20 words
- {client_tone} tone
```

### LinkedIn Message

**Column:** `copy_linkedin`

```
Write a LinkedIn message from {client_name} about {{darkfunnel_content_topic}}.

## Rules
- Maximum 280 characters
- Topic-forward, not company-forward
- Pattern: "Hi {{contact_first_name}}, we just published [content topic] for {{company_industry}} teams. Given {{company_name}}'s work in [area], thought you'd find it useful: {{content_match}}"
- NEVER reference their engagement
- {client_tone} tone
```

---

## Signal Stacking Integration

Dark funnel signals feed into Layer 2 (Intent) of the composite scoring model:

| Signal | Base Points | Decay | Source |
|---|---|---|---|
| Website visit — pricing/comparison | +25 | -5/week | Dealfront, RB2B |
| Website visit — solution page | +15 | -5/week | Dealfront, RB2B |
| LinkedIn engagement with client content (2x+) | +15 | -3/week | Teamfluence |
| LinkedIn engagement with competitor content | +20 | -5/week | Teamfluence, Trigify |
| Community question about problem space | +15 | -3/week | Common Room |
| Community solution comparison | +20 | -5/week | Common Room |

Dark funnel signals are strongest when STACKED with fit and relationship signals:
- Dark funnel (high) + ICP fit (A+) + funding = 🔴 Hot
- Dark funnel (medium) + ICP fit (A) = 🟡 Active (standard dark funnel play)
- Dark funnel (low) + ICP fit (A+) = 🟢 Watching (general outbound with topic tag)

---

## Meta Tagging

**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "dark-funnel"
```

Tag every row for A/B testing attribution. See `shared/references/prompt-iteration-pipeline.md`.
