# Clay Prompts — Funding Trigger

---

## Prompt 1: Growth Priorities Research

**Column name:** `trigger_growth_priorities`
**Column type:** Claygent

```
Research the recent funding round of {{company_name}} ({{company_website}}).

They raised {{trigger_funding_amount}} in a {{trigger_funding_round}} round.

Find:
1. What did the CEO/founders say they'll use the funding for? (check press release, TechCrunch, LinkedIn posts)
2. What areas are they investing in? (product, team, market expansion, infrastructure)
3. Any specific initiatives or product launches mentioned?

Check: {{company_website}}/blog, Crunchbase, Google News, founder LinkedIn.
Maximum 150 words. Direct quotes are valuable.
```

---

## Prompt 2: Funding Relevance Score

**Column name:** `trigger_funding_relevance`
**Column type:** AI prompt

```
You are assessing funding relevance for {client_name}'s outbound.

## {client_name}
{client_one_liner}
We help with: {pain_1}, {pain_2}, {pain_3}

## Funded Company
- {{company_name}} ({{company_industry}}, {{company_size}} employees)
- Round: {{trigger_funding_round}}, Amount: {{trigger_funding_amount}}
- Growth priorities: {{trigger_growth_priorities}}

## Task
How relevant is this funding event for {client_name}'s outbound?

- HIGH: Growth priorities directly involve {client_name}'s domain. Clear use case.
- MEDIUM: Growth priorities tangentially connect. Scaling will create relevant needs.
- LOW: Funding is real but priorities don't connect to {client_name}.

Return the score AND one sentence explaining the connection.
Format: "HIGH — They're investing in [area] which directly needs [client domain]."
```

---

## Prompt 3: Hiring Surge Research

**Column name:** `research_hiring_surge`
**Column type:** Claygent

```
Check {{company_name}}'s current job openings.

Look at their careers page ({{company_website}}/careers) and LinkedIn jobs.

1. How many open roles?
2. Which departments are hiring most?
3. Any roles that relate to {client_product_domain}?

Maximum 80 words.
```

---

## Prompt 3b: Hook Type Classification

**Column name:** `hook_type`
**Column type:** AI prompt

```
Select the optimal hook type for funding-triggered outreach from {client_name}.

## Client Proof Points
- Timeline data: {client_timeline_proof_by_vertical}
- Metric proof points: {client_number_proofs}
- Named case studies: {client_named_case_studies}

## Prospect Context
- Company: {{company_name}} ({{company_industry}}, {{company_size}})
- Round: {{trigger_funding_round}}
- Growth priorities: {{trigger_growth_priorities}}

## Classification Rules
1. If we have a verified post-funding ramp timeline → "timeline" (they have capital, show speed to ROI)
2. If we have quantified ROI metrics → "numbers" (they're evaluating investments, show returns)
3. If we have a post-funding success story in their vertical → "social_proof"
4. Otherwise → "hypothesis"

Funding context: "numbers" is strong — they just raised money, ROI language resonates.

Output ONLY one word: timeline, numbers, social_proof, or hypothesis
```

---

## Prompt 4: Trigger-Based Opener (Funding-Aware, Hook Type Branched)

**Column name:** `copy_opener_trigger`
**Column type:** AI prompt

```
Write a funding-aware opener for {{company_name}}. Writing for {client_name}. Tone: {client_tone}.

## Context
- Round: {{trigger_funding_round}}, Amount: {{trigger_funding_amount}}
- Growth priorities: {{trigger_growth_priorities}}
- Hiring surge: {{research_hiring_surge}}
- Days since announcement: {{trigger_funding_freshness_days}}

## Hook Type: {{hook_type}}

IF hook_type = "timeline":
  Lead with speed to results post-funding.
  Pattern: "Post-[round] teams in {{company_industry}} typically see [result] within [timeframe] when they prioritize [area]."
  Use: {client_timeline_proof_by_vertical}

IF hook_type = "numbers":
  Lead with ROI metric relevant to their growth investment.
  Pattern: "[Industry] teams investing in [area] see [X% return/improvement] — that math gets better at {{company_name}}'s stage."
  Use: {client_number_proofs}

IF hook_type = "social_proof":
  Lead with a peer company that scaled post-funding.
  Pattern: "[Named company] was at a similar stage after their [round] — they [result]."
  Use: {client_named_case_studies}

IF hook_type = "hypothesis":
  Lead with the scaling challenge their growth priorities imply.
  Pattern: "Scaling [function] after a {{trigger_funding_round}} is one of the hardest transitions in {{company_industry}}."

## Universal Rules
- ONE sentence, max 25 words
- Lead with WHAT THEY'RE BUILDING, not the money
- Do NOT say "Congrats on the raise" or "I saw you raised"
- Frame funding as context for a business challenge, not an event
- Do NOT start with "I"

Output ONLY the opener.
```

---

## Prompt 5: Email Body

**Column name:** `copy_body`
**Column type:** AI prompt

```
Write the email body for a funding-triggered outbound from {client_name}. Tone: {client_tone}.

## About {client_name}
{client_one_liner}
Key differentiator: {key_differentiator}

## The Prospect
- {{contact_first_name}} ({{contact_title}}) at {{company_name}} ({{company_industry}})
- Growth priorities: {{trigger_growth_priorities}}
- Funding relevance: {{trigger_funding_relevance}}

## Persona Rules
{persona_messaging_rules_from_icp}

## Rules
- Maximum 60 words
- Connect their growth plan to {client_name}'s specific outcome
- Stage-appropriate language:
  - Seed/A: "move fast", "do more with less", "before you scale"
  - B/C: "infrastructure for scale", "don't let [function] become a bottleneck"
  - D+: "optimize", "enterprise-grade", "at your scale"
- Industry language: {industry_language_guide_from_icp}

Output body only.
```

---

## Prompt 6: LinkedIn Connection Request

**Column name:** `copy_linkedin`
**Column type:** AI prompt

```
Write a LinkedIn connection request to {{contact_first_name}} ({{contact_title}} at {{company_name}}).

They recently raised a {{trigger_funding_round}}.

From: {sender_name} at {client_name}

## Rules
- Maximum 280 characters
- Reference growth or scaling — NOT the funding amount
- Connect on shared interest in {{company_industry}} growth challenges
- Do NOT pitch, do NOT congratulate on funding specifically

Output ONLY the message.
```

---

## Prompt 7: CTA (Signal-Stacking Aware)

**Column name:** `copy_cta`
**Column type:** AI prompt

```
Write a CTA for funding-triggered outbound from {client_name}.

## Prospect Context
- Title: {{contact_title}}
- ICP tier: {{score_icp_tier}}
- Funding round: {{trigger_funding_round}}
- Signal composite score: {{signal_composite_score}} (if available)

## CTA Calibration
- Signal composite ≥100 (Hot) OR A+ tier: Direct but investment-framed ("worth 15 min to see if [area] should be part of the post-[round] roadmap?")
- Signal composite 70-99 (Warm): Growth playbook offer ("I put together a scaling framework for post-[round] {{company_industry}} teams — want me to send it?")
- Standard: Pure value ("happy to share how teams at your stage approach [pain area]")
- Stage-appropriate urgency: Seed/A = "before you scale" | B/C = "while you're building" | D+ = "optimize the investment"

## Rules
- Maximum 20 words
- Question or soft offer
- Match: {preferred_cta_style}

Output ONLY the CTA.
```

---

## Prompt 8: Meta Tagging

**Column type:** Formula / Static

```
meta_hook_type_used = {{hook_type}}
meta_prompt_version = "v1.0"
meta_workflow_name = "funding-trigger"
```
