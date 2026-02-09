# Prompt Iteration Pipeline — Systematic Improvement Process

## Why This Exists

Snowflake achieved **15x improvement in reply rates** through systematic prompt iteration — but it took **140+ prompt iterations** to get there. Most teams iterate 3-5 times and call it done. This pipeline provides a structured process for continuous prompt improvement that compounds over time.

**The 10-80-10 Rule (Jason Bay):**
- First 10%: Human strategy (choosing the hook, narrative, and positioning)
- Middle 80%: AI generation at scale (Clay prompt columns)
- Last 10%: Human refinement (reviewing and improving actual outputs)

---

## The Iteration Cycle

```
┌─────────────────────────────────────────────────────┐
│  1. BASELINE: Deploy initial prompts (from skill)   │
│                    ↓                                │
│  2. SAMPLE: Review 25-50 actual outputs             │
│                    ↓                                │
│  3. CATEGORIZE: Sort failures into failure types    │
│                    ↓                                │
│  4. DIAGNOSE: Identify root cause per failure type  │
│                    ↓                                │
│  5. ITERATE: Make ONE targeted prompt change        │
│                    ↓                                │
│  6. TEST: Run on 10 rows, compare to baseline       │
│                    ↓                                │
│  7. VALIDATE: Better? Roll out. Worse? Revert.      │
│                    ↓                                │
│  8. LOG: Record change + result in iteration log    │
│                    ↓                                │
│  (Repeat from Step 2)                               │
└─────────────────────────────────────────────────────┘
```

---

## Step 1: Establish Baseline

Before iterating, document what you're starting with:

```markdown
## Prompt Iteration Log — {client-slug} / {workflow-name}

### Baseline (v1.0)
- Date: {date}
- Prompt: [paste full prompt]
- Sample size: 25 outputs reviewed
- Quality scores:
  - Specificity (1-5): ___  (references real details vs. generic)
  - Relevance (1-5): ___   (matches prospect's actual situation)
  - Human-ness (1-5): ___  (reads like a human wrote it)
  - Accuracy (1-5): ___    (no hallucinated details)
  - Conciseness (1-5): ___ (within word limits, no filler)
- SKIP rate: ___% (how often it outputs SKIP)
- Pass rate: ___% (outputs that pass quality gate without edits)
```

## Step 2: Review Sample Outputs

Pull 25-50 actual Clay outputs. For each, tag:
- ✅ **Pass** — ready to send without edits
- ⚠️ **Needs edit** — fixable with minor human adjustment
- ❌ **Fail** — would damage brand if sent
- 🔄 **SKIP** — correctly identified insufficient data (this is good)

**Target benchmarks:**
- Pass rate: ≥ 70% (ideally 85%+)
- Fail rate: ≤ 5%
- SKIP rate: 10-20% (lower = aggressive, higher = conservative)

## Step 3: Categorize Failures

Every bad output falls into one of these failure types:

| Failure Type | Example | Root Cause |
|---|---|---|
| **Generic output** | "Your company is growing fast" | Prompt lacks specificity constraints or input data is thin |
| **Hallucination** | References a product feature that doesn't exist | Prompt doesn't enforce "only use provided data" |
| **Wrong tone** | Overly formal when client is conversational | Tone instruction missing or too vague |
| **Too long** | 150-word email when limit is 90 | Word limit not enforced or output structure undefined |
| **Template-sounding** | "I noticed [X] and thought [Y]..." pattern obvious | Prompt needs more variation instructions |
| **Wrong angle** | Talks about a pain the prospect doesn't have | Upstream research column fed weak data |
| **AI slop markers** | "In today's fast-paced world...", "I'd love to..." | Banned phrases list incomplete |
| **Irrelevant research** | Claygent found old news not related to trigger | Research prompt scope too broad |
| **Forced personalization** | Awkward shoehorning of a detail that doesn't fit | Prompt forces a variable even when SKIP would be better |
| **Missing context** | Prompt references a variable that was empty | Prompt lacks IF-THEN fallback handling |

## Step 4: Diagnose Root Cause

For each failure cluster, trace back:

1. **Is the failure in the PROMPT or the INPUT DATA?**
   - If 80% of "generic output" failures happen when `research_observation` is thin → fix upstream research prompt first
   - If generic output happens even WITH good input data → fix the copy prompt

2. **Is the failure structural or instructional?**
   - Structural: prompt lacks a section (e.g., no banned phrases list)
   - Instructional: section exists but instruction is vague

3. **Is the failure consistent or random?**
   - Consistent: the prompt reliably produces the same bad pattern → fix the prompt
   - Random: AI variance → add more constraints or examples

