# Clay Prompts — Job Posting Trigger

---

## Prompt 1: Job Description Pain Extraction

**Column name:** `trigger_job_pain_indicators`
**Column type:** AI prompt

```
You are a B2B sales research analyst for {client_name}.

## About {client_name}
{client_one_liner}
Problems we solve: {pain_1}, {pain_2}, {pain_3}

## Job Posting
- Company: {{company_name}} ({{company_industry}})
- Job title: {{trigger_job_title_posted}}
- Description: {{trigger_job_description_raw}}

## Task
Extract 2-3 pain points or challenges from this job description that {client_name} could address.

Look for:
- "Must have experience with X" = they currently lack X
- "Responsible for building/creating Y" = Y doesn't exist yet
- "Managing Z processes manually" = Z is a pain point
- Tools/technologies mentioned that {client_name} integrates with or replaces
- Responsibilities that overlap with what {client_name} automates

## Output
Return exactly 2-3 bullet points, each one sentence. Format:
- Pain: [specific challenge inferred from JD]

If the job posting has no relevant connection to {client_name}'s domain, output "NO_MATCH".
```

---

## Prompt 2: Tech Stack from JD

**Column name:** `trigger_job_tools_mentioned`
**Column type:** AI prompt

```
Extract all technologies, tools, and platforms mentioned in this job description:

{{trigger_job_description_raw}}

Return as comma-separated list. If none mentioned, return "None found".
```

---

## Prompt 3: Hiring Intent Score

**Column name:** `trigger_hiring_intent_score`
**Column type:** AI prompt

```
Score the relevance of this job posting for {client_name}'s outbound.

## {client_name} solves: {pain_1}, {pain_2}, {pain_3}
## Target personas: {persona_titles}

## The Posting
- Title: {{trigger_job_title_posted}}
- Company: {{company_name}} ({{company_industry}}, {{company_size}} employees)
- Pain indicators: {{trigger_job_pain_indicators}}

## Scoring
- HIGH: Job directly involves the problem space {client_name} solves. The hire will use or benefit from {client_name}.
- MEDIUM: Job is adjacent — the hire's team or function relates to {client_name}'s value prop.
- LOW: Weak connection — stretching to find relevance.
- NO_MATCH: No meaningful connection.

Return ONLY: HIGH, MEDIUM, LOW, or NO_MATCH
```

---

## Prompt 4: Team Context Research

**Column name:** `research_team_context`
**Column type:** Claygent

```
Research the {{company_industry}} team at {{company_name}} ({{company_website}}).

They are hiring a {{trigger_job_title_posted}}.

Find:
1. How big is their current team in this function? (check LinkedIn)
2. Any recent changes — growth, restructuring, new leadership?
3. What challenges might be driving this hire?

Maximum 100 words. Facts only.
```

---

## Prompt 5b: Hook Type Classification

**Column name:** `hook_type`
**Column type:** AI prompt

```
Select the optimal hook type for job-posting triggered outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Hiring: {{trigger_job_title_posted}}
- Pain indicators: {{trigger_job_pain_indicators}}

## Classification Rules
1. If we have a verified ramp timeline for solving {{trigger_job_pain_indicators}} → "timeline"
2. If we have quantified metrics on the JD pain area → "numbers"
3. If we have a named case study of a company that hired similar + used client → "social_proof"
4. Otherwise → "hypothesis"

Job posting context: "timeline" works well — show they can solve the pain faster than hiring alone.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Prompt 6: Trigger-Based Opener (Hiring-Aware, Hook Type Branched)

**Column name:** `copy_opener_trigger`
**Column type:** AI prompt

```
Write a hiring-aware opener. {{company_name}} is hiring a {{trigger_job_title_posted}}. Writing for {client_name}. Tone: {client_tone}.

## Pain Indicators from JD
{{trigger_job_pain_indicators}}

