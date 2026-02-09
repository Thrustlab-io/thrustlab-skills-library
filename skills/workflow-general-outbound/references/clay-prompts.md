# Clay Prompts — General Outbound

Variables are generated in LAYERS. Each layer's output feeds the next layer's prompts.
See `references/example-campaign-patterns.md` for the full pattern reference.

---

## LAYER 1: Foundation — ICP Qualification

### Prompt 1: ICP Angle Classification

**Column:** `icp_angle`
**Column type:** AI prompt

```
You are classifying a prospect into the correct ICP angle for {client_name}'s outbound.

## {client_name}'s ICP Angles
{List each angle/bucket from icp-mapping.md — e.g.:
- "annotation and computer vision": companies working with image/video labeling, object detection, visual inspection
- "visual transformation": companies doing 3D modeling, image processing, visual optimization
- "constraint-aware design": companies doing generative design, CAD, engineering simulation}

## Prospect Data
- Company: {{company_name}} ({{company_industry}}, {{company_size}} employees)
- Person: {{contact_first_name}} {{contact_last_name}}, {{contact_title}}
- Company description: {{company_description}}

## Task
Which ICP angle best fits this prospect? This determines which content, pain framing, and copy approach to use.

## Output
Return: ANGLE_NAME — one sentence explaining why.
If no angle fits, return "NO_FIT".
```

---

## LAYER 2a: Research Variables — Personal & Company Overlap

These are the CORE personalization variables. They're used in the first touches.

### Prompt 2: Personal Overlap (Role-Specific Pain Hypothesis)

**Column:** `overlap_personal`
**Column type:** AI prompt

```
You are generating a role-specific pain hypothesis for {client_name}'s outbound.

## About {client_name}
{client_one_liner}
Problems we solve:
- {pain_1}: {description}
- {pain_2}: {description}
- {pain_3}: {description}

## Prospect
- {{contact_first_name}} {{contact_last_name}}, {{contact_title}} at {{company_name}}
- Industry: {{company_industry}}
- ICP angle: {{icp_angle}}

## Persona Context
{persona_card_from_icp_for_this_title — priorities, day-to-day pains, what they care about}

## Task
Generate a SHORT, SPECIFIC hypothesis about a hands-on challenge this person faces in their role — something {client_name} can help with.

## Rules
- This will be used in copy as: "Would I be right in thinking {{overlap_personal}} is something your team is dealing with?"
- So it must READ as a natural completion of that sentence
- Must be specific to their ROLE + INDUSTRY combination (use persona × vertical matrix)
- Frame as a DOING challenge, not an abstract problem
- Maximum 15 words
- No jargon the prospect wouldn't use — use THEIR industry language

## Good examples:
- "scaling your visual inspection pipeline across multiple production lines"
- "keeping detection models accurate as your product catalog grows"
- "turning prototype 3D workflows into production-grade pipelines"

## Bad examples:
- "improving operational efficiency" (too generic)
- "leveraging AI for competitive advantage" (B2B slop)
- "dealing with data quality issues" (vague, not role-specific)

Output ONLY the pain hypothesis phrase. No quotes, no prefix.
If insufficient data for a specific hypothesis, use the fallback from icp-mapping.md for this persona × vertical: "{fallback_pain}"
```

---

### Prompt 3: Company Overlap (Company-Level Context)

**Column:** `overlap_company`
**Column type:** AI prompt (uses Claygent research as input)

```
You are generating a company-level context snippet for {client_name}'s outbound.

## Prospect Company
- {{company_name}} ({{company_industry}}, {{company_size}} employees)
- ICP angle: {{icp_angle}}
- Company research: {{research_company_deep}}

## Task
In ONE SHORT PHRASE, describe what {{company_name}} actually does that's relevant to {client_name}'s domain.

## Rules
- This will be used in copy as: "Given {{company_name}}'s work in {{overlap_company}}..."
- So it must READ as a natural completion of that phrase
- Reference SPECIFIC work they do (from research), not their whole business
- Maximum 12 words
- Use their industry's language, not ours

## Good examples:
- "automated quality inspection using computer vision on the assembly line"
- "real-time 3D asset generation for your game pipeline"
- "satellite imagery analysis for agricultural monitoring"

## Bad examples:
- "AI and machine learning solutions" (generic)
- "innovative technology development" (says nothing)

Output ONLY the phrase. If research is insufficient, output "SKIP".
```

