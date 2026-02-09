---
name: workflow-dark-funnel
description: Builds the complete dark funnel engagement workflow — captures and acts on anonymous buying signals from LinkedIn engagement (Teamfluence/Trigify), community activity (Common Room), and website visits beyond basic page views (Dealfront/RB2B). Illuminates the 70%+ of B2B buyer journey that happens invisibly. Produces Clay table config, signal classification prompts, and copy that references engagement without surveillance creep. Use when client wants to capture buying intent from social and community channels alongside website signals. Reads from client profile, strategy, and requires dark funnel tooling setup.
---

# Dark Funnel Engagement Workflow

Produces a complete Clay workflow for outreach based on anonymous engagement signals that traditional triggers miss.

**Copy DNA:** Relevance without surveillance. Dark funnel signals tell you a prospect is INTERESTED but they don't know you can see their activity. The copy must feel like serendipity — "I happened to notice" or "we're researching this space" — never "I saw you liked our LinkedIn post at 2:47 PM." The art is using the signal for TIMING and ANGLE selection while keeping the copy naturally relevant.

**Why dark funnel matters:** 70%+ of B2B buyer journeys happen anonymously — prospects research solutions, read content, engage with peers, and compare options before ever filling a form. By the time they're "in-market" by traditional measures, they're already 60-80% through their decision process. Dark funnel signals let you engage EARLIER in the journey.

**Signal sources (preferred tooling):**
- **Website engagement:** Dealfront (EU-strong, company-level) / RB2B (US, person-level)
- **LinkedIn engagement:** Teamfluence (monitors ICP engagement with your content and competitors') / Trigify (LinkedIn event triggers)
- **Community signals:** Common Room (Slack, Discord, GitHub, forums)

**Prerequisites:**
- `client-profiles/{client-slug}/profile.md`
- `strategies/{client-slug}.md`
- At least one dark funnel signal source configured and sending data to Clay
- Content library: client's blogs, case studies, or resources to share (required for value-forward copy)

## Workflow

### Step 1: Load Client Context

Special attention to:
- **profile.md:** Content assets available for value-forward responses
- **strategy.md:** Signal stacking config, ICP criteria for qualifying dark funnel signals
- **Tooling setup:** Which dark funnel sources are active, what data they send

### Step 2: Build Clay Table Schema

**Layer 1: Signal Capture + Classification**

| Column | Type | Purpose |
|---|---|---|
| `company_name` | Import (signal source) | Company identified from signal |
| `contact_first_name` | Import (if person-level) | Person who engaged (RB2B, Teamfluence) |
| `contact_title` | Enrichment | Their role |
| `contact_email` | Enrichment waterfall | Email for outreach |
| `company_industry` | Enrichment | Sector |
| `company_size` | Enrichment | Size qualification |
| `darkfunnel_source` | Import | Which tool detected: "dealfront", "rb2b", "teamfluence", "trigify", "common_room" |
| `darkfunnel_engagement_type` | Import | What they did: "website_visit", "linkedin_like", "linkedin_comment", "community_post", "github_star", etc. |
| `darkfunnel_content_topic` | Import/AI | Topic of the content they engaged with |
| `darkfunnel_recency` | Formula | Hours/days since signal |
| `darkfunnel_intensity` | AI prompt | "high" (pricing page, competitor comparison, multiple touches) / "medium" (solution content, single touch) / "low" (blog, social like) |
| `score_icp_fit` | AI scoring | Standard ICP qualification |
| `hook_type` | AI prompt | Hook type classification |

**Layer 2: Intent Interpretation + Research**

| Column | Type | Purpose |
|---|---|---|
| `research_engagement_context` | AI prompt | What this engagement MEANS — why would someone in their role look at this content? |
| `research_inferred_pain` | AI prompt | Based on what they engaged with, what pain are they likely experiencing? |
| `research_company_context` | Claygent | What's happening at their company that could explain this interest? |
| `content_match` | Lookup | Best content piece from client to share — matched to their engagement topic |
| `signal_composite_score` | Formula/Lookup | From signal aggregation table |

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md`.

**Critical rules for dark funnel outreach:**
- NEVER reference the specific signal directly ("I saw you visited our pricing page" / "I noticed you liked our post")
- Use the signal for TIMING (when to send) and ANGLE (what topic to lead with) — not as the opener
- The email should feel like coincidence: "We just published research on [the exact topic they're exploring]"
- Content-forward approach: lead with value, not curiosity about their browsing
- Dark funnel signals decay FAST — website visits: act within 24-48 hours. LinkedIn engagement: act within 48-72 hours. Community signals: within 1 week.
- If `darkfunnel_intensity` = "low" → add to general outbound nurture, don't trigger dedicated workflow
- Only "medium" and "high" intensity signals warrant a dedicated dark funnel play

### Step 4: Generate Copy Templates

See `references/copy-templates.md`.

**Copy escalation:**
1. **Value-forward** (Day 0) — Share content relevant to what they were exploring. Feel like helpfulness, not surveillance
2. **Insight share** (Day 3) — Share a proprietary insight or data point on the topic
3. **Soft engagement** (Day 7) — Ask a thought-provoking question about the topic

### Step 5: Quality Gate

- [ ] Copy NEVER reveals the specific dark funnel signal — no "I saw you visited" or "I noticed you engaged with"
- [ ] The topic connection feels natural, not forced
- [ ] If signal is "low" intensity → routed to general outbound, not this workflow
- [ ] `darkfunnel_recency` is within window (24-48h for website, 48-72h for LinkedIn, 1 week for community)
- [ ] Content match is relevant to the engagement topic — not just any content
- [ ] hook_type appropriate — timeline and hypothesis hooks work best here (signal gives topic, not proof)
- [ ] Standard copy rules (≤90 words, ≤45 char subject, etc.)
- [ ] Signal stacking: dark funnel + other signals (ICP fit, job posting, funding) → escalate
