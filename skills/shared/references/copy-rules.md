# Thrustlab Copy Rules — Universal Standards

Every skill that produces outbound copy (emails, LinkedIn messages, call scripts) must follow these rules. No exceptions.

## Email Rules

### Subject Line
- ≤45 characters
- Lowercase preferred (no Title Case unless brand-specific)
- No spam triggers: "free", "guaranteed", "act now", "limited time"
- Pattern: observation or question, never a pitch
- Examples: "saw your {trigger}" / "{first_name}, quick q about {topic}" / "re: {company}'s {trigger}"

### Body
- ≤90 words total (opener + body + CTA combined)
- 3 paragraphs max: opener (1-2 sentences) → value bridge (2-3 sentences) → CTA (1 sentence)
- No bullet points in cold emails
- No "I hope this finds you well" or any pleasantry opener
- Write at 8th grade reading level — short sentences, simple words

### Opening Line — Hook Type System
- ALWAYS observational — reference something specific about the person, company, or trigger
- Never generic ("I came across your profile", "I wanted to reach out", "I'm reaching out because")
- Two types per workflow:
  - **Company-based opener**: Based on company research (works without trigger)
  - **Trigger-based opener**: Based on the specific signal that activated the workflow
- Must reference a real, verifiable detail — never fabricate
- **Every opener must be generated in one of four hook types** (see `shared/references/hook-types-guide.md`):
  - **Timeline** (default, highest performer @ 10% reply rate): Compressed achievement milestone
  - **Numbers** (second @ 8.6%): Specific quantified claim relevant to their role
  - **Social proof** (third @ 6.5%): Named reference company + their result
  - **Hypothesis** (fourth @ 4.3%): Pain framed as a question ("Would I be right...")
- The `hook_type` variable determines which pattern to use per prospect — set in Layer 1.5 of the enrichment chain
- Timeline is the default. Fall back to hypothesis only when timeline/numbers/social proof data is unavailable.

### CTA
- Stage-appropriate — match where the prospect is:
  - Cold/unaware: "worth a quick sanity check?" / "open to exploring this?"
  - Warm/triggered: "worth 15 min to see if this fits?" / "happy to share how {similar_company} handled this"
- Never: "book a demo", "schedule a call", "let's connect" (too aggressive for cold)
- Always a question, never a command
- One CTA per email — never two asks

### Cadence Structure
- 3-5 steps over 10-14 business days
- Step 1: Primary value prop + trigger/observation
- Step 2: Different angle (social proof, case study, or different pain)
- Step 3: Breakup or new insight (never "just following up" or "bumping this")
- Optional Step 4-5: LinkedIn touchpoint or new trigger reference
- Each step must stand alone — assume they never read previous steps

## LinkedIn Rules

### Connection Request
- ≤280 characters total
- First line: specific observation (not "I'd love to connect")
- No pitch in connection request — earn the conversation first
- Pattern: observation → soft bridge → reason to connect

### LinkedIn Message (post-connection)
- ≤150 words
- More conversational than email
- Reference something from their profile or recent activity
- Soft CTA only ("curious if..." / "would love your take on...")

## Call Script Opener
- ≤30 seconds when read aloud
- Pattern: "Hi {first_name}, this is {sender} from {client_name}. I noticed {specific_observation} — [one sentence value bridge]. Is that something you're thinking about?"
- Never start with "How are you?" or "Is this a good time?"

## Banned Words & Phrases

Never use in any outbound copy:
- "innovative", "cutting-edge", "game-changing", "synergy", "leverage", "disruptive"
- "circle back", "touch base", "loop in", "ping"
- "I hope this email finds you well"
- "I wanted to reach out because"
- "I came across your profile/company"
- "As a leader in..."
- "We help companies like yours..."
- "Would love to pick your brain"
- "Just checking in" / "Just following up" / "Bumping this"

## Tone Calibration

The client's `profile.md` specifies tone preference. Apply accordingly:
- **Formal**: Professional but not stiff. Full sentences. No slang.
- **Conversational**: Contractions OK. Shorter sentences. Feels like a colleague.
- **Provocative**: Challenge assumptions. Bold claims backed by data. Pattern interrupt.

Default to **conversational** if tone is unspecified.

## Personalization Minimum Bar

Before any copy ships, it must pass this test:
> "Could this email have been sent to 1,000 other people and still make sense?"
> If yes → it's not personalized enough → rewrite.
