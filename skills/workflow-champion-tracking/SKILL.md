---
name: workflow-champion-tracking
description: Builds the complete champion tracking workflow — monitors when past customers, engaged prospects, or known contacts change jobs, then generates outreach leveraging the existing relationship. The #1 highest-converting signal type. Produces Clay table config, enrichment prompts, and copy templates. Use when client has an existing customer base or past prospect list to track. Reads from client profile, strategy, and requires a champion contact list from CRM.
---

# Champion Tracking Workflow

Produces a complete Clay workflow for outreach to past relationships who've changed jobs.

**Copy DNA:** Trust and shared history as the hook. These people already know the client's product — the copy doesn't sell, it reconnects. The "new broom" window (first 100 days in a new role) is critical: 70% of budget is spent here. Copy leverages the existing relationship while respecting the new context. This is the most personal trigger — every email must feel like a genuine reconnection, not an automated notification.

**Why this is the #1 signal:** Former customers convert at dramatically higher rates because trust and product knowledge already exist. Combined with a funding event or hiring surge at the new company, this becomes the highest-converting stack in outbound.

**Prerequisites:**
- `client-profiles/{client-slug}/profile.md`
- `strategies/{client-slug}.md`
- Champion contact list from client's CRM (past customers, past engaged prospects, past meeting contacts)
- UserGems account OR willingness to use Clay's job change detection on an uploaded list

## Workflow

### Step 1: Load Client Context

Special attention to:
- **profile.md:** Case studies and results from existing customers — these become proof points in copy
- **strategy.md:** Signal stacking configuration — champion signals feed composite scores
- **CRM data:** The champion contact list with relationship history

### Step 2: Build Clay Table Schema

**Layer 1: Champion Data + New Company Enrichment**

| Column | Type | Purpose |
|---|---|---|
| `champion_first_name` | Import (UserGems/CRM) | Person's name |
| `champion_previous_company` | Import | Where they came from (your client's customer) |
| `champion_previous_title` | Import | Their old role |
| `champion_relationship_type` | Import/CRM | "customer", "prospect", "meeting_had", "engaged" |
| `champion_product_used` | Import/CRM | Which of client's products they used/evaluated |
| `champion_result_achieved` | Import/CRM | Specific result during past relationship (if known) |
| `company_name` | Import (UserGems) | New company they joined |
| `contact_title` | Import (UserGems) | New role title |
| `champion_days_in_role` | Formula | Days since job change — urgency calibration |
| `company_industry` | Enrichment | New company's industry |
| `company_size` | Enrichment | New company's size |
| `score_icp_fit` | AI scoring | Does the NEW company fit client's ICP? |
| `hook_type` | AI prompt | Which hook type for this champion (see hook-types-guide.md) |

**Layer 2: Research on New Company Context**

| Column | Type | Purpose |
|---|---|---|
| `research_new_company` | Claygent | What the new company does, relevant to client's solution |
| `research_role_context` | AI prompt | How their new role relates to the problem client solves |
| `research_transition_insight` | AI prompt | Specific insight about what moving from {old} to {new} means for them |
| `overlap_new_company` | AI prompt | Where client's product fits in the new company's stack/workflow |
| `signal_composite_score` | Formula/Lookup | If signal aggregation table exists — pull composite score |

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md` for complete prompts.

**Critical rules for champion tracking:**
- NEVER reference the old relationship too heavily — they moved on, so should you
- One subtle reference to shared history + forward-looking value at new company
- Urgency comes from the "new broom" window (first 100 days), not artificial pressure
- If `champion_relationship_type` = "customer" and `champion_result_achieved` is known, reference the result naturally
- If `champion_relationship_type` = "prospect" or "engaged", keep it lighter — you have awareness, not a track record
- Always research the NEW company's context — the email is about their future, not your past

### Step 4: Generate Copy Templates

See `references/copy-templates.md`.

**Copy escalation for champions:**
1. **Reconnection** (Day 0) — Genuine congrats + one line connecting old relationship to new context
2. **Value bridge** (Day 5) — Share something useful for their new role: insight, resource, or intro
3. **Soft ask** (Day 14) — "If [pain that client solves] comes up in your new role, happy to jam on it"

**Relationship-tier personalization:**
- **Former customer:** Can reference specific results. Strongest copy. "When we worked together at {previous_company}, your team saw {result}. Curious if that playbook applies to {new_company}'s {context}."
- **Former prospect (meeting had):** Reference the conversation, not the product. "We chatted about {topic} when you were at {previous_company} — given {new_company}'s {context}, that seems even more relevant now."
- **Engaged (no meeting):** Lightest touch. "Your name popped up from when you explored {client product area} at {previous_company}. Congrats on the move."

### Step 5: Quality Gate

- [ ] Copy references the RELATIONSHIP naturally — not forced or awkward
- [ ] New company context is researched — not just "congrats on the move, buy our stuff"
- [ ] If `champion_days_in_role` > 100, adjust urgency down (the "new broom" window is closing)
- [ ] If new company doesn't fit ICP (`score_icp_fit` < 5), mark as SKIP — don't waste a relationship on a bad-fit company
- [ ] hook_type is appropriate — timeline and social proof hooks work best here (they already know the product)
- [ ] Every email stands alone — don't reference previous emails in the cadence
- [ ] Standard copy rules (≤90 words, ≤45 char subject, etc.)
- [ ] Signal stacking: if champion has additional signals (funding, website visit), escalate to Hot tier
