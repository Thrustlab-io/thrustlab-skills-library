# Signal Stacking Guide — Composite Scoring Framework

## Why Signal Stacking

Thrustlab's trigger plays currently operate as independent workflows: website visit → website trigger, job change → job change trigger. Each fires in isolation.

**Signal stacking** combines multiple signals into composite scores with time-decay weighting. One team targeting accounts with 4+ stacked signals achieved **85% response rate and 57% meeting rate**. The math is simple: more signals = higher buying probability = better conversion.

## How It Works In The Thrustlab System

Signal stacking does NOT replace individual trigger plays. It adds a **scoring layer on top** that:
1. Detects when one account has multiple active signals
2. Calculates a composite score with time-decay
3. Escalates high-score accounts to higher-touch treatment
4. Selects the strongest signal combination for copy angle

---

## The Three Signal Layers

### Layer 1: Fit Signals (Static — scored once)

These don't change often. Scored during market mapping and ICP qualification.

| Signal | Points | Source | Decay |
|---|---|---|---|
| Industry match (exact) | +15 | Clay enrichment | None |
| Industry match (adjacent) | +5 | Clay enrichment | None |
| Company size in sweet spot | +10 | Clay enrichment | None |
| Revenue in target range | +5 | Clay enrichment | None |
| Tech stack match (complementary) | +5 | BuiltWith/Wappalyzer | None |
| Geography match | +5 | Enrichment | None |
| Persona title match (exact) | +10 | Enrichment | None |
| Persona title match (adjacent) | +3 | Enrichment | None |
| **Max Layer 1:** | **58** | | |

**Layer 1 threshold:** Must score ≥ 25 (basic ICP fit) before ANY signal processing. Don't waste signal-detection credits on non-ICP companies.

### Layer 2: Intent Signals (Dynamic — time-decayed)

These indicate active buying behavior. Each has a freshness window — older signals are worth less.

| Signal | Base Points | Time Decay | Source |
|---|---|---|---|
| Website visit — pricing page | +25 | -5/week | RB2B, Dealfront, Snitcher |
| Website visit — solution page | +15 | -5/week | RB2B, Dealfront |
| Website visit — case study page | +10 | -3/week | RB2B, Dealfront |
| Website visit — blog only | +5 | -3/week | RB2B, Dealfront |
| G2/review site category research | +20 | -5/week | Bombora, G2 Buyer Intent |
| Competitor evaluation (Claygent) | +15 | -3/week | Claygent research |
| Content download | +10 | -3/week | Marketing automation |
| Webinar/event attendance | +10 | -3/week | Marketing automation |
| Job posting indicating need | +15 | -2/week | Clay signal, LinkedIn Jobs |
| Tech stack change (relevant) | +15 | -2/week | BuiltWith |
| Community engagement (dark funnel) | +10 | -3/week | Teamfluence, Trigify, Common Room |
| Social engagement with client content | +5 | -2/week | Teamfluence, Trigify |
| **Max Layer 2 (fresh):** | **155** | | |

### Layer 3: Relationship Signals (Highest Value)

These indicate pre-existing trust or active buying motion.

