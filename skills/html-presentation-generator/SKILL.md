---
name: html-presentation-generator
description: Generates a beautiful, modern, dynamic HTML presentation that summarizes the GTM strategy. Uses the strategy.md document as input to create a concise, visually appealing presentation suitable for sharing with prospects. Outputs a self-contained HTML file with embedded styles and interactive elements.
---

# HTML Presentation Generator

Transforms the detailed GTM strategy document into a beautiful, shareable HTML presentation for prospect review.

**Prerequisites:**
- `Prospects/{prospect-slug}/strategy.md` exists
- `Prospects/{prospect-slug}/profile.md` exists (for company branding context)

**Output:**
- `Prospects/{prospect-slug}/presentation.html` — self-contained HTML presentation

---

## Design Philosophy

The presentation should be:
- **Concise** — Distills the strategy into key insights, not a wall of text
- **Modern** — Clean design with smooth animations and responsive layout
- **Dynamic** — Interactive elements, smooth scrolling, progressive disclosure
- **Beautiful** — Professional aesthetics that reflect well on Thrustlab
- **Self-contained** — Single HTML file with embedded CSS and minimal JavaScript

---

## Workflow

### Step 1: Read Source Documents

Load the strategy and profile documents:
1. `Prospects/{prospect-slug}/strategy.md` — full strategy content
2. `Prospects/{prospect-slug}/profile.md` — company name, industry, personas for context

### Step 2: Extract Key Content

Distill the strategy into presentation-ready sections:

**Title slide:**
- Company name + tagline: "GTM Strategy by Thrustlab"
- Date generated

**Executive Summary (1 slide):**
- 3-4 bullet points capturing the core opportunity and approach
- High-level ICP description
- Market context in 1-2 sentences

**ICP & Positioning (1-2 slides):**
- Target personas with brief descriptions
- Key pain points (top 3-4)
- Value proposition summary
- Competitive differentiation in 2-3 points

**Recommended Trigger Plays (2-3 slides):**
- Top 3 recommended triggers with:
  - Trigger name and description
  - Why it's effective for this prospect
  - Example signal/scenario
  - Expected impact

**Signal Strategy (1 slide):**
- Signal stacking approach overview
- Key data sources and tools
- High-value signal combinations (2-3 examples)

**90-Day Roadmap (1 slide):**
- Phase breakdown with key milestones
- Timeline visualization
- Success metrics

**Next Steps (1 slide):**
- Clear action items for prospect approval
- What Thrustlab will deliver
- Timeline expectations

### Step 3: Design & Generate HTML

Create a single HTML file using modern web standards:

**Technical specifications:**
- Single self-contained HTML file
- Embedded CSS (no external stylesheets)
- Minimal vanilla JavaScript for interactions (no frameworks required)
- Responsive design (mobile, tablet, desktop)
- Dark mode option
- Smooth scroll navigation
- Progressive disclosure (expand/collapse sections)

**Visual design:**
- Clean, professional typography (system fonts for speed)
- Generous whitespace
- Subtle animations (fade-in on scroll, hover effects)
- Color scheme: Professional palette (blues/grays, accent color)
- Icons for key points (Unicode emoji or inline SVG)
- Progress indicator showing current section
- Navigation: Fixed header with section links

**Layout structure:**
- Full-screen vertical scrolling (one concept per viewport)
- Section-based navigation
- Smooth scroll-snapping between sections
- Footer with Thrustlab branding

**Interaction patterns:**
- Click to expand detailed explanations
- Hover effects on interactive elements
- Keyboard navigation support (arrow keys to navigate sections)
- Print-friendly CSS (media query for clean PDF export)

### Step 4: Save & Output

1. Write the complete HTML to `Prospects/{prospect-slug}/presentation.html`
2. Validate that the file is self-contained (no external dependencies)
3. Provide the file path to the user
4. Include instructions for:
   - Opening in browser (double-click)
   - Sharing with prospect (email attachment or host on web)
   - Exporting to PDF (browser print → Save as PDF)

---

## Quality Checklist

Before delivering, verify:
- [ ] All content is prospect-specific (no generic placeholders)
- [ ] HTML file is self-contained (no broken references)
- [ ] Presentation looks professional on desktop, tablet, and mobile
- [ ] All interactive elements work (expand/collapse, navigation)
- [ ] Total presentation length is 8-12 "slides" (viewport sections)
- [ ] Load time is fast (<2 seconds on standard connection)
- [ ] Print/PDF export looks clean
- [ ] Thrustlab branding is present but subtle

---

## Example HTML Structure

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{Company Name} - GTM Strategy | Thrustlab</title>
    <style>
        /* Embedded CSS here */
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; }
        section { min-height: 100vh; display: flex; align-items: center; justify-content: center; }
        /* Additional styles... */
    </style>
</head>
<body>
    <nav><!-- Fixed navigation --></nav>

    <section id="title">
        <!-- Title slide content -->
    </section>

    <section id="summary">
        <!-- Executive summary -->
    </section>

    <!-- Additional sections... -->

    <footer>
        <p>Generated by Thrustlab | {Date}</p>
    </footer>

    <script>
        // Minimal vanilla JS for interactions
        // Smooth scroll, expand/collapse, keyboard nav
    </script>
</body>
</html>
```

---

## Notes

- The presentation should feel like a curated executive briefing, not a document dump
- Prioritize clarity over comprehensiveness — leave detailed execution for the full strategy doc
- The prospect should be able to review this in 10-15 minutes and understand the approach
- Design should be modern but not flashy — trustworthy over trendy
