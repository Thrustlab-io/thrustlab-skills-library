---
name: workflow-general-outbound
description: Builds the complete general outbound workflow — Clay table config, deep research enrichment, Claygent prompts, and copy cadence. Use for every client (not trigger-dependent). This is proactive outreach to ICP-fit accounts WITHOUT a specific trigger event. Copy angle is deep research-based personalization — since there's no timing signal, the email must earn attention purely through relevance and insight. The hardest workflow to make effective. Reads from all upstream client files.
---

# General Outbound Workflow

Produces a complete, ready-to-deploy Clay workflow for non-triggered, ICP-based outreach.

**Copy DNA:** Research depth as the hook. Without a trigger, there's no built-in reason for the prospect to care about your email TODAY. The only way to earn attention is through deep personalization that proves you've done your homework. Copy demonstrates genuine understanding of the prospect's company, role, and challenges — then connects to the client's value prop in a way that feels inevitable, not forced. Every email must pass the "would a human SDR write this after 15 minutes of research?" test.

**Prerequisites:**
- `Prospects/{client-slug}/profile.md`
- `Prospects/{client-slug}/strategy.md`
- `Prospects/{client-slug}/icp-mapping.md` (this workflow depends MORE on icp-mapping than any trigger workflow)
- `Prospects/{client-slug}/market-mapping.md` (account list sources)

## Workflow

### Step 1: Load Client Context

Special attention to:
- **icp-mapping.md:** Persona × vertical matrix is CRITICAL here. Without a trigger, the persona card and industry pain mapping ARE your personalization source.
- **strategy.md:** Messaging architecture, top 6 openers, cadence templates
- **profile.md:** All differentiators, case studies, social proof — you'll use more of these than trigger workflows

### Step 2: Build Clay Table Schema — Two-Layer Variable Architecture

See `references/enrichment-sequence.md` and `references/example-campaign-patterns.md` for the complete reference.

**Key difference from trigger workflows:** No trigger columns. Instead, variables are generated in TWO LAYERS — each layer builds on the previous:

**Layer 1: Foundation + ICP Qualification**

| Column | Type | Purpose |
|---|---|---|
| `contact_first_name` | Import | Person's name |
| `contact_title` | Import/Enrichment | Their exact role |
| `company_name` | Import | Company name |
| `company_industry` | Enrichment | Sector classification |
| `icp_angle` | AI prompt | Which ICP bucket/angle they fall into — determines content matching and copy framing |
| `language` | Formula/Enrichment | Language for copy (EN, NL, etc.) based on region |

**Layer 2: Research Variables (AI prompts using Layer 1 as input)**

| Column | Type | Purpose | Example Output |
|---|---|---|---|
| `hook_type` | AI prompt | Which hook pattern to use — timeline/numbers/social_proof/hypothesis. See `shared/references/hook-types-guide.md` | "timeline" |
| `overlap_personal` | AI prompt | Hypothesis about a ROLE-SPECIFIC pain — framed as a question | "scaling your annotation pipeline for real-time detection" |
| `overlap_company` | AI prompt | What the COMPANY actually does that's relevant to client | "automated quality inspection using computer vision on the assembly line" |
| `company_workflow` | Claygent | Specific workflow/process the company operates | "visual inspection workflows across multiple production sites" |
| `tech_signals` | Claygent | Relevant tech they're using | "custom YOLO models for defect detection" |
| `company_insight` | AI prompt | An insightful QUESTION or observation based on deep research | "Curious whether maintaining those models across sites is becoming the bottleneck?" |
| `scale_trigger` | AI prompt | An inflection point / growth challenge hypothesis | "scaling from prototype to production-grade detection across all SKUs" |
| `blog_link` | Lookup | Primary content matched to `icp_angle` | URL to most relevant content piece |
| `blog_link_2` | Lookup | Secondary content for later touches | URL to secondary content piece |
| `meta_hook_type_used` | Formula | Tag for A/B tracking — which hook type was applied | "timeline" |
| `meta_prompt_version` | Static | Prompt version tag for iteration tracking | "v1.2" |

