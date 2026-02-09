# Clay Prompts & Copy Templates — Content Engagement Trigger

## Clay Prompts

### Prompt 1: Content Topic Mapping

**Column:** `trigger_content_topic`

```
Map this content engagement to a pain/topic category for {client_name}.

## {client_name}'s pain categories:
- {pain_category_1}: {description}
- {pain_category_2}: {description}
- {pain_category_3}: {description}

## Engagement
- Content type: {{trigger_content_type}}
- Content title: {{trigger_content_title}}

Return: The pain category that best matches + one sentence explaining the connection.
If no clear match, return "GENERAL_INTEREST".
```

### Prompt 2: Engagement Depth Assessment

**Column:** `trigger_engagement_depth`

```
Assess engagement depth for outbound prioritization.

## Engagement Data
- Type: {{trigger_content_type}}
- Content: {{trigger_content_title}}
- Recency: {{trigger_engagement_recency_hours}} hours ago

Scoring:
- DEEP: Downloaded gated content (whitepaper, report), attended webinar, engaged with 3+ content pieces, or requested a demo/contact
- MODERATE: Downloaded one ungated asset, watched a video, or engaged with 2 content pieces
- LIGHT: Liked a post, read one blog, brief page visit

Return ONLY: DEEP, MODERATE, or LIGHT
```

### Prompt 2b: Hook Type Classification

**Column:** `hook_type`

```
Select optimal hook type for content-engagement outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Content topic: {{trigger_content_topic}}
- Engagement depth: {{trigger_engagement_depth}}

## Classification
1. If we have a timeline proof on the content topic → "timeline"
2. If we have quantified metrics on the content topic → "numbers"
3. If we have a named case study on the content topic → "social_proof"
4. Otherwise → "hypothesis"

Content engagement context: "hypothesis" is actually strong here — extending the conversation they started.
DEEP engagement shifts toward more direct hooks (timeline/numbers); LIGHT stays at hypothesis.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

### Prompt 3: Trigger-Based Opener (Hook Type Branched)

**Column:** `copy_opener_trigger`

```
Write a content-engagement opener for {{contact_first_name}} at {{company_name}}. For {client_name}. Tone: {client_tone}.

## Engagement
- Type: {{trigger_content_type}}
- Content: {{trigger_content_title}}
- Topic: {{trigger_content_topic}}
- Depth: {{trigger_engagement_depth}}
- Recency: {{trigger_engagement_recency_hours}} hours

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed to results on the topic they're exploring.
  Pattern: "{{company_industry}} teams tackling [content topic] typically see [result] within [timeframe]."

IF hook_type = "numbers":
  Lead with a metric on the content topic.
  Pattern: "The data on [content topic] is interesting — {{company_industry}} teams see [X% improvement] when they [action]."

IF hook_type = "social_proof":
  Lead with a peer company's result on the same topic.
  Pattern: "[Named company] tackled [content topic] and saw [result] — their {{company_industry}} setup is similar to {{company_name}}."

IF hook_type = "hypothesis":
  Extend the conversation they started.
  Pattern: "The [content topic] conversation is heating up in {{company_industry}} — curious whether [specific question] resonates with your team."

## Rules
- ONE sentence, max 25 words
- This is WARM outreach — they already showed interest
- Reference the TOPIC, not their specific action
- Do NOT say "I saw you downloaded" or "Thanks for attending"
- IF DEEP: more direct about the problem space
- IF LIGHT: lead with trend, not their action
- IF recency >168 hours: reference topic generally, not specific asset
- Do NOT start with "I"

Output ONLY the opener.
```

### Prompt 4: Email Body

Standard structure adapted for warm outreach:
- Maximum 60 words
- Bridge from content topic → deeper insight or related problem
- Offer next-step value (related content, case study, data point)
- Warmer CTA than cold triggers (they're already interested)
- "Want me to send over [related resource]?" is a valid CTA here

---

## Copy Templates

### Cadence

| Step | Channel | Timing | Angle |
|---|---|---|---|
| 1 | Email | Day 0 | Topic-specific follow-up + related value |
| 2 | LinkedIn | Day 1 | Connect on shared topic interest |
| 3 | Email | Day 4 | Deeper content: case study / data on same topic |
| 4 | Email | Day 8 | Bridge: topic → underlying problem → client solution |
| 5 | Email | Day 13 | Different topic from content library |

### Email 1 — Content Follow-Up (Day 0)
**Subject:** `More on {content_topic} for {{company_industry}}`
- Opener references the topic (not the download/attendance)
- Body offers next-step value: "We just published [related content] that goes deeper on [specific angle]"
- CTA: offer to send the related resource (low commitment)

### Email 2 — Case Study (Day 4)
**Subject:** `How {{peer_company}} tackled {topic}`
- Lead with company-based opener (variety from Email 1)
- Case study on the same topic as original content
- More direct: "Seeing this pattern a lot in {{company_industry}} — happy to share what's working?"

### Email 3 — Problem Bridge (Day 8)
- Transition from content topic → the underlying business problem → client solution
- This is where the cadence shifts from "helpful content" to "here's how we solve this"
- Still warmer than cold outreach

### Key Uniqueness
- Warmest trigger type — they CHOSE to engage
- Email tone closer to "continuing a conversation" than "starting one"
- Can be more direct with CTAs earlier in the cadence
- Content mapping is critical — wrong topic match = feels disconnected

### Gates
- `trigger_engagement_depth` = "LIGHT" → consider lower priority or lighter cadence (3 emails, not 5)
- `trigger_engagement_recency_hours` > 168 (1 week) → reference topic generally, not specific asset
- Standard ICP gate

---

### Prompt 5: CTA (Signal-Stacking Aware)

**Column:** `copy_cta`

```
Write CTA for content-engagement outbound from {client_name}.

## Calibration
- Title: {{contact_title}}, ICP tier: {{score_icp_tier}}
- Engagement depth: {{trigger_engagement_depth}}
- Signal composite score: {{signal_composite_score}} (if available)

- DEEP engagement OR signal composite ≥70: More direct ("worth a quick chat about how {{company_industry}} teams are approaching [topic]?")
- MODERATE engagement: Content-forward ("we just published a deeper dive on [topic] — want me to send it?")
- LIGHT engagement: Pure value, no ask ("thought this might be relevant given {{company_name}}'s work in [area]")

Maximum 20 words. Match: {preferred_cta_style}
```

---

### Prompt 6: Meta Tagging

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "content-trigger"
```