| Signal | Base Points | Time Decay | Source |
|---|---|---|---|
| Champion job change (former customer) | +40 | -5/month | UserGems, Clay |
| Champion job change (former prospect) | +25 | -5/month | UserGems, Clay |
| Past meeting (didn't close) | +20 | -3/month | CRM data |
| Replied positively to previous outreach | +15 | -3/month | Sequencer data |
| Competitor customer (known) | +20 | None | Claygent, BuiltWith |
| Funding round (recent) | +15 | -5/month | Crunchbase, Clay signal |
| Rapid hiring surge (>20%) | +10 | -3/month | Clay signal, LinkedIn |
| Compliance deadline approaching | +15 | -5/month | Industry sources |
| **Max Layer 3 (fresh):** | **160** | | |

---

## Composite Score Calculation

```
signal_composite_score = Layer 1 (fit) + Layer 2 (intent, time-decayed) + Layer 3 (relationship, time-decayed)

Maximum theoretical: 373 (all signals, all fresh)
Realistic high-scorer: 80-150 (strong fit + 2-3 active signals)
```

### Score Tiers → Treatment

| Composite Score | Tier | Treatment | Response Channel |
|---|---|---|---|
| ≥ 100 | 🔴 **Hot** | Immediate manual outreach within 24 hours. AE-level personalization. Multi-thread the account. | Slack alert → AE |
| 70-99 | 🟠 **Warm** | Priority automated sequence + human review of copy. Trigger-specific play with enhanced personalization. | Priority Clay workflow |
| 40-69 | 🟡 **Active** | Standard trigger-based play fires normally. Good signal, standard treatment. | Standard Clay workflow |
| 25-39 | 🟢 **Watching** | General outbound eligible. Monitor for additional signals before trigger play. | General outbound queue |
| < 25 | ⚪ **Cold** | Below ICP threshold. Do not outreach. | Do not contact |

---

## Highest-Converting Signal Combinations (From 2025 Data)

These specific stacks have proven conversion rates. When detected, treat as 🔴 Hot regardless of raw score:

### Stack 1: Champion + Funding
**Signals:** Former customer changed jobs to new company + that company recently raised funding
**Why it works:** Trusted contact with budget authority in a growth moment
**Copy angle:** "Congrats on the move AND the round — when we worked together at [previous], you saw [result]. Happy to explore whether that playbook fits [new company]'s growth plans."
**Expected conversion:** 3-5x baseline

### Stack 2: Competitor Customer + Hiring + Intent
**Signals:** Uses competitor product + hiring for role client product supports + visited review sites
**Why it works:** Active evaluation of alternatives with budget allocated
**Copy angle:** "Noticed [company] is building out [team]. Teams moving from [competitor] to [client product] typically see [result] within [timeline]."
**Expected conversion:** 2.5x baseline (3x when mentioning the competitor customer)

### Stack 3: Website Visit (Pricing) + Job Posting + Tech Stack
**Signals:** Visited pricing page + posted job for relevant role + uses complementary tech
**Why it works:** All three indicators of active buying process
**Copy angle:** Lead with timeline hook about implementation speed given their existing tech stack
**Expected conversion:** 2-3x baseline

### Stack 4: Content Engagement + Champion Job Change
**Signals:** Engaged with client content (webinar, blog, download) + known contact from previous interactions
**Why it works:** Pre-educated + pre-existing relationship
**Copy angle:** Reference their specific engagement + personal history
**Expected conversion:** 2-3x baseline

---

## Implementation In Clay

### Architecture: Signal Aggregation Table

In addition to individual trigger tables, create a **master Signal Aggregation Table** per client:

```
Table: {client-slug} — Signal Aggregation
├── Source columns: company_name, company_website, company_domain
├── Fit Score: score_fit (from ICP scoring, calculated once)
├── Intent Score: score_intent (sum of time-decayed intent signals)
├── Relationship Score: score_relationship (sum of time-decayed relationship signals)
├── Composite Score: score_composite (fit + intent + relationship)
├── Signal List: signals_active (comma-separated list of current signals)
├── Signal Count: signals_count (how many active signals)
├── Top Signal: signal_strongest (highest-weighted individual signal)
├── Signal Stack Match: signal_stack_id (which known high-converting stack, if any)
├── Recommended Play: play_recommended (which workflow to activate)
├── Treatment Tier: tier (hot/warm/active/watching/cold)
├── Last Signal Date: signal_last_date (most recent signal timestamp)
└── Export columns: routed to appropriate workflow or Slack alert
```

### How Signals Feed In

Each trigger play table pushes rows to the aggregation table via:
1. **Clay webhook** — when a trigger fires, push company + signal type + timestamp to aggregation table
2. **Scheduled merge** — weekly, reconcile all trigger tables into aggregation table
3. **CRM sync** — pull relationship signals (past meetings, past replies) from CRM

### Composite Score Formula In Clay

Use a Clay AI Formula or code column:

```
score_composite = score_fit + score_intent + score_relationship

WHERE:
  score_fit = [static ICP score from enrichment]
  score_intent = SUM(each intent signal × time_decay_multiplier)
  score_relationship = SUM(each relationship signal × time_decay_multiplier)

  time_decay_multiplier = MAX(0, 1 - (days_since_signal / decay_period_days))
```

### Routing Logic

```
IF score_composite >= 100 OR signal_stack_id IN ("champion_funding", "competitor_hiring_intent"):
  → Slack notification to AE/SDR with full signal context
  → Route to manual outreach queue
  → Copy tier: Tier 3 (white-glove, 5+ min per prospect)

ELSE IF score_composite >= 70:
  → Activate strongest signal's trigger play with enhanced variables
  → Add signal_context variable to prompts (all active signals, not just the trigger)
  → Copy tier: Tier 2 (semi-manual, 2 min review)

ELSE IF score_composite >= 40:
  → Standard trigger play fires with normal automation
  → Copy tier: Tier 1 (fully automated)

ELSE IF score_composite >= 25:
  → General outbound queue (no trigger play, ICP-fit only)

ELSE:
  → Do not contact. Flag for ICP review.
```

---

## Integration With GTM Strategy Generator

When generating a client strategy (Section 10: Trigger Playbook), the strategy MUST now include:

1. **Signal stacking configuration:** Which signals to monitor and their client-specific weights
2. **Known high-converting stacks:** Which signal combinations matter most for this client's market
3. **Routing rules:** What treatment each tier gets (aligned with client's sales capacity)
4. **Alert configuration:** Who gets notified at each tier, via which channel

The strategy template includes a Signal Stacking section that reads the client's active trigger plays and constructs the composite scoring formula.

---

## Important Constraints

- Signal stacking requires at minimum 2 active trigger plays to be meaningful
- Start simple: fit + 1 intent layer. Add complexity as data proves the model
- Recalibrate weights quarterly based on actual conversion data
- Don't over-index on score — a single 🔴 hot signal (champion + funding) beats a dozen weak signals
- Signal stacking is a ROUTING decision, not a copy decision — the copy still comes from the specific trigger play's templates