---

## LAYER 2b: Deep Research Variables

These are used in later touches (LI Message 2, Email 2) to demonstrate genuine research depth.

### Prompt 4: Deep Company Research

**Column:** `research_company_deep`
**Column type:** Claygent (web research)

```
Research {{company_name}} ({{company_website}}) thoroughly.

Find:
1. What does the company do? (one sentence)
2. Recent news, announcements, or changes (last 3-6 months)
3. What they're building or investing in (product launches, market expansion)
4. Recent content they've published (blog, press, case studies)
5. Technology stack (from job postings, BuiltWith, public info)
6. Recent hires or departures in key roles

Check: {{company_website}}, LinkedIn company page, Google News, Crunchbase, G2.
Maximum 200 words. Prioritize RECENT and SPECIFIC.
If nothing notable: "Standard {{company_industry}} company — no recent distinctive signals."
```

### Prompt 5: Company Workflow

**Column:** `company_workflow`
**Column type:** AI prompt (from Claygent research)

```
Based on the research about {{company_name}}, identify the specific WORKFLOW or PROCESS they operate that's most relevant to {client_name}'s domain.

## Research: {{research_company_deep}}
## ICP angle: {{icp_angle}}
## {client_name} helps with: {pain_1}, {pain_2}, {pain_3}

## Task
Describe the specific workflow/process in 5-10 words.

This will be used as: "Spent some time looking at how {{company_name}} handles {{company_workflow}}."

## Good output: "visual inspection workflows across multiple production sites"
## Bad output: "AI operations" (too vague)

Output ONLY the workflow phrase. If not enough data: "SKIP".
```

### Prompt 6: Tech Signals

**Column:** `tech_signals`
**Column type:** AI prompt (from research)

```
From the research on {{company_name}}, extract the most relevant technology signals.

## Research: {{research_company_deep}}
## ICP angle: {{icp_angle}}

What specific technologies, tools, or technical approaches is {{company_name}} using that relate to {client_name}'s domain?

This will be used as: "it seems you're running {{tech_signals}} in-house."

Output a short phrase (5-10 words). If nothing found: "SKIP".
```

### Prompt 7: Company Insight (Hypothesis Question)

**Column:** `company_insight`
**Column type:** AI prompt

```
Generate an insightful hypothesis QUESTION about {{company_name}} based on research.

## Research: {{research_company_deep}}
## Workflow: {{company_workflow}}
## Tech: {{tech_signals}}
## {client_name} solves: {pain_1}, {pain_2}, {pain_3}

## Task
Write ONE question that demonstrates deep understanding of their business AND connects to a challenge {client_name} could help with.

## Rules
- Must be a genuine, curiosity-driven question — not a leading sales question
- Should make the prospect think "they actually understand what we're doing"
- Maximum 25 words
- Frame as hypothesis: "Curious whether..." or "Is [specific challenge] becoming..."

## Good examples:
- "Curious whether you've hit the point where maintaining those models across sites becomes the bottleneck?"
- "Is the manual annotation step still the main blocker for getting new product lines into production?"

## Bad examples:
- "Would you like to improve your efficiency?" (generic sales question)
- "Are you looking for an AI solution?" (lazy)

Output ONLY the question.
```

### Prompt 8: Scale Trigger (Inflection Point)

**Column:** `scale_trigger`
**Column type:** AI prompt