## Step 5: Make ONE Change

**The cardinal rule: one change at a time.** If you change three things, you can't attribute improvement.

### Common Fixes By Failure Type

**Generic output:**
```diff
- Write a personalized opener for {{contact_first_name}}.
+ Write a 1-2 sentence opener that references {{research_observation}} 
+ and connects it to how {client_name} helps with {pain_1}.
+ The opener MUST mention a specific detail from the research — 
+ company name, product name, initiative, or metric. 
+ If research is too vague, output SKIP.
```

**Hallucination:**
```diff
+ CRITICAL: Only reference information explicitly present in the data columns
+ above. Do NOT infer, assume, or fabricate any details about the company 
+ or person. If a column is empty, treat it as unavailable — do not guess.
```

**Template-sounding:**
```diff
+ Do NOT start with "I noticed..." or "I saw that..." or "I came across..."
+ Instead, lead with the observation ITSELF as if it's common knowledge:
+ WRONG: "I noticed your team is hiring a VP of Engineering"
+ RIGHT: "A VP Eng hire at this stage usually signals [insight]"
```

**AI slop:**
```diff
+ BANNED OPENINGS: "In today's", "As a leader in", "I wanted to reach out",
+ "Hope this finds you", "I'd love to", "I believe", "I think you'll find",
+ "As someone who", "In the rapidly evolving"
+ If your output starts with any of these, rewrite from scratch.
```

## Step 6: Test on Small Batch

- Run the updated prompt on 10 rows (same rows as before if possible)
- Compare v1 output vs. v2 output side by side
- Score using the same 5 criteria from Step 1

## Step 7: Validate & Roll Out

| Outcome | Action |
|---|---|
| Pass rate improved ≥ 5 percentage points | Roll out to full table |
| Pass rate unchanged | The change was neutral — try a different fix |
| Pass rate decreased | Revert immediately, try different approach |
| SKIP rate increased > 30% | Prompt is too conservative — loosen constraints slightly |

## Step 8: Log the Change

```markdown
### Iteration v1.1
- Date: {date}
- Change: [describe exactly what changed in the prompt]
- Failure type addressed: [which failure type]
- Result: Pass rate {old}% → {new}%
- SKIP rate: {old}% → {new}%
- Decision: Rolled out / Reverted
- Notes: [any observations]
```

---

## Iteration Schedule

| Phase | When | Focus |
|---|---|---|
| **Launch week** | Days 1-3 | Fix critical failures (hallucinations, AI slop) — aim for 0% fail rate |
| **Week 1** | Days 4-7 | Improve specificity — aim for 50% pass rate |
| **Week 2** | Days 8-14 | Refine tone and angles — aim for 70% pass rate |
| **Week 3-4** | Days 15-28 | Optimize for conversion — test hook types, CTAs |
| **Monthly** | Ongoing | Review reply data → iterate based on what actually converts |

---

## Qualitative A/B Testing (SalesBread Method)

Beyond output quality, use REPLY DATA to iterate:

### Step 1: Analyze Negative Replies
Don't just track reply rate — READ the negative replies:
- "Not interested" → angle was wrong, not the execution
- "We already have a solution" → need competitive positioning
- "Not the right time" → trigger timing or signal freshness issue
- "Who are you?" → opener wasn't specific enough
- No reply at all → subject line or deliverability issue

### Step 2: Write Next Variant Based on Objection Pattern
If 40% of negative replies say "we already have a solution":
- Add competitive differentiation to the opener
- Switch to numbers hook showing improvement OVER existing approaches
- Add "even if you're already using [competitor approach]" framing

SalesBread went from **9.8% → 18% → 30% reply rate** using this method.

---

## Prompt Versioning Convention

Store prompt versions in the skill's references directory:

```
workflow-{name}/
  references/
    clay-prompts.md             ← current production version
    prompt-iteration-log.md     ← iteration history + results
```

Each Clay table should tag rows with the prompt version that generated the copy, enabling attribution:
- Column: `meta_prompt_version` — e.g., "v1.3"
- This lets you compare reply rates across prompt versions

---

## Credit-Efficient Testing

- **Test on 5 rows first** — catch obvious failures before burning credits
- **Use Clay Sandbox Mode** where available for iteration
- **Bring Your Own API Keys** (BYO) for OpenAI/Anthropic — 90% cost reduction
- **Use Clay AI Formulas** (free) for data manipulation; reserve Claygent/AI columns for genuine research
- **Use the Clay Metaprompter** to optimize prompts before running at scale
- **Conditional runs:** Only fire expensive prompts when upstream data passes quality checks