## Team Context
{{research_team_context}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed to solve the pain they're hiring for.
  Pattern: "Teams building out [function] typically go from [current state] to [result] in [timeframe] — often before the new hire starts."
  Use: {client_timeline_proof_by_vertical}

IF hook_type = "numbers":
  Lead with a metric on the JD pain area.
  Pattern: "{{company_industry}} teams see [X% improvement] in [area from JD] — usually before the [role] req even closes."
  Use: {client_number_proofs}

IF hook_type = "social_proof":
  Lead with a company that solved the same hiring challenge.
  Pattern: "[Named company] was in the same position — growing [department]. They [result]."
  Use: {client_named_case_studies}

IF hook_type = "hypothesis":
  Frame the hiring need as a broader challenge.
  Pattern: "Building out a [function] team at {{company_name}} — the [pain from JD] is usually what drives that."

## Universal Rules
- ONE sentence, max 25 words
- Frame as "building the [function]" or "growing the [capability]", NOT "I saw your job posting"
- The insight: you know WHY they're hiring (from JD), not just THAT they're hiring
- Do NOT start with "I"

Output ONLY the opener.
```

---

## Prompt 7: Email Body

**Column name:** `copy_body`
**Column type:** AI prompt

```
Write the email body for a job-posting triggered outbound from {client_name}. Tone: {client_tone}.

## About {client_name}
{client_one_liner}
Key differentiator: {key_differentiator}

## The Prospect
- {{contact_first_name}} ({{contact_title}}) at {{company_name}} ({{company_industry}})
- They're hiring: {{trigger_job_title_posted}}
- Pain indicators from JD: {{trigger_job_pain_indicators}}
- Team context: {{research_team_context}}

## Persona Rules
{persona_messaging_rules_from_icp}

## Task
Connect the hiring need to {client_name}'s outcome. Two valid framings:
1. "Augment" — {client_name} helps the new hire (or existing team) do more, faster
2. "Complement" — {client_name} handles [part of the JD] so the team can focus on [higher value work]

NEVER frame as "you don't need to hire" — that's insulting to the hiring manager.

## Rules
- Maximum 60 words
- Reference a specific pain from the JD, not generic challenges
- Industry language: {industry_language_guide_from_icp}
- Natural transition to CTA

Output body only.
```

---

## Prompt 8: LinkedIn Connection Request

**Column name:** `copy_linkedin`
**Column type:** AI prompt

```
Write a LinkedIn connection request to {{contact_first_name}} ({{contact_title}} at {{company_name}}).

Their team is hiring a {{trigger_job_title_posted}}.

From: {sender_name} at {client_name}

## Rules
- Maximum 280 characters
- Reference team growth or department expansion — shows you pay attention
- Do NOT pitch, do NOT mention the specific job posting URL
- Connect on the topic of building teams / scaling operations in {{company_industry}}

Output ONLY the message.
```

---

## Prompt 9: CTA (Signal-Stacking Aware)

**Column name:** `copy_cta`
**Column type:** AI prompt

```
Write a CTA for job-posting triggered outbound from {client_name}.

## Prospect Context
- Title: {{contact_title}}
- ICP tier: {{score_icp_tier}}
- Hiring: {{trigger_job_title_posted}}
- Signal composite score: {{signal_composite_score}} (if available)

## CTA Calibration
- Signal composite ≥100 (Hot): Direct ("worth 15 min to see if [client product] could complement the [role] hire?")
- Signal composite 70-99 (Warm): Resource share ("we put together a [function] scaling framework — want me to send it?")
- Standard: Insight-led ("happy to share how [industry] teams handle [JD pain] — might save some interview cycles")
- NEVER frame as "you don't need to hire" — augment, don't replace

## Rules
- Maximum 20 words
- Question or soft offer
- Match: {preferred_cta_style}

Output ONLY the CTA.
```

---

## Prompt 10: Meta Tagging

**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "job-posting-trigger"
```