```
Identify the most likely INFLECTION POINT or scaling challenge for {{company_name}}.

## Research: {{research_company_deep}}
## Workflow: {{company_workflow}}
## Company size: {{company_size}}
## ICP angle: {{icp_angle}}

## Task
In one short phrase, describe the scaling challenge they're most likely approaching.

This will be used as: "Is {{scale_trigger}} something {{company_name}} is working on?"

## Rules
- Specific to THEIR situation, not a generic growth challenge
- Frame as an action/transition, not a state
- Maximum 15 words

## Good: "scaling from prototype to production-grade detection across all SKUs"
## Bad: "growing the business" (meaningless)

Output ONLY the phrase. If research is thin: use industry-level fallback from icp-mapping.md.
```

---

## LAYER 2c: Content Matching

### Prompt 9: Blog/Content Matching

**Column:** `blog_link` and `blog_link_2`
**Column type:** Lookup/Formula

```
No AI prompt needed — this is a LOOKUP TABLE.

Based on {{icp_angle}}, match to the appropriate content:

{content_mapping_table_from_client}

Example structure:
| Angle | Primary Content ({{blog_link}}) | Secondary Content ({{blog_link_2}}) |
|---|---|---|
| {angle_1} | {url_1a} | {url_1b} |
| {angle_2} | {url_2a} | {url_2b} |
| {angle_3} | {url_3a} | {url_3b} |
```

---

## LAYER 2d: Prospect Activity (Optional Enhancement)

### Prompt 10: Recent Activity

**Column:** `research_recent_activity`
**Column type:** Claygent

```
Research {{contact_first_name}} {{contact_last_name}} ({{contact_title}} at {{company_name}}).

Check LinkedIn profile:
1. Recent posts or articles (last 3 months)
2. Topics they discuss
3. Speaking engagements or publications
4. Professional communities

Maximum 100 words. If low-activity: "Low public activity — use company/role-based angle."
```

---

## LAYER 3: Copy Generation

Copy prompts ASSEMBLE the variables generated above. They don't do research — they frame.

### Prompt 10b: Hook Type Classification

**Column:** `hook_type`
**Column type:** AI prompt
**Purpose:** Select optimal hook pattern. See `shared/references/hook-types-guide.md`.

```
Select the optimal hook type for general outbound from {client_name} to {{contact_first_name}} at {{company_name}}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Title: {{contact_title}}
- ICP angle: {{icp_angle}}
- Personal overlap: {{overlap_personal}}
- Company overlap: {{overlap_company}}

## Classification Rules
1. If we have a verified achievement timeline for {{company_industry}} or {{icp_angle}} → "timeline"
2. If we have quantified metrics relevant to {{contact_title}}'s KPIs → "numbers"
3. If we have a named case study in {{company_industry}} or matching {{icp_angle}} → "social_proof"
4. Otherwise → "hypothesis"

General outbound context: No trigger signal means the hook does ALL the heavy lifting.
Timeline is default and strongest (10% reply vs 4.3% hypothesis). Only fall back to hypothesis when other proof points unavailable.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

### Prompt 11: LinkedIn Invite

**Column:** `copy_linkedin_invite`

```
Write a LinkedIn connection request from {sender_name} at {client_name} to {{contact_first_name}} ({{contact_title}} at {{company_name}}).

## Available context:
- Industry: {{company_industry}}
- Company overlap: {{overlap_company}}
- ICP angle: {{icp_angle}}

## Rules
- Maximum 280 characters
- Frame as: "came across {{company_name}} while researching [industry] teams working on [angle topic]"
- OR: "Given {{company_name}}'s work in {{overlap_company}}, thought there might be shared interests"
- Peer curiosity tone — NOT pitching, just connecting
- Include a reason to connect (share research, exchange notes)

## Language: {{language}}

Output ONLY the message.
```

### Prompt 12: Hook-Typed Email/Message (Touch 2)

**Column:** `copy_hypothesis_touch`

```
Write a hook-typed outbound message for {{contact_first_name}} at {{company_name}}. For {client_name}. Tone: {client_tone}.

