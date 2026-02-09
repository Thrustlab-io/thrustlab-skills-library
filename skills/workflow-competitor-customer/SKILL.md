---
name: workflow-competitor-customer
description: Builds the complete competitor customer targeting workflow — identifies companies using competitor products and generates competitive displacement outreach. Produces Clay table config with competitor detection enrichment, competitive positioning prompts, and copy that leverages existing customer social proof. 2.5x higher conversion when targeting competitor users, 3x when mentioning your existing customer. Use when client has identifiable competitors with detectable usage. Reads from client profile, strategy, and competitive battlecards.
---

# Competitor Customer Targeting Workflow

Produces a complete Clay workflow for outreach to companies currently using a competitor product.

**Copy DNA:** Competitive intelligence as the hook. These prospects already understand the problem space — they're actively paying for a solution. The copy doesn't need to educate on the pain; it needs to create doubt about the current choice and offer a credible alternative. Data shows **3x conversion when mentioning an existing customer who switched**.

**The key insight:** Competitor customers are NOT cold. They've already bought in the category. Your job is to make them re-evaluate, not discover.

**Prerequisites:**
- `client-profiles/{client-slug}/profile.md` (especially competitor section and case studies)
- `strategies/{client-slug}.md` (especially competitive battlecards in Appendix A)
- `client-profiles/{client-slug}/icp-mapping.md`
- Knowledge of client's top 3-5 competitors
- At least 1 case study of a customer who switched FROM a competitor

## Workflow

### Step 1: Load Client Context

Special attention to:
- **profile.md → Competitors section:** Names, positioning differences, where client wins
- **strategy.md → Appendix A (Competitive Battlecards):** What to say per competitor, where client wins, proof points
- **Case studies:** Customers who switched FROM a competitor — gold for social proof hooks

### Step 2: Build Clay Table Schema

**Layer 1: Company + Competitor Detection**

| Column | Type | Purpose |
|---|---|---|
| `company_name` | Import/Enrichment | Target company |
| `company_industry` | Enrichment | Sector classification |
| `company_size` | Enrichment | Size qualification |
| `contact_first_name` | Import/Enrichment | Decision maker |
| `contact_title` | Import/Enrichment | Must be decision-maker for the tool category |
| `competitor_product_used` | Claygent/BuiltWith | Which competitor product they use |
| `competitor_confidence` | AI prompt | "confirmed" / "likely" / "possible" / "none" |
| `competitor_detection_source` | Formula | How we know: BuiltWith, job posting, Claygent, G2 review |
| `score_icp_fit` | AI scoring | Standard ICP qualification |
| `hook_type` | AI prompt | Hook type classification (social proof strongest here) |

**Layer 2: Competitive Intelligence Research**

| Column | Type | Purpose |
|---|---|---|
| `research_competitor_usage` | Claygent | How they use the competitor product — workflows, scale |
| `research_competitor_pain` | AI prompt | Known limitations of {{competitor_product_used}} relevant to this company |
| `research_switch_angle` | AI prompt | Most compelling reason THIS company would switch |
| `competitor_customer_name` | Lookup/Static | {client_name} customer who switched from same competitor |
| `competitor_customer_result` | Lookup/Static | Result the switcher achieved |
| `signal_composite_score` | Formula/Lookup | From signal aggregation table |

### Step 3: Generate Clay Prompts

See `references/clay-prompts.md`.

**Critical rules for competitor targeting:**
- NEVER trash the competitor — prospects feel defensive about their choice
- Acknowledge competitor's strength, differentiate on specifics
- If `competitor_confidence` = "possible" or "none" → route to general outbound instead
- Best angle: "companies who moved from [competitor] to [client] saw [result]"
- Pair switch story with timeline hook — "took X weeks, not months"
- No switch case study? Use numbers hook with differentiated metric

### Step 4: Generate Copy Templates

See `references/copy-templates.md`.

**Copy escalation for competitor customers:**
1. **Competitive seed** (Day 0) — Plant doubt without attacking: share a result competitor can't match
2. **Switch story** (Day 4) — Name-drop a customer who switched + their result
3. **Differentiation deep dive** (Day 8) — Specific capability comparison relevant to THEIR use case
4. **Soft close** (Day 14) — "Even if you're happy with [competitor], worth seeing what's changed"

### Step 5: Quality Gate

- [ ] `competitor_confidence` is "confirmed" or "likely" — never send competitive copy to "possible"/"none"
- [ ] Copy never disparages the competitor — acknowledge, differentiate, prove
- [ ] Named customer in social proof is REAL and result VERIFIED
- [ ] Switch angle specific to THIS prospect's use case, not generic
- [ ] hook_type appropriate — social_proof strongest when switch case study exists
- [ ] Standard copy rules (≤90 words, ≤45 char subject, etc.)
- [ ] Signal stacking: competitor + other signals (funding, hiring, website visit) → escalate
