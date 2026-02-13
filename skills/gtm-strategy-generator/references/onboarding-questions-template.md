# Onboarding Questions Template

This template is used to generate the `onboarding-questions.md` file for each new client. The questions are derived from the auto-discovered profile (`[To be confirmed by client]` fields) and the strategy (play-specific validation).

**Every onboarding-questions.md must follow this structure exactly.** Adapt the specific questions based on the client's profile and strategy — do NOT use generic placeholders. Each question should reference what was discovered or assumed.

---

## Document Structure

```markdown
# {Company Name} — Onboarding Meeting Questions

**Purpose:** Questions to cover during the client onboarding meeting. These confirm assumptions from the auto-discovered profile, validate the strategy, and gather information that can only come from the client directly.

**Meeting with:** [Client contact name(s)]
**Date:** [Meeting date]

---

## 1. Profile Confirmation

These fields were marked `[To be confirmed by client]` in the profile. Quick confirmations needed:

### Sales Motion
{Generate questions for every [Client] and [Infer] field in the Sales Motion section of profile.md. Reference what was assumed.}

### Target Market
{Generate questions for every [Client] and [Infer] field in the Target Market section. Reference what was assumed.}

### Personas
{Generate questions for every [Client] and [Infer] field in the Personas section. Reference what was assumed.}

### Engagement Parameters
{Generate questions about outreach preferences, blackout periods, meeting preferences, response SLAs.}

---

## 2. Strategy Validation

### Trigger Plays — Confirm Our Approach
{Generate questions for each of the 3 recommended plays from strategy.md. Ask:
- Does this trigger/approach resonate?
- Any concerns or adjustments?
- Do you have access to the required data sources?
- Any prerequisites we should know about?}

### Messaging & Tone
{Questions about:
- Outreach language (which languages needed?)
- Sender persona (who should outreach come from?)
- Tone validation (does our recommended tone resonate?)
- Competitor mention policy}

---

## 3. Product & Proof Points

{Questions about:
- Current customer count and growth trajectory
- Customer success stories available for outreach
- Specific testimonials that can be used
- Product capabilities relevant to the strategy
- Pricing confirmation and any upcoming changes
- Trial/demo/guarantee offers}

---

## 4. Tech Stack & Infrastructure

{Questions about:
- CRM system and data quality
- Historical contacts in CRM (reactivation opportunity)
- Email sequencer availability
- LinkedIn access and team profiles
- Existing signal/intent tools
- Databases in use
- Other GTM tooling}

---

## 5. Existing Content & Assets

{Questions about:
- Content library size and topics
- Publishing cadence and who creates content
- Lead magnets available or needed
- Social media presence beyond LinkedIn
- Reviews on G2/Capterra/Trustpilot/Google
- Sales collateral (decks, battlecards, FAQs)
- Customer logos available for outreach}

---

## 6. Operations & Capacity

{Questions about:
- Who handles booked meetings
- Meeting capacity per week
- Meeting booking tool/link
- Follow-up process after meetings
- Physical mail capacity (if letter outreach planned)
- Budget for outbound tools
- Who manages daily outbound operations}

---

## 7. Quick Wins & Priorities

{Questions about:
- Existing warm leads or pipeline
- Past trial users for reactivation
- Referral programs or network opportunities
- Upcoming events, deadlines, or launches}

---

## 8. Success Metrics & Expectations

{Questions about:
- Primary KPI at 3 months and 6 months
- Reporting cadence preference
- Reporting format preference
- Decision-maker for strategy changes}

---

## Action Items After Meeting

| Action | Owner | Due Date |
|--------|-------|----------|
| Confirm/correct profile.md fields | Client | |
| Share CRM access / API keys | Client | |
| Provide customer testimonials & logos | Client | |
| Share booking link | Client | |
| Identify sender persona for outreach | Client | |
| Set up Clay workspace | Thrustlab | |
| {Play-specific setup action 1} | Thrustlab | |
| {Play-specific setup action 2} | Thrustlab | |
| Draft outreach sequences for review | Thrustlab | |

---

*Generated: {date}*
*Based on: {Company Name} GTM Strategy*
```

---

## Key Principles

1. **Be specific, not generic.** Every question should reference what was actually discovered or assumed. "We assumed your sales cycle is X — is that correct?" not "What is your sales cycle?"

2. **Checkbox format.** Every question is a `- [ ]` checkbox so it can be tracked during the meeting.

3. **Prioritize.** The most critical questions (blocking implementation) should come first within each section.

4. **Include context.** Where a question might seem odd, add brief context: "We need this because..."

5. **Action items table.** Always end with a concrete action items table. Pre-populate with known actions based on the strategy.

6. **Play-specific questions.** The Strategy Validation section must include questions specific to each recommended play — not generic trigger questions.