## Variables to use:
- Personal overlap: {{overlap_personal}}
- Company overlap: {{overlap_company}}
- Blog link: {{blog_link}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Structure:
  1. Brief warm opening
  2. Timeline-based hook: "Teams in {{company_industry}} tackling {{overlap_personal}} typically go from [phase 1] to [result] in [timeframe]."
     Use: {client_timeline_proof_by_vertical}
  3. Bridge: "Curious whether that matches what you're seeing at {{company_name}}."
  4. Value offer: link to {{blog_link}}

IF hook_type = "numbers":
  Structure:
  1. Brief warm opening
  2. Numbers-based hook: "[Persona] teams see [X% improvement] in [metric related to overlap_personal]."
     Use: {client_number_proofs}
  3. Bridge: "Worth exploring whether that applies to {{company_name}}'s setup."
  4. Value offer: link to {{blog_link}}

IF hook_type = "social_proof":
  Structure:
  1. Brief warm opening
  2. Social-proof hook: "[Named company in industry] tackled [similar challenge] and saw [specific result]."
     Use: {client_named_case_studies}
  3. Bridge: "{{company_name}}'s work in {{overlap_company}} looks like a similar situation."
  4. Value offer: link to {{blog_link}}

IF hook_type = "hypothesis":
  Structure:
  1. Brief warm opening
  2. Hypothesis question: "Would I be right in thinking {{overlap_personal}} is something your team is dealing with?"
  3. Value offer: link to {{blog_link}}
  4. Curiosity close: "Curious what you're seeing on your end."

## Rules
- ≤90 words for email, ≤280 chars for LinkedIn message
- The hook is the CORE — it must feel specific and insightful
- The blog link is the VALUE — not a pitch, a genuinely useful resource
- Signal composite calibration (if available):
  - {{signal_composite_score}} ≥70: can be slightly more direct in the bridge, include a soft CTA
  - {{signal_composite_score}} <40 or unavailable: keep pure value, curiosity close only
- Language: {{language}}

Output the message.
```

### Prompt 13: Deep Research Email/Message (Touch 3)

**Column:** `copy_deep_research_touch`

```
Write a deep-research message for {{contact_first_name}} at {{company_name}}. For {client_name}. Tone: {client_tone}.

## Variables to use:
- Company workflow: {{company_workflow}}
- Tech signals: {{tech_signals}}
- Company insight: {{company_insight}}
- Scale trigger: {{scale_trigger}}
- Blog link 2: {{blog_link_2}}

## Structure:
1. Research demonstration: "Spent some time looking at how {{company_name}} handles {{company_workflow}}."
2. EITHER observation + insight question: "{{company_insight}}"
   OR tech signal + scale question: "From what I can see, you're running {{tech_signals}} in-house. Is {{scale_trigger}} something {{company_name}} is working on?"
3. PS with secondary blog: "PS: thought this might be useful as well: {{blog_link_2}}"

## Rules
- This touch PROVES you've done real research — it's the credibility builder
- Must reference specific details from the research variables
- ≤90 words for email, shorter for LinkedIn
- Language: {{language}}

Output the message.
```

### Prompt 14: Soft Close (Final Touch)

**Column:** `copy_soft_close`

```
Write a soft close message for {{contact_first_name}} at {{company_name}}. For {client_name}. Tone: {client_tone}.

## Variables:
- Personal overlap: {{overlap_personal}}
- Company name: {{company_name}}

## Structure:
1. Open door: "If {{overlap_personal}} or any related topic becomes a priority at {{company_name}}, we're here."
2. Low-pressure offer: "Sometimes an outside perspective helps. Happy to take a look if useful."
3. Warm sign-off: "All the best with what you're building."
4. PS: redirect to ongoing content (LinkedIn page, newsletter, etc.)

## Rules
- This is NOT a sales push — it's leaving the door open
- Tone: warm, genuine, zero pressure
- Short: 50-60 words max
- Language: {{language}}

Output the message.
```

---

### Prompt 15: Meta Tagging

**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "general-outbound"
```

These columns are for A/B tracking and prompt iteration attribution. See `shared/references/prompt-iteration-pipeline.md`.
