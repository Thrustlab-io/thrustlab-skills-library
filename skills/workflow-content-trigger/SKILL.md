---
name: workflow-content-trigger
description: Builds the complete content engagement trigger workflow. Use when the strategy recommends content engagement (downloads, webinar attendance, social engagement, podcast listens) as a trigger play. Copy angle is demonstrated interest — the prospect already engaged with the client's content, making outreach feel like a natural next step rather than cold outreach. The warmest trigger type. Reads from all upstream client files.
---

# Content Engagement Trigger Workflow

Produces a complete, ready-to-deploy Clay workflow for content engagement signals.

**Copy DNA:** Demonstrated interest. This is the warmest trigger — the prospect CHOSE to engage with the client's content. They downloaded a whitepaper, attended a webinar, engaged with a LinkedIn post, or visited multiple content pieces. Copy extends the conversation they've already started, offering deeper value on the topic they're interested in. Never feels like cold outreach because it isn't — it's follow-up.

**Prerequisites:**
- All standard upstream files
- **CRITICAL:** Client must have content assets (blog, whitepapers, webinars, podcast, case studies). This trigger doesn't work without content to reference.
- `client-profiles/{client-slug}/tooling-setup.md` (confirms HubSpot/Marketo, LinkedIn Analytics, or Phantom Buster)

## Workflow

### Step 1: Load Client Context

Special attention to:
- **profile.md:** Existing content assets, content topics, social proof materials
- **icp-mapping.md:** Which content topics map to which persona pain points

### Step 2: Map Content → Pain Points

Before building Clay columns, map which content assets indicate which buyer signals:

| Content Asset | Topic | Pain Indicated | Best Persona Match |
|---|---|---|---|
| {whitepaper_1} | {topic} | {pain} | {persona} |
| {webinar_1} | {topic} | {pain} | {persona} |
| {blog_post_1} | {topic} | {pain} | {persona} |

### Step 3: Build Clay Table Schema

**Unique columns:**

| Column | Type | Purpose |
|---|---|---|
| `trigger_content_type` | Import | DOWNLOAD / WEBINAR / LINKEDIN_ENGAGEMENT / MULTI_PAGE_VISIT |
| `trigger_content_title` | Import | Specific content they engaged with |
| `trigger_content_topic` | AI prompt | Map content to a pain/topic category |
| `trigger_engagement_recency_hours` | Formula | Time since engagement |
| `trigger_engagement_depth` | AI prompt | DEEP (download, webinar, multiple assets) / LIGHT (single like, one blog) |
| `research_content_context` | AI prompt | What does this engagement tell us about their current challenges? |

### Step 4: Generate Prompts + Copy

**Critical rules:**
- This is warm outreach, not cold — the tone should reflect that they're already interested
- Reference the SPECIFIC content they engaged with (title, topic), not just "your recent interest"
- Offer the NEXT logical piece of value on the same topic
- Engagement depth determines approach: deep engagement = more direct, light engagement = softer
- Recency matters: <24 hours = very timely, 24-72 hours = timely, >72 hours = reference topic, not specific asset

### Step 5: Cadence

- **Email 1 (Day 0):** Content-specific follow-up + related value offer
- **LinkedIn (Day 1):** Connect on the topic of the content, share related insight
- **Email 2 (Day 4):** Deeper content: case study or data on the same topic
- **Email 3 (Day 8):** Bridge from content topic → how client solves the underlying problem
- **Email 4 (Day 13):** Breakup: different topic from content library

### Step 6: Quality Gate

- [ ] Specific content asset referenced (not generic "your interest")
- [ ] Next-step content/value is genuinely related to what they consumed
- [ ] Tone is warm (they chose to engage) — not cold outreach energy
- [ ] Engagement depth reflected in directness of CTA
- [ ] Recency reflected in specificity of reference
- [ ] Standard copy rules

### Hook Type & Signal Stacking Integration

**Hook Type:** Add `hook_type` column after ICP scoring. Classification reads client proof points and prospect context to select timeline/numbers/social_proof/hypothesis. Timeline is default when data supports it. See `shared/references/hook-types-guide.md`.

All copy generation prompts should branch on `hook_type`.

**Signal Stacking:** This trigger's signals feed into the Signal Aggregation Table (if active). See `shared/references/signal-stacking-guide.md` for point values and decay rates. When composite score crosses tier thresholds, escalate treatment.

**Prompt Iteration:** Tag all rows with `meta_prompt_version` and `meta_hook_type_used`. See `shared/references/prompt-iteration-pipeline.md`.