**CRITICAL INSIGHT:** Variables do the personalization. Copy templates do the framing. The copy itself is relatively standard — the VARIABLES make each email feel unique.

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md`.

**Critical rules for non-triggered outbound:**
- Research is your ONLY differentiator — invest 3-4 Clay columns in research before generating copy
- The `personalization_angle` column is the most important — it selects the best angle BEFORE the opener prompt runs
- Three angle types, in priority order:
  1. **Company-specific:** Something happening at THEIR company (news, initiative, content they published)
  2. **Role-specific:** Something about THEIR role in THEIR industry (from persona × vertical matrix)
  3. **Industry-specific:** A trend in THEIR vertical that creates a challenge (fallback — least personalized)
- Never default to "I saw your company..." — that's lazy. Be specific or don't send.
- General outbound has the HIGHEST bar for personalization. If the research can't find a genuine angle, SKIP the lead.

### Step 4: Generate Copy Templates — Scenario-Based Branching

See `references/copy-templates.md` and `references/example-campaign-patterns.md`.

General outbound isn't one linear cadence. It BRANCHES based on channel availability and prospect response:

**Scenario 1:** Has email + LinkedIn accepted early → LinkedIn-led, email reinforces
**Scenario 2:** Has email + LinkedIn accepted late → Email-led, LinkedIn follows
**Scenario 3:** Has email + LinkedIn never accepted → Email-only sequence
**Scenario 4:** No email found → LinkedIn-only sequence

The copy ESCALATION across all scenarios follows the same pattern, with the **opening touch branching by hook_type** (see `shared/references/hook-types-guide.md`):

1. **Opening** — Branched by `hook_type`:
   - **Timeline**: "Teams in {{company_industry}} typically go from [phase 1] to [result] in [timeframe]." Variables: `overlap_company`, `industry`
   - **Numbers**: "{{company_industry}} teams using [approach] see [X% improvement] in [metric]." Variables: `overlap_company`, `industry`
   - **Social proof**: "[Named company] tackled [challenge] and saw [result]." Variables: `overlap_company`
   - **Hypothesis** (fallback): Peer curiosity, connect on shared interest. Variables: `overlap_company`, `industry`
2. **Hypothesis Question** — "Would I be right in thinking [pain] is something your team is dealing with?" Variables: `overlap_personal`, `blog_link`
3. **Deep Research** — "Spent some time looking at how [company] handles [workflow]." Variables: `company_workflow`, `tech_signals`, `company_insight`, `scale_trigger`, `blog_link_2`
4. **Soft Close** — "If [topic] comes up, our door is always open." Non-threatening, routes to ongoing nurture.

**KEY:** Copy doesn't pitch harder over time. It goes DEEPER on research, then SOFTER on close.

**KEY 2:** Timeline hooks should be the DEFAULT for most prospects. Fall back to hypothesis only when timeline/numbers/social proof data is unavailable for this prospect's vertical. See `shared/references/hook-types-guide.md` for performance data (timeline: 10% reply rate vs. hypothesis: 4.3%).

### Step 5: A/B Testing Framework

General outbound benefits most from testing because there's no trigger timing to optimize:

**Test matrix (ranked by impact):**
- **Hook type** (HIGHEST LEVER): timeline vs. numbers vs. social_proof vs. hypothesis — see `shared/references/hook-types-guide.md`
  - Phase 1: 40% timeline, 20% numbers, 20% social_proof, 20% hypothesis
  - Phase 2: 60% winner, 20% each to top 2 runners-up
  - Phase 3: 70% winner, 30% rotation
- Opener type: Company-specific vs. role-specific vs. industry-specific
- CTA type: Meeting ask vs. resource share vs. question
- Subject line: Name mention vs. company mention vs. topic-only
- Send time: Test across time zones and days of week

**Tracking:** Every row must be tagged with `meta_hook_type_used` and `meta_prompt_version` for attribution. See `shared/references/prompt-iteration-pipeline.md` for the full iteration process.

### Step 6: Quality Gate (STRICTER than trigger workflows)

- [ ] Every email references something SPECIFIC to this prospect — no generic angles
- [ ] `personalization_angle` is company-specific or role-specific (not just industry-level)
- [ ] Research columns contain real, verifiable details (not AI hallucinations)
- [ ] If research is thin (no news, no activity, no content), mark as SKIP — don't force bad personalization
- [ ] Each email in the cadence uses a DIFFERENT angle — no repeating the same pain point
- [ ] Copy reads like a human wrote it after genuine research
- [ ] Standard copy rules (≤90 words, ≤45 char subject, etc.)
- [ ] "Would a human SDR write this?" test — if no, rewrite
